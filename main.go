package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var settings = LoadSettings()

const gridBoxSize = 50

var gridWidth int32 = settings.ViewportWidth / gridBoxSize
var gridHeight int32 = settings.ViewportHeight / gridBoxSize
var scaleX int32 = settings.ViewportWidth / gridWidth
var scaleY int32 = settings.ViewportHeight / gridHeight

var textures map[string]rl.Texture2D
var audio map[string]rl.Sound
var gameState int8 = -1
var score int8 = 0

func initialize() {
	rl.InitWindow(settings.ViewportWidth, settings.ViewportHeight, "Snake")
	rl.SetTargetFPS(settings.TargetFPS)
	rl.InitAudioDevice()

	LoadAudio("audio")
	LoadTextures("sprites")
}

func close() {
	UnloadSound()
	UnloadTexture()
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func draw(snakeParts []rl.Vector2, snakeHeadDirection rl.Vector2, fruitLocation rl.Vector2) {
	rl.BeginDrawing()
	DrawGrid()

	if gameState == 0 {
		DrawSnake(snakeParts, snakeHeadDirection)
		DrawFruit(fruitLocation)
		rl.DrawText(strconv.Itoa(int(score)), int32((gridWidth/2)*gridBoxSize)-10, gridBoxSize, 50, rl.White)
	} else {
		rl.DrawText("Press \"R\" to play ", int32((gridWidth/6)*gridBoxSize), gridBoxSize*2, 45, rl.White)
	}

	if settings.DevMode {
		rl.DrawFPS(10, 10)
	}

	rl.EndDrawing()
}

func gameLoop(snakeHeadDirection *rl.Vector2, snakeParts []rl.Vector2, fruitLocation *rl.Vector2) []rl.Vector2 {
	switch rl.GetKeyPressed() {
	case settings.UpKey:
		snakeHeadDirection.Y = -1
		snakeHeadDirection.X = 0

	case settings.DownKey:
		snakeHeadDirection.Y = 1
		snakeHeadDirection.X = 0

	case settings.LeftKey:
		snakeHeadDirection.X = -1
		snakeHeadDirection.Y = 0

	case settings.RightKey:
		snakeHeadDirection.X = 1
		snakeHeadDirection.Y = 0

	}

	CalculateSnakePosition(*snakeHeadDirection, snakeParts)
	collision := CheckCollisions(snakeParts, *fruitLocation)

	switch collision {
	case "wall":
		gameState = -1
		rl.PlaySound(audio["Hit.wav"])

	case "snake":
		gameState = -1
		rl.PlaySound(audio["Hit.wav"])

	case "fruit":
		rl.PlaySound(audio["Swallow.wav"])
		snakeParts = append(snakeParts, snakeParts[len(snakeParts)-1])
		*fruitLocation = SpawnFruit(snakeParts)
		score++

		if int32(score) == (gridHeight*gridWidth)-3 {
			gameState = 1
		}
	}

	return snakeParts
}

func main() {
	initialize()

	var snakeHeadDirection rl.Vector2
	var snakeParts []rl.Vector2
	var fruitLocation = SpawnFruit(snakeParts)

	for !rl.WindowShouldClose() {
		if gameState == -1 {
			snakeHeadDirection = rl.Vector2{X: 0, Y: -1}
			snakeParts = []rl.Vector2{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7}}
			fruitLocation = SpawnFruit(snakeParts)

			if rl.GetKeyPressed() == settings.RestartKey {
				gameState = 0
				score = 0
			}

		} else {
			snakeParts = gameLoop(&snakeHeadDirection, snakeParts, &fruitLocation)
		}

		draw(snakeParts, snakeHeadDirection, fruitLocation)
	}

	close()
}
