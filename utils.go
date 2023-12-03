package main

import (
	"bufio"
	"image/color"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GridToViewport(X int32, Y int32) rl.Vector2 {
	return rl.Vector2{X: float32(X * scaleX), Y: float32(Y * scaleY)}
}

func ViewportToGrid(viewportCoord rl.Vector2) rl.Vector2 {
	gridX := (viewportCoord.X + 25) / float32(scaleX)
	gridY := (viewportCoord.Y + 25) / float32(scaleY)

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

func ReadFromFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			data[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func HexToColor(hex string) (rl.Color, error) {
	// Remove the "#" prefix if present
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	// Parse hex string to a 32-bit integer
	value, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return color.RGBA{}, err
	}

	// Extract RGB components
	red := uint8((value >> 16) & 0xFF)
	green := uint8((value >> 8) & 0xFF)
	blue := uint8(value & 0xFF)

	// Return color.RGBA struct
	return color.RGBA{R: red, G: green, B: blue, A: 255}, nil
}

func LoadSettings() Settings {
	keys, e := ReadFromFile("Settings.txt")

	if e != nil {
		log.Fatal(e)
	}

	ViewportWidth, e := strconv.Atoi(keys["width"])
	ViewportHeight, e := strconv.Atoi(keys["height"])
	TargetFPS, e := strconv.Atoi(keys["fps"])
	UpKey, e := strconv.Atoi(keys["up"])
	DownKey, e := strconv.Atoi(keys["down"])
	LeftKey, e := strconv.Atoi(keys["left"])
	RightKey, e := strconv.Atoi(keys["right"])
	RestartKey, e := strconv.Atoi(keys["restart"])
	BlackSquareColor, e := HexToColor(keys["black_square_color"])
	LightSquareColor, e := HexToColor(keys["light_square_color"])

	if e != nil {
		log.Fatal(e)
	}

	var DevMode bool

	if keys["dev_mode"] == "true" {
		DevMode = true
	} else {
		DevMode = false
	}

	return Settings{
		ViewportWidth:    int32(ViewportWidth),
		ViewportHeight:   int32(ViewportHeight),
		TargetFPS:        int32(TargetFPS),
		UpKey:            int32(UpKey),
		DownKey:          int32(DownKey),
		LeftKey:          int32(LeftKey),
		RightKey:         int32(RightKey),
		RestartKey:       int32(RestartKey),
		BlackSquareColor: BlackSquareColor,
		LightSquareColor: LightSquareColor,
		DevMode:          DevMode,
	}
}

type Settings struct {
	ViewportWidth    int32
	ViewportHeight   int32
	TargetFPS        int32
	DevMode          bool
	BlackSquareColor rl.Color
	LightSquareColor rl.Color
	RestartKey       int32
	UpKey            int32
	DownKey          int32
	LeftKey          int32
	RightKey         int32
}
