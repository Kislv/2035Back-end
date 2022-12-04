package usrrepository

import (
	"eventool/internal/pkg/database"
	"eventool/internal/pkg/domain"
	"eventool/internal/pkg/utils/cast"
	"eventool/internal/pkg/utils/log"
)

type dbUserRepository struct {
	dbm *database.DBManager
}

func InitUsrRep(manager *database.DBManager) domain.UserRepository {
	return &dbUserRepository{
		dbm: manager,
	}
}

func (ur *dbUserRepository) GetById(id uint64) (domain.User, error) {
	resp, err := ur.dbm.Query(queryGetById, id)
	if len(resp) == 0 {
		log.Warn("{GetById}")
		log.Error(domain.Err.ErrObj.NoUser)
		return domain.User{}, domain.Err.ErrObj.NoUser
	}
	if err != nil {
		log.Warn("{GetById} in query: " + queryGetById)
		log.Error(err)
		return domain.User{}, domain.Err.ErrObj.InternalServer
	}

	row := resp[0]
	out := domain.User{
		Id:             cast.ToUint64(row[0]),
		Username:       cast.ToString(row[1]),
		Password:       "",
		RepeatPassword: "",
		Email:          cast.ToString(row[2]),
		Imgsrc:         cast.ToString(row[3]),
		PhoneNumber:    cast.ToString(row[4]),
		Age:            cast.ToUint64(row[5]),
	}

	return out, nil
}

func (ur *dbUserRepository) UpdateAvatar(clientId uint64, url string) (domain.User, error) {
	_, err := ur.dbm.Query(queryUpdAvatarByUsID, clientId, url)
	if err != nil {
		return domain.User{}, err
	}

	updated, err := ur.GetById(clientId)
	if err != nil {
		return domain.User{}, err
	}

	return updated, err
}

func (ur *dbUserRepository) GetCategory(id uint64) ([]string, error) {
	query := queryGetCategory
	resp, err := ur.dbm.Query(query, id)

	if len(resp) == 0 {
		log.Warn("{GetCategory}")
		log.Error(domain.Err.ErrObj.NoCategory)
		return nil, domain.Err.ErrObj.NoCategory
	}
	if err != nil {
		log.Warn("{GetCategory} in query: " + query)
		log.Error(err)
		return nil, domain.Err.ErrObj.InternalServer
	}


	category := make([]string, 0)
	for i := range resp {
		category = append(category, cast.ToString(resp[i][0]))
	}

	return category, nil
}
