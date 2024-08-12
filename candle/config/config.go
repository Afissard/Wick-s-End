package config

type Config struct {
	// Debug
	DebugEnable bool

	// Screen
	ScreenWidth     int
	ScreenHeight    int
	ScreenScaleCoef int

	// Tile
	TileSize int
}

var Global Config

func Init(screenWidth, screenHeight, screenScaleCoef int) {
	Global = Config{
		// Screen
		ScreenWidth:     screenWidth,
		ScreenHeight:    screenHeight,
		ScreenScaleCoef: screenScaleCoef,

		// Tile
		TileSize: 8,
	}
}
