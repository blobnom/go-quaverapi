package quaverapi

import (
	"io/ioutil"
	"net/http"
)

var (
	BaseURL string = "https://api.quavergame.com/v1"
)

func ApiCall(path string) ([]byte, error) {
	req, err := http.Get(BaseURL + path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	return data, err
}
