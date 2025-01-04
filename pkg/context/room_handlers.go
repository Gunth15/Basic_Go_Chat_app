package context

import (
	"log"
	"net/http"

	"github.com/chat_app/pkg/database"
	"github.com/gorilla/websocket"
)

// upgrader is the private and represents the read and write buffer size for the ws connecton
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewChatRoomMux(ctxt *Ctxt) *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

func (ctxt *Ctxt) WsChatRoom(w http.ResponseWriter, r *http.Request, user database.User) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "error: Could not create websocket", http.StatusInternalServerError)
	}
	// Add to connection to online users
	ctxt.Online[user.ID] = conn

	for {
		conn.NextReader()
	}
}
