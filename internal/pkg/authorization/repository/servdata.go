package autrepository

import (
	"eventool/internal/pkg/database"
	"eventool/internal/pkg/domain"
	"eventool/internal/pkg/utils/cast"
	"eventool/internal/pkg/utils/log"

	"golang.org/x/crypto/bcrypt"
)

type dbAuthRepository struct {
	dbm *database.DBManager
}

func InitAutRep(manager *database.DBManager) domain.AuthRepository {
	return &dbAuthRepository{
		dbm: manager,
	}
}

func (ur *dbAuthRepository) GetByEmail(email string, isFound bool) (domain.User, error) {
	resp, err := ur.dbm.Query(queryGetByEmail, email)
	if len(resp) == 0 {
		if (!isFound) {
			return domain.User{}, nil
		}
		log.Warn("{GetByEmail}")
		log.Error(domain.Err.ErrObj.NoUser)
		return domain.User{}, domain.Err.ErrObj.NoUser
	} else {
		if (!isFound){
			log.Warn("{GetByEmail}")
			log.Error(domain.Err.ErrObj.EmailExists)
			return domain.User{}, domain.Err.ErrObj.EmailExists
		}
	}
	
	if err != nil {
		log.Warn("{GetByEmail} in query: " + queryGetByEmail)
		log.Error(err)
		return domain.User{}, domain.Err.ErrObj.InternalServer
	}

	row := resp[0]
	out := domain.User{
		Id:             cast.ToUint64(row[0]),
		Username:       cast.ToString(row[1]),
		Password:       cast.ToString(row[4]),
		Email:          cast.ToString(row[2]),
		Imgsrc:         cast.ToString(row[3]),
		RepeatPassword: cast.ToString(row[4]),
		PhoneNumber:    cast.ToString(row[5]),
	}

	return out, nil
}

func (ur *dbAuthRepository) GetById(id uint64) (domain.User, error) {
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
		Email:          cast.ToString(row[2]),
		Imgsrc:         cast.ToString(row[3]),
		RepeatPassword: "",
	}

	return out, nil
}

func (ur *dbAuthRepository) AddUser(us domain.User) (uint64, error) {
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Warn("{AddUser}")
		log.Error(err)
		return 0, domain.Err.ErrObj.InternalServer
	}

	resp, err := ur.dbm.Query(queryAddUser, us.Username, us.Email, passwordByte, us.PhoneNumber, us.Age)
	if err != nil {
		log.Warn("{AddUser} in query: " + queryAddUser)
		log.Error(err)
		return 0, err
	}
	userId := cast.ToUint64(resp[0][0])

	for i, _ := range(us.Categories) {
		_, err = ur.dbm.Query(queryCreateUserCategory, userId, us.Categories[i])
		if err != nil {
			log.Warn("{queryCreateUserCategory} in query: " + queryCreateUserCategory)
			log.Error(err)
			return userId, err
		}
	}

	return userId, nil
}
