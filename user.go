package quaverapi

import (
	"encoding/json"
	"fmt"

	"github.com/elainaaa/go-quaverapi/structs"
)

func GetUser(id int) (*structs.User, error) {
	return nil, nil
}

func SearchUsers(username string) (*structs.Users, error) {
	rawData, err := ApiCall(fmt.Sprintf("/users/search/%s", username))
	if err != nil {
		return nil, err
	}

	users := structs.Users{}
	err = json.Unmarshal(rawData, &users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
