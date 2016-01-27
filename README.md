[![GoDoc](https://godoc.org/github.com/ashwanthkumar/slack-go-webhook?status.svg)](https://godoc.org/github.com/ashwanthkumar/slack-go-webhook)

# slack-go-webhook

Go Lang library to send messages to Slack via Incoming Webhooks. 

## Usage
```go
package main

import "github.com/ashwanthkumar/slack-go-webhook"

func main() {
  attachment1 := slack.Attachment {}
  attachment1.
    AddField(slack.Field { Title: "Author", Value: "Ashwanth Kumar" }).
    AddField(slack.Field { Title: "Status", Value: "Completed" })

  payload := slack.Payload("Hello from <https://github.com/ashwanthkumar/slack-go-webhook|slack-go-webhook>, a Go-Lang library to send slack webhook messages.", 
                           "golang-bot", 
                           "",
                           "golang-test",
                           []slack.Attachment { attachment1 })


  slack.Send("https://hooks.slack.com/services/foo/bar/baz", payload)
}

```

## License
Licensed under the Apache License, Version 2.0: http://www.apache.org/licenses/LICENSE-2.0

