package message

import (
	"fmt"
	"strings"

	"github.com/ibrahKrep/lolbot/lib"
	"github.com/ibrahKrep/lolbot/utils"

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
	pushname := simple.Msg.Info.PushName
	commands := strings.Split(simple.GetCmd(), " ")
	cmd := strings.Join(commands[:1], "")
	args := strings.Join(commands[1:], " ")

	fmt.Printf("From: %v\nSender: %v\nMessage: %v\n", from, sender, simple.GetCmd())
	//fmt.Println(simple.Msg.Message)

	switch cmd {
	case prefix + "menu":
		menu := fmt.Sprintf(`
*Halo,* %v.
*Silakan gunakan fitur di bawah ini.*

*『DOWNLOADER』*

➠ .youtube <url>
    ➥ Untuk mendownload video youtube.
➠ .youtubeaudio <url>
    ➥ Untuk mendownload audio video youtube.
➠ .tiktok <url>
    ➥ Untuk mendownload video tiktok.

*『MAKER』*

➠ .sticker
    ➥ Untuk membuat stiker.
    kirim gambar dengan caption .sticker

*『OTHER』*

➠ .ssweb <url>
    ➥ Screenshot web page.

*『HTTP METHOD』*

➠ .get <url>
    ➥ Untuk mendapatkan response data dari website.
`, pushname)
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
		case "video/mp4":
			randm := lib.RandStr(4)
			lib.SaveMedia("./tmp/"+randm+".mp4", body)
			simple.SendVideo(from, "", "./tmp/"+randm+".mp4", true)

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
		res := utils.GetStream(args, randm, "video")
		msg := fmt.Sprintf(`
*TITLE:* %v
*CHANNEL:* %v
*DURATION:* %v
`, res.Title, res.Author, res.Duration.String())
		simple.SendVideo(from, msg, randm, true)
	case prefix + "youtubeaudio":
		randm := "./tmp/" + lib.RandStr(4) + ".mp4"
		res := utils.GetStream(args, randm, "audio")
		radm := "./tmp/" + lib.RandStr(4) + ".png"
		lib.SaveMediaFromUrl(res.Thumbnails[0].URL, radm)
		msg := fmt.Sprintf(`
*TITLE:* %v
*CHANNEL:* %v
*DURATION:* %v
`, res.Title, res.Author, res.Duration.String())
		simple.SendImage(from, msg, radm, true)
		simple.SendAudio(from, randm, false, true)
	case prefix + "sticker":
		data, _ := simple.Wcli.Download(simple.Msg.Message.GetImageMessage())
		random := "./tmp/" + lib.RandStr(4) + ".png"
		random2 := "./tmp/" + lib.RandStr(4) + ".webp"
		lib.SaveMedia(random, data)
		lib.Exec("cwebp", []string{random, "-o", random2})
		simple.SendSticker(from, random2, true)
	case prefix + "tiktok":
		res := utils.Tiktok(args)
		simple.SendVideo(from, "*DONE*", res, true)
	case prefix + "ssweb":
		res := utils.SsWeb(args)
		simple.SendImage(from, "", res, true)
	}
}
