package main

import (
	"fmt"
	"os/exec"
	"strings"
  "time"
)

// Funktion zur Überprüfung der Installation von Wine und Proton
func checkInstallation(executor string) bool {
	_, err := exec.LookPath(executor)
	return err == nil
}

// Funktion zum Testen der Performance-Komponenten
func checkPerformanceComponents() map[string]bool {
	requiredComponents := []string{"dxvk", "vkd3d"}
	componentStatus := make(map[string]bool)

	for _, comp := range requiredComponents {
		componentStatus[comp] = checkInstallation(comp)
	}

	return componentStatus
}

// Funktion zur Erfassung von Systemmetriken
func measureSystemMetrics(executor string) (map[string]string, error) {
	metrics := make(map[string]string)

	// Beispiel für die Erfassung von CPU- und Speicherverbrauch
	cpuUsage, err := exec.Command("sh", "-c", "top -bn1 | grep 'Cpu(s)'").Output()
	if err != nil {
		return nil, err
	}
	metrics["CPU Usage"] = strings.TrimSpace(string(cpuUsage))

	memUsage, err := exec.Command("sh", "-c", "free -m").Output()
	if err != nil {
		return nil, err
	}
	metrics["Memory Usage"] = strings.TrimSpace(string(memUsage))

	// Metriken zur Laufzeitanalyse
	startTime := time.Now()
	cmd := exec.Command(executor, "--version") // Dummy-Befehl zur Messung des Startaufwands
	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	duration := time.Since(startTime)
	metrics["Start Time"] = duration.String()

	return metrics, nil
}

// Funktion zur Vergleichsanalyse
func compareMetrics(metrics1, metrics2 map[string]string) {
	fmt.Println("\nVergleich der Leistungsmetriken:")

	for key, val1 := range metrics1 {
		if val2, exists := metrics2[key]; exists {
			fmt.Printf("%s:\n  Wine: %s\n  Proton: %s\n", key, val1, val2)
		}
	}
}

func main() {
	executors := []string{"wine", "proton"}

	fmt.Println("Überprüfung der Installation von Wine und Proton...")

	for _, execName := range executors {
		if checkInstallation(execName) {
			fmt.Printf("%s ist installiert.\n", execName)
		} else {
			fmt.Printf("%s ist nicht installiert.\n", execName)
		}
	}

	fmt.Println("\nÜberprüfung der Performance-Komponenten...")
	performanceComponents := checkPerformanceComponents()
	for comp, status := range performanceComponents {
		if status {
			fmt.Printf("%s ist installiert.\n", comp)
		} else {
			fmt.Printf("%s ist nicht installiert.\n", comp)
		}
	}

	fmt.Println("\nErfassung der Systemmetriken für Wine...")
	wineMetrics, err := measureSystemMetrics("wine")
	if err != nil {
		fmt.Printf("Fehler beim Erfassen der Systemmetriken für Wine: %v\n", err)
		return
	}

	fmt.Println("\nErfassung der Systemmetriken für Proton...")
	protonMetrics, err := measureSystemMetrics("proton")
	if err != nil {
		fmt.Printf("Fehler beim Erfassen der Systemmetriken für Proton: %v\n", err)
		return
	}

	// Vergleich der Metriken
	compareMetrics(wineMetrics, protonMetrics)

	fmt.Println("\nBenchmark-Test abgeschlossen.")
}

