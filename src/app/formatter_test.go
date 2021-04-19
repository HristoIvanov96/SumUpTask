package app

import (
	"reflect"
	"testing"
)

func TestFormatSortedTasks(t *testing.T) {
	input := []Task{
		Task{Name: "1", Command: "first", RequiredTasks: []string{"2"}},
		Task{Name: "2", Command: "second"},
	}

	result := FormatSortedTasks(input)
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
	var input []Task

	result := FormatSortedTasks(input)
	var expected []FormattedTask
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FormatSortedTasks failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("FormatSortedTasks success")
	}
}

func TestFormatCommands(t *testing.T) {
	input := []Task{
		Task{Name: "1", Command: "first", RequiredTasks: []string{"2"}},
		Task{Name: "2", Command: "second"},
	}

	result := FormatCommands(input)
	expected := "#!/usr/bin/env bash \n\nfirst\nsecond\n"
	if result != expected {
		t.Errorf("FormatCommands failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("FormatCommands success")
	}
}

func TestFormatCommandsEmpty(t *testing.T) {
	var input []Task

	result := FormatCommands(input)
	expected := ""
	if result != expected {
		t.Errorf("FormatCommands failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("FormatCommands success")
	}
}
