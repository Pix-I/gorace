package game

import (
	"fmt"
	"math/rand"
	"time"
)

type GameMap struct {
	Matrix             [11][11]string
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
	var newMatrix [11][11]string
	placeSouthTrees(&newMatrix)
	for i := 0; i < 10; i++ {
		newMatrix[i+1] = cg.Matrix[i]
	}
	newMatrix[5][5] = "C"
	newMatrix[6][5] = ""
	cg.Matrix = newMatrix
}

func (cg *GameMap) MoveDown() {
	if cg.Matrix[cg.X][cg.Y-1] == "T" {
		return
	}
	var newMatrix [11][11]string
	placeNorthTrees(&newMatrix)
	for i := 0; i < 10; i++ {
		newMatrix[i] = cg.Matrix[i+1]
	}
	newMatrix[5][5] = "C"
	newMatrix[4][5] = ""
	cg.Matrix = newMatrix
}

func (cg *GameMap) MoveLeft() {
	if cg.Matrix[cg.X-1][cg.Y] == "T" {
		return
	}
	var newMatrix [11][11]string
	placeWestTrees(&newMatrix)
	for i := 0; i < 11; i++ {
		for j := 0; j < 10; j++ {
			newMatrix[i][j] = cg.Matrix[i][j+1]
		}
	}
	newMatrix[5][5] = "C"
	newMatrix[5][4] = ""
	cg.Matrix = newMatrix
}

func (cg *GameMap) MoveRight() {
	if cg.Matrix[cg.X+1][cg.Y] == "T" {
		return
	}
	var newMatrix [11][11]string
	placeEastTrees(&newMatrix)
	for i := 0; i < 11; i++ {
		for j := 0; j < 9; j++ {
			newMatrix[i][j+1] = cg.Matrix[i][j]
		}
	}
	newMatrix[5][5] = "C"
	newMatrix[5][6] = ""
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

func addCar(matrix *[11][11]string) {
	matrix[5][5] = "C"
}

func initTrees(matrix *[11][11]string) {
	placeNorthTrees(matrix)
	placeSouthTrees(matrix)
	placeEastTrees(matrix)
	placeWestTrees(matrix)
}

func placeSouthTrees(matrix *[11][11]string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 4; i++ {
		var placeTree = r1.Intn(9)
		matrix[0][placeTree+1] = "T"
	}
}

func placeNorthTrees(matrix *[11][11]string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 4; i++ {
		var placeTree = r1.Intn(8)
		matrix[9][placeTree+1] = "T"
	}
}
func placeWestTrees(matrix *[11][11]string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 4; i++ {
		var placeTree = r1.Intn(8)
		matrix[placeTree+1][0] = "T"
	}
}

func placeEastTrees(matrix *[11][11]string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 4; i++ {
		var placeTree = r1.Intn(8)
		matrix[placeTree+1][9] = "T"
	}
}
