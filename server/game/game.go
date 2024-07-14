package game

type Game struct {
	AllPlayers []player
}

type player struct {
	id   int
	name string
}
