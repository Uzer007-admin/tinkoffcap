package handler

import (
	"encoding/json"
	"log"
	"time"
	"tinkoff/internal/api_db"
	"tinkoff/internal/model"

	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
	"github.com/valyala/fasthttp"
)

func TeamIDSearch(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	queryArgs := ctx.QueryArgs()
	paramValue := model.ParamSearch{Field: string((queryArgs.Peek("team_id")))}
	var Inter api_db.SearchIn
	find, err := Inter.SearchDB.TeamSearchDB(paramValue, db)
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
