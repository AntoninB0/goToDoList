package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Date struct {
	Day     int
	Month   int
	Year    int
	Hour    int
	Minutes int
}

type Task struct {
	taskName     string
	taskDescribe string
	taskDate     Date
}

func taskToJson(task Task) string {
	return fmt.Sprintf(`{"taskName": "%s", "taskDescribe": "%s", "date": {"Day": %d, "Month": %d, "Year": %d, "Hour": %d, "Minutes": %d}}`, task.taskName, task.taskDescribe, task.taskDate.Day, task.taskDate.Month, task.taskDate.Year, task.taskDate.Hour, task.taskDate.Minutes)
}

func saveInFile(tasks []Task) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		return err
	}

	return nil
}

func addTask(taskList []Task) []Task {
	var task Task
	var taskDateList []string
	var taskDateRow string

	fmt.Print("Entrez un nom de tache : ")
	fmt.Scanln(&task.taskName)

	fmt.Print("Entrez un descriptif : ")
	fmt.Scanln(&task.taskDescribe)

	fmt.Print("Entrez une Date (DD/MM/YY/HH/mm) : ")
	fmt.Scanln(&taskDateRow)

	taskDateList = strings.Split(taskDateRow, "/")

	fmt.Sscanf(taskDateList[0], "%d", &task.taskDate.Day)
	fmt.Sscanf(taskDateList[1], "%d", &task.taskDate.Month)
	fmt.Sscanf(taskDateList[2], "%d", &task.taskDate.Year)
	fmt.Sscanf(taskDateList[3], "%d", &task.taskDate.Hour)
	fmt.Sscanf(taskDateList[4], "%d", &task.taskDate.Minutes)

	taskList = append(taskList, task)

	return taskList
}

func loadFile() []Task {
	var tasks []Task

	file, err := os.Open("tasks.json")
	if err != nil {
		return tasks
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	for decoder.More() {
		var task Task
		if err := decoder.Decode(&task); err != nil {
			return tasks
		}
		tasks = append(tasks, task)
	}

	fmt.Println("Tâches chargées avec succès depuis le fichier tasks.json")
	return tasks
}

func main() {
	var taskList []Task

	taskList = addTask(taskList)

	err := saveInFile(taskList)
	if err != nil {
		fmt.Println("Erreur lors de l'enregistrement des tâches :", err)
		return
	}
}
