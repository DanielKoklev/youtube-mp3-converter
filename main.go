package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	if _, err := os.Stat("urls.txt"); err == nil {
		file, err := os.Open("urls.txt")
		if err != nil {
			fmt.Println("Error opening urls.txt:", err)
			return
		}
		defer file.Close()

		var urls []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading urls.txt:", err)
			return
		}

		outputDir := "./output"
        if err := os.MkdirAll(outputDir, 0755); err != nil {
            fmt.Println("Error creating output directory:", err)
            return
        }

		for _, url := range urls {
			fmt.Println("Downloading and converting:", url)
			cmd := exec.Command("yt-dlp", "-f", "bestaudio", "-o", outputDir+"/%(title)s.%(ext)s", "--extract-audio", "--audio-format", "mp3", url)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			cmd.Path = "/usr/bin/yt-dlp"
			if err := cmd.Run(); err != nil {
				fmt.Println("Error processing URL", url, ":", err)
			} else {
				fmt.Println("Successfully processed URL:", url)
			}
		}
	}

	fmt.Println("Download and conversion complete!")
}