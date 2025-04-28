package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
			cmd := exec.Command("yt-dlp", "-f", "bestaudio", "-o", outputDir+"/%(title)s.%(ext)s", "--extract-audio", url)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if _, err := os.Stat("/usr/bin/yt-dlp"); os.IsNotExist(err) {
				fmt.Println("yt-dlp not found in /usr/bin/. Please ensure yt-dlp is installed and accessible.")
				return
			}

			if _, err := exec.LookPath("ffmpeg"); err != nil {
				fmt.Println("ffmpeg not found. Please ensure ffmpeg is installed and accessible.")
				return
			}

			fmt.Printf("Running command: %s\n", cmd.String())

			if err := cmd.Run(); err != nil {
				fmt.Println("Error processing URL", url, ":", err)
			} else {
				fmt.Println("Successfully processed URL:", url)

				files, err := filepath.Glob(outputDir + "/*.opus")
				if err != nil {
					fmt.Println("Error finding downloaded files:", err)
					return
				}

				for _, file := range files {
					outputFile := strings.TrimSuffix(file, filepath.Ext(file)) + ".mp3"
					convertCmd := exec.Command("ffmpeg", "-i", file, outputFile)
					convertCmd.Stdout = os.Stdout
					convertCmd.Stderr = os.Stderr

					fmt.Printf("Converting %s to %s\n", file, outputFile)

					if err := convertCmd.Run(); err != nil {
						fmt.Println("Error converting file", file, ":", err)
					} else {
						fmt.Println("Successfully converted", file, "to", outputFile)

						if err := os.Remove(file); err != nil {
							fmt.Println("Error removing original file", file, ":", err)
						}
					}
				}
			}
		}
	} else {
		fmt.Println("urls.txt not found.")
	}

	fmt.Println("Download and conversion complete!")
}
