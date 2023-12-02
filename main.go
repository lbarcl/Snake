package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const gridBoxSize = 50
const viewportWidth = 500
const viewportHeight = 500
const gridWidth = viewportWidth / gridBoxSize
const gridHeight = viewportHeight / gridBoxSize
const scaleX = viewportWidth / gridWidth
const scaleY = viewportHeight / gridHeight

var textures map[string]rl.Texture2D
var audio map[string]rl.Sound
var gameState int8 = -1
var score int8 = 0

func initialize() {
	rl.InitWindow(viewportWidth, viewportHeight, "Snake")
	rl.SetTargetFPS(2)
	rl.InitAudioDevice()

	LoadAudio("audio/")
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
		rl.DrawCircle(int32(fruitLocation.X), int32(fruitLocation.Y), gridBoxSize/2, rl.Blue)
		rl.DrawText(strconv.Itoa(int(score)), int32((gridWidth/2)*gridBoxSize)-10, gridBoxSize, 50, rl.White)
	} else {
		rl.DrawText("Press \"R\" to play ", int32((gridWidth/6)*gridBoxSize), gridBoxSize*2, 45, rl.White)
	}
	rl.EndDrawing()
}

func gameLoop(snakeHeadDirection *rl.Vector2, snakeParts []rl.Vector2, fruitLocation *rl.Vector2) {
	switch rl.GetKeyPressed() {
	case rl.KeyUp:
		snakeHeadDirection.Y = -1
		snakeHeadDirection.X = 0
	case rl.KeyDown:
		snakeHeadDirection.Y = 1
		snakeHeadDirection.X = 0
	case rl.KeyLeft:
		snakeHeadDirection.X = -1
		snakeHeadDirection.Y = 0
	case rl.KeyRight:
		snakeHeadDirection.X = 1
		snakeHeadDirection.Y = 0
	}

	CalculateSnakePosition(*snakeHeadDirection, snakeParts)
	collision := CheckCollisions(snakeParts, ViewportToGrid(*fruitLocation))

	switch collision {
	case "wall":
		gameState = -1

	case "snake":
		gameState = -1
	case "fruit":
		rl.PlaySound(audio["swallow.waw"])
		score++
		newTail := rl.Vector2{X: -1, Y: -1}
		snakeParts = append(snakeParts, newTail)
		*fruitLocation = SpawnFruit(snakeParts)

		if score == (gridHeight*gridWidth)-1 {
			gameState = 1
		}
	}
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

			if rl.GetKeyPressed() == rl.KeyR {
				gameState = 0
				score = 0
			}

		} else {
			gameLoop(&snakeHeadDirection, snakeParts, &fruitLocation)
		}

		draw(snakeParts, snakeHeadDirection, fruitLocation)
	}

	close()
}
