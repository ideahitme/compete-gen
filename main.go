package main

import (
	"os"
	"fmt"
	"log"
	"github.com/ideahitme/compete-gen/templates"
	"os/exec"
)

const (
	runCPP = "g++ -pipe -std=c++11 code.cpp -o code && ./code < input"
	runGo = "cat input | go run code.go"
)


func help() {
	fmt.Println("Usage example:")
	fmt.Println("Create cpp directory: gen cpp")
	fmt.Println("Create go directory: gen go")
	fmt.Println("Run cpp with input: gen cpp (same command). Equivalent to `g++ -pipe -std=c++11 code.cpp -o code && ./code < input`")
	fmt.Println("Run go with input: gen go (same command). Equivalent to `cat input | go run code.go`")
}

// checks if command now should run instead of create
func isRun() bool {
	if _, err := os.Stat("input"); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Fatalf("Failed to check if project already initialized: %v", err)
	}
	return true
}

func createInput() {
	if _, err := os.Create("input"); err != nil {
		log.Fatalf("Failed to create input file: %v", err)
	}
}

func run(cmd string) {
	shCmd := exec.Command("sh", "-c", fmt.Sprintf(`%s`,cmd))
	output, err := shCmd.CombinedOutput();
	if err != nil {
		log.Fatalf("Failed to run %s: %v", cmd, err)
	}
	fmt.Println(string(output))
}

func createCpp() {
	file, err := os.Create("code.cpp")
	if err != nil {
		log.Fatalf("Error creating code.cpp file: %v", err)
	}
	if _, err := file.Write([]byte(templates.CPP)); err != nil {
		log.Fatalf("Error writing to code.cpp file: %v", err)
	}
}

func createGo() {
	file, err := os.Create("code.go")
	if err != nil {
		log.Fatalf("Error creating code.go file: %v", err)
	}
	if _, err := file.Write([]byte(templates.GO)); err != nil {
		log.Fatalf("Error writing to code.go file: %v", err)
	}
}

func main () {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}
	lang := os.Args[1]
	switch lang {
	case "cpp":
		if isRun() {
			run(runCPP)
			return
		}
		createInput()
		createCpp()
	case "go":
		if isRun() {
			run(runGo)
			return
		}
		createInput()
		createGo()
	default:
		help()
		os.Exit(1)
	}
	fmt.Println("done!")
}


