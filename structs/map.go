package structs

type Mapsets struct {
	Status  int      `json:"status"`
	Mapsets []Mapset `json:"mapsets"`
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
	Bpm              int     `json:"bpm"`
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
