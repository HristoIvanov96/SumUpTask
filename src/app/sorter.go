package app


func SortTasks(tasks []Task) []Task {
	var sortedTasks []Task
	addedTaskNames := make(map[string]bool)
	i := 0
	for len(tasks) > 0 {
		if tasks[i].RequiredTasks == nil ||
			checkRequiredTasksAreDone(tasks[i].RequiredTasks, addedTaskNames) {
			sortedTasks = append(sortedTasks, tasks[i])
			addedTaskNames[tasks[i].Name] = true
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

func checkRequiredTasksAreDone(required []string, addedTaskNames map[string]bool) bool {
	for _, task := range required {
		if _, ok := addedTaskNames[task]; !ok {
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