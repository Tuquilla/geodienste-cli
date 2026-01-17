package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/kglaus/stac-client/configs"
)

func Config(path string) configs.Config {
	configInput, err := os.Open(path)
	if err != nil {
		fmt.Println("Error open config json")
		fmt.Println(err)
	}

	body, err := io.ReadAll(configInput)
	if err != nil {
		fmt.Println("Error read config json")
	}

	var config configs.Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		fmt.Printf("Error unmarshalling json: %s", err)
	}
	return config
}
