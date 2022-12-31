package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type PlayGround struct {
	direction  string
	snake      [][]int
	food       [][]int
	background [][]int
	width      int
	height     int
}

func newPlayGround() *PlayGround {
	snake := [][]int{{0, 2}, {0, 1}, {0, 0}}
	food := [][]int{{1, 2}}
	width := 25
	height := 32
	var background [][]int

	//init background
	for i := 0; i < width; i++ {
		var row []int
		for j := 0; j < height; j++ {
			row = append(row, 0)
		}
		background = append(background, row)
	}

	//init snake
	for _, v := range snake {
		background[v[0]][v[1]] = 1
	}

	//init food
	for _, v := range food {
		background[v[0]][v[1]] = 2
	}

	return &PlayGround{snake: snake, food: food, background: background, width: width, height: height}
}

func (p *PlayGround) move(direction string) bool {
	//头部和尾部
	head := p.snake[0]
	//尾部
	x := p.snake[len(p.snake)-1][0]
	y := p.snake[len(p.snake)-1][1]
	tail := []int{x, y}

	//移动身体
	if len(p.snake) > 1 {
		for i := len(p.snake) - 1; i > 0; i-- {
			p.snake[i][0] = p.snake[i-1][0]
			p.snake[i][1] = p.snake[i-1][1]
		}
	}

	switch direction {
	case "w":
		head[0]--
	case "s":
		head[0]++
	case "a":
		head[1]--
	case "d":
		head[1]++
	}

	//移动头部
	p.snake[0] = head

	//是否碰撞
	if p.isCollision() {
		return false
	}
	//是否吃到食物
	if p.eatFood() {
		p.snake = append(p.snake, tail)
		p.food = [][]int{}
		p.randomFood()
	}
	//更新背景
	for i := 0; i < p.width; i++ {
		for j := 0; j < p.height; j++ {
			p.background[i][j] = 0
		}
	}
	for _, v := range p.snake {
		p.background[v[0]][v[1]] = 1
	}
	for _, v := range p.food {
		p.background[v[0]][v[1]] = 2
	}
	return true
}

//碰撞检测
func (p *PlayGround) isCollision() bool {
	//边界碰撞
	if p.snake[0][0] < 0 || p.snake[0][0] > p.width || p.snake[0][1] < 0 || p.snake[0][1] > p.height {
		return true
	}
	//自身碰撞
	for i := 1; i < len(p.snake); i++ {
		if p.snake[0][0] == p.snake[i][0] && p.snake[0][1] == p.snake[i][1] {
			return true
		}
	}
	return false
}

//吃食物
func (p *PlayGround) eatFood() bool {
	if p.snake[0][0] == p.food[0][0] && p.snake[0][1] == p.food[0][1] {
		return true
	}
	return false
}

//随机生成食物
func (p *PlayGround) randomFood() {
	//TODO
	//0-24随机数
	x := rand.Intn(p.width)
	y := rand.Intn(p.height)
	//判断是否与蛇重合
	for _, v := range p.snake {
		if v[0] == x && v[1] == y {
			p.randomFood()
			return
		}
	}
	p.food = append(p.food, []int{x, y})
}

func print(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += 1
	}
}

func (p *PlayGround) render(background [][]int) {
	w, h := termbox.Size()
	l := w/2 - p.height
	t := h/2 - p.width/2
	//在正中央渲染
	for i := 0; i < len(background); i++ {
		for j := 0; j < len(background[i]); j++ {
			if background[i][j] == 0 {
				print(l+j*2, i+t, termbox.ColorWhite, termbox.ColorWhite, "  ")
			}
			if background[i][j] == 1 {
				print(l+j*2, i+t, termbox.ColorGreen, termbox.ColorGreen, "██")
			}
			if background[i][j] == 2 {
				print(l+j*2, i+t, termbox.ColorRed, termbox.ColorRed, "██")
			}
		}
	}
	//渲染边界
	//for i := 0; i < 25; i++ {
	//	print(i*2+5, 25, termbox.ColorYellow, termbox.ColorYellow, "  ")
	//}
	//for i := 0; i < 25; i++ {
	//	print(i*2+5, 0, termbox.ColorYellow, termbox.ColorYellow, "  ")
	//}
	//for i := 0; i < 25; i++ {
	//	print(5, i, termbox.ColorYellow, termbox.ColorYellow, "  ")
	//}
	//for i := 0; i < 25; i++ {
	//	print(55, i, termbox.ColorYellow, termbox.ColorYellow, "  ")
	//}
	termbox.Flush()
}

func keyTransfer(key termbox.Key) string {
	//TODO wasd
	switch key {
	case termbox.KeyArrowUp:
		return "w"
	case termbox.KeyArrowDown:
		return "s"
	case termbox.KeyArrowLeft:
		return "a"
	case termbox.KeyArrowRight:
		return "d"
	default:
		return " "
	}
}

func main() {
	//全局随机做种
	rand.Seed(time.Now().UnixNano())
	playGround := newPlayGround()

	if err := termbox.Init(); err != nil {
		panic(err)
	}

	//w, h := termbox.Size()

	playGround.render(playGround.background)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				termbox.Close()
				return
			default:
				if playGround.move(keyTransfer(ev.Key)) {
					playGround.render(playGround.background)
				} else {
					termbox.Close()
					fmt.Println("Game Over")
					return
				}
			}
			time.Sleep(time.Second / 10)
		}
	}
}
