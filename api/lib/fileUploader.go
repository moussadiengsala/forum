package lib

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func ImageUploader(w http.ResponseWriter, imageData []byte) (string, error) {
	uploadsDir := "./static/uploads/"
	_, errs := os.Stat(uploadsDir)
	if os.IsNotExist(errs) {
		os.MkdirAll(uploadsDir, os.ModePerm)
	}

	// You can use any filename generation logic you prefer here
	fileName := time.Now().Format("20060102150405") + ".jpg"
	filepath := filepath.Join(uploadsDir, fileName)
	f, err := os.Create(filepath)

	if err != nil {
		// Handle the error
		return "", err
	}

	defer f.Close()

	_, copyErr := f.Write(imageData)

	if copyErr != nil {
		// Handle the error
		return "", copyErr
	}

	return fileName, nil
}

// UploadImages takes an array of image data and returns an array of uploaded file names
func UploadImages(w http.ResponseWriter, images [][]byte) (string, error) {
	var filenames []string

	for _, imageData := range images {
		fileName, err := ImageUploader(w, imageData)
		if err != nil {
			// Handle the error
			return "", err
		}
		filenames = append(filenames, fileName)
	}

	imagesJSON, err := json.Marshal(filenames)
	if err != nil {
		return "", err
	}
	return string(imagesJSON), nil
}
