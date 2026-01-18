package configs

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed config.json
var configJSON []byte

func ReadConfig() Config {
	var config Config

	err := json.Unmarshal(configJSON, &config)
	if err != nil {
		fmt.Printf("Error unmarshalling json: %s\n", err)
	}

	return config
}
