package Task

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task-manager/errors"
)

// type Task struct {
// }

func CreateTask(task string) (string, error) {
	file, err := os.OpenFile("tasks.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err1 := file.WriteString(task + "\n")
	if err1 != nil {
		return "", err1
	}
	file1, err2 := os.Open("tasks.txt")
	if err2 != nil {
		return "", err2
	}
	var tasks []string
	scannner := bufio.NewScanner(file1)
	for scannner.Scan() {
		tasks = append(tasks, scannner.Text())
	}
	var result string
	for i, task := range tasks {
		result += fmt.Sprintf("%d.%s", i+1, task) + "\n"
	}
	return result, nil

}

func CompletedTask(assignedTaskIndex int) (string, error) {
	//open the AssignedTask file
	assfile, err := os.Open("AssignedTask.txt")
	if err != nil {
		return "", err
	}
	//close the file
	defer assfile.Close()
	//create a slice to store the assigned task
	var assigTask []string
	//Scan in bytes (reading)
	scanner := bufio.NewScanner(assfile)
	for scanner.Scan() {
		assigTask = append(assigTask, scanner.Text())
	}
	// making sure the index is the actual available indexes
	if assignedTaskIndex > len(assigTask)+1 {
		return "", errors.Taskerr{Message: "Out of index"}
	}
	//Get that particular task you want to markAscompleted
	markcomTask := assigTask[assignedTaskIndex-1]
	//split to get you to be able to change the pending part to completed
	Splitted := strings.Split(markcomTask, "|")
	Splitted[2] = "Completed"
	//join back the splitted which turned array to string form seperated by "|"
	assigTask[assignedTaskIndex-1] = strings.Join(Splitted, "|")
	//We now create a erase the content of the last file and copy the editedSlice into it
	file, err2 := os.OpenFile("AssignedTask.txt", os.O_TRUNC|os.O_WRONLY, 0644)
	if err2 != nil {
		return "", err2
	}
	defer file.Close()
	//Now we copy the editedSlice by looping through it to the file after erasing it
	for _, task := range assigTask {
		_, err3 := file.WriteString(task + "\n")
		if err3 != nil {
			return "", err3
		}
	}
	return Splitted[0] + " has completed the Task " + Splitted[1], nil

}
func AllocateTask(index int, user_index int) (string, error) {
	//first we open the task.txt to get the task to be allocated through its index
	file, err := os.Open("tasks.txt")
	if err != nil {
		return "", err
	}
	// Create an array we arrange the task into cause of the index we about to use
	var Taskarr []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Taskarr = append(Taskarr, scanner.Text())

	}
	task := Taskarr[index-1]
	// then we open the user.txt to get the user to be given a task
	userfile, err1 := os.Open("Students.txt")
	if err1 != nil {
		return "", err1
	}
	defer file.Close()
	var userarr []string
	userScanner := bufio.NewScanner(userfile)
	for userScanner.Scan() {
		userarr = append(userarr, userScanner.Text())
	}
	user := userarr[user_index-1]
	//Now we create the assignTask.txt to save the students assigned to a task and it current status
	assignFile, asserr := os.OpenFile("AssignedTask.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if asserr != nil {
		return "", asserr
	}
	defer assignFile.Close()
	_, err2 := assignFile.WriteString(user + "|" + task + "|" + "Pending" + "\n")
	if err2 != nil {
		return "", err2
	}
	return user + " has been allocated a task", nil

}
func ViewTask() (string, int, error) {
	file1, err := os.Open("tasks.txt")
	if err != nil {
		return "", 0, err
	}
	var tasks []string
	var result string
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())

	}
	var count int
	for i, task := range tasks {
		result += fmt.Sprintf("%d.%s", i+1, task) + "\n"
		count = i + 1
	}
	return "Here are the list of tasks " + "\n" + result, count, nil
}
func ViewassignTask() (string, int, error) {
	file1, err := os.Open("AssignedTask.txt")
	if err != nil {
		return "", 0, err
	}
	var tasks []string
	var result string
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())

	}
	var count int
	for i, task := range tasks {
		result += fmt.Sprintf("%d.%s", i+1, task) + "\n"
		count = i + 1
	}
	return "Here are the list of tasks " + "\n" + result, count, nil
}
