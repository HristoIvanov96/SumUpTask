package main

import "errors"

type FormattedTask struct {
	Name string
	Command string
}


func FormatSortedTasks(tasks []Task) ([]FormattedTask, error) {
	var formattedTasks []FormattedTask
	for _, task := range tasks {
		formattedTasks = append(formattedTasks, FormattedTask{task.Name, task.Command})
	}
	if formattedTasks == nil {
		return nil, errors.New("please add tasks in the request body")
	}
	return formattedTasks, nil
}

func FormatCommands(tasks []Task) (string, error) {
	if len(tasks) == 0 {
		return "", errors.New("please add tasks in the request body")
	}
	commands := "#!/usr/bin/env bash \n\n"
	for _, task := range tasks {
		commands += task.Command + "\n"
	}
	return commands, nil
}
