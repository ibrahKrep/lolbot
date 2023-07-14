package message

import (
	"fmt"
	"strings"

	"lol/lib"
	"lol/utils"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func Message(cli *whatsmeow.Client, msg *events.Message) {
	var (
		owner  string = "62895605887712@s.whatsapp.net"
		prefix string = "."
	)

	simple := lib.NewSimple(cli, msg)
	from := simple.Msg.Info.Chat
	sender := simple.Msg.Info.Sender
	commands := strings.Split(simple.GetCmd(), " ")
	cmd := strings.Join(commands[:1], "")
	args := strings.Join(commands[1:], " ")

	fmt.Printf("From: %v\nSender: %v\n", from, sender)
	fmt.Println(simple.GetCmd())
	fmt.Println(simple.Msg.Message)

	switch cmd {
	case prefix + "menu":
		menu := `
*『DOWNLOADER』*

➠ .youtube <url>
  ➥ Untuk mendownload video youtube.
➠ youtubeaudio <url>
  ➥ Untuk mendownload audio video youtube.
`
		simple.Reply(from, strings.Trim(menu, "\n"))
	case prefix + "get":
		body, res := lib.GetHttp(args)
		switch res.Header.Get("Content-Type") {
		case "image/png":
			randm := lib.RandStr(4)
			lib.SaveMedia("./tmp/"+randm+".png", body)
			simple.SendImage(from, "", "./tmp/"+randm+".png", true)
		case "image/jpeg":
			randm := lib.RandStr(4)
			lib.SaveMedia("./tmp/"+randm+".jpeg", body)
			simple.SendImage(from, "", "./tmp/"+randm+".jpeg", true)
		default:
			simple.Reply(from, string(res.Header.Get("Content-Type")))
			simple.Reply(from, string(body))
		}
	case "EX":
		if own := owner == sender.String(); own {
			output := lib.Exec(strings.Join(commands[1:2], ""), commands[2:])
			simple.Reply(from, strings.Trim(output, "\n"))
		}
	case prefix + "youtube":
		randm := "./tmp/" + lib.RandStr(4) + ".mp4"
		utils.GetStream(args, randm, "video")
		simple.SendVideo(from, "*DONE*", randm, true)
	case prefix + "youtubeaudio":
		randm := "./tmp/" + lib.RandStr(4) + ".mp4"
		utils.GetStream(args, randm, "audio")
		simple.SendAudio(from, randm, false, true)
	}
}
