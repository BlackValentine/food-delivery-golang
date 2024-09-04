package skuser

import (
	"log"
	"server/common"

	socketio "github.com/googollee/go-socket.io"
	"gorm.io/gorm"
)

type SmallAppContext interface {
	GetMainDBConnection() *gorm.DB
}

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnUserUpdateLocation(appCtx SmallAppContext, requester common.Requester) func(s socketio.Conn, location LocationData) {
	return func(s socketio.Conn, location LocationData) {
		log.Println("User update location: user id is", requester.GetUserId(), "at location:", location)
	}
}
