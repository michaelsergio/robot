package main

import (
	"fmt"
	"github.com/michaelsergio/robot"
)

func main() {
	var rob robot.XRobot
	rob.NewRobot()
	rob.KeyPress(65)
	rob.MouseMove(100, 100)
	fmt.Println(rob)
	rob.Version()
}
