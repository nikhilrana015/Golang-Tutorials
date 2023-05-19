package main

import "fmt"

type Movie struct {
	ID          int64
	MovieName   string
	MovieLength int64
	Rating      float64
}

func main() {

	movie1 := Movie{1, "Dhoom-2", 128, 4.5}
	fmt.Printf("Details of the movie: %+v\n", movie1)

	movie1.updateRating()

	fmt.Printf("Details of the movie: %+v\n", movie1)

}

// pass by refernce here
func (m *Movie) updateRating() {
	m.Rating = 3.5
	fmt.Println("New ratnig of the movie is: ", m.Rating)
}
