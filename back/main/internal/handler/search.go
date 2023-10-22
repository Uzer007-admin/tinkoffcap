package handler

import (
	"Hackthon2023/internal/api_db"
	"Hackthon2023/internal/model"
	"encoding/json"
	"github.com/restream/reindexer/v3"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

func TeamIDSearch(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	queryArgs := ctx.QueryArgs()
	paramValue := model.ParamSearch{Field: string(queryArgs.Peek("team_id"))}
	find, err := api_db.TeamIDSearchDB(paramValue, db)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
		response := model.ResposeError{
			Code:        503,
			Description: "Сервис взаимодействия с базой данных недоступен",
			Error:       err,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Response.SetBody(jsonData)
		return
	}
	response := model.ResposeSuccess{
		Code:   200,
		Result: find,
		Time:   time.Since(start).Nanoseconds(),
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Response.SetBody(jsonData)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)

}

func TeamCreateNew(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	var teamNew model.TeamCreate
	if string(ctx.Method()) == "POST" {
		// Проверяем заголовок Content-Type
		contentType := string(ctx.Request.Header.ContentType())
		if contentType != "application/json" {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			ctx.SetBodyString("Invalid Content-Type")
			return
		}

		err := json.Unmarshal(ctx.Request.Body(), &teamNew)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
			ctx.SetBodyString("Failed to unmarshal request body")
			return
		}

	} else {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.SetBodyString("Only POST method is allowed")
		return
	}

	err := api_db.TeamCreateNewDB(teamNew, db)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
		response := model.ResposeError{
			Code:        503,
			Description: "Сервис взаимодействия с базой данных недоступен",
			Error:       err,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Response.SetBody(jsonData)
		return
	}
	response := model.ResposeSuccess{
		Code:   200,
		Result: nil,
		Time:   time.Since(start).Nanoseconds(),
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Response.SetBody(jsonData)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
}

func TeamSearch(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	queryArgs := ctx.QueryArgs()
	paramValue := model.ParamSearch{Field: string(queryArgs.Peek("team_name"))}
	find, err := api_db.TeamSearchDB(paramValue, db)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
		response := model.ResposeError{
			Code:        503,
			Description: "Сервис взаимодействия с базой данных недоступен",
			Error:       err,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Response.SetBody(jsonData)
		return
	}
	response := model.ResposeSuccess{
		Code:   200,
		Result: find,
		Time:   time.Since(start).Nanoseconds(),
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Response.SetBody(jsonData)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)

}

func ConditionalSearch(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	queryArgs := ctx.QueryArgs()
	paramValue := model.ParamSearch{Field: string(queryArgs.Peek("conditional_name"))}
	find, err := api_db.ConditionalSearchDB(paramValue, db)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
		response := model.ResposeError{
			Code:        503,
			Description: "Сервис взаимодействия с базой данных недоступен",
			Error:       err,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Response.SetBody(jsonData)
		return
	}
	response := model.ResposeSuccess{
		Code:   200,
		Result: find,
		Time:   time.Since(start).Nanoseconds(),
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Response.SetBody(jsonData)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)

}

func TournamentSearch(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	queryArgs := ctx.QueryArgs()
	filterOpen := string(queryArgs.Peek("opened"))
	filterActiv := string(queryArgs.Peek("active"))
	filterFinish := string(queryArgs.Peek("finished"))
	paramValue := model.FilterSearch{
		Field:  string(queryArgs.Peek("tournament_name")),
		Filter: []string{filterOpen, filterActiv, filterFinish},
	}
	find, err := api_db.TournamentSearchDB(paramValue, db)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
		response := model.ResposeError{
			Code:        503,
			Description: "Сервис взаимодействия с базой данных недоступен",
			Error:       err,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Response.SetBody(jsonData)
		return
	}
	response := model.ResposeSuccess{
		Code:   200,
		Result: find,
		Time:   time.Since(start).Nanoseconds(),
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Response.SetBody(jsonData)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)

}

func UserTeam(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	queryArgs := ctx.QueryArgs()
	paramValue := model.ParamSearch{Field: string(queryArgs.Peek("user_id"))}
	find, err := api_db.UserTeamDB(paramValue, db)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
		response := model.ResposeError{
			Code:        503,
			Description: "Сервис взаимодействия с базой данных недоступен",
			Error:       err,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Response.SetBody(jsonData)
		return
	}
	response := model.ResposeSuccess{
		Code:   200,
		Result: find,
		Time:   time.Since(start).Nanoseconds(),
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Response.SetBody(jsonData)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
}

func GameTeams(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	queryArgs := ctx.QueryArgs()
	paramValue := model.ParamSearch{Field: string(queryArgs.Peek("game_id"))}
	find, err := api_db.GameTeamsDB(paramValue, db)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
		response := model.ResposeError{
			Code:        503,
			Description: "Сервис взаимодействия с базой данных недоступен",
			Error:       err,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Response.SetBody(jsonData)
		return
	}
	response := model.ResposeSuccess{
		Code:   200,
		Result: find,
		Time:   time.Since(start).Nanoseconds(),
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Response.SetBody(jsonData)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
}

func TeamStatistic(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	queryArgs := ctx.QueryArgs()
	paramValue := model.ParamSearch{Field: string(queryArgs.Peek("team_id"))}
	find, err := api_db.TeamStatisticDB(paramValue, db)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
		response := model.ResposeError{
			Code:        503,
			Description: "Сервис взаимодействия с базой данных недоступен",
			Error:       err,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Response.SetBody(jsonData)
		return
	}
	response := model.ResposeSuccess{
		Code:   200,
		Result: find,
		Time:   time.Since(start).Nanoseconds(),
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Response.SetBody(jsonData)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
}

func TournamentStatusActive(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	queryArgs := ctx.QueryArgs()
	paramValue := model.ParamSearch{Field: string(queryArgs.Peek("tournament_id"))}
	err := api_db.TournamentStatusActiveDB(paramValue, db)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
		response := model.ResposeError{
			Code:        503,
			Description: "Сервис взаимодействия с базой данных недоступен",
			Error:       err,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Response.SetBody(jsonData)
		return
	}
	response := model.ResposeSuccess{
		Code:   200,
		Result: nil,
		Time:   time.Since(start).Nanoseconds(),
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Response.SetBody(jsonData)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
}

func GameConfirm(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	var result model.MatchFinished
	if string(ctx.Method()) == "POST" {
		// Проверяем заголовок Content-Type
		contentType := string(ctx.Request.Header.ContentType())
		if contentType != "application/json" {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			ctx.SetBodyString("Invalid Content-Type")
			return
		}

		err := json.Unmarshal(ctx.Request.Body(), &result)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
			ctx.SetBodyString("Failed to unmarshal request body")
			return
		}

	} else {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.SetBodyString("Only POST method is allowed")
		return
	}

	err := api_db.GameStatusDB(result, db)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
		response := model.ResposeError{
			Code:        503,
			Description: "Сервис взаимодействия с базой данных недоступен",
			Error:       err,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Response.SetBody(jsonData)
		return
	}
	response := model.ResposeSuccess{
		Code:   200,
		Result: nil,
		Time:   time.Since(start).Nanoseconds(),
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Response.SetBody(jsonData)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
}

func TournamentLast(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	queryArgs := ctx.QueryArgs()
	paramValue := model.ParamSearch{Field: string(queryArgs.Peek("team_id"))}
	find, err := api_db.TournamentLastDB(paramValue, db)
	if err != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
		response := model.ResposeError{
			Code:        503,
			Description: "Сервис взаимодействия с базой данных недоступен",
			Error:       err,
		}
		jsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		ctx.Response.SetBody(jsonData)
		return
	}
	response := model.ResposeSuccess{
		Code:   200,
		Result: find,
		Time:   time.Since(start).Nanoseconds(),
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	ctx.Response.SetBody(jsonData)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)

}
