package main

import (
	"fmt"
	"sort"
)

func main() {

	var subjects = []string{"maths", "social science"}
	fmt.Println("Subjects list are: ", subjects)

	subjects = append(subjects, "english", "computers")

	fmt.Println("Updated Subjects list are: ", subjects)
	fmt.Printf("Type of subjectList: %T\n", subjects)

	// start from 2nd index to 3rd index. Last value always not included.
	fmt.Println("indexing in slices: ", subjects[2:4])

	// initialization using make
	scores := make([]int, 3)

	scores[0] = 1
	scores[1] = 2
	scores[2] = 3

	fmt.Println("Scores list: ", scores)
	fmt.Printf("Type of scores list: %T\n", scores)

	scores = append(scores, 34, 45, 78)

	fmt.Println("Scores list: ", scores)

	fmt.Println("Is list sorted: ", sort.IntsAreSorted(scores))

	scores = append(scores, 1)

	sort.Ints(scores)
	fmt.Println("Newly sorted: ", scores)

	// remove elements from slice
	var idx int = 3
	scores = append(scores[:idx], scores[idx+1:]...)
	fmt.Println("Scores after deletion of 3rd index", scores)

	//range of deletion
	start, end := 2, 4
	scores = append(scores[:start], scores[end+1:]...)
	fmt.Println("Scores after range of deletion", scores)

}
