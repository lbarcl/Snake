package main

import rl "github.com/gen2brain/raylib-go/raylib"

func DrawGrid() {
	for x := 0; x < gridWidth; x++ {
		for y := 0; y < gridHeight; y++ {
			if (x%2 == 0 && y%2 != 0) || (x%2 != 0 && y%2 == 0) {
				rl.DrawRectangle(int32(x*gridBoxSize), int32(y*gridBoxSize), gridBoxSize, gridBoxSize, rl.DarkGreen)
			} else {
				rl.DrawRectangle(int32(x*gridBoxSize), int32(y*gridBoxSize), gridBoxSize, gridBoxSize, rl.Green)
			}
		}
	}
}

func DrawSnake(parts []rl.Vector2) {
	for i := range parts {
		vpCords := GridToViewport(int32(parts[i].X), int32(parts[i].Y))
		rl.DrawCircle(int32(vpCords.X), int32(vpCords.Y), float32((gridBoxSize)/2-(i*2)), rl.Red)
	}
}
