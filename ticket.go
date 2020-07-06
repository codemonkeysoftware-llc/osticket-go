package osticket

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// CreateTicketCommand contains the fields allowed for making commands. Alert
// and Autorespond are optional states and can be set using Should() or ShouldNot().
// If they are not set, osTicket will use default values.
type CreateTicketCommand struct {
	Email       string        `json:"email"`
	Name        string        `json:"name"`
	Subject     string        `json:"subject"`
	Message     Message       `json:"message"`
	Alert       OptionalState `json:"alert,omitempty"`
	Autorespond OptionalState `json:"autorespond,omitempty"`
	IPAddress   string        `json:"ip,omitempty"`
	Priority    string        `json:"priority,omitempty"`
	Source      string        `json:"source,omitempty"`
	TopicID     string        `json:"topicId,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

// Message is used to create the data url for sending a message. If ContentType
// Type is not set, the request will fail.
type Message struct {
	ContentType ContentType
	Body        string
}

// MarshalJSON constructs the data URL for the Message field.
func (m *Message) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("data:%s,%s", m.ContentType, m.Body)

	return json.Marshal(s)
}

// ContentType is used to determine the data url content type part
type ContentType string

// ContentType options
const (
	ContentTypeHTML  ContentType = "text/html"
	ContentTypePlain ContentType = "text/plain"
)

// OptionalState represents a state that if not set will use a default
type OptionalState *bool

// Should sets an OptionalState to true
func Should() OptionalState {
	b := true
	return &b
}

// ShouldNot sets an OptionalState to false
func ShouldNot() OptionalState {
	b := false
	return &b
}

// NewAttachment creates a new Attachments for creating a ticket. This method
// should be preferred over creating an attachment by hand, since it handles
// setting Encoding, which will always be `base64`
func NewAttachment(name string, mimeType string, data []byte) *Attachment {
	return &Attachment{
		Name:     name,
		Data:     data,
		MimeType: mimeType,
		Encoding: "base64",
	}
}

// Attachment is a file to be attached to a request
type Attachment struct {
	Name     string
	Data     []byte
	MimeType string
	Encoding string
}

func (a *Attachment) MarshalJSON() ([]byte, error) {
	s := &strings.Builder{}
	fmt.Fprintf(s, "data:%s;%s,", a.MimeType, a.Encoding)
	e := base64.NewEncoder(base64.StdEncoding, s)
	_, err := e.Write(a.Data)
	if err != nil {
		return nil, err
	}

	err = e.Close()
	if err != nil {
		return nil, err
	}

	m := map[string]string{
		a.Name: s.String(),
	}

	return json.Marshal(m)
}
