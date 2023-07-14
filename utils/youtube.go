package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func GetStream(url, filename, mediatype string) {
	if mediatype == "video" {
		video_id, err := youtube.ExtractVideoID(url)
		if err != nil {
			panic(err)
		}
		client := youtube.Client{}
		video, err := client.GetVideo(video_id)
		if err != nil {
			panic(err)
		}
		formats := video.Formats.WithAudioChannels()
		fmt.Println(formats)

		stream, _, err := client.GetStream(video, &formats[0])
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = io.Copy(file, stream)
		if err != nil {
			panic(err)
		}
		fmt.Println("SUCCESS GET VIDEO FROM YOUTUBE")
	} else if mediatype == "audio" {
		video_id, err := youtube.ExtractVideoID(url)
		if err != nil {
			panic(err)
		}
		client := youtube.Client{}
		video, err := client.GetVideo(video_id)
		if err != nil {
			panic(err)
		}
		formats := video.Formats.WithAudioChannels().Type("audio/mp4").Quality("tiny")
		fmt.Println(formats)
		stream, _, err := client.GetStream(video, &formats[1])
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = io.Copy(file, stream)
		if err != nil {
			panic(err)
		}
		fmt.Println("SUCCESS GET AUDIO FROM YOUTUBE")
	}
}
