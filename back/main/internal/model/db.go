package model

import "github.com/restream/reindexer/v3"

type TournamentItem struct {
	ID        int64         `json:"id" reindex:"id,hash,pk"`
	Name      string        `json:"name" reindex:"name,text"`
	Date      DateType      `json:"date"`
	Status    int64         `json:"status" reindex:"status,hash"`
	Places    []int64       `json:"places" reindex:"places,tree,sparse"`
	Teams     []int64       `json:"teams" reindex:"teams,hash,sparse"`
	Limit     int32         `json:"limit" reindex:"limit,tree"`
	StatusArr []*StatusItem `reindex:"status_arr,,joined"`
	TeamsArr  []*TeamItem   `reindex:"teams_arr,,joined"`
	PlacesArr []*TeamItem   `reindex:"places_arr,,joined"`
}

type UserItem struct {
	ID            int64           `json:"id" reindex:"id,hash,pk"`
	Name          string          `json:"name" reindex:"name,hash"`
	Surname       string          `json:"surname" reindex:"surname,hash"`
	Patronymic    string          `json:"patronymic" reindex:"patronymic,hash"`
	Email         string          `json:"email" reindex:"email,hash"`
	Authenticated int32           `json:"authenticated" reindex:"authenticated,hash"`
	Login         string          `json:"login" reindex:"login,hash"`
	Password      string          `json:"password" reindex:"password,hash"`
	IP            string          `json:"ip" reindex:"ip,hash"`
	Birthday      int64           `json:"birthday" reindex:"birthday,tree"`
	City          string          `json:"city" reindex:"city,hash"`
	Position      reindexer.Point `json:"position" reindex:"position,rtree"`
	Date          DateType        `json:"date"`
}

type TeamItem struct {
	ID       int64       `json:"id" reindex:"id,hash,pk"`
	Name     string      `json:"name" reindex:"name,text"`
	Date     int64       `json:"date" reindex:"date,tree"`
	Rating   float64     `json:"rating" reindex:"rating,tree"`
	Users    []int64     `json:"users" reindex:"users,hash,sparse"`
	UsersArr []*UserItem `reindex:"users_arr,,joined"`
}

type GameItem struct {
	ID            int64             `json:"id" reindex:"id,hash,pk"`
	LocalID       int64             `json:"localID" reindex:"localID,hash"`
	Team          TeamType          `json:"team"`
	Score         ScoreType         `json:"score"`
	Stage         int64             `json:"stage" reindex:"stage,hash"`
	ParentFirst   int64             `json:"parentFirst" reindex:"parentFirst,hash"`
	ParentSecond  int64             `json:"parentSecond" reindex:"parentSecond,hash"`
	Tournament    int64             `json:"tournament" reindex:"tournament,hash"`
	Date          int64             `json:"date" reindex:"date,tree"`
	Winner        int64             `json:"winner" reindex:"winner,hash"`
	Status        int64             `json:"status" reindex:"status,hash"`
	TournamentArr []*TournamentItem `reindex:"tournament_arr,,joined"`
	WinnerArr     []*TeamItem       `reindex:"winner_arr,,joined"`
	StatusArr     []*StatusItem     `reindex:"status_arr,,joined"`
}

type StatusItem struct {
	ID   int64  `json:"id" reindex:"id,hash,pk"`
	Name string `json:"name" reindex:"name,hash"`
}

type DateType struct {
	Update int64 `json:"update" reindex:"update,tree"`
	Create int64 `json:"create" reindex:"create,tree"`
}

type TeamType struct {
	First     int64       `json:"first" reindex:"team_first,hash"`
	Second    int64       `json:"second" reindex:"team_second,hash"`
	FirstArr  []*TeamItem `reindex:"first_arr,,joined"`
	SecondArr []*TeamItem `reindex:"second_arr,,joined"`
}

type ScoreType struct {
	First  int64 `json:"first" reindex:"score_first,tree"`
	Second int64 `json:"second" reindex:"score_second,tree"`
}
