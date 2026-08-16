package main

// import Student
// A package is simply a folder containing Go files that belong together.
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"task-manager/Student"
	"task-manager/Task"
	// "task-manager/Student"
)

// var s Student.Student(name:"hamid")
// var s Student.Student{Name:"Hamid"}

func main() {
	fmt.Println("WELCOME TO TESTING MY MINI TASK MANAGER CLI")
	menus := []string{
		"Add Student",
		"View Students",
		"Add Task",
		"View Tasks",
		"Assign Task",
		"View Assigned Tasks",
		"Mark Task Completed",
		"Exit",
	}
	for i, menu := range menus {
		fmt.Println(i+1, ".", menu)
	}
	var choice int
	var err error
	//making sure if the input is wrong, it is able to redisplay choose your option
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Choose your option: ")
		scanner.Scan()
		input := scanner.Text()
		choice, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter an integer")
			continue
		}
		if choice > len(menus) || choice <= 0 {
			fmt.Println("Number not in range or out of range")
			continue
		}
		//leave the for loop
		break
	}
	switch choice {
	case 1:
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter the student name: ")
		scanner.Scan()
		name := scanner.Text()
		studentt := Student.Student{Name: name}
		msg, err := studentt.AddStudent()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
	case 2:
		std, _, err := Student.ViewStudent()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(std)
	case 3:
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter the new task: ")
		scanner.Scan()
		input := scanner.Text()

		Tasks, err := Task.CreateTask(input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(Tasks)
	case 4:
		task, _, err := Task.ViewTask()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(task)
	case 5:
		var scanner *bufio.Scanner
		var count int
		var Taskchoice int
		var taskcount int
		var err error
		var std string
		//Showing the available tasks
		fmt.Println("Below are the Tasks available")
		task, taskcount, err := Task.ViewTask()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(task)

		for {
			//Choosing the task we want to allocate
			scanner = bufio.NewScanner(os.Stdin)
			fmt.Println("Choose the task to be allocated: ")
			scanner.Scan()
			input := scanner.Text()
			//convert the input from the ascii it is to integer
			Taskchoice, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Please enter an integer")
				continue
			}
			//Making sure the input is within range
			if Taskchoice > taskcount || Taskchoice <= 0 {
				fmt.Println("Number not in range or out of range")
				continue
			} else {
				//show students
				fmt.Println("Below are the available Students")
				std, count, err = Student.ViewStudent()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(std)
				break
			}

		}
		for {
			//choosing the student to be allocated a task
			fmt.Println("Choose the student to be given Task: ")
			scanner.Scan()
			stuinput := scanner.Text()
			Studentchoice, err := strconv.Atoi(stuinput)
			//Making sure it is an integer and in range
			if err != nil {
				fmt.Println("Please enter an integer")
				continue
			} else if Studentchoice > count || Studentchoice <= 0 {
				fmt.Println("Number not in range or out of range")
				continue
			} else {
				msg, err := Task.AllocateTask(Taskchoice, Studentchoice)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(msg)
				break
			}
		}
	case 6:
		task, _, err := Task.ViewassignTask()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(task)

	case 7:
		//Show the assigned tasks to chose from
		fmt.Println("These are the assigned task: ")
		task, count, err := Task.ViewassignTask()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(task)
		for {
			//choose the assigned task to be mark completed
			fmt.Println("Choose the assigned task to be Marked completed: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()
			Studentchoice, err := strconv.Atoi(input)
			//Making sure it is an integer and in range
			if err != nil {
				fmt.Println("Please enter an integer")
				continue
			} else if Studentchoice > count || Studentchoice <= 0 {
				fmt.Println("Number not in range or out of range")
				continue
			} else {
				msg, err := Task.CompletedTask(Studentchoice)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(msg)
				break
			}
		}
	case 8:
		fmt.Println("Ctrl + C ")

	}

	//os.Stdin cause it has read method which satisfies the i.o reader interface
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("What is your name: ")
	// scanner.Scan()
	// inputName := scanner.Text()
	// fmt.Println("Hello, " + inputName)
	// s := Student.Student{Name: "Hamid"}
	// fmt.Println(s.AddStudent())
	// t, err := task.CompletedTask(12)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(t)
	// s, err := Student.ViewStudent()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(s)

}
