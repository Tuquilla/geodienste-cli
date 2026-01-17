package stac

import (
	"slices"
	"testing"

	"github.com/kglaus/stac-client/internal/stac/models"
)

func TestCreateCollections(t *testing.T) {
	collections, _ := createCollections([]byte(collectionsTestdata))
	if collections.Collections[0].Id != "klaeranlagen_mit_finanzkennzahlen" {
		t.Errorf("Collection Id was not correct")
	}

	index := slices.IndexFunc(collections.Collections[0].Links, func(links models.Link) bool {
		return links.Rel == "items"
	})

	if index == -1 {
		t.Errorf("No Link to items found")
	}
}

func TestCreateFeatureCollection(t *testing.T) {
	assetObjectKey := "csv_zip"
	featureCollection, _ := createFeatureCollection([]byte(featureCollectionDatasetFixture))

	firstAssets := featureCollection.Features[0].Assets
	_, assetKeyExists := firstAssets[assetObjectKey]

	if assetKeyExists != true {
		t.Errorf("AssetObject key %s does not exist", assetObjectKey)
	}
}
