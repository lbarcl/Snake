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

func DrawSnake(parts []rl.Vector2, headDirection rl.Vector2) {
	var rotation float32

	if headDirection.Y == -1 {
		rotation = 0
	} else if headDirection.Y == 1 {
		rotation = 180
	} else if headDirection.X == -1 {
		rotation = 270
	} else {
		rotation = 90
	}

	for i := range parts {
		vpCords := GridToViewport(int32(parts[i].X), int32(parts[i].Y))
		if i == 0 {
			rl.DrawTextureEx(textures["Head.png"], vpCords, rotation, 1, rl.White)
		} else if i != len(parts)-1 {
			var dX [2]int = [2]int{int(parts[i-1].X - parts[i].X), int(parts[i].X - parts[i+1].X)}
			var dY [2]int = [2]int{int(parts[i-1].Y - parts[i].Y), int(parts[i].Y - parts[i+1].Y)}

			if dX[0] == dX[1] && dY == dY {
				rl.DrawTextureEx(textures["Mid.png"], vpCords, rotation, 1, rl.White)
			} else {

			}
		}
	}
}
