package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	// islands
	islandmap = make([]string, screena)

	// selected menu
	selectedmenuon                    bool
	item1, item2, item3, item4, item5 string
	selectedmenuh, selectedmenuw      int32
	// map
	screenw, screenh, screena int
	// cursor
	activeblock, selectedblock int
	// sizes
	lrg, sml  bool
	txtlrg    = int32(20)
	txtsml    = int32(10)
	linelrg   = txtlrg
	linesml   = txtsml
	txtsize   int32
	linespace int32
	// core
	monh32, monw32                 int32
	monitorh, monitorw, monitornum int
	grid16on, grid4on, debugon     bool
	framecount                     int
	mousepos                       rl.Vector2
	camera                         rl.Camera2D
)

func updateall() { // MARK: updateall()

	if grid4on {
		grid4()
	}
	if grid16on {
		grid16()
	}
	if debugon {
		debug()
	}

	getactiveblock()
}
func getactiveblock() { // MARK: getactiveblock()

	xchange := 0
	ychange := 0

	ycount := 0

	for b := 0; b < screenh; b++ {
		if mousepos.Y < float32(monitorh-ychange) && mousepos.Y > float32((monitorh-16)-ychange) {
			for a := 0; a < screenw; a++ {
				if mousepos.X > float32(0+xchange) && mousepos.X < float32(16+xchange) {
					activeblock = a + ycount
				}
				xchange += 16
			}
		}
		ychange += 16
		ycount += screenw
	}
}
func main() { // MARK: main()
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides INFO window

	startsettings()
	raylib()
}
func setscreen() { // MARK: setscreen()
	monitornum = rl.GetMonitorCount()
	monitorh = rl.GetScreenHeight()
	monitorw = rl.GetScreenWidth()
	monh32 = int32(monitorh)
	monw32 = int32(monitorw)
	rl.SetWindowSize(monitorw, monitorh)
	setsizes()
}
func setsizes() { // MARK: setsizes()
	if monitorw >= 1600 {
		txtsize = txtlrg
		linespace = linelrg
		lrg = true
		sml = false
	} else if monitorw < 1600 && monitorw >= 1280 {
		txtsize = txtsml
		linespace = linesml
		lrg = false
		sml = true
	}

	screenh = monitorh / 16
	screenw = monitorw / 16
	screena = screenh * screenw
}
func startsettings() { // MARK: start
	camera.Zoom = 1.0
	camera.Target.X = 0.0
	camera.Target.Y = 0.0
	activeblock = -1
	selectedblock = -1
	debugon = true
	grid16on = true
	selectedmenuon = true
}

func selectedmenu() { // MARK: selectedmenu()
	item1 = "delete"
	item2 = "add action"
	item3 = "change"
	item4 = "change"
	item5 = ""

	count := int32(0)

	if item1 != "" {
		count++
	}
	if item2 != "" {
		count++
	}
	if item3 != "" {
		count++
	}
	if item4 != "" {
		count++
	}
	if item5 != "" {
		count++
	}
	if lrg {
		selectedmenuh = (count * linelrg) + 8
		selectedmenuw = 200
	} else if sml {
		selectedmenuh = (count * linesml) + 8
		selectedmenuw = 100
	}

}

