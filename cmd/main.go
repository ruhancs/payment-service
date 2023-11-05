package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"payment-service/internal/application/dto"
	"payment-service/internal/application/factory"
	"payment-service/internal/infra/web"
	"payment-service/pkg/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	conn,err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DSN"))
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVER"),
		"group.id": "payment",
	}
	kafkaConsumer := kafka.NewConsumer(&configMap,[]string{"create-order"})

	msgChan := make(chan *ckafka.Message)
	kafkaConsumer.Consume(msgChan)

	createOrderUseCase := factory.CreateOrderUsecaseFactory(conn)

	go func ()  {
		for msg := range msgChan{
			var createOrderPayload dto.CreateOrderInputDto
			err := json.Unmarshal(msg.Value,&createOrderPayload)
			fmt.Println(createOrderPayload)
			if err != nil {
				fmt.Println("invalid payment data")
				continue
			}
			createOrderUseCase.Execute(context.Background(),createOrderPayload)
		} 	
	}()

	getOrderUseCase := factory.GetOrderUsecaseFactory(conn)
	updateOrderUseCase := factory.UpdateOrderStatusUsecaseFactory(conn)
	app := web.NewApplication(getOrderUseCase,updateOrderUseCase)

	app.Server()
}