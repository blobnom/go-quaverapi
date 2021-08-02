package quaverapi

import (
	"encoding/json"
)

func GetServerStats() (ServerStats, error) {
	rawData, err := ApiCall("/stats")
	if err != nil {
		return Server{}, err
	}

	server := Server{}
	err = json.Unmarshal(rawData, &server)
	if err != nil {
		return ServerStats{}, err
	}

	return server.Stats, nil
}
