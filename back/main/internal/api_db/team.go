package api_db

import (
	"Hackthon2023/internal/model"
	"encoding/json"
	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
	"log"
	"math"
	"sort"
	"strconv"
)

type ratingForMap struct {
	ID   int64
	Rate float64
}

func indexOf(word int64, data []int64) int32 {
	for k, v := range data {
		if word == v {
			return int32(k)
		}
	}
	return -1
}

func TeamIDSearchDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.TeamSearchItem, error) {
	err := db.OpenNamespace("team", reindexer.DefaultNamespaceOptions(), model.TeamItem{})
	if err != nil {
		panic(err)
	}
	query := db.Query("team").Where("id", reindexer.EQ, obj.Field).Limit(1)
	var response []model.TeamSearchItem
	iterator := query.Exec()
	if iterator.Error() != nil {
		panic(err)
	}
	for iterator.Next() {
		elem := iterator.Object().(*model.TeamItem)
		response = append(response, model.TeamSearchItem{
			ID:        elem.ID,
			Name:      elem.Name,
			Date:      elem.Date,
			UserCount: len(elem.Users),
		})
	}
	return response, err
}
func UserTeamDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.TeamSearchItem, error) {
	err := db.OpenNamespace("team", reindexer.DefaultNamespaceOptions(), model.TeamItem{})
	if err != nil {
		panic(err)
	}
	query := db.Query("team").Where("users", reindexer.SET, obj.Field)
	var response []model.TeamSearchItem
	iterator := query.Exec()
	if iterator.Error() != nil {
		panic(err)
	}
	for iterator.Next() {
		elem := iterator.Object().(*model.TeamItem)
		response = append(response, model.TeamSearchItem{
			ID:        elem.ID,
			Name:      elem.Name,
			Date:      elem.Date,
			UserCount: len(elem.Users),
		})
	}
	return response, err
}
func TeamCreateNewDB(obj model.TeamCreate, db *reindexer.Reindexer) error {
	err := db.OpenNamespace("team", reindexer.DefaultNamespaceOptions(), model.TeamItem{})
	if err != nil {
		return err
	}
	err = db.Upsert("team", &model.TeamItem{
		ID:       0,
		Name:     obj.Name,
		Date:     0,
		Users:    obj.Users,
		UsersArr: []*model.UserItem{},
	}, "date=now(NSEC)", "id=serial()")
	return err
}

func GameTeamsDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.GameScore, error) {
	var response []model.GameScore

	err := db.OpenNamespace("game", reindexer.DefaultNamespaceOptions(), model.GameItem{})
	if err != nil {
		return response, err
	}
	query := db.Query("game").Where("id", reindexer.EQ, obj.Field).Limit(1)
	iterator := query.Exec()
	if iterator.Error() != nil {
		return response, iterator.Error()
	}
	for iterator.Next() {
		elem := iterator.Object().(*model.GameItem)
		response = append(response, model.GameScore{
			FirstTeamID:     elem.Team.First,
			SecondTeamID:    elem.Team.Second,
			FirstTeamScore:  elem.Score.First,
			SecondTeamScore: elem.Score.Second,
			Winner:          elem.Winner,
		})
	}
	return response, err
}

