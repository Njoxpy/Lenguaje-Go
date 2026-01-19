package main

type Student struct {
	Name  string
	Age   int
	Grade float64
}

var studentsList []Student

func addStudent(name string, age int, grade float64) {
	newStudent := Student{Name: name, Age: age, Grade: grade}

	studentsList = append(studentsList, newStudent)
}

func main() {

}
