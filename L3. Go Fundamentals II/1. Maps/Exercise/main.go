package main

import (
	"fmt"
	"strings"
)

func main() {
	courses := map[int]string{
		1: "Calculus",
		2: "Biology",
		3: "Chemistry",
		4: "Computer Science",
		5: "Communications",
		6: "English",
		7: "Cantonese",
	}

	for _, course := range courses {
		if strings.HasPrefix(course, "C") {
			fmt.Println(course)
		}
	}

	courses[4] = "Algorithms"

	courses[9] = "Spanish"

	delete(courses, 1)

	fmt.Println(courses)
}
