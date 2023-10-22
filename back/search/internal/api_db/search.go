package api_db

import (
	"tinkoff/internal/model"

	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
)

func ConditionalSearchDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.ConditionaltSearchItem, error) {
	err := db.OpenNamespace("tournaments", reindexer.DefaultNamespaceOptions(), model.TournamentItem{})
	if err != nil {
		panic(err)
	}
	query := db.Query("tournaments").Match("name", obj.Field+"*~").Limit(20)
	if err != nil {
		panic(err)
	}
	var response []model.ConditionaltSearchItem
	iterator := query.Exec()
	if err != nil {
		panic(err)
	}
	for iterator.Next() {
		elem := iterator.Object().(*model.TournamentItem)
		response = append(response, model.ConditionaltSearchItem{Name: elem.Name})
	}
	return response, err
}
func TournamentSearchDB(obj model.FilterSearch, db *reindexer.Reindexer) ([]model.TournamentSearchItem, error) {
	err := db.OpenNamespace("status", reindexer.DefaultNamespaceOptions(), model.StatusItem{})
	if err != nil {
		panic(err)
	}
	err = db.OpenNamespace("tournaments", reindexer.DefaultNamespaceOptions(), model.TournamentItem{})
	if err != nil {
		panic(err)
	}
	query := db.Query("tournaments").Match("name", obj.Field+"*~").Limit(20)
	queryDouble := db.Query("status")
	flag := false
	for index, value := range obj.Filter {
		if (value != "") && (value != "0") {
			switch index {
			case 0:
				queryDouble = queryDouble.Where("name", reindexer.EQ, "OPENED")
				flag = true
			case 1:
				if flag {
					queryDouble = queryDouble.Or().Where("name", reindexer.EQ, "ACTIVE")
				} else {
					queryDouble = queryDouble.Where("name", reindexer.EQ, "ACTIVE")
					flag = true
				}
			case 2:
				if flag {
					queryDouble = queryDouble.Or().Where("name", reindexer.EQ, "FINISHED")
				} else {
					queryDouble = queryDouble.Where("name", reindexer.EQ, "FINISHED")
				}
			}
		}
	}
	query = query.InnerJoin(queryDouble, "status_arr").On("status", reindexer.EQ, "id")
	var response []model.TournamentSearchItem
	iterator := query.Exec()
	if iterator.Error() != nil {
		panic(err)
	}
	for iterator.Next() {
		elem := iterator.Object().(*model.TournamentItem)
		response = append(response, model.TournamentSearchItem{
			ID:     elem.ID,
			Name:   elem.Name,
			Status: elem.StatusArr[0].Name,
			Update: elem.Date.Update,
			Create: elem.Date.Create,
		})
	}
	return response, err
}
func TeamSearchDB(obj model.ParamSearch, db *reindexer.Reindexer) ([]model.TeamSearchItem, error) {
	err := db.OpenNamespace("team", reindexer.DefaultNamespaceOptions(), model.TeamItem{})
	if err != nil {
		panic(err)
	}
	query := db.Query("team").Match("name", obj.Field+"*~").Limit(50)
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
