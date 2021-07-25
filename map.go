package quaverapi

import (
	"encoding/json"
	"fmt"

	"github.com/elainaaa/go-quaverapi/structs"
)

func GetMapset(id int) (*structs.Mapset, error) {
	rawData, err := ApiCall(fmt.Sprintf("/mapsets/%d", id))
	if err != nil {
		return nil, err
	}

	mapset := structs.Mapset{}
	err = json.Unmarshal(rawData, &mapset)
	if err != nil {
		return nil, err
	}

	return &mapset, nil
}

func GetMap(id int) (*structs.Map, error) {
	rawData, err := ApiCall(fmt.Sprintf("/maps/%d", id))
	if err != nil {
		return nil, err
	}

	maps := structs.Map{}
	err = json.Unmarshal(rawData, &maps)
	if err != nil {
		return nil, err
	}

	return &maps, nil
}
