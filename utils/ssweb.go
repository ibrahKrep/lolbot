package utils

import (
	"fmt"

	"github.com/ibrahKrep/lolbot/lib"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func SsWeb(url string) (random string) {
	random = fmt.Sprintf("./tmp/%v.png", lib.RandStr(4))

	u := launcher.New().
		Bin("/usr/bin/chromium-browser"). //ubuntu proot-distro "/usr/bin/chromium" as root nosandbox true
		NoSandbox(false).
		MustLaunch()

	page := rod.New().ControlURL(u).MustConnect().MustPage(url)
	page.MustWindowFullscreen()
	page.MustWaitStable().MustScreenshot(random)

	return
}
