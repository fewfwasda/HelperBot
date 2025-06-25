package notemanager

import (
	texts "HelperBot/Data/textsUI"
	"fmt"
	"strconv"
	"strings"
)

var UserNotes = make(map[int][]string)

func AddNote(userID int, addNote string) string {
	UserNotes[userID] = append(UserNotes[userID], addNote)

	return fmt.Sprintf(texts.ReplyToUserDeleteNote, addNote)
}

func DeleteNote(userID int, number string) string {
	noteNumber, _ := strconv.Atoi(number)

	index := noteNumber - 1
	userNotes, err := UserNotes[userID]
	if !err || index < 0 || index >= len(userNotes) {
		return fmt.Sprintf(texts.ErrFailDeleteNote, noteNumber)
	}
	removedNote := userNotes[index]

	UserNotes[userID] = append(userNotes[:index], userNotes[index+1:]...)

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
	delete(UserNotes, userID)
}

func formatLine(number int, note string) string {
	return fmt.Sprintf("%v. %v\n", number, note)
}
