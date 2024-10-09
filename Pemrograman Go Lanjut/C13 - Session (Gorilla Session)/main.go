package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/kidstuff/mongostore"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
)

const SESSION_ID = "id"

func newMongoStore() *mongostore.MongoStore {
	mgoSession, err := mgo.Dial("localhost:27123")
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}

	dbCollection := mgoSession.DB("learnwebgolang").C("session")
	maxAge := 86400 * 7
	ensureTTL := true
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store := mongostore.NewMongoStore(
		dbCollection,
		maxAge,
		ensureTTL,
		authKey,
		encryptionKey,
	)
	return store
}

func newPostgresStore() *pgstore.PGStore {
	url := "postgres://novalagung:@127.0.0.1:5432/novalagung?sslmode=disable"
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}

	return store
}

func newCookieStore() *sessions.CookieStore {
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7
	store.Options.HttpOnly = true

	return store
}

func main() {
	e := echo.New()
	store := newMongoStore()
	// store := sessions.NewCookieStore([]byte("secret-key"))
	// SESSION_ID := "session-id"

	e.GET("/set", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Values["message1"] = "hello"
		session.Values["message2"] = "world"
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.GET("/get", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)

		if len(session.Values) == 0 {
			return c.String(http.StatusOK, "empty result")
		}

		return c.String(http.StatusOK, fmt.Sprintf(
			"%s %s",
			session.Values["message1"],
			session.Values["message2"],
		))
	})

	e.GET("/delete", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Options.MaxAge = -1
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
