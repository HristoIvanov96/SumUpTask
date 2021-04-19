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

func FormatCommands(tasks []Task) string {
	if len(tasks) == 0 {
		return ""
	}
	commands := "#!/usr/bin/env bash \n\n"
	for _, task := range tasks {
		commands += task.Command + "\n"
	}
	return commands
}
