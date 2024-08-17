package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Funktion zum Überprüfen, ob eine Datei eine Grafik-Engine verwendet
func detectGraphicsEngine(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Fehler beim Öffnen der Datei: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Typische Zeichenfolgen, die auf eine Grafik-Engine hinweisen könnten
	signatures := []string{
		"Direct3D",  // DirectX
		"OpenGL",    // OpenGL
		"Vulkan",    // Vulkan
		"D3D",       // Allgemeine DirectX Signatur
		"DXGI",      // DirectX Graphics Infrastructure
		"shader",    // Shader-Programmen
		"texture",   // Texturen, typisches Grafikelement
		"gameengine", // Häufiges Wort in Spiel-Engines
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, signature := range signatures {
			if strings.Contains(line, signature) {
				return true // Grafik-Engine erkannt
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Fehler beim Lesen der Datei: %v\n", err)
		os.Exit(1)
	}

	return false // Keine Grafik-Engine gefunden
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Nutzung: go run detect.go <Pfad zur .exe-Datei>")
		return
	}

	appPath := os.Args[1]

	if detectGraphicsEngine(appPath) {
		fmt.Printf("Die Anwendung %s verwendet eine Grafik-Engine. Empfohlen: Proton.\n", appPath)
	} else {
		fmt.Printf("Die Anwendung %s verwendet keine bekannte Grafik-Engine. Empfohlen: Wine.\n", appPath)
	}
}

