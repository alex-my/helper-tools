package utils

import (
	"fmt"
	"sync"
	"time"
)

// 终端进度条
// 使用:
// bar := utils.NewBar(5)
// bar.Run()
//
// 每执行完一个任务，调用一次 bar.Add()
//
// 结束时，执行 bar.WaitClose()

// Bar 进度条
type Bar struct {
	// total 总进度
	total int

	// curr 当前进度
	curr int

	// ext 额外自定义消息
	ext string

	// spend 耗时(秒)
	spend int64

	// refresh 刷新进度条
	refresh chan bool

	// closeFlag 完成
	closeFlag chan bool

	closeOnce sync.Once

	// doneProcess 已完成的进度条
	doneProcess string

	// remainProcess 未完成的进度条
	remainProcess string
}

// NewBar 创建一个进度条
func NewBar(total int) *Bar {
	bar := &Bar{
		total:     total,
		refresh:   make(chan bool),
		closeFlag: make(chan bool),
	}

	return bar
}

// Run 执行进度条
func (bar *Bar) Run() {
	go bar.spendTime()
	go bar.move()
}

// Add 完成一个进度
func (bar *Bar) Add(ext ...string) {
	bar.curr++

	if len(ext) > 0 {
		bar.ext = ext[0]
	}

	// 计算进度条
	rate := bar.curr * 100 / bar.total
	bar.doneProcess = ""
	for i := 0; i < rate; i++ {
		bar.doneProcess += ">"
	}
	bar.remainProcess = ""
	for i := 0; i < (100 - rate); i++ {
		bar.remainProcess += "-"
	}

	bar.refresh <- true
}

// WaitClose 等待进度条输出
func (bar *Bar) WaitClose() {
	for {
		select {
		case <-bar.closeFlag:
			return
		case <-time.After(2 * time.Second):
			return
		}
	}
}

func (bar *Bar) init() {
	width := 100

	bar.doneProcess = ""
	for i := 0; i < width; i++ {
		bar.remainProcess += "-"
	}
}

func (bar *Bar) spendTime() {
	for {
		select {
		case <-time.After(time.Second):
			bar.spend++
			bar.refresh <- true
		case <-bar.closeFlag:
			return
		}
	}
}

func (bar *Bar) move() {
	for range bar.refresh {
		message := fmt.Sprintf("%s%s [%d:%d][%s] %s", bar.doneProcess, bar.remainProcess, bar.curr, bar.total, bar.timeToString(), bar.ext)

		fmt.Printf("\r\x1b[32;1m[%s\x1b[39;22m", message)

		if bar.curr == bar.total {
			// 等待一段时间再结束，等 fmt 输出
			<-time.After(time.Second)
			bar.close()
		}
	}
}

func (bar *Bar) timeToString() string {
	hour := bar.spend / 3600
	minute := (bar.spend - hour*3600) / 60
	second := bar.spend - hour*3600 - minute*60

	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func (bar *Bar) close() {
	bar.closeOnce.Do(func() {
		bar.closeFlag <- true
	})
}
