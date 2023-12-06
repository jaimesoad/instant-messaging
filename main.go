package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"messaging/src/qrs"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	opt "github.com/jaimesoad/go-optional"
	_ "github.com/mattn/go-sqlite3"
)

type onlineUser struct {
	User   string `json:"user"`
	Online bool   `json:"online"`
}

//go:embed sql/schema.sql
var schema string
var conns = make(map[string]*websocket.Conn)
var onlines = make(map[string]bool)
var ctx, q = CreateDB()
var m, onlineM sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func CreateDB() (context.Context, *qrs.Queries) {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	fmt.Println(db.ExecContext(ctx, schema))

	q := qrs.New(db)

	return ctx, q
}

func main() {

	http.HandleFunc("/ws", directMessaging)
	http.HandleFunc("/online", checkUserOnline)
	http.HandleFunc("/allOnline", getAllOnlineUsers)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func directMessaging(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	username := r.URL.Query().Get("user")
	recepient := r.URL.Query().Get("recepient")

	m.Lock()
	conns[username] = conn
	m.Unlock()

	if !opt.Get(q.ExistUser(ctx, username)).Default(false) {
		q.NewUser(ctx, username)
	}

	sendPendingMessages(username)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			m.Lock()
			delete(conns, username)
			m.Unlock()
			return
		}
		fmt.Println(string(msg))

		sendDirectMessage(username, recepient, string(msg))
	}
}

func sendDirectMessage(sender, recepient, message string) {
	m.Lock()
	conn, ok := conns[recepient]
	m.Unlock()

	exist := opt.Get(q.ExistUser(ctx, recepient)).Default(false)

	switch {
	case !(ok || exist):
		return

	case !ok && exist:
		msg := qrs.SaveMessageParams{
			Content:   message,
			Recepient: recepient,
			Sender:    sender,
		}
		q.SaveMessage(ctx, msg)
		return
	}

	err := conn.WriteMessage(websocket.TextMessage, []byte(message))

	if err != nil {
		conn.Close()
		m.Lock()
		delete(conns, recepient)
		m.Unlock()
	}
}

func sendPendingMessages(username string) {
	m.Lock()
	conn := conns[username]
	m.Unlock()

	if !opt.Get(q.PendingMessages(ctx, username)).Default(false) {
		return
	}

	messages := opt.Get(q.SavedMessages(ctx, username)).Default([]qrs.SavedMessagesRow{})

	for _, pending := range messages {
		err := conn.WriteMessage(websocket.TextMessage, []byte(pending.Sender+": "+pending.Content))

		if err != nil {
			conn.Close()
			m.Lock()
			delete(conns, username)
			m.Unlock()
			return
		}
	}

	q.DeletePending(ctx, username)
}

func checkUserOnline(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	m.Lock()
	_, ok := conns[user]
	m.Unlock()

	if !ok {
		return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		users := opt.Get(q.FindUsers(ctx, string(msg))).Default([]string{})

		var foundUsers []onlineUser

		for _, foundUser := range users {
			foundUsers = append(foundUsers, onlineUser{
				User:   foundUser,
				Online: isOnline(foundUser),
			})
		}

		err = conn.WriteJSON(foundUsers)
		if err != nil {
			return
		}
	}
}

func getAllOnlineUsers(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	onlineM.Lock()
	onlines[user] = true
	onlineM.Unlock()

	for {

		var onlineUsers []string
		for userOnline := range onlines {
			if userOnline != user {
				onlineUsers = append(onlineUsers, userOnline)
			}
		}

		err := conn.WriteJSON(onlineUsers)
		if err != nil {
			log.Println(err)

			onlineM.Lock()

			delete(onlines, user)
			onlineM.Unlock()
			return
		}

		time.Sleep(1 * time.Second)
	}
}

func isOnline(user string) bool {
	m.Lock()
	_, online := conns[user]
	m.Unlock()

	return online
}
