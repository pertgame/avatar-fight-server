package game

import (
	"github.com/0990/goserver/server"
)

var Server *server.Server
var GMgr *GameMgr
var UMgr *UserMgr

func Init(serverID int32) error {
	s, err := server.NewServer(serverID)
	if err != nil {
		return err
	}
	Server = s
	registerHandler()
	GMgr = NewGameMgr(s.Worker())
	UMgr = NewUserMgr()
	return nil
}

func Run() {
	Server.Run()
	GMgr.Start()
}
