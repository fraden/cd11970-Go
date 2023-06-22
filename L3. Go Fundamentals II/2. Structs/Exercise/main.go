package main

import "fmt"

type Student struct {
	id   int
	name string
}

type Classroom struct {
	id          int
	capacity    uint
	subject     string
	studentList []Student
}

func main() {

	c1 := Classroom{
		id:       1,
		capacity: 20,
		subject:  "Math",
		studentList: []Student{
			{
				id:   234,
				name: "Daniel",
			},
			{
				id:   235,
				name: "Jennifer",
			},
		},
	}

	c2 := new(Classroom)
	c2.id = 2
	c2.capacity = 30
	c2.subject = "Music"
	c2.studentList = []Student{
		{
			id:   50,
			name: "Sven",
		},
		{
			id:   51,
			name: "Ida",
		},
	}

	fmt.Println(c1)
	fmt.Println(c2)
}
