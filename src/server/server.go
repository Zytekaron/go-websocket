package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zytekaron/gotil/random"
	"go-websocket/src/types"
	"log"
	"net/http"
	"strings"
)

var cache = types.NewClientCache()
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Run the web server that connects clients
func Run() {
	r := gin.New()

	ch := makeHandler()

	r.GET("/ws", auth, func(c *gin.Context) {
		ws(c, ch)
	})

	err := http.ListenAndServe(":1337", r)
	log.Fatal(err)
}

// The websocket endpoint handler
func ws(c *gin.Context, sh chan<- *types.ServerMessage) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	id := random.MustSecureString(16, []rune("0123456789abcdef"))
	client := types.NewClient(id, conn)
	cache.Set(id, client)

	for {
		var msg *types.Message
		err = conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("S Read error:", err)
			break
		}

		sh <- &types.ServerMessage{
			Message: msg,
			Client:  client,
		}
	}

	// Client has disconnected
	cache.Delete(id)
}

// The authentication header middleware
func auth(c *gin.Context) {
	header := c.GetHeader("Authorization")
	header = strings.TrimPrefix(header, "Bearer")
	header = strings.TrimLeft(header, " ")

	query := c.Query("token")

	var token string
	if header != "" {
		token = header
	} else if query != "" {
		token = query
	} else {
		c.AbortWithStatusJSON(403, types.NewResponseError(403, "Missing authentication header"))
		return
	}

	if !isValid(token) {
		c.AbortWithStatusJSON(401, types.NewResponseError(401, "Missing access"))
	}
}

// Check if a string is valid websocket authentication
func isValid(auth string) bool {
	return auth == "i_am_god" // fixme proper auth
}
