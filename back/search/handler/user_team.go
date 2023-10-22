package handler

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
	"tinkoff/internal/api_db"
	"tinkoff/internal/model"
	"tinkoff/service"

	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
	"github.com/valyala/fasthttp"
)

func UserTeam(ctx *fasthttp.RequestCtx, db *reindexer.Reindexer, start time.Time) {
	var UserData model.Auth
	if string(ctx.Method()) == "POST" {
		// Проверяем заголовок Content-Type
		contentType := string(ctx.Request.Header.ContentType())
		if contentType != "application/json" {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			ctx.SetBodyString("Invalid Content-Type")
			return
		}

		// Отправляем успешный ответ
		var err error
		UserData, err = service.Client(ctx.Request.Body(), UserData)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.SetBodyString("Only POST method is allowed")
		return
	}

	var Inter api_db.TeamIn
	find, err := Inter.TeamDB.UserTeamDB(model.ParamSearch{Field: (strconv.FormatInt(UserData.Id, 10))}, db)
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
