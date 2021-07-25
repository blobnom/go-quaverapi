package structs

type Leaderboard struct {
	Users []LeaderboardUser `json:"users"`
}

type LeaderboardUser struct {
	Id           int         `json:"id"`
	SteamId      string      `json:"steam_id"`
	Username     string      `json:"username"`
	Country      string      `json:"country"`
	Allowed      int         `json:"allowed"`
	Privileges   int         `json:"privileges"`
	UserGroups   int         `json:"usergroups"`
	AvatarUrl    string      `json:"avatar_url"`
	RegDate      string      `json:"time_registered"`
	LastActivity string      `json:"latest_activity"`
	Stats        RatingStats `json:"stats"`
}

type LeaderboardHitsUser struct {
	Id           int         `json:"id"`
	SteamId      string      `json:"steam_id"`
	Username     string      `json:"username"`
	Country      string      `json:"country"`
	Allowed      int         `json:"allowed"`
	Privileges   int         `json:"privileges"`
	UserGroups   int         `json:"usergroups"`
	AvatarUrl    string      `json:"avatar_url"`
	RegDate      string      `json:"time_registered"`
	LastActivity string      `json:"latest_activity"`
	Stats        RatingStats `json:"stats"`
}

type RatingStats struct {
	Rank        int     `json:"rank"`
	RankedScore int     `json:"ranked_score"`
	Accuracy    float64 `json:"overall_accuracy"`
	Rating      float64 `json:"overall_performance_rating"`
	Playcount   int     `json:"play_count"`
	MaxCombo    int     `json:"max_combo"`
}

type HitsStats struct {
	Rank       int `json:"rank"`
	TotalHits  int `json:"total_hits"`
	TotalPlays int `json:"total_ranked_plays"`
	TotalFails int `json:"total_ranked_failures"`
}
