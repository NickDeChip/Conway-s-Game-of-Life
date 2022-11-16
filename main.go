package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	winWidth  = 800
	winHeight = 600
	scl       = 10
	rows      = winWidth / scl
	cols      = winHeight / scl
)

var curGen = [rows][cols]int8{}
var nextGen = [rows][cols]int8{}

var timer float32

var cycles int
var amountAlive int
var displayAmountAlive int

func main() {
	rl.InitWindow(winWidth, winHeight, "LIFE?!")
	rl.SetTargetFPS(60)

	restart()

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeyR) {
			restart()
		}

		timer += rl.GetFrameTime()

		if timer >= 0.1 {
			timer = 0
			cycles += 1
			gen()
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		draw()

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func draw() {
	for i := int32(0); i < rows; i++ {
		for j := int32(0); j < cols; j++ {
			switch curGen[i][j] {
			case 0:
				rl.DrawRectangle(i*scl, j*scl, int32(scl), int32(scl), rl.White)
			case 1:
				rl.DrawRectangle(i*scl, j*scl, int32(scl), int32(scl), rl.Black)
			}
			rl.DrawRectangleLines(i*scl, j*scl, int32(scl), int32(scl), rl.Black)
		}
	}

	rl.DrawText(fmt.Sprintf("Cycles: %d", cycles), 20, 0, 30, rl.Red)
	rl.DrawText(fmt.Sprintf("Alive: %d", displayAmountAlive), winWidth*0.8, 0, 30, rl.Red)
}

func gen() {
	friends := int8(0)
	for i := int32(0); i < rows; i++ {
		for j := int32(0); j < cols; j++ {
			friends = countFriends(i, j)

			alive := curGen[i][j]

			if alive == 1 {
				amountAlive += 1
			}

			if friends < 2 && alive == 1 {
				nextGen[i][j] = 0
			} else if friends > 3 && alive == 1 {
				nextGen[i][j] = 0
			} else if friends == 3 && alive == 0 {
				nextGen[i][j] = 1
			} else {
				nextGen[i][j] = curGen[i][j]
			}
		}

	}
	displayAmountAlive = amountAlive
	amountAlive = 0
	curGen = nextGen
}

func countFriends(x int32, y int32) int8 {
	sum := int8(0)
	for i := int32(-1); i <= 1; i++ {
		for j := int32(-1); j <= 1; j++ {

			row := (x + i + rows) % rows
			col := (y + j + cols) % cols

			sum += curGen[row][col]
		}
	}

	sum -= curGen[x][y]

	return sum
}

func restart() {
	timer = 0
	cycles = 0
	for i := int32(0); i < rows; i++ {
		for j := int32(0); j < cols; j++ {
			curGen[i][j] = int8(rl.GetRandomValue(0, 1))
		}
	}
}
