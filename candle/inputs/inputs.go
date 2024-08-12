package inputs

type Inputs struct {
	Up, Down, Left, Right bool // directional input
	A, B, X, Y            bool // in-game inputs
	Menu, Enter, Debug    bool // special  for menu and other
}
