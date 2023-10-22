package api_db

import (
	"tinkoff/internal/model"

	"github.com/restream/reindexer/v3"
)

type SearchIn struct {
	SearchDB
}
type TeamIn struct {
	TeamDB
}
type TournamentIn struct {
	TournamentDB
}
type SearchDB interface {
	ConditionalSearchDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.ConditionaltSearchItem, error)
	TournamentSearchDB(obj model.FilterSearch, db *reindexer.Reindexer) ([]model.TournamentSearchItem, error)
	TeamSearchDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.TeamSearchItem, error)
}
type TeamDB interface {
	TeamIDSearchDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.TeamSearchItem, error)
	UserTeamDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.TeamSearchItem, error)
	TeamCreateNew(obj model.TeamCreate) error
}
type TournamentDB interface {
	ConditionalSearchDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.ConditionaltSearchItem, error)
	TournamentLastDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.TournamentStatistics, error)
	TeamSearchDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.TeamSearchItem, error)
}
