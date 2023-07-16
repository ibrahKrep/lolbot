package utils

import (
	"fmt"

	"lol/lib"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func SsWeb(url string) (random string) {
	random = fmt.Sprintf("./tmp/%v.png", lib.RandStr(4))

	u := launcher.New().
		Bin("/usr/bin/chromium").
		NoSandbox(true).
		MustLaunch()

	page := rod.New().ControlURL(u).MustConnect().MustPage(url)
	page.MustWindowFullscreen()
	page.MustWaitStable().MustScreenshot(random)

	return
}
