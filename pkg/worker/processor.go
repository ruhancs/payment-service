package worker

import (
	"context"
	//"database/sql"
	"encoding/json"
	"fmt"
	"payment-service/internal/domain/gateway"
	"payment-service/internal/infra/mail"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type TaskProcessorInterface interface {
	Start() error
	ProcessTaskSendEmail(ctx context.Context, task *asynq.Task) error
}

// processador das tarefas na fila do redis
type RedisTaskProcessor struct {
	server             *asynq.Server
	customerRepository gateway.CustomerRepositoryInterface
	mailer             mail.EmailSenderInterface
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, custRepository gateway.CustomerRepositoryInterface, mailer mail.EmailSenderInterface) TaskProcessorInterface {
	logger := NewLogger()
	redis.SetLogger(logger)
	server := asynq.NewServer(
		redisOpt,
		//controle dos parametros dos processos asincronos,ex: maximo de processos em paralelo, quantidade de tentaivas
		asynq.Config{
			//tratamento de erros retornado apos todas tentativas de reprocessamento
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Error().Err(err).Str("type", task.Type()).
					Bytes("payload", task.Payload()).Msg("process task failed")
			}),
			Logger: logger,
		},
	)
	return &RedisTaskProcessor{
		server:             server,
		customerRepository: custRepository,
		mailer:             mailer,
	}
}

func (processor *RedisTaskProcessor) ProcessTaskSendEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendEmail
	err := json.Unmarshal(task.Payload(), &payload)
	if err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	user, err := processor.customerRepository.GetByEmail(ctx, payload.Email)
	if err != nil {
		//if err == sql.ErrNoRows {
		//	return fmt.Errorf("user not found: %w", asynq.SkipRetry)
		//}
		return fmt.Errorf("failed to get user: %w", err)
	}

	//ENVIAR EMAIL PARA O USUARIO
	subject := "payment successfuly"
	content := fmt.Sprintf(`Hello %s,<br/>
	Thank you your payment was registering succesfuly!<br/>
	`, user.FirstName)
	to := user.Email
	err = processor.mailer.SendEmail(subject,content,to,nil)
	if err != nil {
		return fmt.Errorf("failed to send confimation payment email: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", user.Email).Msg("processed task")

	return nil
}

// registrar tarefaz e acao da tarefa
func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	//TaskSendEmail=nome da tarefa, processor.ProcessTaskSendEmail = execucao da tarefa
	mux.HandleFunc(TaskSendEmail, processor.ProcessTaskSendEmail)

	return processor.server.Start(mux)
}
