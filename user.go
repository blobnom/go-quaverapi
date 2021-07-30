package quaverapi

import (
	"encoding/json"
	"fmt"
)

func GetUser(id int) (*UserData, error) {
	rawData, err := ApiCall(fmt.Sprintf("/users/full/%d", id))
	if err != nil {
		return nil, err
	}

	user := User{}
	err = json.Unmarshal(rawData, &user)
	if err != nil {
		return nil, err
	}

	return &user.User, nil
}

func SearchUsers(username string) (*[]UserInfo, error) {
	rawData, err := ApiCall(fmt.Sprintf("/users/search/%s", username))
	if err != nil {
		return nil, err
	}

	users := Users{}
	err = json.Unmarshal(rawData, &users)
	if err != nil {
		return nil, err
	}

	return &users.Users, nil
}

func GetUserRankGraph(id int, mode int) (*RankGraph, error) {
	rawData, err := ApiCall(fmt.Sprintf("/users/graph/rank?id=%d&mode=%d", id, mode))
	if err != nil {
		return nil, err
	}

	rankgraph := RankGraph{}
	err = json.Unmarshal(rawData, &rankgraph)
	if err != nil {
		return nil, err
	}

	return &rankgraph, nil
}

func GetUserAchievements(id int) (*Achievements, error) {
	rawData, err := ApiCall(fmt.Sprintf("/users/%d/achievements", id))
	if err != nil {
		return nil, err
	}

	achievements := Achievements{}
	err = json.Unmarshal(rawData, &achievements)
	if err != nil {
		return nil, err
	}

	return &achievements, nil
}

func GetUserMapsets(id int) (*UserMapsets, error) {
	rawData, err := ApiCall(fmt.Sprintf("/users/mapsets/%d", id))
	if err != nil {
		return nil, err
	}

	mapsets := UserMapsets{}
	err = json.Unmarshal(rawData, &mapsets)
	if err != nil {
		return nil, err
	}

	return &mapsets, nil
}
