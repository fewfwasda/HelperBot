package weather

import (
	texts "HelperBot/Data/textsUI"
	"database/sql"
	"fmt"
	"time"
)

func GetUserCity(db *sql.DB, userID int64) (string, bool) {
	const query = `
	SELECT city
	FROM user_cities
	WHERE user_id=$1
	`
	var city string
	err := db.QueryRow(query, userID).Scan(&city)

	if err == sql.ErrNoRows {
		return "", false
	}
	if err != nil {
		return "", false
	}
	return city, true
}

func SetCity(db *sql.DB, userID int64, city string) error {
	const query = `
	INSERT INTO user_cities (user_id, city, updated_at)
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id) DO UPDATE
		SET city = EXCLUDED.city,
			updated_at = EXCLUDED.updated_at
	`
	_, err := db.Exec(query, userID, city, time.Now())
	if err != nil {
		return fmt.Errorf("%v: %v", texts.ErrSaveCity, err)
	}
	return nil
}
