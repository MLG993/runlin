package main

import (
	"fmt"
	"os/exec"
	"strings"
  "time"
)

func checkInstallation(executor string) bool {
	_, err := exec.LookPath(executor)
	return err == nil
}

func checkPerformanceComponents() map[string]bool {
	requiredComponents := []string{"dxvk", "vkd3d"}
	componentStatus := make(map[string]bool)

	for _, comp := range requiredComponents {
		componentStatus[comp] = checkInstallation(comp)
	}

	return componentStatus
}

func measureSystemMetrics(executor string) (map[string]string, error) {
	metrics := make(map[string]string)

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

	startTime := time.Now()
	cmd := exec.Command(executor, "--version") 
	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	duration := time.Since(startTime)
	metrics["Start Time"] = duration.String()

	return metrics, nil
}

func compareMetrics(metrics1, metrics2 map[string]string) {
	fmt.Println("\nComparing powermetrics:")

	for key, val1 := range metrics1 {
		if val2, exists := metrics2[key]; exists {
			fmt.Printf("%s:\n  Wine: %s\n  Proton: %s\n", key, val1, val2)
		}
	}
}

func main() {
	executors := []string{"wine", "proton"}

	fmt.Println("Checking Proton and Wine installation...")

	for _, execName := range executors {
		if checkInstallation(execName) {
			fmt.Printf("%s is not installed.\n", execName)
		} else {
			fmt.Printf("%s is not installed.\n", execName)
		}
	}

	fmt.Println("\nChecking the performance components...")
	performanceComponents := checkPerformanceComponents()
	for comp, status := range performanceComponents {
		if status {
			fmt.Printf("%s is not installed.\n", comp)
		} else {
			fmt.Printf("%s is not installed.\n", comp)
		}
	}

	fmt.Println("\nGetting your system-stats for wine...")
	wineMetrics, err := measureSystemMetrics("wine")
	if err != nil {
		fmt.Printf("Error while getting your pc specs: %v\n", err)
		return
	}

	fmt.Println("\nGetting your system-stats for proton...")
	protonMetrics, err := measureSystemMetrics("proton")
	if err != nil {
		fmt.Printf("Error while getting your pc specs: %v\n", err)
		return
	}

	compareMetrics(wineMetrics, protonMetrics)

	fmt.Println("\nBenchmark is done.")
}

