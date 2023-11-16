package mail

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	err := godotenv.Load("../../../.env")
	require.NoError(t, err)

	sender := NewGmailSender(os.Getenv("EMAIL_SENDER_NAME"), os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_PASS"))

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="http://test.guru">Payment Successfuly</a></p>
	`
	to := "test.guru@gmail.com"
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject,content,to,attachFiles)
	require.NoError(t, err)
}