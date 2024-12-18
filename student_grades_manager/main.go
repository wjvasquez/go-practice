package main

import (
	"fmt"
	"slices"
)

func main() {
	var numberStudents, sumGrades, highestGrade,
		lowestGrade, studentGrade int
	var averageGrade float64
	var studentGrades []int
	var studentName, highestGradeStudent, lowestGradeStudent string
	var studentNames []string

	fmt.Print("Enter the number of students: ")
	fmt.Scan(&numberStudents)

	if numberStudents <= 0 {
		fmt.Println("Error: The number of students must be a positive integer.")
		return
	}

	for i := 0; i < numberStudents; i++ {
		fmt.Printf("Enter student %d name: ", i + 1)
		fmt.Scan(&studentName)

		fmt.Printf("Enter %s's grade: ", studentName)
		fmt.Scan(&studentGrade)

		if studentGrade < 0 || studentGrade > 100 {
			for {
				fmt.Print("Error, please enter a grade between 0-100, try again: ")
				fmt.Scan(&studentGrade)

				if studentGrade >= 0 && studentGrade <= 100 {
					break
				}
			}
		}

		studentNames = append(studentNames, studentName)
		studentGrades = append(studentGrades, studentGrade)
	}

	for i := 0; i < len(studentGrades); i++ {
		sumGrades = sumGrades + studentGrades[i]
	}
	averageGrade = float64(sumGrades) / float64(len(studentGrades))

	highestGrade = slices.Max(studentGrades)
	lowestGrade = slices.Min(studentGrades)

	highestGradeStudent = studentNames[slices.Index(studentGrades, highestGrade)]
	lowestGradeStudent = studentNames[slices.Index(studentGrades, lowestGrade)]

	fmt.Printf("Average grade: %.2f\n", averageGrade)
	fmt.Printf("Highest grade: %d (%s) \n", highestGrade, highestGradeStudent)
	fmt.Printf("Lowest grade: %d (%s) \n", lowestGrade, lowestGradeStudent)
}
