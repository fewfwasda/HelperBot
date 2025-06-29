package notemanager

import (
	texts "HelperBot/Data/textsUI"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	UserNotes   = make(map[int][]string)
	storagePath = "UserDataStorage/userNotes.json"
	mu          sync.Mutex
)

func AddNote(userID int, addNote string) string {
	mu.Lock()
	UserNotes[userID] = append(UserNotes[userID], addNote)
	mu.Unlock()

	saveToDisk()
	return fmt.Sprintf(texts.ReplyToUserAddNote, addNote)
}

func DeleteNote(userID int, number string) string {
	mu.Lock()
	noteNumber, _ := strconv.Atoi(number)

	index := noteNumber - 1
	userNotes, ok := UserNotes[userID]
	if !ok || index < 0 || index >= len(userNotes) {
		return fmt.Sprintf(texts.ErrFailDeleteNote, noteNumber)
	}
	removedNote := userNotes[index]

	UserNotes[userID] = append(userNotes[:index], userNotes[index+1:]...)
	mu.Unlock()

	saveToDisk()
	return fmt.Sprintf(texts.ReplyToUserDeleteNote, removedNote)
}

func ShowNoteList(userID int) string {
	notes := UserNotes[userID]
	if len(notes) == 0 {
		return texts.ErrEmptyNoteList
	}

	var sb strings.Builder
	sb.WriteString(texts.ReplyToUserShowAllNotes)

	for i, note := range notes {
		sb.WriteString(formatLine(i+1, note))
	}
	return sb.String()
}

func ClearNoteList(userID int) {
	mu.Lock()
	delete(UserNotes, userID)
	mu.Unlock()
	saveToDisk()
}

func LoadFromDisk() error {
	mu.Lock()
	defer mu.Unlock()

	info, _ := os.Stat(storagePath)

	if info.Size() == 0 {
		return nil
	}

	file, err := os.ReadFile(storagePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &UserNotes)
}

func formatLine(number int, note string) string {
	return fmt.Sprintf("%v. %v\n", number, note)
}

func saveToDisk() error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(UserNotes, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(storagePath, data, 0644)
}
