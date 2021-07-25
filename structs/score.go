package structs

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
	MaxCombo     int     `json:"count_combo"`
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
	MaxCombo     int      `json:"count_combo"`
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
