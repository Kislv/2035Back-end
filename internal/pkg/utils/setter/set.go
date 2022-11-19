package setter

import (
	"eventool/internal/pkg/database"
	"eventool/internal/pkg/utils/config"
	"eventool/internal/pkg/utils/log"

	autmcs "eventool/internal/pkg/authorization/delivery/grpc"
	autdelivery "eventool/internal/pkg/authorization/delivery/rest"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Data struct {
	Db  *database.DBManager
	Api *mux.Router
}

type Services struct {
	Act Data
	Mov Data
	Usr Data
	Col Data
	Gen Data
	Ann Data
	Ser Data
	Pla Data

	Rat Data
	Aut Data
	Com Data
}

func setAutMcs() autmcs.AutherClient {
	autconn, err := grpc.Dial(":"+config.DevConfigStore.Mcs.Auth.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Warn("{setAutMcs} mcs Dial")
	}

	return autmcs.NewAutherClient(autconn)
}

func SetHandlers(svs Services) {
	autdelivery.SetAutHandlers(svs.Aut.Api, setAutMcs())
}
