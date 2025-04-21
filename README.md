# MP3 Converter

MP3 Converter is a simple and efficient tool for converting audio files to MP3 format. It supports various input formats and provides a user-friendly interface for seamless conversion.

## Features

- Convert audio files to MP3 format.
- Support for multiple input formats (e.g., WAV, FLAC, AAC).
- High-quality audio output.
- Batch conversion support.
- Lightweight and easy to use.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/DanielKoklev/youtube-mp3-converter.git
    cd youtube-mp3-converter
    ```

## Usage

1. Place your url links in the `urls.txt` file.
2. Run the converter:
    ```bash
    go run main.go
    ```
3. Converted MP3 files will be saved in the `output` directory.

## Requirements

- GoLang 1.24 or higher
- Required libraries ffmpeg, yt-dlp

