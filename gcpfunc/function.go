package gcpfunc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	msgTextType = "text"
)

type Viber struct {
	Key string
}

type MessageReply struct {
	Receiver     string    `json:"receiver,omitempty"`
	Sender       *User     `json:"sender,omitempty"`
	TrackingData string    `json:"tracking_data,omitempty"`
	Type         string    `json:"type,omitempty"`
	Text         string    `json:"text,omitempty"`
	Media        string    `json:"media,omitempty"`
	Thumbnail    string    `json:"thumbnail,omitempty"`
	Keyboard     *Keyboard `json:"keyboard,omitempty"`
}

type ViberRequest struct {
	Event        string           `json:"event,omitempty"`
	Timestamp    int64            `json:"timestamp,omitempty"`
	ChatHostname string           `json:"chat_hostname,omitempty"`
	MessageToken int64            `json:"message_token,omitempty"`
	Type         string           `json:"type,omitempty"`
	User         *User            `json:"user,omitempty"`
	Sender       *User            `json:"sender,omitempty"`
	Subscribed   bool             `json:"subscribed,omitempty"`
	Message      *MessageReceived `json:"message,omitempty"`
}

type User struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	Language   string `json:"language,omitempty"`
	Country    string `json:"country,omitempty"`
	APIVersion int    `json:"api_version,omitempty"`
}

type Location struct {
	Lat float64 `json:"lat,omitempty"`
	Lon float64 `json:"lon,omitempty"`
}

type MessageReceived struct {
	Type         string    `json:"type,omitempty"`
	Text         string    `json:"text,omitempty"`
	Media        string    `json:"media,omitempty"`
	Location     *Location `json:"location,omitempty"`
	TrackingData string    `json:"tracking_data,omitempty"`
}

type RegResp struct {
	Status        int      `json:"status"`
	StatusMessage string   `json:"status_message"`
	EventTypes    []string `json:"event_types"`
}

type Keyboard struct {
	Type          string   `json:"Type"`
	DefaultHeight bool     `json:"DefaultHeight"`
	Buttons       []Button `json:"Buttons"`
}

type Button struct {
	ActionType string `json:"ActionType"`
	ActionBody string `json:"ActionBody"`
	Text       string `json:"Text"`
	TextSize   string `json:"TextSize"`
	Columns    int    `json:"Columns"`
	Rows       int    `json:"Rows"`
	TextHAlign string `json:"TextHAlign"`
	TextVAlign string `json:"TextVAlign"`
	BgColor    string `json:"BgColor"`
	Image      string `json:"Image"`
	Silent     bool   `json:"Silent"`
}

func NewKeyboard() *Keyboard {
	return &Keyboard{
		Type:          "keyboard",
		DefaultHeight: true,
		Buttons: []Button{
			{
				Columns:    6,
				Rows:       1,
				ActionType: "reply",
				ActionBody: "1",
				Text:       `<font color="#494E67">1η) Συνάθροιση Κυριακή 12:00-13:30</font>`,
				TextSize:   "large",
				BgColor:    "#dd8157",
				TextHAlign: "center",
				TextVAlign: "middle",
				Silent:     true,
			},
			{
				Columns:    6,
				Rows:       1,
				ActionType: "reply",
				ActionBody: "2",
				Text:       `<font color="#494E67">2η) Συνάθροιση Κυριακή 14:00-15:30</font>`,
				TextSize:   "large",
				BgColor:    "#f7bb3f",
				TextHAlign: "center",
				TextVAlign: "middle",
				Silent:     true,
			},
			{
				Columns:    6,
				Rows:       1,
				ActionType: "reply",
				ActionBody: "3",
				Text:       `<font color="#494E67">Θες την ευλογία του Θεού;</font>`,
				TextSize:   "large",
				BgColor:    "#a8aaba",
				TextHAlign: "center",
				TextVAlign: "middle",
				Silent:     true,
			},
			{
				Columns:    6,
				Rows:       1,
				ActionType: "reply",
				ActionBody: "4",
				Text:       `<font color="#494E67">Θες να πας στον ουρανό;</font>`,
				TextSize:   "large",
				BgColor:    "#7eceea",
				TextHAlign: "center",
				TextVAlign: "middle",
				Silent:     true,
			},
		},
	}
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

// Inquire prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func (v Viber) Inquire(w http.ResponseWriter, r *http.Request) {
	var req ViberRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		switch err {
		case io.EOF:
			log.Printf("EOF in json.Decoder: %v\n", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		default:
			log.Printf("json.NewDecoder: %v\n", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	// Does not look its necessary for webhook callbacks
	v.setToken(w.Header())

	switch req.Event {
	case "webhook":
		if err := json.NewEncoder(w).Encode(newRegResp()); err != nil {
			log.Printf("web hook -> json.Encoder: %v\n", err)
		}
		return
	case "message":
		log.Printf("senderID: %s, SenderName: %s, Event: %s, Subscribed: %t, Timestamp: %d, MessageType: %s, "+
			"MsgText: %s, TrackingData: %v,  Media:%v\n", req.Sender.ID, req.Sender.Name, req.Event, req.Subscribed,
			req.Timestamp, req.Message.Type, req.Message.Text, req.Message.TrackingData, req.Message.Media)

		userChoice, err := strconv.Atoi(req.Message.Text)
		if err != nil && userChoice == 0 {
			if err := v.sendTextMsg(fmt.Sprintf("Εδάφιο για σένα:\n"+
				"«εγένετο η βασιλεία του κόσμου του Κυρίου ημών και του Χριστού αυτού, και "+
				"βασιλεύσει εις τους αιώνας των αιώνων»Αποκ Ια:15\n\nΕίπατε %q\n", req.Message.Text), req.Sender.ID); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		} else {
			switch userChoice {
			case 1:
				if err := v.sendTextMsg("Αμήν σας περιμένουμε την Κυριακή στις 12:00μμ!", req.Sender.ID); err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			case 2:
				if err := v.sendTextMsg("Αμήν σας περιμένουμε την Κυριακή στις 2:00μμ!", req.Sender.ID); err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			case 3:
				if err := v.sendTextMsg("Είπατε Αμήν, Ο θεός να σας ευλογεί!", req.Sender.ID); err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			case 4:
				if err := v.sendTextMsg("Είπατε Αμήν, Ο Χριστός θέλει και δύναται να σας ετοιμάσει να ζήσετε αιώνια μαζί του!", req.Sender.ID); err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}
		}
		return

	case "conversation_started":
		if err := json.NewEncoder(w).Encode(newWelcomeMsg()); err != nil {
			log.Printf("json.Encoder: %v\n", err)
		}

		return
	default:
		log.Printf("unhandled event: %s, req: %v\n", req.Event, req)
	}
}

func (v Viber) setToken(header http.Header) {
	header.Add("X-Viber-Auth-Token", v.Key)
}

func (v Viber) sendTextMsg(msgText, receiverID string) error {
	bytesMsg, err := json.Marshal(newAckMsg(msgText, receiverID))
	if err != nil {
		log.Printf("post message -> user, json.Marshal error: %v\n", err)
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
