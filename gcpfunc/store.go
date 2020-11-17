package gcpfunc

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
)

const (
	users    = "users"
	meetings = "meetings"
	bookings = "bookings"
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
	Subscribed  bool      `firestore:"subscribed"`
	PhoneNumber string    `firestore:"phoneNumber"`
	Type        string    `firestore:"type"`
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

func (s *CloudStore) addUserIfNotExists(p *person) {
	if p == nil {
		return
	}
	personsCol := s.db.Collection(users)
	q := personsCol.Where("id", "==", p.ID)
	iter := q.Documents(context.Background())
	defer iter.Stop()
	doc, err := iter.Next()
	if err != nil && err != iterator.Done {
		log.Printf("Firestore addUserIfNotExists iter error: %s\n", err)
		return
	}
	if doc != nil {
		if dbPerson := doc.Data(); dbPerson != nil && dbPerson["id"] == p.ID {
			return
		}
	}
	s.upsertPerson(personsCol, p)
}

func (s *CloudStore) upsertPerson(personsCol *firestore.CollectionRef, p *person) {
	if p.ID == "" {
		return
	}
	_, _, err := personsCol.Add(context.Background(), p)
	if err != nil {
		log.Printf("Firestore addUserIfNotExists error: %s\n", err)
	}
}
