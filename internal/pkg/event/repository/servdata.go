package eventrepository

import (
	"eventool/internal/pkg/database"
	"eventool/internal/pkg/domain"
	"eventool/internal/pkg/utils/cast"
	"eventool/internal/pkg/utils/log"
	"strings"
	"time"
)

type dbeventrepository struct {
	dbm *database.DBManager
}

func InitEventRep(manager *database.DBManager) domain.EventRepository {
	return &dbeventrepository{
		dbm: manager,
	}
}

func (er *dbeventrepository) CreateEvent(event domain.EventCreatingRequest) (domain.EventCreatingResponse, error) {
	resp, err := er.dbm.Query(queryCreateEvent, event.Title, event.Description,
								event.UserId, event.Longitude, event.Latitude, event.MaxMembersQuantity,
								event.MinMembersQuantity, time.Now().Format("2006.01.02 15:04:05"), event.StartDate,
								event.EndDate, event.MinAge, event.MaxAge, event.Price)
	if err != nil {
		log.Warn("{CreateEvent} in query: " + queryCreateEvent)
		log.Error(err)
		return domain.EventCreatingResponse{}, err
	}

	return domain.EventCreatingResponse{
		Id:     					cast.ToUint64(resp[0][0]),
		PosterPath:    				cast.ToString(resp[0][1]),
		Title:  					cast.ToString(resp[0][2]),
		Rating:  					cast.ToString(resp[0][3]),
		VotesNum:  					cast.ToUint64(resp[0][4]),
		Description:  				cast.ToString(resp[0][5]),
		UserId:  					cast.ToString(resp[0][6]),
		Longitude:  				cast.ToString(resp[0][7]),
		Latitude:  					cast.ToString(resp[0][8]),
		CurrentMembersQuantity:  	cast.ToUint64(resp[0][9]),
		MaxMembersQuantity:  		cast.ToUint64(resp[0][10]),
		MinMembersQuantity:  		cast.ToUint64(resp[0][11]),
		CreatingDate:  				cast.ToString(resp[0][12]),
		StartDate:  				cast.ToString(resp[0][13]),
		EndDate:  					cast.ToString(resp[0][14]),
		MinAge:  					cast.ToString(resp[0][15]),
		MaxAge:  					cast.ToString(resp[0][16]),
		Price:  					cast.ToString(resp[0][17]),
	}, nil
}

func (er *dbeventrepository) CreateEventCategory(eventId uint64, categories []string) ([]string, error) {
	// var sb strings.Builder
	// sb.WriteString(queryCreateEventCategoryFirstPart)
	// for i, el := range(categories) {
	// 	sb.WriteString(queryCreateEventCategorySecondPart)
	// 	sb.WriteString(string(i * 2))
	// 	sb.WriteString(queryCreateEventCategoryThirdPart)
	// 	sb.WriteString(string(i * 2 + 1))
	// 	sb.WriteString(queryCreateEventCategoryForthPart)
	// 	sb.WriteString(",")
	// }
	// sb.WriteString(queryCreateEventCategoryFifthPart)
	// var resp []database.DBbyterow
	var err error
	for i, _ := range(categories) {
		_, err = er.dbm.Query(queryCreateEventCategory, eventId, categories[i])
		if err != nil {
			log.Warn("{CreateEventCategory} in query: " + queryCreateEventCategory)
			log.Error(err)
			return nil, err
		}
	}

	return categories, nil
}

func (er *dbeventrepository) EventAlreadyExist(event domain.EventCreatingRequest) (bool, error) {
	resp, err := er.dbm.Query(queryCheckEvent, event.Title, event.Longitude, event.Latitude)
	if err != nil {
		log.Warn("{EventCreating} in query: " + queryCheckEvent)
		log.Error(err)
		return false, err
	}

	if cast.ToUint64(resp[0][0]) != 0 {
		return true, nil
	}
	return false, nil
}


