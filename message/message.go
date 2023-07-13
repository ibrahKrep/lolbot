package message

import (
	"fmt"
	"strings"

	"lol/lib"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func Message(cli *whatsmeow.Client, msg *events.Message) {

	const (
		prefix string = "."
	)

	simple := lib.NewSimple(cli, msg)
	from := simple.Msg.Info.Chat
	commands := strings.Split(simple.GetCmd(), " ")
	cmd := strings.Join(commands[:1], "")
	args := strings.Join(commands[1:], " ")

	fmt.Println(simple.GetCmd())
	fmt.Println(simple.Msg.Message)

	switch cmd {
	case prefix + "menu":
		simple.Reply(from, "gak ada menu")
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
	}
}
