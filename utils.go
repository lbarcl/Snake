package main

import (
	"io/fs"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GridToViewport(X int32, Y int32) rl.Vector2 {
	return rl.Vector2{X: float32(X) * scaleX, Y: float32(Y) * scaleY}
}

func ViewportToGrid(viewportCoord rl.Vector2) rl.Vector2 {
	gridX := (viewportCoord.X + 25) / scaleX
	gridY := (viewportCoord.Y + 25) / scaleY

	return rl.Vector2{X: gridX, Y: gridY}
}

func LoadTextures(path string) {
	textures = make(map[string]rl.Texture2D)

	filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			textures[info.Name()] = rl.LoadTexture(path)
		}

		return nil
	})
}

func LoadAudio(path string) {
	audio = make(map[string]rl.Sound)

	filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			audio[info.Name()] = rl.LoadSound(path)
		}

		return nil
	})
}

func UnloadSound() {
	for k := range audio {
		rl.UnloadSound(audio[k])
	}
}

func UnloadTexture() {
	for k := range textures {
		rl.UnloadTexture(textures[k])
	}
}
