package main

func SortTasks(tasks []Task) []Task {
	var sortedTasks []Task
	doneTasks := make(map[string]bool)
	i := 0

	//Iterate over the tasks and check if the prerequisite ones are done
	//If they are add it to done and remove from list
	for len(tasks) > 0 {
		if checkRequiredTasksAreDone(tasks[i].RequiredTasks, doneTasks) {
			sortedTasks = append(sortedTasks, tasks[i])
			doneTasks[tasks[i].Name] = true
			tasks = removeTask(tasks, i)
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}
	return sortedTasks
}

//Check if all required tasks for a given one are finished
func checkRequiredTasksAreDone(required []string, doneTasks map[string]bool) bool {
	for _, task := range required {
		if !doneTasks[task] {
			return false
		}
	}
	return true
}

func removeTask(tasks []Task, index int) []Task {
	tasks[index] = tasks[len(tasks)-1]
	tasks[len(tasks)-1] = Task{}
	tasks = tasks[:len(tasks)-1]
	return tasks
}