package skuser

import (
	"200lab-project-1/common"
	"200lab-project-1/component/appctx"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnUserUpdateLocation(appCtx appctx.AppContext, requester common.Requester) func(s socketio.Conn, location LocationData) {
	return func(s socketio.Conn, location LocationData) {
		log.Println("User", requester.GetUserId(), "update location", location)
	}
}
