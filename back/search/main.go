package main

import (
	"fmt"
	"time"

	"tinkoff/handler"

	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
	"github.com/valyala/fasthttp"
)

// надо дописать чтение конфига тип
var db = reindexer.NewReindex("cproto://rx.web-gen.ru:6534/tinkoff")

func handleRequest(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	start := time.Now()
	switch string(ctx.Path()) {
	case "/team_search":
		//test(ctx)TeamSearch
		handler.TeamSearch(ctx, db, start)
	case "/conditional_search":
		handler.ConditionalSearch(ctx, db, start)
	case "/tournament_search":
		handler.TournamentSearch(ctx, db, start)
	default:
		//full_search(ctx)
	}
	// Отправляем ответ клиенту
}

// Обработка корневого маршрута "/"

func main() {
	if db.Status().Err != nil {
		panic(db.Status().Err)

	}
	server := fasthttp.Server{
		Handler: handleRequest, // Обработчик запросов
	}

	// Слушаем порт 8080 и обрабатываем запросы
	err := server.ListenAndServe(":3300")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
