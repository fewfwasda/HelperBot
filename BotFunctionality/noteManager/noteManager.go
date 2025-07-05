package notemanager

import (
	texts "HelperBot/Data/textsUI"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func AddNote(db *sql.DB, userID int64, note string) string {

	const query = `INSERT INTO notes (user_id, note, created_at)
VALUES ($1, $2, $3);
`
	_, err := db.Exec(query, userID, note, time.Now())
	if err != nil {
		return texts.ErrAddNote
	}
	return fmt.Sprintf(texts.ReplyToUserAddNote, note)

}

func DeleteNote(db *sql.DB, userID int64, ordinal string) string {
	const query = `DELETE FROM notes
        WHERE id = (
            SELECT id
            FROM notes
            WHERE user_id = $1
            ORDER BY created_at
            LIMIT 1
            OFFSET $2
        )`
	indexNote, _ := strconv.Atoi(ordinal)

	res, err := db.Exec(query, userID, indexNote-1)

	if err != nil {
		return err.Error()
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err.Error()
	}

	if n == 0 {
		return fmt.Sprintf(texts.ErrDeleteNote, indexNote)
	}

	return fmt.Sprintf(texts.ReplyToUserDeleteNote, indexNote)
}

func NoteList(db *sql.DB, userID int64) string {

	const query = `SELECT note
	FROM notes
	WHERE user_id=$1
	ORDER BY created_at;
	`
	rows, err := db.Query(query, userID)
	if err != nil {
		return err.Error()
	}
	defer rows.Close()

	var sb strings.Builder
	counter := 1

	for rows.Next() {
		var note string
		if err := rows.Scan(&note); err != nil {
			return err.Error()
		}
		sb.WriteString(fmt.Sprintf("%d. %s\n", counter, note))
		counter++
	}

	noteList := sb.String()

	return noteList
}

func ClearNoteList(db *sql.DB, userID int64) string {
	const query = `DELETE FROM notes
	 WHERE user_id = $1;`
	_, err := db.Exec(query, userID)
	if err != nil {
		return err.Error()
	}
	return texts.ReplyToUserClearNote
}
