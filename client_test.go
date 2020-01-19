package osticket_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codemonkeysoftware/osticket-go"
	"github.com/stretchr/testify/assert"
)

func TestCreateTicket(t *testing.T) {
	cmd := &osticket.CreateTicketCommand{
		Email:   "test@example.com",
		Name:    "Example Person",
		Subject: "Test Subject",
		Message: osticket.Message{
			ContentType: osticket.ContentTypePlain,
			Body:        "This is a test message from the osticket-go package. Does this work? 1 + 1 = 2",
		},
	}

	s := httptest.NewServer(createdHandler)
	defer s.Close()
	apiClient := osticket.NewAPIClient(http.DefaultClient, s.URL, `APIKEY`)
	err := apiClient.CreateTicket(cmd)
	assert.NoError(t, err)
}

type staticHandler int

const (
	createdHandler staticHandler = http.StatusCreated
)

func (s staticHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(int(s))
}
