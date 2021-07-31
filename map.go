package quaverapi

import (
	"encoding/json"
	"fmt"
)

func GetMapset(id int) (Mapsets, error) {
	rawData, err := ApiCall(fmt.Sprintf("/mapsets/%d", id))
	if err != nil {
		return Mapsets{}, err
	}

	mapset := Mapsets{}
	err = json.Unmarshal(rawData, &mapset)
	if err != nil {
		return Mapsets{}, err
	}

	return mapset, nil
}

func GetMap(id int) (Maps, error) {
	rawData, err := ApiCall(fmt.Sprintf("/maps/%d", id))
	if err != nil {
		return Maps{}, err
	}

	maps := Maps{}
	err = json.Unmarshal(rawData, &maps)
	if err != nil {
		return Maps{}, err
	}

	return maps, nil
}
