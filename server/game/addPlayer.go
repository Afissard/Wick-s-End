package game

func (g *Game) AddPlayer(name string){
	newPlayer := player{
		id : len(g.AllPlayers),
		name: name,
	}
	g.AllPlayers = append(g.AllPlayers, newPlayer)
}