func TournamentStatusActiveDB(obj model.ParamSearch, db *reindexer.Reindexer) error {
	var response []model.GameItem
	err := db.OpenNamespace("tournaments", reindexer.DefaultNamespaceOptions(), model.TournamentItem{})
	if err != nil {
		return err
	}
	err = db.OpenNamespace("status", reindexer.DefaultNamespaceOptions(), model.StatusItem{})
	if err != nil {
		return err
	}
	err = db.OpenNamespace("team", reindexer.DefaultNamespaceOptions(), model.TeamItem{})
	if err != nil {
		return err
	}
	err = db.OpenNamespace("game", reindexer.DefaultNamespaceOptions(), model.GameItem{})
	if err != nil {
		return err
	}
	queryTeam := db.Query("team")
	query := db.Query("tournaments").Where("id", reindexer.EQ, obj.Field).Join(queryTeam, "teams_arr").On("teams", reindexer.SET, "id")
	iterator := query.Exec()
	if iterator.Error() != nil {
		return iterator.Error()
	}
	for iterator.Next() {
		elem := iterator.Object().(*model.TournamentItem)
		rating := make(map[int64]float64)
		degree := math.Log2(float64(len(elem.Teams)))
		degreeInt := 0.0
		if int(degree*10)%10 < 4 {
			degreeInt = math.Floor(degree)
		} else {
			degreeInt = math.Ceil(degree)
		}
		firstStageTeams := int(math.Pow(2.0, degreeInt))

		for i := 0; i < len(elem.TeamsArr); i++ {
			rating[elem.TeamsArr[i].ID] = elem.TeamsArr[i].Rating
		}

		ratingForMapArr := make([]ratingForMap, 0, len(rating))
		for key, value := range rating {
			ratingForMapArr = append(ratingForMapArr, ratingForMap{Rate: value, ID: key})
		}
		sort.Slice(ratingForMapArr, func(i, j int) bool {
			return ratingForMapArr[i].Rate > ratingForMapArr[j].Rate
		})

		var sortedTeams []int64

		for _, elem := range ratingForMapArr {
			sortedTeams = append(sortedTeams, elem.ID)
		}

		if len(rating) != firstStageTeams {
			for i := 0; i < firstStageTeams*2-len(rating); i++ {
				sortedTeams = append(sortedTeams, -1)
			}
		}

		for i := 0; i < len(sortedTeams)/2; i++ {
			item := model.GameItem{
				LocalID:    int64(len(response)),
				Tournament: elem.ID,
				Team:       model.TeamType{First: sortedTeams[i], Second: sortedTeams[len(sortedTeams)-i-1]},
				Stage:      1,
				Status:     0,
			}

			if sortedTeams[len(sortedTeams)-i-1] == -1 {
				item.Winner = sortedTeams[i]
			}

			response = append(response, item)
		}

		gameCount := len(response) / 2
		var stage, parent int64 = 2, 0

		for gameCount > 0 {
			for i := 0; i < gameCount; i++ {
				response = append(response, model.GameItem{
					ParentFirst:  parent,
					ParentSecond: parent + 1,
					LocalID:      int64(len(response)),
					Tournament:   elem.ID,
					Stage:        stage,
				})
				parent += 2
			}
			gameCount /= 2
			stage++
		}
	}

	for _, elem := range response {
		_, err = db.Insert("game", &elem, "id=serial()")
		if err != nil {
			return err
		}
	}

	return nil
}

func GameStatusDB(obj model.MatchFinished, db *reindexer.Reindexer) error {
	err := db.OpenNamespace("game", reindexer.DefaultNamespaceOptions(), model.GameItem{})
	if err != nil {
		return err
	}
	err = db.OpenNamespace("status", reindexer.DefaultNamespaceOptions(), model.StatusItem{})
	if err != nil {
		return err
	}
	query := db.Query("status").Where("name", reindexer.EQ, "FINISHED")

	var statusInfo model.StatusItem
	statusJSON, _ := query.GetJson()
	err = json.Unmarshal(statusJSON, &statusInfo)

	_, err = db.Update("game", model.GameItem{
		ID:     obj.ID,
		Score:  model.ScoreType{First: obj.ScoreFirst, Second: obj.ScoreSecond},
		Winner: obj.Winner,
		Status: statusInfo.ID,
	})
	return err
}

func TournamentLastDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.TournamentStatistics, error) {
	TeamID, err := strconv.ParseInt(obj.Field, 10, 64) // второй аргумент - это основание системы счисления, третий - размер типа в битах
	if err != nil {
		log.Fatal(err)
	}
	err = db.OpenNamespace("tournaments", reindexer.DefaultNamespaceOptions(), model.TournamentItem{})
	if err != nil {
		panic(err)
	}
	err = db.OpenNamespace("game", reindexer.DefaultNamespaceOptions(), model.GameItem{})
	if err != nil {
		panic(err)
	}

	query := db.Query("tournaments").Where("status", reindexer.EQ, 2).Where("teams", reindexer.SET, obj.Field)
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
		countGame := int32(iterGame.TotalCount())

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
			Game:       countGame,
			Winner:     int32(*aggSumRes.Value),
			Loser:      countGame - int32(*aggSumRes.Value),
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

	que = db.Query("game").Where("tournament", reindexer.EQ, TournamentID).Where("team_second", reindexer.EQ, TeamID)
	que.AggregateSum("score_second")
	que.ReqTotal()
	iterGame = que.Exec()
	var sumGoal float64
	if iterGame.Error() != nil {
		log.Fatal(iterGame.Error())
	}
	aggSumResSecond := iterGame.AggResults()[0]
	sumGoal = *aggSumResFirst.Value + *aggSumResSecond.Value
	return int32(sumGoal), iterGame.Error()
}

func TournamentOTournamentDBneStatDB(TeamID int64, TournamentID int64, db *reindexer.Reindexer) ([]model.TournamentStatistics, error) {

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
			Loser:      coutGame - int32(*aggSumRes.Value),
		})
	}
	return response, err
}
