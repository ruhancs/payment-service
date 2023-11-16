package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"payment-service/internal/application/factory"
	"payment-service/internal/domain/gateway"
	"payment-service/internal/infra/mail"
	"payment-service/internal/infra/repository"
	"payment-service/internal/infra/web"
	"payment-service/pkg/worker"

	"github.com/hibiken/asynq"
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

	customerRepo := repository.NewCustomerRepository(conn)

	getOrderUseCase := factory.GetOrderUsecaseFactory(conn)
	paymentUseCase := factory.PaymentUseCaseUsecaseFactory(conn)
	var srv = &http.Server{}
	redisOpt := asynq.RedisClientOpt{
		Addr: os.Getenv("REDIS_ADDRESS"),
	}
	taskDist := worker.NewRedisTaskDistributor(redisOpt)
	app := web.NewApplication(getOrderUseCase,paymentUseCase,srv,taskDist)
	
	go runTaskProcessor(redisOpt,customerRepo)
	errChan := make(chan error)
	go app.GracefullyShutdown(errChan)

	if err := <- errChan; err != nil {
		fmt.Println(err)
	}
}

func runTaskProcessor(redisOpt asynq.RedisClientOpt, customerRepo gateway.CustomerRepositoryInterface) {
	mailer := mail.NewGmailSender(os.Getenv("EMAIL_SENDER_NAME"),os.Getenv("EMAIL_ADDRESS"),os.Getenv("EMAIL_PASS"))
	processor := worker.NewRedisTaskProcessor(redisOpt,customerRepo,mailer)
	log.Println("Start task processor...")
	err := processor.Start()
	if err != nil {
		log.Fatal(err)
	}
}