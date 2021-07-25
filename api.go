package quaverapi

import (
	"io/ioutil"
	"net/http"
	"fmt"
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

func main() {
	users, err := SearchUsers("lewdloli")
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%+v", users)
}