package setter

import (
	"eventool/internal/pkg/database"
	"eventool/internal/pkg/utils/config"
	"eventool/internal/pkg/utils/log"

	autmcs "eventool/internal/pkg/authorization/delivery/grpc"
	autdelivery "eventool/internal/pkg/authorization/delivery/rest"


	usrdelivery "eventool/internal/pkg/user/delivery/rest"
	usrrepository "eventool/internal/pkg/user/repository"
	usrusecase "eventool/internal/pkg/user/usecase"

	eventdelivery "eventool/internal/pkg/event/delivery"
	eventrepository "eventool/internal/pkg/event/repository"
	eventusecase "eventool/internal/pkg/event/usecase"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Data struct {
	Db  *database.DBManager
	Api *mux.Router
}

type Services struct {
	// Act Data
	// Mov Data
	User Data
	// Col Data
	// Gen Data
	// Ann Data
	// Ser Data
	Event Data

	// Rat Data
	Aut Data
	// Com Data
}

func setAutMcs() autmcs.AutherClient {
	autconn, err := grpc.Dial(":"+config.DevConfigStore.Mcs.Auth.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Warn("{setAutMcs} mcs Dial")
	}

	return autmcs.NewAutherClient(autconn)
}

func SetHandlers(svs Services) {
	userRep := usrrepository.InitUsrRep(svs.User.Db)
	eventRep := eventrepository.InitEventRep(svs.Event.Db)

	userUsc := usrusecase.InitUsrUsc(userRep)
	eventUsc := eventusecase.InitEventUsc(eventRep)

	usrdelivery.SetUsrHandlers(svs.User.Api, userUsc)
	eventdelivery.SetEventHandlers(svs.Event.Api, eventUsc)

	autdelivery.SetAutHandlers(svs.Aut.Api, setAutMcs())
}
