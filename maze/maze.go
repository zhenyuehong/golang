package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row) //注意这里 ［］的类型是［］int
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(dir point) point {
	return point{p.i + dir.i, p.j + dir.j}
}

//判断是否越界
func (p point) pointAt(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		//如果到终点，退出
		if cur == end {
			break
		}
		//maze at next is 0
		//and steps at next is 0
		//and next != start

		//当用于遍历数组和切片的时候，range函数返回索引和元素；
		//当用于遍历字典的时候，range函数返回字典的键和值。
		for _, dir := range dirs {
			next := cur.add(dir)

			val, ok := next.pointAt(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.pointAt(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			curStep, _ := cur.pointAt(steps)

			steps[next.i][next.j] = curStep + 1
			Q = append(Q, next)
		}

	}

	return steps
}

func main() {
	maze := readMaze("maze/maze.in")
	//for _, row := range maze {
	//	for _, val := range row {
	//		fmt.Printf("%d ", val)
	//	}
	//	fmt.Println()
	//}
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

}
