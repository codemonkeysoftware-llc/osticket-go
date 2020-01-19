[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/codemonkeysoftware/osticket-go) [![GitHub](https://img.shields.io/github/license/codemonkeysoftware/osticket-go)](https://raw.githubusercontent.com/codemonkeysoftware/osticket-go/master/LICENSE)

# osticket-go
Go client for [osTicket](https://osticket.com/) API


## Usage

Below is an example of how to create a ticket using osticket-go
```go
func newTicket() error {
  message := `Hi Support,
  I've been having trouble with your application. Can you please help me?

  Thanks,
  Example Person`

  cmd := &osticket.CreateTicketCommand{
    Email:   "test@example.com",
    Name:    "Example Person",
    Subject: "Test Subject",
    Message: osticket.Message{
      ContentType: osticket.ContentTypePlain,
      Body:        message,
    },
  }

  apiClient := osticket.NewAPIClient(http.DefaultClient, `https://example.com/support`, `APIKEY`)
  return apiClient.CreateTicket(cmd)
}
```

## Notes

The api client currently calls {baseurl}/api/http.php/{endpoint}.json, since calling the endpoint directly doesn't work on a clean install of osTicket.