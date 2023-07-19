# LOLBOT
![golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white)

> This project was written in golang language and used whatsmeow library.

> This project support for Termux (proot-distro) & Linux.

## Features

* Downloader (youtube & tiktok)
* Maker (sticker)
* Tools (ssweb)
* And more.

## How to install
Download golang [in here](https://go.dev/doc/install)

### Installing dependencies
#### Termux
```bash
pkg install golang -y
pkg install proot-distro -y
pkg install git -y
pkg install libwebp
proot-distro install ubuntu
proot-distro login ubuntu
apt install gnupg -y
echo 'deb http://ftp.de.debian.org/debian bullseye main' > /etc/apt/sources.list.d/debian.list
apt-key adv --keyserver keyserver.ubuntu.com --recv-keys DCC9EFBF77E11517
apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 648ACFD622F3D138
apt-key adv --keyserver keyserver.ubuntu.com --recv-keys AA8E81B4331F7F50
apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 112695A0E562B32A
printf 'Package: *\nPin: release a=focal\nPin-Priority: 500\n\n\nPackage: *\nPin: origin "ftp.debian.org"\nPin-Priority: 300\n\n\nPackage: chromium*\nPin: origin "ftp.debian.org"\nPin-Priority: 700\n' >> /etc/apt/preferences.d/chromium.pref
apt update
apt install chromium
rm /etc/apt/sources.list.d/debian.list
rm /etc/apt/preferences.d/chromium.pref
apt update
```

#### Ubuntu
```bash
apt update
apt upgrade
sudo apt install webp
sudo apt install chromium-browser
sudo snap install golang
```

### Installing project

```bash
git clone https://github.com/ibrahKrep/lolbot
cd lolbot
go run .
```

## Use this bot
after installing and run, type `.menu` and send to bot.

this example sending message
### Sending Message

```
// sending text
simple.Send(jid, "Hello")

// reply message
simple.Reply(jid, "How are you")

// send image
simple.SendImage(jid types.JID, text string, source string, quoted bool)

simple.SendImage(jid, "Hello", "./tmp/random.png", true)

// send audio
simple.SendAudio(jid types.JID, source string, ptt bool, quoted bool)

// send video
simple.SendVideo(jid types.JID, text string, source string, quoted bool)

// send sticker
simple.SendSticker(jid types.JID, source string, quoted bool)

```
