package tasks

import (
	"fmt"
	"github.com/fatih/color"
)

type Task struct {
	ID     int
	Title  string
	Status bool
}

var taskList []Task

func AddTask(title string) {
	err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	id := len(taskList) + 1
	task := Task{ID: id, Title: title, Status: false}
	taskList = append(taskList, task)

	err = SaveTasks()
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}


func ShowTasks(filter ...string) {

	err := LoadTasks()
	if err != nil {
	fmt.Println("Error loading tasks:", err)
	return
	}


	for _, task := range taskList {
		match := true
		if len(filter) > 0 {
			switch filter[0] {
			case "done":
				match = task.Status
			case "pending":
				match = !task.Status
			}
		}

		if match {
			var state string
			var titleColored string

			if task.Status {
				state = color.HiGreenString("✅")
				titleColored = color.New(color.FgHiBlack).Sprint(task.Title)
			} else {
				state = color.HiRedString("❌")
				titleColored = color.New(color.FgHiWhite).Sprint(task.Title)
			}

			fmt.Printf("[%d] %s - %s\n", task.ID, state, titleColored)
		}
	}
}

func MarkDone(id int) {
	err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for i, task := range taskList {
		if task.ID == id {
			taskList[i].Status = true

			err = SaveTasks()
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			}

			fmt.Printf("Status of Task %d updated successfully\n", id)
			return
		}
	}
	fmt.Println("Task ID not found.")
}


func DeleteTask(id int) {
	err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	index := -1
	for i, task := range taskList {
		if task.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Task not found.")
		return
	}

	taskList = append(taskList[:index], taskList[index+1:]...)

	for i := range taskList {
		taskList[i].ID = i + 1
	}

	err = SaveTasks()
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Deleted task #%d\n", id)
}


func EditTask(id int, newTitle string) {
	err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for i := range taskList {
		if taskList[i].ID == id {
			taskList[i].Title = newTitle

			err = SaveTasks()
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			}

			fmt.Printf("Task #%d updated successfully\n", id)
			return
		}
	}
	fmt.Println("Task ID not found.")
}

