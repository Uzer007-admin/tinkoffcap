package api_db

import (
	"log"
	"strconv"
	"tinkoff/internal/model"

	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
)

func indexOf(word int64, data []int64) int32 {
	for k, v := range data {
		if word == v {
			return int32(k)
		}
	}
	return -1
}
func TournamentLastDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.TournamentStatistics, error) {
	TeamID, err := strconv.ParseInt(obj.Field, 10, 64) // второй аргумент - это основание системы счисления, третий - размер типа в битах
	if err != nil {
		log.Fatal(err)
	}
	err = db.OpenNamespace("tournament", reindexer.DefaultNamespaceOptions(), model.TournamentItem{})
	if err != nil {
		panic(err)
	}
	err = db.OpenNamespace("game", reindexer.DefaultNamespaceOptions(), model.GameItem{})
	if err != nil {
		panic(err)
	}

	query := db.Query("tournament").Where("status", reindexer.EQ, 3).Where("teams", reindexer.SET, TeamID)
	var response []model.TournamentStatistics
	iterator := query.Exec()
	if iterator.Error() != nil {
		panic(iterator.Error())
	}
	for iterator.Next() {
		elem := iterator.Object().(*model.TournamentItem)
		que := db.Query("game").Where("tournament", reindexer.EQ, elem.ID).Where("team", reindexer.SET, TeamID)
		que.AggregateSum("winner")
		que.ReqTotal()
		iterGame := que.Exec()
		if iterGame.Error() != nil {
			panic(iterGame.Error())
		}
		aggSumRes := iterGame.AggResults()[0]
		coutGame := int32(iterGame.TotalCount())

		Goal, err := GoalS(db, TeamID, elem.ID)
		if err != nil {
			panic(err)
		}
		response = append(response, model.TournamentStatistics{
			ID:         elem.ID,
			Name:       elem.Name,
			DateCreate: elem.Date.Create,
			DateUpdate: elem.Date.Update,
			Place:      indexOf(elem.ID, elem.Places),
			Goal:       Goal,
			Game:       coutGame,
			Winner:     int32(*aggSumRes.Value),
			Loser:      int32(coutGame - int32(*aggSumRes.Value)),
		})
	}
	return response, err
}
func GoalS(db *reindexer.Reindexer, TeamID int64, TournamentID int64) (int32, error) {
	que := db.Query("game").Where("tournament", reindexer.EQ, TournamentID).Where("team_first", reindexer.EQ, TeamID)
	que.AggregateSum("score_first")
	que.ReqTotal()
	iterGame := que.Exec()
	if iterGame.Error() != nil {
		panic(iterGame.Error())
	}
	aggSumResFirst := iterGame.AggResults()[0]
	//coutGameFirst := int32(iterGame.TotalCount())

	que = db.Query("game").Where("tournament", reindexer.EQ, TournamentID).Where("team_second", reindexer.EQ, TeamID)
	que.AggregateSum("score_second")
	que.ReqTotal()
	iterGame = que.Exec()
	var sumGoal float64
	if iterGame.Error() != nil {
		log.Fatal(iterGame.Error())
		sumGoal = 0
		panic(iterGame.Error())
	}
	aggSumResSecond := iterGame.AggResults()[0]
	//coutGameSecond := int32(iterGame.TotalCount())
	sumGoal = *aggSumResFirst.Value + *aggSumResSecond.Value
	return int32(sumGoal), iterGame.Error()
}

func TournamentOneStatDB(TeamID int64, TournamentID int64, db *reindexer.Reindexer) ([]model.TournamentStatistics, error) {

	err := db.OpenNamespace("tournament", reindexer.DefaultNamespaceOptions(), model.TournamentItem{})
	if err != nil {
		panic(err)
	}
	err = db.OpenNamespace("game", reindexer.DefaultNamespaceOptions(), model.GameItem{})
	if err != nil {
		panic(err)
	}

	query := db.Query("tournament").Where("status", reindexer.EQ, 3).Where("teams", reindexer.SET, TeamID).Where("id", reindexer.EQ, TournamentID)
	var response []model.TournamentStatistics
	iterator := query.Exec()
	if iterator.Error() != nil {
		panic(iterator.Error())
	}
	for iterator.Next() {
		elem := iterator.Object().(*model.TournamentItem)
		que := db.Query("game").Where("tournament", reindexer.EQ, elem.ID).Where("team", reindexer.SET, TeamID)
		que.AggregateSum("winner")
		que.ReqTotal()
		iterGame := que.Exec()
		if iterGame.Error() != nil {
			panic(iterGame.Error())
		}
		aggSumRes := iterGame.AggResults()[0]
		coutGame := int32(iterGame.TotalCount())

		Goal, err := GoalS(db, TeamID, elem.ID)
		if err != nil {
			panic(err)
		}
		response = append(response, model.TournamentStatistics{
			ID:         elem.ID,
			Name:       elem.Name,
			DateCreate: elem.Date.Create,
			DateUpdate: elem.Date.Update,
			Place:      indexOf(elem.ID, elem.Places),
			Goal:       Goal,
			Game:       coutGame,
			Winner:     int32(*aggSumRes.Value),
			Loser:      int32(coutGame - int32(*aggSumRes.Value)),
		})
	}
	return response, err
}
func TournamentGameOneStatDB(TeamID int64, TournamentID int64, db *reindexer.Reindexer) ([]model.TournamentStatistics, error) {

	err := db.OpenNamespace("tournament", reindexer.DefaultNamespaceOptions(), model.TournamentItem{})
	if err != nil {
		panic(err)
	}
	err = db.OpenNamespace("game", reindexer.DefaultNamespaceOptions(), model.GameItem{})
	if err != nil {
		panic(err)
	}

	query := db.Query("game").Where("status", reindexer.EQ, 3).Where("teams", reindexer.SET, TeamID).Where("id", reindexer.EQ, TournamentID)
	var response []model.TournamentStatistics
	iterator := query.Exec()
	if iterator.Error() != nil {
		panic(iterator.Error())
	}
	for iterator.Next() {
		elem := iterator.Object().(*model.TournamentItem)
		que := db.Query("game").Where("tournament", reindexer.EQ, elem.ID).Where("team", reindexer.SET, TeamID)
		que.AggregateSum("winner")
		que.ReqTotal()
		iterGame := que.Exec()
		if iterGame.Error() != nil {
			panic(iterGame.Error())
		}
		aggSumRes := iterGame.AggResults()[0]
		coutGame := int32(iterGame.TotalCount())

		Goal, err := GoalS(db, TeamID, elem.ID)
		if err != nil {
			panic(err)
		}
		response = append(response, model.TournamentStatistics{
			ID:         elem.ID,
			Name:       elem.Name,
			DateCreate: elem.Date.Create,
			DateUpdate: elem.Date.Update,
			Place:      indexOf(elem.ID, elem.Places),
			Goal:       Goal,
			Game:       coutGame,
			Winner:     int32(*aggSumRes.Value),
			Loser:      int32(coutGame - int32(*aggSumRes.Value)),
		})
	}
	return response, err
}
