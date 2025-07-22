package message

import (
	"database/sql"
	"fmt"
	"log"
)

type Message struct {
	ID      int
	Title   string
	Content string
}

type MessageInterface interface {
	GetMessageTitleByTitle(string) (string, error)
}

func GetMessageContext(conn *sql.DB, title string) (string, error) {
	if conn == nil {
		return "", fmt.Errorf("database connection is closed")
	}
	query := `SELECT content FROM text_ui WHERE title = $1`
	row := conn.QueryRow(query, title)

	var message Message
	err := row.Scan(&message.Content)
	if err != nil {
		log.Printf("Scan error (title: %s): %v", title, err)
		return "", err
	}

	return message.Content, nil
}
