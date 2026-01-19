package main

type Student struct {
    Name   string
    Scores []int
}

func addStudent(students []Student, name string, scores ...int) []Student {
	newStudent := Student {
		Name: name,
		Scores: scores
	}

	students := append(students, newStudent)

	return students

}
