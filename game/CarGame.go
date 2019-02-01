package game

import (
	"fmt"
	"math/rand"
	"time"
)

type GameMap struct {
	Matrix             [10][10]string
	GoalX, GoalY, X, Y int
	GameName           string
}

func (cg *GameMap) InitGame() {
	addCar(&cg.Matrix)
	initTrees(&cg.Matrix)
}

func (cg *GameMap) MoveUp() {
	if cg.Matrix[cg.X][cg.Y+1] == "T" {
		return
	}
	var newMatrix [10][10]string
	placeSouthTrees(&newMatrix)
	for i := 0; i < 9; i++ {
		newMatrix[i+1] = cg.Matrix[i]
	}
	newMatrix[4][4] = "C"
	newMatrix[5][4] = ""
	cg.Matrix = newMatrix
}

func (cg *GameMap) MoveDown() {
	if cg.Matrix[cg.X][cg.Y-1] == "T" {
		return
	}
	var newMatrix [10][10]string
	placeNorthTrees(&newMatrix)
	for i := 0; i < 9; i++ {
		newMatrix[i] = cg.Matrix[i+1]
	}
	newMatrix[4][4] = "C"
	newMatrix[3][4] = ""
	cg.Matrix = newMatrix
}

func (cg *GameMap) MoveLeft() {
	if cg.Matrix[cg.X-1][cg.Y] == "T" {
		return
	}
	var newMatrix [10][10]string
	placeWestTrees(&newMatrix)
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			newMatrix[i][j] = cg.Matrix[i][j+1]
		}
	}
	newMatrix[4][4] = "C"
	newMatrix[4][3] = ""
	cg.Matrix = newMatrix
}

func (cg *GameMap) MoveRight() {
	if cg.Matrix[cg.X+1][cg.Y] == "T" {
		return
	}
	var newMatrix [10][10]string
	placeEastTrees(&newMatrix)
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			newMatrix[i][j+1] = cg.Matrix[i][j]
		}
	}
	newMatrix[4][4] = "C"
	newMatrix[4][5] = ""
	cg.Matrix = newMatrix
}

func (cg *GameMap) GameEnded() bool {
	if cg.X == cg.GoalX && cg.Y == cg.GoalY {
		return true
	}
	return false
}

func (cg *GameMap) PrintMap() {
	fmt.Print(cg.Matrix)
}

func addCar(matrix *[10][10]string) {
	matrix[4][4] = "C"
}

func initTrees(matrix *[10][10]string) {
	placeNorthTrees(matrix)
	placeSouthTrees(matrix)
	placeEastTrees(matrix)
	placeWestTrees(matrix)
}

func placeSouthTrees(matrix *[10][10]string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 4; i++ {
		var placeTree = r1.Intn(8)
		matrix[0][placeTree+1] = "T"
	}
}

func placeNorthTrees(matrix *[10][10]string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 4; i++ {
		var placeTree = r1.Intn(8)
		matrix[9][placeTree+1] = "T"
	}
}
func placeWestTrees(matrix *[10][10]string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 4; i++ {
		var placeTree = r1.Intn(8)
		matrix[placeTree+1][0] = "T"
	}
}

func placeEastTrees(matrix *[10][10]string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 4; i++ {
		var placeTree = r1.Intn(8)
		matrix[placeTree+1][9] = "T"
	}
}
