package botstates

var BotStates = make(map[int64]string)

func SetState(userID int64, state string) {
	BotStates[userID] = state
}
func GetState(userID int64) string {
	return BotStates[userID]
}
func ClearState(userID int64) {
	delete(BotStates, userID)
}
