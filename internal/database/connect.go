package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/cloyop/tiniest/internal/types"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	var TURSO_TOKEN = os.Getenv("TURSO_TOKEN")
	var TURSO_URL = os.Getenv("TURSO_URL")
	DB, err := sql.Open("libsql", TURSO_URL+"?authToken="+TURSO_TOKEN)
	if err != nil {
		log.Fatal(err)
	}
	db = DB
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	keepAlive()
	return DB
}

func keepAlive() {
	go func() {
		db.Query("SELECT * FROM users WHERE id=?", "")
	}()
	time.AfterFunc(time.Minute*10, keepAlive)
}
func InsertFeedBack(f types.FeedBack) error {
	return db.QueryRow(`INSERT INTO feedbacks (name,email,subject,description,date) VALUES (?,?,?,?,?)`, f.Name, f.Email, f.Subject, f.Description, f.Date).Err()
}
