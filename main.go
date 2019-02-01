package main

import (
	"be.pixelogical/gorace/game"
)

func main() {
	example := game.GameMap{GoalX: 25, GoalY: 20, X: 4, Y: 4, GameName: "Awesome example"}
	example.InitGame()
	example.MoveDown()
	example.MoveDown()
	example.MoveDown()
	example.MoveDown()
	example.MoveLeft()
	example.MoveRight()
	example.MoveDown()
	example.MoveDown()
	example.PrintMap()
}
