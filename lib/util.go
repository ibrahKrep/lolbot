package lib

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os/exec"
)

func SaveMedia(name string, data []byte) bool {
	err := ioutil.WriteFile(name, data, 0666)
	if err != nil {
		return false
	}
	return true
}

func GetHttp(url string) ([]byte, *http.Response) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(errors.New("Cannot get response"))
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(errors.New("Cannot read response"))
	}
	return data, res
}

func SaveMediaFromUrl(url, source string) bool {
	data, _ := GetHttp(url)
	ok := SaveMedia(source, data)
	return ok
}

func RandStr(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvfxyz")
	rn := make([]rune, length)

	for i := range rn {
		rn[i] = letters[rand.Intn(len(letters))]
	}

	return string(rn)
}

func Exec(name string, args []string) string {
	cmd := exec.Command(name, args...)
	output, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(output)
}
