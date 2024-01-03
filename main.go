package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/go-vgo/robotgo"
)

type Position struct {
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
	IsDrag bool    `json:"is_drag"`
}

var mtx sync.Mutex

func main() {
	fmt.Println("Press 'Ctrl + C' to exit.")
	// robotgo.MouseSleep = 300
	robotgo.ActiveName("Prodrafts")

	http.HandleFunc("/moveMouse", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		// 创建一个 MyData 结构体实例
		var requestData []Position

		// 将请求的 body 数据解析为 JSON 并存入 requestData 中
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		// 在服务器端打印解析后的数据
		// fmt.Printf("Received JSON data: %+v\n", requestData)
		mtx.Lock()
		for _, v := range requestData {
			if v.IsDrag {
				dragMouse(int(v.X), int(v.Y))
			} else {
				robotgo.Move(int(v.X), int(v.Y))
			}
		}
		mtx.Unlock()

		// 返回响应给客户端
		fmt.Fprint(w, "JSON data received and processed successfully!")
	})

	// 启动HTTP服务器，监听在本地的8080端口
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 模拟鼠标移动
	// robotgo.MoveSmooth(200, 200)

	// // 等待一段时间
	// // time.Sleep(2 * time.Second)

	// // 模拟鼠标点击
	// // robotgo.Click("left", true)
	// // time.Sleep(time.Second)
	// // robotgo.Click("left", true)

	// // 等待一段时间
	// // time.Sleep(2 * time.Second)

	// // 模拟鼠标移动到另一个位置
	// dragMouse(300, 500)
	// // 等待一段时间
	// time.Sleep(2 * time.Second)
}

// func moveMouse(x, y int) {
// 	// 获取当前鼠标位置
// 	currentX, currentY := robotgo.GetMousePos()

// 	fmt.Printf("Current Mouse Position: (%d, %d)\n", currentX, currentY)

// 	// 移动鼠标到指定位置
// 	robotgo.MoveMouseSmooth(x, y)
// 	// robotgo.DragSmooth(x, y)
// 	// robotgo.MoveMouse(x, y)
// 	fmt.Printf("Moved Mouse To: (%d, %d)\n", x, y)
// }

func dragMouse(x, y int) {
	// 获取当前鼠标位置
	// currentX, currentY := robotgo.GetMousePos()

	// fmt.Printf("Current Mouse Position: (%d, %d)\n", currentX, currentY)

	// 移动鼠标到指定位置
	// robotgo.MoveMouseSmooth(x, y)
	// robotgo.DragSmooth(x, y)
	// robotgo.MouseToggle()
	// robotgo.Click("left", false)
	robotgo.Toggle("left")
	robotgo.Drag(x, y)
	robotgo.MilliSleep(100)
	// robotgo.MoveSmooth(x, y, 2.0, 3.0)
	// robotgo.MilliSleep(500)
	robotgo.Toggle("left", "up")

	// robotgo.MoveMouse(x, y)
	// fmt.Printf("Moved Mouse To: (%d, %d)\n", x, y)
}
