package app

type Task struct {
	Name string `json:"name"`
	Command string `json:"command"`
	RequiredTasks []string `json:"requires"`
}

//func (e Employee) LeavesRemaining() {
//	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
//}