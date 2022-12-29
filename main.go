package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type PlayGround struct {
	direction  rune
	snake      [][]int
	food       [][]int
	background [][]int
}

func newPlayGround() *PlayGround {
	snake := [][]int{{0, 0}}
	food := [][]int{{1, 2}}
	var background [][]int

	//init background 25*25
	for i := 0; i < 25; i++ {
		var row []int
		for j := 0; j < 25; j++ {
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

	return &PlayGround{snake: snake, food: food, background: background}
}

func (p *PlayGround) move(direction rune) bool {
	//头部和尾部
	head := p.snake[0]
	//tail := p.snake[len(p.snake)-1]
	switch direction {
	case 'w':
		head[0]--
	case 's':
		head[0]++
	case 'a':
		head[1]--
	case 'd':
		head[1]++
	}
	//移动身体
	//for i := len(p.snake) - 1; i > 0; i-- {
	//	p.snake[i] = p.snake[i-1]
	//}

	//是否碰撞
	if p.isCollision() {
		fmt.Println("碰撞")
		return false
	}
	//是否吃到食物
	if p.eatFood() {
		p.food = [][]int{}
		p.randomFood()
	}
	//更新背景
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
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
	if p.snake[0][0] < 0 || p.snake[0][0] > 24 || p.snake[0][1] < 0 || p.snake[0][1] > 24 {
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

//追加蛇身
func (p *PlayGround) appendSnake() {
	var newSnake [][]int
	for _, v := range p.snake {
		newSnake = append(newSnake, v)
	}
	newSnake = append(newSnake, []int{0, 0})
	p.snake = newSnake
}

//随机生成食物
func (p *PlayGround) randomFood() {
	//TODO
	//0-24随机数
	x := rand.Intn(25)
	y := rand.Intn(25)
	p.food = append(p.food, []int{x, y})
}

func (p *PlayGround) print() {
	for _, v := range p.background {
		for _, v2 := range v {
			if v2 == 0 {
				fmt.Print("□")
			}
			if v2 == 1 {
				fmt.Print("■")
			}
			if v2 == 2 {
				fmt.Print("★")
			}
		}
		fmt.Println()
	}
}

func main() {
	//全局随机做种
	rand.Seed(time.Now().UnixNano())
	playGround := newPlayGround()
	//print background
	playGround.move('d')
	playGround.print()
	//开始游戏
	reader := bufio.NewReader(os.Stdin)
	for {
		//w a s d
		fmt.Println("请输入方向：")
		char, _, _ := reader.ReadRune()
		if !playGround.move(char) {
			fmt.Println("game over")
			break
		}
		playGround.print()
	}
}
