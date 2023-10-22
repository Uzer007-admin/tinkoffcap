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
type GametStatistics struct {
	ID   int64
	Name string
	//sopernik int64
	S     int64 //две штуки
	Place int32
	//Статус, Имя соперника, раунд игры, название турнира
	Goal   int32
	Game   int32
	Winner int32
	Loser  int32 //партии общие
	//партии выигранные и проигранные
	//количество выигранных и проигранных
	//ОБЩАЯ СТАТИСТИКА Преимущество 12.  Количество выиграшей Кол партии. Рейтинг. Количество побед и поражений.
}

type Score struct {
	Our int32
	In  int32
}
type TeamCreate struct {
	Name  string
	Date  int64
	Users []int64
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
