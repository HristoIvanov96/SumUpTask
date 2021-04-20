package utils

import (
	"SumUpTask/models"
	"errors"
)

//Using a struct to format the response
type FormattedTask struct {
	Name string
	Command string
}


func FormatSortedTasks(tasks []models.Task) ([]FormattedTask, error) {
	var formattedTasks []FormattedTask

	//Format sorted tasks
	for _, task := range tasks {
		formattedTasks = append(formattedTasks, FormattedTask{task.Name, task.Command})
	}
	if formattedTasks == nil {
		return nil, errors.New("please add tasks in the request body")
	}
	return formattedTasks, nil
}

func FormatCommands(tasks []models.Task) (*string, error) {
	if len(tasks) == 0 {
		return nil, errors.New("please add tasks in the request body")
	}
	//Start with the first line of the bash script and add all commands
	commands := "#!/usr/bin/env bash \n\n"
	for _, task := range tasks {
		commands += task.Command + "\n"
	}
	return &commands, nil
}
