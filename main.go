package main

func main() {
	CallAPIWithRetries(6)
	tasks := []string{"Task1", "Task2", "Task3", "Task4", "Task5", "Task6", "Task7", "Task8", "Task9"}
	DistributeTasks(tasks, 0)
	DistributeTasks(tasks, 3)
}
