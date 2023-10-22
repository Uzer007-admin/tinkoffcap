package model

type ConditionaltSearchItem struct {
	Name string
}
type TournamentSearchItem struct {
	ID     int64
	Name   string
	Status string
	Update int64
	Create int64
}
type ParamSearch struct {
	Field string
}
type FilterSearch struct {
	Field  string
	Filter []string
}
type TeamSearchItem struct {
	ID        int64
	Name      string
	Date      int64
	UserCount int
}
type TournamentStatistics struct {
	ID         int64
	Name       string
	DateCreate int64
	DateUpdate int64 //две штуки
	Place      int32
	Score      int32
	Goal       int32
	Game       int32
	Winner     int32
	Loser      int32 //партии общие
	//партии выигранные и проигранные
	//количество выигранных и проигранных
}
type TeamCreate struct {
	Name  string  `json:"name"`
	Date  int64   `json:"date"`
	Users []int64 `json:"users"`
}
type ResposeSuccess struct {
	Code   int
	Result interface{}
	Time   int64
}
type Auth struct {
	AccessToken string `json:"accessToken"`
	Id          int64  `json:"id"`
	Login       string `json:"login"`
}
type ResposeError struct {
	Code        int
	Description string
	Error       error
}

type GameScore struct {
	FirstTeamID     int64
	SecondTeamID    int64
	FirstTeamScore  int64
	SecondTeamScore int64
	Winner          int64
}

type CommandStatistic struct {
	Score       int64
	Matches     []int64
	Tournaments []int64
}

type MatchFinished struct {
	ID          int64 `json:"id"`
	ScoreFirst  int64 `json:"score_first"`
	ScoreSecond int64 `json:"score_second"`
	Winner      int64 `json:"winner"`
}
