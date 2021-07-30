package quaverapi

import (
	"encoding/json"
	"fmt"
)

func GetUserBest(id int, mode int) (*[]Score, error) {
	rawData, err := ApiCall(fmt.Sprintf("/users/scores/best?id=%d&mode=%d", id, mode))
	if err != nil {
		return nil, err
	}

	scores := Scores{}
	err = json.Unmarshal(rawData, &scores)
	if err != nil {
		return nil, err
	}

	return &scores.Scores, nil
}

func GetUserRecent(id int, mode int) (*[]Score, error) {
	rawData, err := ApiCall(fmt.Sprintf("/users/scores/recent?id=%d&mode=%d", id, mode))
	if err != nil {
		return nil, err
	}

	scores := Scores{}
	err = json.Unmarshal(rawData, &scores)
	if err != nil {
		return nil, err
	}

	return &scores.Scores, nil
}

func GetUserFirstPlace(id int, mode int) (*[]Score, error) {
	rawData, err := ApiCall(fmt.Sprintf("/users/scores/firstplace?id=%d&mode=%d", id, mode))
	if err != nil {
		return nil, err
	}

	scores := Scores{}
	err = json.Unmarshal(rawData, &scores)
	if err != nil {
		return nil, err
	}

	return &scores.Scores, nil
}

func GetMapScores(id int) (*[]MapScore, error) {
	rawData, err := ApiCall(fmt.Sprintf("/scores/map/%d", id))
	if err != nil {
		return nil, err
	}

	scores := MapScores{}
	err = json.Unmarshal(rawData, &scores)
	if err != nil {
		return nil, err
	}

	return &scores.Scores, nil
}
