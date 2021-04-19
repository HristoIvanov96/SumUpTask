package main

import (
	"reflect"
	"testing"
)


func TestSortTasksAreSorted(t *testing.T) {
	input := []Task{
		Task{Name: "1", Command: "first", RequiredTasks: []string{"2"}},
		Task{Name: "2", Command: "second"},
	}

	result := SortTasks(input)
	expected := []Task{
		Task{Name: "2", Command: "second"},
		Task{Name: "1", Command: "first", RequiredTasks: []string{"2"}},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortTasks failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("SortTasks success")
	}
}

func TestSortTasksAreAlreadySorted(t *testing.T) {
	input := []Task{
		Task{Name: "1", Command: "first"},
		Task{Name: "2", Command: "second", RequiredTasks: []string{"1"}},
	}

	result := SortTasks(input)
	expected := []Task{
		Task{Name: "1", Command: "first"},
		Task{Name: "2", Command: "second", RequiredTasks: []string{"1"}},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortTasks failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("SortTasks success")
	}
}

func TestSortTasksReturnsEmpty(t *testing.T) {
	var input []Task

	result := SortTasks(input)
	var expected []Task
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortTasks failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("SortTasks success")
	}
}

func TestCheckRequiredTasksAreDone(t *testing.T) {
	inputRequired := []string{"1", "3"}
	inputAddedTaskNames := map[string]bool{"1": true, "2": true, "3": true}

	result := checkRequiredTasksAreDone(inputRequired, inputAddedTaskNames)
	expected := true
	if result != expected {
		t.Errorf("CheckRequiredTasksAreDone failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("CheckRequiredTasksAreDone success")
	}
}

func TestCheckRequiredTasksAreDoneReturnsFalse(t *testing.T) {
	inputRequired := []string{"1", "3"}
	inputAddedTaskNames := map[string]bool{"1": true, "2": true, "3": false}

	result := checkRequiredTasksAreDone(inputRequired, inputAddedTaskNames)
	expected := false
	if result != expected {
		t.Errorf("CheckRequiredTasksAreDone failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("CheckRequiredTasksAreDone success")
	}
}

func TestCheckRequiredTasksAreDoneEmpty(t *testing.T) {
	inputRequired := []string{"1", "3"}
	inputAddedTaskNames := map[string]bool{}

	result := checkRequiredTasksAreDone(inputRequired, inputAddedTaskNames)
	expected := false
	if result != expected {
		t.Errorf("CheckRequiredTasksAreDone failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("CheckRequiredTasksAreDone success")
	}
}

func TestRemoveTask(t *testing.T) {
	inputTasks := []Task{
		Task{Name: "1", Command: "first"},
		Task{Name: "2", Command: "second", RequiredTasks: []string{"1"}},
	}

	result := removeTask(inputTasks, 0)
	expected := []Task{
		Task{Name: "2", Command: "second", RequiredTasks: []string{"1"}},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("RemoveTask failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("RemoveTask success")
	}
}
