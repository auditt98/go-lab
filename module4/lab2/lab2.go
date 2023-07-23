package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearConsole() {
	var clearCmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		clearCmd = exec.Command("cmd", "/c", "cls")
	default:
		clearCmd = exec.Command("clear")
	}

	clearCmd.Stdout = os.Stdout
	_ = clearCmd.Run()
}

func printLight(color string, seconds int) {
	clearConsole()
	fmt.Println("Color:", color)
	fmt.Println("Seconds Left:", seconds)
	time.Sleep(1 * time.Second)
}

func RandomTrafficLight() {
	colors := []string{"red", "yellow", "green"}
	for {
		for _, color := range colors {
			clearConsole()
			switch color {
			case "red":
				for i := 15; i > 0; i-- {
					printLight(color, i)
				}
			case "green":
				for i := 30; i > 0; i-- {
					printLight(color, i)
				}
			case "yellow":
				for i := 3; i > 0; i-- {
					printLight(color, i)
				}
			}
		}
	}
}

func main() {
	go RandomTrafficLight()
	select {}
}
