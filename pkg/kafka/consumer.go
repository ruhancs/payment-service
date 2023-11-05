package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

//topicos é onde o consumidor se inscrevera para receber os resultados
func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(msgChan chan *ckafka.Message) {
	fmt.Println("START CONSUMER")
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}
	//todos resultados recebidos sao enviados para os inscritos no topico
	go func() {
		for {
			msg, err := consumer.ReadMessage(-1)
			if err == nil {
				fmt.Println("MSG RECEIVEDE ON TOPIC")
				//envia para um canal de thread para ler uma thread consome os dados do kafka e outra sobe o webserver
				msgChan <- msg
				continue 
			}
			fmt.Println("ERROR TO RECEIVE")
		}
	}()
}