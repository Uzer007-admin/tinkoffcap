package api_db

import (
	"tinkoff/internal/model"

	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
)

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
func TeamCreateNew(obj model.TeamCreate, db *reindexer.Reindexer) error {
	err := db.OpenNamespace("team", reindexer.DefaultNamespaceOptions(), model.TeamItem{})
	if err != nil {
		panic(err)
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
