package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("Press 'Ctrl + C' to exit.")
	robotgo.ActiveName("Prodrafts")

	// 模拟鼠标移动
	moveMouse(200, 200)

	// 等待一段时间
	// time.Sleep(2 * time.Second)

	// 模拟鼠标点击
	robotgo.Click("left", true)
	// time.Sleep(time.Second)
	robotgo.Click("left", true)

	// 等待一段时间
	// time.Sleep(2 * time.Second)

	// 模拟鼠标移动到另一个位置
	dragMouse(300, 300)
	moveMouse(201, 201)
	dragMouse(300, 400)
	moveMouse(301, 401)
	time.Sleep(time.Second)
	dragMouse(350, 450)
	moveMouse(351, 451)
	dragMouse(550, 850)
	// 等待一段时间
	time.Sleep(2 * time.Second)
}

func moveMouse(x, y int) {
	// 获取当前鼠标位置
	currentX, currentY := robotgo.GetMousePos()

	fmt.Printf("Current Mouse Position: (%d, %d)\n", currentX, currentY)

	// 移动鼠标到指定位置
	robotgo.MoveMouseSmooth(x, y)
	// robotgo.DragSmooth(x, y)
	// robotgo.MoveMouse(x, y)
	fmt.Printf("Moved Mouse To: (%d, %d)\n", x, y)
}

func dragMouse(x, y int) {
	// 获取当前鼠标位置
	// currentX, currentY := robotgo.GetMousePos()

	// fmt.Printf("Current Mouse Position: (%d, %d)\n", currentX, currentY)

	// 移动鼠标到指定位置
	// robotgo.MoveMouseSmooth(x, y)
	robotgo.Drag(x, y)
	// robotgo.MouseToggle()
	// robotgo.Click("left", false)
	// robotgo.Toggle("left")

	// robotgo.MilliSleep(100)
	// robotgo.MoveSmooth(x, y, 2.0, 3.0)
	// robotgo.MilliSleep(500)
	// robotgo.Toggle("left", "up")

	// robotgo.MoveMouse(x, y)
	fmt.Printf("Moved Mouse To: (%d, %d)\n", x, y)
}
