package main

import (
	"Hackthon2023/internal/handler"
	"fmt"
	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
	"github.com/valyala/fasthttp"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type YML struct {
	Scheme   string
	Hostname string
	Port     string
	Path     string
}

/*func handleRequest(ctx *fasthttp.RequestCtx) {
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
}*/

// Обработка корневого маршрута "/"

func main() {
	yfile, err := os.ReadFile("config/config.yml")
	if err != nil {
		fmt.Println("Ошибка чтения файла")
	}

	data := make(map[string]YML)
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatal(err)
	}

	var db = reindexer.NewReindex(fmt.Sprintf("%v://%v:%v/%v",
		data["db"].Scheme, data["db"].Hostname, data["db"].Port, data["db"].Path))

	if db.Status().Err != nil {
		panic(db.Status().Err)

	}

	h := handler.NewHandler(db)

	server := fasthttp.Server{
		Handler: h.InitRoutes, // Обработчик запросов
	}

	// Слушаем порт 8080 и обрабатываем запросы
	err = server.ListenAndServe(data["http"].Port)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
