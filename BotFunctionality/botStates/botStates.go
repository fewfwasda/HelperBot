package botstates

var BotStates = make(map[int]string)

func SetState(userID int, state string) {
	BotStates[userID] = state
}
func GetState(userID int) string {
	return BotStates[userID]
}
func ClearState(userID int) {
	delete(BotStates, userID)
}
