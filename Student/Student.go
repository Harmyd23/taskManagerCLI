package Student

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	Name string
}

func (s Student) AddStudent() (string, error) {
	file, err := os.OpenFile(
		"Students.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err1 := file.WriteString(s.Name + "\n")
	if err1 != nil {
		return "", err1
	}
	file1, _ := os.Open("Students.txt")
	last := ""
	var names []string
	var result string
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		names = append(names, scanner.Text())
		last = scanner.Text()
	}
	for i, name := range names {
		result += fmt.Sprintf("%d.%s", i+1, name) + "\n"
	}
	return last + " has been successfully added, Here are the updated list of names " + "\n" + result, nil
}
func ViewStudent() (string, int, error) {
	file1, err := os.Open("Students.txt")
	if err != nil {
		return "", 0, err
	}
	var names []string
	var result string
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		names = append(names, scanner.Text())

	}
	var count int
	for i, name := range names {
		result += fmt.Sprintf("%d.%s", i+1, name) + "\n"
		count = i + 1
	}
	return "Here are the list of names " + "\n" + result, count, nil
}
