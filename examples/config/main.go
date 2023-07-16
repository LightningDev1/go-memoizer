package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/LightningDev1/go-memoizer"
)

// ProgramConfig is the config for the program.
type ProgramConfig map[string]any

// GetConfig gets the config for the program.
func GetConfig() (config *ProgramConfig, err error) {
	content, err := os.ReadFile("./config.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	fmt.Println("Loaded config!")

	return config, nil
}

// ConfigMemoizer is the memoizer for the config.
var ConfigMemoizer = memoizer.NewMemoizer(GetConfig, time.Second*5)

func doSomething(i int) {
	config, err := ConfigMemoizer.Get()
	if err != nil {
		panic(err)
	}

	// Do something with the config.

	fmt.Println(i, config)
}

func main() {
	for i := 0; i < 50; i++ {
		doSomething(i)

		time.Sleep(time.Second)
	}
}
