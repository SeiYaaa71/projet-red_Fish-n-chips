package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"runtime"
)

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	fmt.Println("=== Jeu RPG Console ===")
}

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Magenta= "\033[35m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
	Purple = "\033[35m"
)

func waitForEnter() {
	fmt.Print("\n" + Cyan + "Appuyez sur Entrée pour continuer..." + Reset)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
