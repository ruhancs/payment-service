package worker

//identificador do tipo da tarefa no redis
const TaskSendEmail = "task:send_email"

// dados da tarefa para armazenar no redis
type PayloadSendEmail struct {
	Email string `json:"email"`
}