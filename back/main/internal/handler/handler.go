package handler

import (
	"github.com/restream/reindexer/v3"
	"github.com/valyala/fasthttp"
	"time"
)

type Handler struct {
	db *reindexer.Reindexer
}

func NewHandler(db *reindexer.Reindexer) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) InitRoutes(ctx *fasthttp.RequestCtx) {
	start := time.Now()
	switch string(ctx.Path()) {
	case "/team/id":
		TeamIDSearch(ctx, h.db, start)
	case "/team/name":
		TeamSearch(ctx, h.db, start)
	case "/tournament/conditional":
		ConditionalSearch(ctx, h.db, start)
	case "/tournament/filter":
		TournamentSearch(ctx, h.db, start)
	case "/user/team":
		UserTeam(ctx, h.db, start)
	case "/team/new":
		TeamCreateNew(ctx, h.db, start)
	case "/team/tournament":
		TournamentLast(ctx, h.db, start)
	case "/team/game":
		GameTeams(ctx, h.db, start)
	case "/team/statistic":
		TeamStatistic(ctx, h.db, start)
	case "/tournament/active":
		TournamentStatusActive(ctx, h.db, start)
	case "/game/complete":
		GameConfirm(ctx, h.db, start)
	}
}
