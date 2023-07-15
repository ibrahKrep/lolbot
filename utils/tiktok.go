package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"lol/lib"
)

func Tiktok(urltiktok string) string {
	var sourcefile string
	if urltiktok == "" {
		fmt.Println(errors.New("[TIKTOK] zero url value"))
	} else {
		var tiktok = struct {
			Token string `json:"token"`
			Id    string `json:"id"`
		}{}
		resp, err := http.PostForm(
			"https://api.tikmate.app/api/lookup",
			url.Values{"url": {urltiktok}},
		)
		if err != nil {
			fmt.Println(errors.New("[TIKTOK] Cannot get response."))
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &tiktok)
		if err != nil {
			panic(err)
		}

		// json marshal
		//jsm, _ := json.Marshal(tiktok)
		//fmt.Sprintf("%+v", tiktok)

		urldownload := fmt.Sprintf("https://tikmate.app/download/%v/%v.mp4", tiktok.Token, tiktok.Id)
		sourcefile = fmt.Sprintf("./tmp/%v.mp4", lib.RandStr(4))
		lib.SaveMediaFromUrl(urldownload, sourcefile)
	}
	return sourcefile
}
