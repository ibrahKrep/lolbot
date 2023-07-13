package lib

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"google.golang.org/protobuf/proto"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

type Simple struct {
	Wcli *whatsmeow.Client
	Msg  *events.Message
}

func NewSimple(cli *whatsmeow.Client, msg *events.Message) *Simple {
	return &Simple{
		Wcli: cli,
		Msg:  msg,
	}
}

func (simple *Simple) Send(jid types.JID, text string) {
	msg := &waProto.Message{
		Conversation: proto.String(text),
	}
	res, err := simple.Wcli.SendMessage(context.Background(), jid, msg)
	if err != nil {
		fmt.Println(res)
	}
}

func (simple *Simple) Reply(jid types.JID, text string) {
	msg := &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Text: proto.String(text),
			ContextInfo: &waProto.ContextInfo{
				StanzaId:      proto.String(simple.Msg.Info.ID),
				Participant:   proto.String(simple.Msg.Info.Sender.String()),
				QuotedMessage: simple.Msg.Message,
			},
		},
	}
	res, err := simple.Wcli.SendMessage(context.Background(), jid, msg)
	if err != nil {
		fmt.Println(res)
	}
}

func (simple *Simple) SendImage(jid types.JID, text string, source string, quoted bool) {
	data, err := os.ReadFile(source)
	if err != nil {
		return
	} else {
		fmt.Println(data)
	}

	uploaded, err := simple.Wcli.Upload(context.Background(), data, whatsmeow.MediaImage)
	if err != nil {
		return
	} else {
		fmt.Println(uploaded)
	}

	if quoted {
		msg := &waProto.Message{
			ImageMessage: &waProto.ImageMessage{
				Caption:       proto.String(text),
				Url:           proto.String(uploaded.URL),
				DirectPath:    proto.String(uploaded.DirectPath),
				MediaKey:      uploaded.MediaKey,
				Mimetype:      proto.String(http.DetectContentType(data)),
				FileEncSha256: uploaded.FileEncSHA256,
				FileSha256:    uploaded.FileSHA256,
				FileLength:    proto.Uint64(uint64(len(data))),
				ContextInfo: &waProto.ContextInfo{
					StanzaId:      proto.String(simple.Msg.Info.ID),
					Participant:   proto.String(simple.Msg.Info.Sender.String()),
					QuotedMessage: simple.Msg.Message,
				},
			},
		}
		res, err := simple.Wcli.SendMessage(context.Background(), jid, msg)
		if err != nil {
			fmt.Println("err sending img")
		} else {
			fmt.Println(res)
		}
	} else {
		msg := &waProto.Message{
			ImageMessage: &waProto.ImageMessage{
				Caption:       proto.String(text),
				Url:           proto.String(uploaded.URL),
				DirectPath:    proto.String(uploaded.DirectPath),
				MediaKey:      uploaded.MediaKey,
				Mimetype:      proto.String(http.DetectContentType(data)),
				FileEncSha256: uploaded.FileEncSHA256,
				FileSha256:    uploaded.FileSHA256,
				FileLength:    proto.Uint64(uint64(len(data))),
			},
		}
		res, err := simple.Wcli.SendMessage(context.Background(), jid, msg)
		if err != nil {
			fmt.Println("err sending img")
		} else {
			fmt.Println(res)
		}
	}
}

func (simple *Simple) SendVideo(jid types.JID, text string, source string, quoted bool) {
	data, err := os.ReadFile(source)
	if err != nil {
		return
	}
	uploaded, err := simple.Wcli.Upload(context.Background(), data, whatsmeow.MediaVideo)
	if err != nil {
		return
	}
	if quoted {
		msg := &waProto.Message{
			VideoMessage: &waProto.VideoMessage{
				Caption:       proto.String(text),
				Url:           proto.String(uploaded.URL),
				DirectPath:    proto.String(uploaded.DirectPath),
				MediaKey:      uploaded.MediaKey,
				Mimetype:      proto.String(http.DetectContentType(data)),
				FileEncSha256: uploaded.FileEncSHA256,
				FileSha256:    uploaded.FileSHA256,
				FileLength:    proto.Uint64(uint64(len(data))),
				ContextInfo: &waProto.ContextInfo{
					StanzaId:      proto.String(simple.Msg.Info.ID),
					Participant:   proto.String(simple.Msg.Info.Sender.String()),
					QuotedMessage: simple.Msg.Message,
				},
			},
		}
		res, err := simple.Wcli.SendMessage(context.Background(), jid, msg)
		if err != nil {
			fmt.Println("error sending video")
		}
		fmt.Println(res)
	} else {
		msg := &waProto.Message{
			VideoMessage: &waProto.VideoMessage{
				Caption:       proto.String(text),
				Url:           proto.String(uploaded.URL),
				DirectPath:    proto.String(uploaded.DirectPath),
				MediaKey:      uploaded.MediaKey,
				Mimetype:      proto.String(http.DetectContentType(data)),
				FileEncSha256: uploaded.FileEncSHA256,
				FileSha256:    uploaded.FileSHA256,
				FileLength:    proto.Uint64(uint64(len(data))),
				ContextInfo: &waProto.ContextInfo{
					StanzaId:      proto.String(simple.Msg.Info.ID),
					Participant:   proto.String(simple.Msg.Info.Sender.String()),
					QuotedMessage: simple.Msg.Message,
				},
			},
		}
		res, err := simple.Wcli.SendMessage(context.Background(), jid, msg)
		if err != nil {
			fmt.Println("error sending video")
		}
		fmt.Println(res)

	}
}

func (simple *Simple) GetCmd() string {
	conversation := simple.Msg.Message.GetConversation()
	image := simple.Msg.Message.GetImageMessage().GetCaption()
	video := simple.Msg.Message.GetVideoMessage().GetCaption()
	extended := simple.Msg.Message.GetExtendedTextMessage().GetText()
	document := simple.Msg.Message.GetDocumentMessage().GetCaption()
	var cmd string

	if conversation != "" {
		cmd = conversation
	}
	if image != "" {
		cmd = image
	}
	if video != "" {
		cmd = video
	}
	if extended != "" {
		cmd = extended
	}
	if document != "" {
		cmd = document
	}

	return cmd

}
