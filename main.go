package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"lol/message"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"google.golang.org/protobuf/proto"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var cli *whatsmeow.Client

func main() {
	dbLog := waLog.Stdout("Database", "ERROR", true)
	container, err := sqlstore.New("sqlite3", "file:store.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	// Store sesi
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("cli", "ERROR", true)
	cli = whatsmeow.NewClient(deviceStore, clientLog)
	cli.AddEventHandler(eventHandler)

	if cli.Store.ID == nil {
		qrChan, _ := cli.GetQRChannel(context.Background())
		err = cli.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("QR CODE: ")
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("LOGIN EVENT: ", evt.Event)
			}
		}
	} else {
		err = cli.Connect()
		if err != nil {
			panic(err)
		}
	}
	c := make(chan os.Signal, 1)
	go signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	cli.Disconnect()
}

func init() {
	store.DeviceProps.PlatformType = waProto.DeviceProps_ANDROID_PHONE.Enum()
	store.DeviceProps.Os = proto.String("SLAMET BOT")
}

func eventHandler(evt interface{}) {
	switch msg := evt.(type) {
	case *events.Message:
		message.Message(cli, msg)
	}
}
