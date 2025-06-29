package main

import "fmt"

func main() {

	allTasks := []Subtask{
		Subtask{Status: "incomplete"}, 
		Subtask{Status: "completed"},
	}
	
	for _, x := range allTasks {
		if x.Status != "completed" {
			fmt.Printf("Task is not completed\n")
		}
	}
}

type Subtask struct {
	Param1 string
	Param2 string
	Status string
}
