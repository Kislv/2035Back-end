package mcsauth

import (
	proto "eventool/internal/pkg/authorization/delivery/grpc"
	autrepository "eventool/internal/pkg/authorization/repository"
	autusecase "eventool/internal/pkg/authorization/usecase"
	"eventool/internal/pkg/database"
	"eventool/internal/pkg/utils/config"
	"eventool/internal/pkg/utils/log"

	"net"

	"google.golang.org/grpc"
)

func RunServer() {
	db := database.InitDatabase()
	db.Connect()
	defer db.Disconnect()

	autRep := autrepository.InitAutRep(db)
	autUsc := autusecase.InitAutUsc(autRep)

	s := grpc.NewServer()

	proto.RegisterAutherServer(s, autUsc)

	l, err := net.Listen(config.DevConfigStore.Mcs.Auth.ConnType, ":"+config.DevConfigStore.Mcs.Auth.Port)
	if err != nil {
		log.Warn("{RunServer} mcs auth")
		log.Error(err)
	}

	err = s.Serve(l)
	if err != nil{
		log.Error(err)
	}
}
