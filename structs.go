package quaverapi

// Leaderboard structures
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

// Map structures
type Mapsets struct {
	Status int    `json:"status"`
	Mapset Mapset `json:"mapset"`
}

type UserMapsets struct {
	Status  int      `json:"status"`
	Mapsets []Mapset `json:"mapsets"`
}

type Maps struct {
	Status int `json:"status"`
	Map    Map `json:"map"`
}

type Mapset struct {
	Id               int    `json:"id"`
	Md5              string `json:"package_md5"`
	CreatorId        int    `json:"creator_id"`
	CreatorUsername  string `json:"creator_username"`
	CreatorAvatarUrl string `json:"creator_avatar_url"`
	Artist           string `json:"artist"`
	Title            string `json:"title"`
	Source           string `json:"source"`
	Tags             string `json:"tags"`
	Description      string `json:"description"`
	SubmitDate       string `json:"date_submitted"`
	LastUpdateDate   string `json:"date_last_updated"`
	Visible          int    `json:"visible"`
	RankedStatus     int    `json:"ranked_status"`
	Maps             []Map
}

type Map struct {
	Id               int     `json:"id"`
	MapsetId         int     `json:"mapset_id"`
	Md5              string  `json:"md5"`
	CreatorId        int     `json:"creator_id"`
	CreatorUsername  string  `json:"creator_username"`
	Mode             int     `json:"game_mode"`
	RankedStatus     int     `json:"ranked_status"`
	Artist           string  `json:"artist"`
	Title            string  `json:"title"`
	Source           string  `json:"source"`
	Tags             string  `json:"tags"`
	Description      string  `json:"description"`
	DifficultyName   string  `json:"difficulty_name"`
	Length           int     `json:"length"`
	Bpm              float64 `json:"bpm"`
	DifficultyRating float64 `json:"difficulty_rating"`
	ObjectsNormal    int     `json:"count_hitobject_normal"`
	ObjectsLong      int     `json:"count_hitobject_long"`
	Playcount        int     `json:"play_count"`
	Failcount        int     `json:"fail_count"`
	PendingMods      int     `json:"mods_pending"`
	AcceptedMods     int     `json:"mods_accepted"`
	DeniedMods       int     `json:"mods_denied"`
	IgnoredMods      int     `json:"mods_ignored"`
	OnlineOffset     int     `json:"online_offset"`
}

// Score structures
type Scores struct {
	Status int     `json:"status"`
	Scores []Score `json:"scores"`
}

type MapScores struct {
	Status int        `json:"status"`
	Scores []MapScore `json:"scores"`
}

type Score struct {
	ScoreId      int     `json:"id"`
	Time         string  `json:"time"`
	Mode         int     `json:"mode"`
	Mods         int     `json:"mods"`
	ModsString   string  `json:"mods_string"`
	Rating       float64 `json:"performance_rating"`
	PersonalBest bool    `json:"personal_best"`
	Score        int     `json:"total_score"`
	Accuracy     float64 `json:"accuracy"`
	Grade        string  `json:"grade"`
	MaxCombo     int     `json:"max_combo"`
	CountMarv    int     `json:"count_marv"`
	CountPerf    int     `json:"count_perf"`
	CountGreat   int     `json:"count_great"`
	CountGood    int     `json:"count_good"`
	CountOkay    int     `json:"count_okay"`
	CountMiss    int     `json:"count_miss"`
	ScrollSpeed  int     `json:"scroll_speed"`
	TourGameId   int     `json:"tournament_game_id"`
	Ratio        float64 `json:"ratio"`
	Map          Map     `json:"map"`
}

type MapScore struct {
	ScoreId      int      `json:"id"`
	MapMd5       string   `json:"map_md5"`
	Time         string   `json:"time"`
	Mode         int      `json:"mode"`
	Mods         int      `json:"mods"`
	ModsString   string   `json:"mods_string"`
	Rating       float64  `json:"performance_rating"`
	PersonalBest bool     `json:"personal_best"`
	Score        int      `json:"total_score"`
	Accuracy     float64  `json:"accuracy"`
	Grade        string   `json:"grade"`
	MaxCombo     int      `json:"max_combo"`
	CountMarv    int      `json:"count_marv"`
	CountPerf    int      `json:"count_perf"`
	CountGreat   int      `json:"count_great"`
	CountGood    int      `json:"count_good"`
	CountOkay    int      `json:"count_okay"`
	CountMiss    int      `json:"count_miss"`
	User         UserInfo `json:"user"`
}

type ReplayData struct {
	Status int      `json:"status"`
	Hits   []string `json:"hits"`
}

// User structures
type User struct {
	User UserData `json:"user"`
}

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
	Id           int         `json:"id"`
	SteamId      string      `json:"steam_id"`
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
	Id   int             `json:"id"`
	Type int             `json:"type"`
	Date string          `json:"timestamp"`
	Map  UserActivityMap `json:"map"`
}

type UserActivityMap struct {
	Id   int    `json:"id"`
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

type RankGraph struct {
	Status     int    `json:"status"`
	Statistics []Rank `json:"statistics"`
}

type Rank struct {
	Rank      int    `json:"rank"`
	Timestamp string `json:"timestamp"`
}

type Achievements struct {
	Status       int           `json:"status"`
	Achievements []Achievement `json:"achievements"`
}

type Achievement struct {
	Id           int    `json:"id"`
	SteamApiName string `json:"steam_api_name"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Difficulty   string `json:"difficulty"`
	Unlocked     bool   `json:"unlocked"`
}

// Server Stats structures
type Server struct {
	Status int         `json:"status"`
	Stats  ServerStats `json:"stats"`
}

type ServerStats struct {
	OnlineUsers  int64 `json:"total_online_users"`
	TotalUsers   int64 `json:"total_users"`
	TotalMapsets int64 `json:"total_mapsets"`
	TotalScores  int64 `json:"total_scores"`
}