func raylib() { // MARK: raylib()
	rl.InitWindow(monw32, monh32, "boom island")
	setscreen()
	rl.CloseWindow()
	rl.InitWindow(monw32, monh32, "boom island")
	setscreen()

	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window
	//	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() { // MARK: WindowShouldClose

		mousepos = rl.GetMousePosition()
		framecount++
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.BeginMode2D(camera)

		// MARK: draw map layer 1

		drawx := int32(0)
		drawy := int32(monitorh - 16)
		screencount := 0
		for a := 0; a < screena; a++ {

			if selectedblock == a {
				rl.DrawRectangle(drawx, drawy, 16, 16, rl.Fade(rl.Red, 0.2))
				if selectedmenuon {
					selectedmenu()
					rl.DrawRectangle(drawx+18, drawy, selectedmenuw, selectedmenuh, rl.Fade(rl.Green, 0.4))
					rl.DrawText(item1, drawx+24, drawy+2, txtsize, rl.White)
					rl.DrawText(item2, drawx+24, drawy+2+linespace, txtsize, rl.White)
					rl.DrawText(item3, drawx+24, drawy+2+linespace*2, txtsize, rl.White)
					rl.DrawText(item4, drawx+24, drawy+2+linespace*3, txtsize, rl.White)
				}
			}

			drawx += 16
			screencount++
			if screencount == screenw {
				screencount = 0
				drawx = 0
				drawy -= 16
			}

		}

		// MARK: draw map layer 2

		rl.EndMode2D() // MARK: draw no camera
		input()
		updateall()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
func debug() { // MARK: debug
	rl.DrawRectangle(monw32-300, 0, 500, monw32, rl.Fade(rl.Black, 0.7))
	rl.DrawFPS(monw32-290, monh32-100)

	mouseposXTEXT := fmt.Sprintf("%.0f", mousepos.X)
	mouseposYTEXT := fmt.Sprintf("%.0f", mousepos.Y)
	screenhTEXT := strconv.Itoa(screenh)
	screenwTEXT := strconv.Itoa(screenw)
	activeblockTEXT := strconv.Itoa(activeblock)
	selectedblockTEXT := strconv.Itoa(selectedblock)
	screenaTEXT := strconv.Itoa(screena)

	rl.DrawText(mouseposXTEXT, monw32-290, linespace, txtsize, rl.White)
	rl.DrawText("mouseposX", monw32-200, linespace, txtsize, rl.White)
	rl.DrawText(mouseposYTEXT, monw32-290, linespace*2, txtsize, rl.White)
	rl.DrawText("mouseposY", monw32-200, linespace*2, txtsize, rl.White)
	rl.DrawText(screenhTEXT, monw32-290, linespace*3, txtsize, rl.White)
	rl.DrawText("screenh", monw32-200, linespace*3, txtsize, rl.White)
	rl.DrawText(screenwTEXT, monw32-290, linespace*4, txtsize, rl.White)
	rl.DrawText("screenw", monw32-200, linespace*4, txtsize, rl.White)
	rl.DrawText(activeblockTEXT, monw32-290, linespace*5, txtsize, rl.White)
	rl.DrawText("activeblock", monw32-200, linespace*5, txtsize, rl.White)
	rl.DrawText(selectedblockTEXT, monw32-290, linespace*6, txtsize, rl.White)
	rl.DrawText("selectedblock", monw32-200, linespace*6, txtsize, rl.White)
	rl.DrawText(screenaTEXT, monw32-290, linespace*7, txtsize, rl.White)
	rl.DrawText("screena", monw32-200, linespace*7, txtsize, rl.White)

}
func input() { // MARK: keys input
	if rl.IsKeyPressed(rl.KeyF1) {
		if selectedmenuon {
			selectedmenuon = false
		} else {
			selectedmenuon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF2) {
		if grid16on {
			grid16on = false
		} else {
			grid16on = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF3) {
		if grid4on {
			grid4on = false
		} else {
			grid4on = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF3) {
		if selectedmenuon {
			selectedmenuon = false
		} else {
			selectedmenuon = true
		}
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		selectedblock = activeblock
	}
	if rl.IsKeyPressed(rl.KeyRightControl) {
		camera.Target.X = 0
		camera.Target.Y = 0
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		camera.Target.X -= 16
	}
	if rl.IsKeyDown(rl.KeyRight) {

	}
	if rl.IsKeyDown(rl.KeyUp) {
		camera.Target.Y -= 16
	}
	if rl.IsKeyDown(rl.KeyDown) {
		camera.Target.Y += 16
	}
	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpAdd) {
		if camera.Zoom < 2.1 {
			camera.Zoom += 0.1
		}
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		if camera.Zoom > 0.5 {
			camera.Zoom -= 0.1
		}
	}
}
func grid16() { // MARK: grid16()
	for a := 0; a < monitorw; a += 16 {
		a32 := int32(a)
		rl.DrawLine(a32, 0, a32, monh32, rl.Fade(rl.Green, 0.1))
	}
	for a := monitorh - 16; a > 0; a -= 16 {
		a32 := int32(a)
		rl.DrawLine(0, a32, monw32, a32, rl.Fade(rl.Green, 0.1))
	}
}
func grid4() { // MARK: grid4()
	for a := 0; a < monitorw; a += 4 {
		a32 := int32(a)
		rl.DrawLine(a32, 0, a32, monh32, rl.Fade(rl.White, 0.1))
	}
	for a := monitorh - 16; a > 0; a -= 4 {
		a32 := int32(a)
		rl.DrawLine(0, a32, monw32, a32, rl.Fade(rl.White, 0.1))
	}
}

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
