package stac

import (
	"testing"
)

func TestTransfromJsonToInternal(t *testing.T) {
	collections := transfromJsonToInternalDataStructure([]byte(collections_testdata))
	if collections.Collections[0].Id != "klaeranlagen_mit_finanzkennzahlen" {
		t.Errorf("Collection Id was not correct")
	}
}
