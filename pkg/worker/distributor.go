package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

type TaskDistributorInterface interface {
	DistributeTaskSendEmail(ctx context.Context,payload *PayloadSendEmail,opts ...asynq.Option) error
}

// envia tarefaz para serem executadas no redis, asincronamente
type RedisTaskDistributor struct {
	client *asynq.Client
}

func(distributor *RedisTaskDistributor) DistributeTaskSendEmail(ctx context.Context,payload *PayloadSendEmail,opts ...asynq.Option) error{
	jsonPayload,err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w",err)
	}
	//opts sao opcoes de como a tarefa sera distribuida
	task := asynq.NewTask(TaskSendEmail,jsonPayload, opts...)
	//enviar a tarefa para a fila do redis
	taskInfo,err := distributor.client.EnqueueContext(ctx,task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task")
	}
	log.Println("type: "+ taskInfo.Type+ " payload: " +string(taskInfo.Payload)+ " queue: "+taskInfo.Queue+ " enqueued task")
	return nil
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributorInterface {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client: client,
	}
}