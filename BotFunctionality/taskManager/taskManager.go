package taskmanager

// func AddTask(tasks map[int][]string, userID int, task string) {
// 	tasks[userID] = append(tasks[userID], task)
// }

// func ShowTaskList(tasks []string) string {
// 	if len(tasks) == 0 {
// 		return "Ваш список задач пуст. Вы можете добавить задачу написав любое сообщение"
// 	}

// 	var sb strings.Builder
// 	sb.WriteString("Все ваши задачи: \n")

// 	for i, task := range tasks {
// 		sb.WriteString(formatLine(i+1, task))
// 	}
// 	return sb.String()
// }

// func RemoveTasks(tasks map[int][]string, userID int, number int) (string, error) {
// 	index := number - 1
// 	userTasks, err := tasks[userID]
// 	if !err || index < 0 || index >= len(userTasks) {
// 		return "", fmt.Errorf("невозможно удалить задачу с номером %v", number)
// 	}
// 	removedTask := userTasks[index]

// 	tasks[userID] = append(userTasks[:index], userTasks[index+1:]...)

// 	return fmt.Sprintf("Задача %v удалена %v", number, removedTask), nil
// }

// func formatLine(number int, task string) string {
// 	return fmt.Sprintf("%v. %v\n", number, task)
// }

// func ClearTasks(tasks map[int][]string, userID int) {
// 	delete(tasks, userID)
// }
