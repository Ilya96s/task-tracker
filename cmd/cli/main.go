package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Info struct {
	GoVersion string `json:"go_version"`
	Module    string `json:"module"`
	Version   string `json:"version"`

	CommitHash string `json:"commit_hash"`
	Revision   string `json:"revision"`
	Modified   string `json:"modified"`

	StartedAt string `json:"started_at"`
	Uptime    string `json:"uptime"`
}

func main() {
	addr := flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()

	url := *addr + "/debug/info"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("request failed: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status: %s\n", resp.Status)
		os.Exit(1)
	}

	var info Info

	if err = json.NewDecoder(resp.Body).Decode(&info); err != nil {
		fmt.Printf("decoding failed: %v\n", err)
		os.Exit(1)
	}

	printInfo(info)
}

func printInfo(info Info) {
	fmt.Println("=== SERVICE INFO ===")
	fmt.Printf("Module      : %s\n", info.Module)
	fmt.Printf("Version     : %s\n", info.Version)
	fmt.Printf("Go Version  : %s\n", info.GoVersion)
	fmt.Printf("Commit Hash : %s\n", info.CommitHash)
	fmt.Printf("Modified    : %s\n", info.Modified)
	fmt.Printf("Started At  : %s\n", info.StartedAt)
	fmt.Printf("Uptime      : %s\n", info.Uptime)
}
