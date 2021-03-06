package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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
)

const service = "dd-delta-prof"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var flagF = flag.Bool("version", false, "Print version and exit")
	flag.Parse()
	if *flagF {
		fmt.Printf("%s\n", version)
		return nil
	}

	jsonData, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		return err
	}

	err = profiler.Start(
		profiler.WithService(service),
		profiler.WithEnv("prod"),
		profiler.WithVersion(version),
		profiler.WithAgentlessUpload(),
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

	allocLoop(jsonData)

	return nil
}

func allocLoop(jsonData []byte) {
	for {
		start := time.Now()
		var dst interface{}
		if err := json.Unmarshal(jsonData, &dst); err != nil {
			panic(err)
		}
		time.Sleep(time.Since(start))
	}
}
