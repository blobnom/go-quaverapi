package structs

// The Response for GetUser function
type User struct {
	User UserData `json:"user"`
}

// The Response for SearchUsers function
type Users struct {
	Users []UserInfo `json:"users"`
}

type UserData struct {
	Info     UserInfo       `json:"info"`
	Activity []UserActivity `json:"activity_feed"`
	Keys4    UserMode       `json:"keys4"`
	Keys7    UserMode       `json:"keys7"`
}

type UserInfo struct {
	ID           int         `json:"id"`
	SteamID      string      `json:"steam_id"`
	RegDate      string      `json:"time_registered"`
	Allowed      int         `json:"allowed"`
	Privileges   int         `json:"privileges"`
	UserGroups   int         `json:"usergroups"`
	MuteEnd      string      `json:"mute_endtime"`
	LastActivity string      `json:"latest_activity"`
	Username     string      `json:"username"`
	Country      string      `json:"country"`
	AvatarUrl    string      `json:"avatar_url"`
	Socials      UserSocials `json:"information"`
	Online       bool        `json:"online"`
}

type UserSocials struct {
	Discord string `json:"discord"`
	Twitter string `json:"twitter"`
	Twitch  string `json:"twitch"`
	Youtube string `json:"youtube"`
}

type UserActivity struct {
	ID   int             `json:"id"`
	Type int             `json:"type"`
	Date string          `json:"timestamp"`
	Map  UserActivityMap `json:"map"`
}

type UserActivityMap struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserMode struct {
	GlobalRank  int       `json:"globalRank"`
	CountryRank int       `json:"countryRank"`
	MultiRank   int       `json:"multiplayerWinRank"`
	Stats       UserStats `json:"stats"`
}

type UserStats struct {
	TotalScore     int     `json:"total_score"`
	RankedScore    int     `json:"ranked_score"`
	Accuracy       float64 `json:"overall_accuracy"`
	Rating         float64 `json:"overall_performance_rating"`
	Playcount      int     `json:"play_count"`
	Failcount      int     `json:"fail_count"`
	MaxCombo       int     `json:"max_combo"`
	ReplaysWatched int     `json:"replays_watched"`
	TotalMarv      int     `json:"total_marv"`
	TotalPerf      int     `json:"total_perf"`
	TotalGreat     int     `json:"total_great"`
	TotalGood      int     `json:"total_good"`
	TotalOkay      int     `json:"total_okay"`
	TotalMiss      int     `json:"total_miss"`
	TotalPauses    int     `json:"total_pauses"`
	MultiWins      int     `json:"multiplayer_wins"`
	MultiLosses    int     `json:"multiplayer_losses"`
	MultiTies      int     `json:"multiplayer_ties"`
}
