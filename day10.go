package main

import "fmt"

func main() {

	userGrades := make(map[string]int)

	userGrades["alice"] = 90
	userGrades["bob"] = 75
	userGrades["alice"] = 92
	fmt.Println("Grades:", userGrades)

	bobsGrade := userGrades["bob"]
	fmt.Println("Bob's Grade:", bobsGrade)

	charliesGrade := userGrades["charlie"]
	fmt.Println("Charlie's Grade:", charliesGrade)

	grade, ok := userGrades["charlie"]
	if ok {
		fmt.Println("Charlie's grade is:", grade)
	} else {
		fmt.Println("Charlie is not in the system.")
	}

	delete(userGrades, "bob")
	fmt.Println("After deleting Bob:", userGrades)

	fmt.Println("\nLooping over grades:")
	for key, val := range userGrades {
		fmt.Println("User:", key, "Grade:", val)
	}
}