func (cr *dbeventrepository) GetEvent(categoriesName []string) (domain.EventListResponse, error) {
	var resp []database.DBbyterow 
	var err error
	query := ""
	if categoriesName[0] != "all"{
		// query = queryGetEventListFirstPart
		var sb strings.Builder
		sb.WriteString(queryGetEventListFirstPart)
		sb.WriteString("(")
		for i, _ := range(categoriesName){
			if i != len(categoriesName) - 1{
				sb.WriteString("'")
				sb.WriteString(categoriesName[i])
				sb.WriteString("'")
				sb.WriteString(",")
			} else {
				sb.WriteString("'")
				sb.WriteString(categoriesName[i])
				sb.WriteString("'")

			}
		}
		sb.WriteString(")")
		sb.WriteString(queryGetEventListSecondPart)
		query = sb.String()
		// TODO: sql injection
		resp, err = cr.dbm.Query(sb.String())
	} else {
		query = queryGetAllEventList
		resp, err = cr.dbm.Query(query)
	}

	if err != nil {
		log.Warn("{GetEvent} in query: " + query)
		log.Error(err)
		return domain.EventListResponse{}, domain.Err.ErrObj.InternalServer
	}

	if len(resp) == 0 {
		log.Warn("{GetMovies}")
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.EventListResponse{}, domain.Err.ErrObj.SmallDb
	}
	
	events := make([]domain.EventCreatingResponse, 0)
	for i := range resp {
		events = append(events, domain.EventCreatingResponse{
			Id:     					cast.ToUint64(resp[i][0]),
			PosterPath:    				cast.ToString(resp[i][1]),
			Title:  					cast.ToString(resp[i][2]),
			Rating:  					cast.FlToStr((cast.ToFloat64(resp[i][3]))),
			VotesNum:  					cast.ToUint64(resp[i][4]),
			Description:  				cast.ToString(resp[i][5]),
			UserId:  					cast.ToString(resp[i][6]),
			Longitude:  				cast.FlToStr((cast.ToFloat64(resp[i][7]))),
			Latitude:  					cast.FlToStr((cast.ToFloat64(resp[i][8]))),
			CurrentMembersQuantity:  	cast.ToUint64(resp[i][9]),
			MaxMembersQuantity:  		cast.ToUint64(resp[i][10]),
			MinMembersQuantity:  		cast.ToUint64(resp[i][11]),
			CreatingDate:  				cast.ToString(resp[i][12]),
			StartDate:  				cast.ToString(resp[i][13]),
			EndDate:  					cast.ToString(resp[i][14]),
			MinAge:  					cast.ToString(resp[i][15]),
			MaxAge:  					cast.ToString(resp[i][16]),
			Price:  					cast.ToString(resp[i][17]),
		})
	}

	out := domain.EventListResponse{
		EventList: events,
	}

	return out, nil
}

func (cr *dbeventrepository) GetCertainEvent(eventId uint64) (domain.EventCreatingResponse, error){
	var resp []database.DBbyterow 
	var err error
	query := queryGetCertainEvent
	resp, err = cr.dbm.Query(query, eventId)

	if err != nil {
		log.Warn("{GetCertainEvent} in query: " + query)
		log.Error(err)
		return domain.EventCreatingResponse{}, domain.Err.ErrObj.InternalServer
	}

	if len(resp) == 0 {
		log.Warn("{GetCertainEvent}")
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.EventCreatingResponse{}, domain.Err.ErrObj.SmallDb
	}
	
	event := domain.EventCreatingResponse{
		Id:     					cast.ToUint64(resp[0][0]),
		PosterPath:    				cast.ToString(resp[0][1]),
		Title:  					cast.ToString(resp[0][2]),
		Rating:  					cast.FlToStr((cast.ToFloat64(resp[0][3]))),
		VotesNum:  					cast.ToUint64(resp[0][4]),
		Description:  				cast.ToString(resp[0][5]),
		UserId:  					cast.ToString(resp[0][6]),
		Longitude:  				cast.FlToStr((cast.ToFloat64(resp[0][7]))),
		Latitude:  					cast.FlToStr((cast.ToFloat64(resp[0][8]))),
		CurrentMembersQuantity:  	cast.ToUint64(resp[0][9]),
		MaxMembersQuantity:  		cast.ToUint64(resp[0][10]),
		MinMembersQuantity:  		cast.ToUint64(resp[0][11]),
		CreatingDate:  				cast.ToString(resp[0][12]),
		StartDate:  				cast.ToString(resp[0][13]),
		EndDate:  					cast.ToString(resp[0][14]),
		MinAge:  					cast.ToString(resp[0][15]),
		MaxAge:  					cast.ToString(resp[0][16]),
		Price:  					cast.ToString(resp[0][17]),
	}

	out := event

	return out, nil
}


func (cr *dbeventrepository) GetCategory() (domain.CategoryListResponse, error) {
	var resp []database.DBbyterow 
	var err error
	resp, err = cr.dbm.Query(queryGetCategoryList)

	if err != nil {
		log.Warn("{GetCategory} in query: " + queryGetCategoryList)
		log.Error(err)
		return domain.CategoryListResponse{}, domain.Err.ErrObj.InternalServer
	}

	if len(resp) == 0 {
		log.Warn("{GetCategory}")
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.CategoryListResponse{}, domain.Err.ErrObj.SmallDb
	}
	
	categoryList := make([]domain.CategoryResponse, 0)
	for i := range resp {
		categoryList = append(categoryList, domain.CategoryResponse{
			Name:    				cast.ToString(resp[i][0]),
			ImagePath:    			cast.ToString(resp[i][1]),
		})
	}

	out := domain.CategoryListResponse{
		CategoryList: categoryList,
	}

	return out, nil
}
