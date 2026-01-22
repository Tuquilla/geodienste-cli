package helpers

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"runtime"
	"slices"
	"strings"

	"github.com/kglaus/stac-client/internal/stac/models"
)

func SaveAsset(assetObject models.AssetObject) error {
	dir, assetRootPath, err := OperatingSystem()
	if err != nil {
		return fmt.Errorf("error getting operating system: %v", err)
	}

	path := os.DirFS(dir)
	entries, _ := fs.Glob(path, "stac_downloads")
	if !slices.Contains(entries, "stac_downloads") {
		if err := os.Mkdir(assetRootPath, os.ModePerm); err != nil {
			return fmt.Errorf("error making dir: %v", err)
		}
	}

	resp, err := http.Get(assetObject.Href)
	if err != nil {
		return fmt.Errorf("error downloading asset: %v", err)
	}
	defer resp.Body.Close()

	assetUrlSplit := strings.Split(assetObject.Href, "/")
	assetTitle := assetUrlSplit[len(assetUrlSplit)-1]
	save, err := os.Create(assetRootPath + assetTitle)
	if err != nil {
		return fmt.Errorf("error saving downloaded asset at creation: %v", err)
	}
	defer save.Close()

	_, err = io.Copy(save, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving downloaded asset at copy: %v", err)
	}
	return nil
}

func OperatingSystem() (string, string, error) {
	var dir string
	var assetRootPath string
	switch runtime.GOOS {
	case "windows":
		dir = os.Getenv("USERPROFILE") + "\\Documents"
		assetRootPath = dir + "\\stac_downloads\\"
	case "linux":
		// TODO Implement linux paths
		dir = "~"
		return dir, assetRootPath, errors.New("error linux is not implemented")
	default:
		return dir, assetRootPath, errors.New("error unknown operating system")
	}
	return dir, assetRootPath, nil
}
