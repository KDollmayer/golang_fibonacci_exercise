package main

import "fmt"

var priorityWeights = map[int]int{
	1: 8,
	2: 5,
	3: 3,
	4: 2,
	5: 1,
}

func DistributeTasks(tasks []string, priority int) {
	weight, ok := priorityWeights[priority]
	if !ok {
		fmt.Println("Invalid priority level")
		return
	}

	fmt.Printf("Distributing %d tasks with priority %d\n", weight, priority)
	for i := 0; i < weight; i++ {
		if i >= len(tasks) {
			break
		}
		fmt.Printf("Processing task: %s\n", tasks[i])
	}
}
