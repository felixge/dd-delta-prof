package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "embed"

	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

var (
	//go:embed version.txt
	rawVersion string
	version    = strings.TrimSpace(rawVersion)
	//go:embed example.json
	exampleJSON []byte
)

const service = "dd-delta-prof"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	err := profiler.Start(
		profiler.WithService(service),
		profiler.WithEnv("prod"),
		profiler.WithVersion(version),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
			profiler.BlockProfile,
			profiler.MutexProfile,
			profiler.GoroutineProfile,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

	log.Printf("Started %s\n", service)

	allocLoop()

	return nil
}

func allocLoop() {
	for {
		start := time.Now()
		var dst interface{}
		if err := json.Unmarshal(exampleJSON, &dst); err != nil {
			panic(err)
		}
		time.Sleep(time.Since(start))
	}
}
