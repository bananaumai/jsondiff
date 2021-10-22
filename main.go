package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/go-cmp/cmp"
)

func main()  {
	args := os.Args
	if len(args) < 3 {
		log.Fatalf("specify two files")
	}

	f1 := args[1]
	f2 := args[2]

	j1, err := load(f1)
	if err != nil {
		log.Fatal(err)
	}

	j2, err := load(f2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", cmp.Diff(j1, j2))
}

func load(filePath string) (interface{}, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", filePath, err)
	}
	defer func() { _ = f.Close() }()

	var j interface{}
	dec := json.NewDecoder(f)
	if err := dec.Decode(&j); err != nil {
		return nil, fmt.Errorf("failed to decode %s: %w", filePath, err)
	}

	return j, nil
}
