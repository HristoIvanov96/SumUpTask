package app

type FormattedTask struct {
	Name string
	Command string
}


func FormatSortedTasks(tasks []Task) []FormattedTask {
	var formattedTasks []FormattedTask
	for _, task := range tasks {
		formattedTasks = append(formattedTasks, FormattedTask{task.Name, task.Command})
	}
	return formattedTasks
}
