FROM golang:1.24-alpine 

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

RUN apk add --no-cache yt-dlp ffmpeg

CMD ["./main"]