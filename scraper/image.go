package scraper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// DownloadFile downloads locally the file stored at the given URL. It returns the name of the downloaded file with the
// correct file extension.
func DownloadFile(imageURL, fileName string) (string, error) {
	//Get the response bytes from the url
	resp, err := http.Get(imageURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("received non 200 resp code")
	}

	// Detect mime
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		fileName = fmt.Sprintf("%s.jpg", fileName)
	case "image/png":
		fileName = fmt.Sprintf("%s.png", fileName)
	default:
		log.Printf("no mime detected for image at URL %s\n", imageURL)
	}

	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = ioutil.WriteFile(fileName, bytes, 0666)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
