package stac

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kglaus/stac-client/internal/stac/models"
)

func GetCollections(baseUrl string) (models.Collections, error) {
	resp, err := http.Get(baseUrl + "/collections")
	if err != nil {
		return models.Collections{}, fmt.Errorf("error invalid url: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Collections{}, fmt.Errorf("error reading response body: %v", err)
	}

	collections, err := createCollections(body)
	if err != nil {
		return models.Collections{}, fmt.Errorf("error creating collections: %v", err)
	}
	return collections, nil
}

func GetItems(itemUrl string) (models.FeatureCollection, error) {
	resp, err := http.Get(itemUrl)
	if err != nil {
		return models.FeatureCollection{}, fmt.Errorf("error getting item url: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.FeatureCollection{}, fmt.Errorf("error reading response body: %v", err)
	}

	featureCollection, err := createFeatureCollection(body)
	if err != nil {
		return models.FeatureCollection{}, fmt.Errorf("error creating feature collections: %v", err)
	}
	return featureCollection, nil
}

func createCollections(body []byte) (models.Collections, error) {
	var collections models.Collections
	err := json.Unmarshal(body, &collections)
	if err != nil {
		fmt.Printf("Error unmarshalling json: %s", err)
		return models.Collections{}, fmt.Errorf("error unmarshalling collection json: %v", err)
	}
	return collections, nil
}

func createFeatureCollection(body []byte) (models.FeatureCollection, error) {
	var featureCollections models.FeatureCollection
	err := json.Unmarshal(body, &featureCollections)
	if err != nil {
		return models.FeatureCollection{}, fmt.Errorf("error unmarshalling feature collection json: %v", err)
	}
	return featureCollections, nil
}
