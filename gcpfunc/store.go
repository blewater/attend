package gcpfunc

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
)

const (
	root      = "church"
	personCol = "person"
)

func newFirebaseCfg() *firebase.Config {
	return &firebase.Config{
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		ProjectID:     os.Getenv("PROJECT_ID"),
		StorageBucket: os.Getenv("STORAGE_BUCKET"),
	}
}

type person struct {
	ID          string    `firestore:"id,omitempty"`
	First       string    `firestore:"first,omitempty"`
	Last        string    `firestore:"last,omitempty"`
	Subscribed  bool      `firestore:"subscribed,omitempty"`
	PhoneNumber string    `firestore:"phoneNumber,omitempty"`
	Type        string    `firestore:"type,omitempty"`
	Updated     time.Time `firestore:"updated,omitempty"`
}

type CloudStore struct {
	app *firebase.App
	db  *firestore.Client
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

func (s *CloudStore) userToDB(v ViberRequest) {
	u := v.Sender
	if v.Sender == nil {
		u = v.User
		if u == nil {
			return
		}
	}
	names := strings.Split(u.Name, " ")
	p := &person{
		ID:         u.ID,
		First:      names[0],
		Last:       names[1],
		Subscribed: v.Subscribed,
		Updated:    time.Now().UTC(),
	}
	// upsert
	_, err := s.db.Collection(root).Doc(personCol).Set(context.Background(), p)
	if err != nil {
		log.Printf("Firestore userToDB error has occurred: %s\n", err)
	}
}
