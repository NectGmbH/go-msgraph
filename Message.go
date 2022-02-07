package msgraph

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Message struct {
	From         Recipient   `json:"from"`
	Subject      string      `json:"subject"`
	Importance   string      `json:"importance"`
	IsRead       bool        `json:"isRead"`
	Body         Body        `json:"body"`
	ToRecipients []Recipient `json:"toRecipients"`
}
type Body struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}
type MsgEmailAddress struct {
	Address string `json:"address"`
}
type Recipient struct {
	MsgEmailAddress MsgEmailAddress `json:"emailAddress"`
}

func (g *GraphClient) CreateMessage(user string, msg Message, opts ...UpdateQueryOption) error {
	resource := fmt.Sprintf("/users/%s/mailFolders/INBOX/messages", user)

	bodyBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(bodyBytes)
	err = g.makePOSTAPICall(resource, compileUpdateQueryOptions(opts), reader, nil)
	return err
}
