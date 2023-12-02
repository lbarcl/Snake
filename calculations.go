package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func CalculateSnakePosition(direction rl.Vector2, parts []rl.Vector2) {
	for i := len(parts) - 1; i >= 0; i-- {
		if i == 0 {
			parts[0].X += direction.X
			parts[0].Y += direction.Y
		} else {
			parts[i].X = parts[i-1].X
			parts[i].Y = parts[i-1].Y
		}
	}

}

func SpawnFruit(snakeParts []rl.Vector2) rl.Vector2 {
	var X int32 = 0
	var Y int32 = 0
l:
	for {
		X = rand.Int31n(gridWidth)
		Y = rand.Int31n(gridHeight)

		if X == 0 || Y == 0 {
			continue
		}

		for i := range snakeParts {
			if snakeParts[i].X == float32(X) && snakeParts[i].Y == float32(Y) {
				continue l
			}
		}

		break
	}

	return GridToViewport(X, Y)
}

func CheckCollisions(snakeParts []rl.Vector2, fruit rl.Vector2) string {
	for i := 1; i < len(snakeParts); i++ {
		if snakeParts[i].X == snakeParts[0].X && snakeParts[i].Y == snakeParts[0].Y {
			return "snake"
		}
	}

	if snakeParts[0].X == fruit.X && snakeParts[0].Y == fruit.Y {
		return "fruit"
	}

	if snakeParts[0].X < 0 || snakeParts[0].X > gridWidth || snakeParts[0].Y < 0 || snakeParts[0].Y > gridHeight {
		return "wall"
	}

	return ""
}
