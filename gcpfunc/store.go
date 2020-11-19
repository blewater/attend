package gcpfunc

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/google/uuid"
)

const (
	users       = "users"
	meetings    = "meetings"
	attendances = "attendances"
)

func newFirebaseCfg() *firebase.Config {
	return &firebase.Config{
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		ProjectID:     os.Getenv("PROJECT_ID"),
		StorageBucket: os.Getenv("STORAGE_BUCKET"),
	}
}

type person struct {
	ID          string    `firestore:"id"`
	First       string    `firestore:"first"`
	Last        string    `firestore:"last"`
	Subscribed  bool      `firestore:"subscribed"`
	PhoneNumber string    `firestore:"phoneNumber"`
	Role        string    `firestore:"role"`
	ExtraFamily int       `firestore:"extraFamily"`
	Created     time.Time `firestore:"created"`
}

type meeting struct {
	ID      uuid.UUID    `firestore:"id"`
	Day     time.Weekday `firestore:"day"`
	Minutes int          `firestore:"minutes"`
	Time    time.Time    `firestore:"time"`
	Count   int          `firestore:"count"`
	Created time.Time    `firestore:"created"`
}

type CloudStore struct {
	app *firebase.App
	db  *firestore.Client
}

func NewUuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func NewStore() *CloudStore {
	cfg := newFirebaseCfg()

	ctx := context.Background()
	db, err := firebase.NewApp(ctx, cfg)
	if err != nil {
		log.Fatalf("error initializing db: %v\n", err)
	}

	store, err := db.Firestore(context.Background())
	if err != nil {
		log.Println("Firestore connection error:", err)
		return nil
	}
	fireStore := &CloudStore{
		app: db,
		db:  store,
	}

	return fireStore
}

func (s *CloudStore) addAttendanceIfNotExists(m *meeting) {
	if m == nil {
		return
	}
	attendancesCol := s.db.Collection(attendances)
	attendance := attendancesCol.Doc(m.ID.String())
	attendanceRef, err := attendance.Get(context.Background())
	if err != nil || !attendanceRef.Exists() {
		// s.upsertAttendance(meetingsCol, m)
	}
}

func (s *CloudStore) addMeetingIfNotExists(m *meeting) {
	if m == nil {
		return
	}
	meetingsCol := s.db.Collection(meetings)
	meeting := meetingsCol.Doc(m.ID.String())
	meetingRef, err := meeting.Get(context.Background())
	if err != nil || !meetingRef.Exists() {
		// s.upsertMeeting(meetingsCol, m)
	}
}

func (s *CloudStore) addUserIfNotExists(p *person) {
	if p == nil {
		return
	}
	usersCol := s.db.Collection(users)
	user := usersCol.Doc(p.ID)
	userRef, err := user.Get(context.Background())
	if err != nil || !userRef.Exists() {
		s.upsertUser(usersCol, p)
	}
}

func (s *CloudStore) upsertUser(personsCol *firestore.CollectionRef, p *person) {
	if p.ID == "" {
		return
	}
	_, err := personsCol.Doc(p.ID).Set(context.Background(), p)
	if err != nil {
		log.Printf("Firestore addUserIfNotExists error: %s\n", err)
	}
}
