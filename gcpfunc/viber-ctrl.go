package gcpfunc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	msgTextType = "text"
)

type Viber struct {
	key string
	db  *CloudStore
}

func NewViberApp(key string) Viber {
	s := NewStore()
	v := Viber{
		key: key,
		db:  s,
	}

	return v
}

func (v Viber) Meeting(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/meeting" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			errW := fmt.Sprintf("meeting parseform() err: %v", err)
			log.Println(errW)
			if _, fErr := fmt.Fprintf(w, errW); fErr != nil {
				log.Println("fPrintf error:", fErr)
			}
		}

		return
	}
	fmt.Printf("Post from website! r.PostFrom = %v\n", r.PostForm)
	day, err := strconv.Atoi(r.FormValue("day"))
	if err != nil {
		dayErrStr := fmt.Sprintf("Bad day param:%s, error:%s", r.FormValue("day"), err)
		_, _ = fmt.Fprintf(w, dayErrStr)
		fmt.Println(dayErrStr)
		return
	}
	minutes, err := strconv.Atoi(r.FormValue("minutes"))
	if err != nil {
		minutesErrStr := fmt.Sprintf("Bad minutes param:%s, error:%s", r.FormValue("minutes"), err)
		_, _ = fmt.Fprintf(w, minutesErrStr)
		fmt.Println(minutesErrStr)
		return
	}
	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil {
		countErrStr := fmt.Sprintf("Bad count param:%s, error:%s", r.FormValue("count"), err)
		_, _ = fmt.Fprintf(w, countErrStr)
		fmt.Println(countErrStr)
		return
	}
	timeVal, err := time.Parse(time.RFC3339, r.FormValue("time"))
	if err != nil {
		timeErrStr := fmt.Sprintf("Bad time param:%s, error:%s", r.FormValue("time"), err)
		_, _ = fmt.Fprintf(w, timeErrStr)
		fmt.Println(timeErrStr)
		return
	}
	meeting := &meeting{
		ID:      NewUuid(),
		Day:     time.Weekday(day - 1),
		Minutes: minutes,
		Time:    timeVal,
		Count:   count,
		Created: time.Now().UTC(),
	}

	v.db.addMeetingIfNotExists(meeting)
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

	v.userToDB(req)
	// Does not look its necessary for webhook callbacks
	v.setToken(w.Header())

	switch req.Event {
	case "webhook":
		if err := json.NewEncoder(w).Encode(newRegResp()); err != nil {
			log.Printf("web hook -> error: %v\n", err)
			return
		}
		log.Println("Webhook registration success.")

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
				if err := v.sendTextMsg("Είπατε Αμήν, μελετάτε καθημερινά την Αγία Γραφή και θα γευτείτε την ευλογία του Θεού!", req.Sender.ID); err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			case 4:
				if err := v.sendTextMsg("Είπατε Αμήν, ο Χριστός θέλει και δύναται να σας ετοιμάσει να ζήσετε αιώνια μαζί του!", req.Sender.ID); err != nil {
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
	case "unsubscribed":
		log.Println("Received unsubscribed msg:", req)
	default:
		log.Printf("unhandled event: %s, req: %v\n", req.Event, req)
	}
}

func (v Viber) setToken(header http.Header) {
	header.Add("X-Viber-Auth-Token", v.key)
}

func (v Viber) userToDB(req ViberRequest) {
	u := req.Sender
	if req.Sender == nil {
		u = req.User
		if u == nil {
			return
		}
	}
	names := strings.Split(u.Name, " ")
	p := &person{
		ID:         u.ID,
		First:      names[0],
		Last:       names[1],
		Subscribed: req.Subscribed,
		Created:    time.Now().UTC(),
	}

	// no need to wait
	go v.db.addUserIfNotExists(p)
}
