package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawGrid() {
	var x int32 = 0
	for x < gridWidth {
		var y int32 = 0
		for y < gridHeight {
			if (x%2 == 0 && y%2 != 0) || (x%2 != 0 && y%2 == 0) {
				rl.DrawRectangle(int32(x*gridBoxSize), int32(y*gridBoxSize), gridBoxSize, gridBoxSize, settings.BlackSquareColor)
			} else {
				rl.DrawRectangle(int32(x*gridBoxSize), int32(y*gridBoxSize), gridBoxSize, gridBoxSize, settings.LightSquareColor)
			}
			if settings.DevMode {
				rl.DrawText(strconv.Itoa(int(x))+", "+strconv.Itoa(int(y)), int32(x*gridBoxSize), int32(y*gridBoxSize), 10, rl.White)
			}
			y++
		}
		x++
	}
}

func DrawSnake(parts []rl.Vector2, headDirection rl.Vector2) {
	for i := 0; i < len(parts); i++ {
		vpCords := GridToViewport(int32(parts[i].X), int32(parts[i].Y))

		if i == 0 {
			drawTexture("Head.png", DirectionToRotation(headDirection), vpCords)
		} else if i == len(parts)-1 {
			tailDirection := rl.Vector2Subtract(parts[i-1], parts[i])
			drawTexture("Tail.png", DirectionToRotation(tailDirection), vpCords)
		} else {
			dNext := rl.Vector2Subtract(parts[i], parts[i+1])
			dbefore := rl.Vector2Subtract(parts[i-1], parts[i])

			if dNext.X == dbefore.X && dNext.Y == dbefore.Y {
				// Straight segment
				directionToNext := rl.Vector2Subtract(parts[i], parts[i+1])
				drawTexture("Mid.png", DirectionToRotation(directionToNext), vpCords)
			} else if dbefore.X < 0 {
				if dNext.Y == -1 {
					// Turn to the left from the left
					drawTexture("To_Left.png", 270, vpCords)
				} else if dNext.Y == 1 {
					// Turn to the left from the left
					drawTexture("To_Right.png", 270, vpCords)
				}
			} else if dbefore.X > 0 {
				// Similar logic for turns to the right
				if dNext.Y == -1 {
					drawTexture("To_Right.png", 90, vpCords)
				} else if dNext.Y == 1 {
					drawTexture("To_Left.png", 90, vpCords)
				}
			} else if dbefore.Y < 0 {
				if dNext.Y == 1 {
					drawTexture("To_Left.png", 0, vpCords)
				} else if dNext.Y == -1 {
					drawTexture("To_Right.png", 0, vpCords)
				} else if dNext.X == 1 {
					drawTexture("To_Left.png", 0, vpCords)
				} else {
					drawTexture("To_Right.png", 0, vpCords)
				}
			} else if dbefore.Y > 0 {
				// Similar logic for turns downward
				if dNext.X == 1 {
					drawTexture("To_Right.png", 180, vpCords)
				} else if dNext.X == -1 {
					drawTexture("To_Left.png", 180, vpCords)
				}
			}

			//fmt.Printf("db: %f, %f | dn: %f, %f \n", dbefore.X, dbefore.Y, dNext.X, dNext.Y)
		}
	}
}

func drawTexture(partName string, rotation float32, cords rl.Vector2) {
	rl.DrawTexturePro(
		textures[partName],
		rl.Rectangle{X: 0, Y: 0, Width: float32(textures[partName].Width), Height: float32(textures[partName].Height)},
		rl.Rectangle{X: cords.X - 25, Y: cords.Y - 25, Width: float32(scaleX), Height: float32(scaleY)},
		rl.Vector2{X: float32(scaleX / 2), Y: float32(scaleY / 2)},
		rotation,
		rl.White,
	)
}

func DrawFruit(fruitPosition rl.Vector2) {
	vpCords := GridToViewport(int32(fruitPosition.X), int32(fruitPosition.Y))
	drawTexture("Fruit.png", 0, vpCords)
}
