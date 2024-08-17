package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func detectGraphicsEngine(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Fehler beim Öffnen der Datei: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	signatures := []string{
		"Direct3D",  
		"OpenGL",   
		"Vulkan",    
		"D3D",       
		"DXGI",     
		"shader",   
		"texture",   
		"gameengine",
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, signature := range signatures {
			if strings.Contains(line, signature) {
				return true 
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while opening the file: %v\n", err)
		os.Exit(1)
	}

	return false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run detect.go <Path to .exe-file>")
		return
	}

	appPath := os.Args[1]

	if detectGraphicsEngine(appPath) {
		fmt.Printf("The application %s uses a graphic-engine. Recomendation: Proton.\n", appPath)
	} else {
    fmt.Printf("The application %s uses a graphic-engine. Recomendation: Wine.\n", appPath)
	}
}

