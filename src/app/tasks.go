package app

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name string `json:"name"`
	Command string `json:"command"`
	RequiredTasks []string `json:"requires"`
}