package gcpfunc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func newWelcomeMsg() *MessageReply {
	return &MessageReply{
		Sender: &User{
			Name:     "Εκκλησία της Κηφισιάς",
			Country:  "GR",
			Language: "EL",
		},
		TrackingData: "tracking data",
		Type:         "picture",
		Text:         "Καλώς Ήρθατε στην εφαρμογή επισκέψεων",
		Media:        "https://upload.wikimedia.org/wikipedia/el/9/9b/Eaep_out.jpg",
		Thumbnail:    "https://upload.wikimedia.org/wikipedia/el/9/9b/Eaep_out.jpg",
	}
}

func newAckMsg(msgText string, receiverID string) *MessageReply {
	return &MessageReply{
		Receiver: receiverID,
		Sender: &User{
			Name:     "Εκκλησία της Κηφισιάς",
			Country:  "GR",
			Language: "EL",
		},
		Type:     msgTextType,
		Text:     msgText,
		Keyboard: NewKeyboard(),
	}
}

func (v Viber) sendTextMsg(msgText, receiverID string) error {
	bytesMsg, err := json.Marshal(newAckMsg(msgText, receiverID))
	if err != nil {
		log.Printf("post message -> person, json.Marshal error: %v\n", err)
		return err
	}

	req, err := http.NewRequest("POST", "https://chatapi.Viber.com/pa/send_message", bytes.NewBuffer(bytesMsg))
	if err != nil {
		log.Println("err:", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	v.setToken(req.Header)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("sent message Status: %s, headers: %v\n", resp.Status, resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	log.Println("body:", string(body))

	return nil
}

func newRegResp() *RegResp {
	return &RegResp{
		Status:        0,
		StatusMessage: "unhandled message",
		EventTypes: []string{
			"failed",
			"subscribed",
			"unsubscribed",
			"conversation_started",
		},
	}
}
