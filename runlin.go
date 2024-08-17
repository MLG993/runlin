package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func isInstalled(cmdName string) bool {
	_, err := exec.LookPath(cmdName)
	return err == nil
}

func installDependencies(deps []string) {
	fmt.Println("Überprüfen und Installieren der Abhängigkeiten...")

	for _, dep := range deps {
		if !isInstalled(dep) {
			fmt.Printf("Installiere %s...\n", dep)
			installCmd := exec.Command("sudo", "yay", "-S", "--needed", "--noconfirm", dep)
			installCmd.Stdout = os.Stdout
			installCmd.Stderr = os.Stderr
			err := installCmd.Run()
			if err != nil {
				fmt.Printf("Fehler bei der Installation von %s: %v\n", dep, err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("%s ist bereits installiert.\n", dep)
		}
	}
	fmt.Println("Alle Abhängigkeiten sind installiert.")
}

func installPerformanceEnhancers() {
	fmt.Println("Überprüfen auf optionale Performance-Verbesserungen...")

	performanceDeps := []string{"dxvk-bin", "vkd3d-bin"} 
	for _, dep := range performanceDeps {
		if !isInstalled(dep) {
			fmt.Printf("%s ist nicht installiert. Möchten Sie es installieren? (j/n): ", dep)
			var response string
			fmt.Scanln(&response)
			if strings.ToLower(response) == "j" {
				installDependencies([]string{dep})
			}
		} else {
			fmt.Printf("%s ist bereits installiert.\n", dep)
		}
	}
	fmt.Println("Alle Performance-Verbesserungen sind installiert.")
}

func runWithWine(appPath string) {
	fmt.Printf("Starte %s mit Wine...\n", appPath)
	cmd := exec.Command("wine", appPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Fehler beim Starten der Anwendung mit Wine: %v\n", err)
		os.Exit(1)
	}
}

func runWithProton(appPath string) {
	fmt.Printf("Starte %s mit Proton...\n", appPath)
	protonPath := "/usr/bin/proton" 

	cmd := exec.Command(protonPath, "run", appPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Fehler beim Starten der Anwendung mit Proton: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Nutzung: go run runlin.go [proton | wine | install | performance] <Pfad zur .exe-Datei>")
		return
	}

	command := os.Args[1]
	appPath := filepath.Clean(os.Args[2])

	switch strings.ToLower(command) {
	case "install":
		installDependencies([]string{"wine", "proton-ge-custom-bin"})
	case "wine":
		if isInstalled("wine") {
			runWithWine(appPath)
		} else {
			fmt.Println("Wine ist nicht installiert. Bitte führen Sie 'go run runlin.go install' aus, um es zu installieren.")
		}
	case "proton":
		if isInstalled("proton") {
			runWithProton(appPath)
		} else {
			fmt.Println("Proton ist nicht installiert. Bitte führen Sie 'go run runlin.go install' aus, um es zu installieren.")
		}
	case "performance":
		installPerformanceEnhancers()
		if isInstalled("proton") {
			runWithProton(appPath)
		} else if isInstalled("wine") {
			runWithWine(appPath)
		} else {
			fmt.Println("Kein unterstützter Executor gefunden. Bitte installieren Sie Proton oder Wine.")
		}
	default:
		fmt.Println("Ungültiger Befehl. Nutzung: go run runlin.go [proton | wine | install | performance] <Pfad zur .exe-Datei>")
	}
}

