package slack

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/parnurzeal/gorequest"
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Attachment struct {
	Fallback   *string  `json:"fallback"`
	Color      *string  `json:"color"`
	PreText    *string  `json:"pretext"`
	AuthorName *string  `json:"author_name"`
	AuthorLink *string  `json:"author_link"`
	AuthorIcon *string  `json:"author_icon"`
	Title      *string  `json:"title"`
	TitleLink  *string  `json:"title_link"`
	Text       *string  `json:"text"`
	ImageUrl   *string  `json:"image_url"`
	Fields     []*Field `json:"fields"`
}

func (attachment *Attachment) AddField(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, &field)
	return attachment
}

func Payload(text, username, imageOrIcon, channel string, attachments []Attachment) map[string]interface{} {
	payload := make(map[string]interface{})
	payload["parse"] = "full"
	if username != "" {
		payload["username"] = username
	}

	if strings.HasPrefix("http", imageOrIcon) {
		payload["icon_url"] = imageOrIcon
	} else if imageOrIcon != "" {
		payload["icon_emoji"] = imageOrIcon
	}

	if channel != "" {
		payload["channel"] = channel
	}

	if text != "" {
		payload["text"] = text
	}

	if len(attachments) > 0 {
		payload["attachments"] = attachments
	}

	return payload
}

func Send(webhookUrl string, proxy string, payload map[string]interface{}) []error {
	data, _ := json.Marshal(payload)
	request := gorequest.New().Proxy(proxy)
	_, _, err := request.
		Post(webhookUrl).
		Send(string(data)).
		End()

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}
}
