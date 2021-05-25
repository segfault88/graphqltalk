package store

type Movie struct {
	ID         int
	Name       string
	DirectorID int
	Stars      []int
	Runtime    int
}

var movies = []Movie{
	{ID: 1, Name: "Shaun of the Dead", DirectorID: 1, Stars: []int{1, 2, 3}, Runtime: 99},
	{ID: 2, Name: "Hot Fuzz", DirectorID: 1, Stars: []int{1, 2, 4}, Runtime: 121},
	{ID: 3, Name: "The World's End", DirectorID: 1, Stars: []int{1, 2, 4}, Runtime: 109},
	{ID: 4, Name: "Star Trek (2009)", DirectorID: 2, Stars: []int{1, 5, 6}, Runtime: 127},
	{ID: 5, Name: "Mission: Impossible - Ghost Protocol", DirectorID: 3, Stars: []int{1, 8, 9}, Runtime: 132},
}

type Director struct {
	ID     int
	Name   string
	Movies []int
}

var directors = []Director{
	{ID: 1, Name: "Edger Wright", Movies: []int{1, 2, 3}},
	{ID: 2, Name: "J.J. Abrams", Movies: []int{4}},
	{ID: 3, Name: "Brad Bird", Movies: []int{5}},
}

type Star struct {
	ID     int
	Name   string
	Movies []int
}

var stars = []Star{
	{ID: 1, Name: "Simon Pegg", Movies: []int{1, 2, 3, 4, 5}},
	{ID: 2, Name: "Nick Frost", Movies: []int{1, 2, 3}},
	{ID: 3, Name: "Kate Ashfield", Movies: []int{1}},
	{ID: 4, Name: "Martin Freeman", Movies: []int{2, 3}},
	{ID: 5, Name: "Chris Pine", Movies: []int{4}},
	{ID: 6, Name: "Zachary Quinto", Movies: []int{4}},
	{ID: 8, Name: "Tom Cruise", Movies: []int{5}},
	{ID: 9, Name: "Jeremy Renner", Movies: []int{5}},
}
