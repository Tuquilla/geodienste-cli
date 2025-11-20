package stac

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kglaus/geodienste-cli/pkg/stac/models"
)

var BASE_URL string = "https://geodienste.ch/stac/collections"

func TestStac() {
	resp, err := http.Get(BASE_URL)
	if err != nil {
		fmt.Println("Error calling service")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading Response Body")
	}

	collections := transfromJsonToInternalDataStructure(body)
	for _, collection := range collections.Collections {
		fmt.Println(collection.Id)
	}
}

func transfromJsonToInternalDataStructure(body []byte) models.Collections {
	var collections models.Collections
	json.Unmarshal(body, &collections)
	return collections
}
