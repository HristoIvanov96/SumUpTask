package utils

import (
	"SumUpTask/models"
	"errors"
	"reflect"
	"testing"
)

func TestFormatSortedTasks(t *testing.T) {
	input := []models.Task{
		models.Task{Name: "1", Command: "first", RequiredTasks: []string{"2"}},
		models.Task{Name: "2", Command: "second"},
	}

	result, _ := FormatSortedTasks(input)
	expected := []FormattedTask{
		FormattedTask{Name: "1", Command: "first"},
		FormattedTask{Name: "2", Command: "second"},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FormatSortedTasks failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("FormatSortedTasks success")
	}
}

func TestFormatSortedTasksEmpty(t *testing.T) {
	var input []models.Task

	_, err := FormatSortedTasks(input)
	expected := errors.New("please add tasks in the request body")
	if err.Error() != expected.Error() {
		t.Errorf("FormatSortedTasks failed, expected %v, got %v", expected, err)
	} else {
		t.Logf("FormatSortedTasks success")
	}
}

func TestFormatCommands(t *testing.T) {
	input := []models.Task{
		models.Task{Name: "1", Command: "first", RequiredTasks: []string{"2"}},
		models.Task{Name: "2", Command: "second"},
	}

	result, _:= FormatCommands(input)
	expected := "#!/usr/bin/env bash \n\nfirst\nsecond\n"
	if *result != expected {
		t.Errorf("FormatCommands failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("FormatCommands success")
	}
}

func TestFormatCommandsEmpty(t *testing.T) {
	var input []models.Task

	_, err := FormatCommands(input)
	expected := errors.New("please add tasks in the request body")
	if err.Error() != expected.Error() {
		t.Errorf("FormatCommands failed, expected %v, got %v", expected, err)
	} else {
		t.Logf("FormatCommands success")
	}
}
