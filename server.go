package revealgo

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"golang.org/x/crypto/bcrypt"
)

//go:embed assets/revealjs
var revealjs embed.FS

//go:embed assets/templates/slide.html
var slideTemplate string

type Server struct {
	port int
}

type ServerParam struct {
	Path          string
	Theme         string
	OriginalTheme bool
	Transition    string
	Multiplex     MultiplexParam
}

type MultiplexParam struct {
	IsMaster   bool
	Secret     string
	Identifier string
}

type RevealMultiplexData struct {
	Secret   string                 `json:"secret,omitempty"`
	SocketId string                 `json:"socketId"`
	State    map[string]interface{} `json:"state"`
}

func (server *Server) Serve(param ServerParam) {
	port := 3000
	if server.port > 0 {
		port = server.port
	}
	fmt.Printf("accepting connections at http://*:%d/\n", port)
	http.Handle("/", contentHandler(param, http.FileServer(http.Dir("."))))
	http.Handle("/revealjs/", assetsHandler("/assets/", http.FileServer(http.FS(revealjs))))

	if param.Multiplex.Secret != "" {
		socketioServer := setupSocketIO()
		go func() {
			if err := socketioServer.Serve(); err != nil {
				log.Fatalf("socketio listen error: %s\n", err)
			}
		}()
		defer socketioServer.Close()

		param.Multiplex.IsMaster = true
		http.Handle("/master/", http.StripPrefix("/master", contentHandler(param, http.FileServer(http.Dir(".")))))
		http.Handle("/socket.io/", socketioServer)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func setupSocketIO() *socketio.Server {
	var allowOriginFunc = func(r *http.Request) bool {
		return true
	}

	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())

		return nil
	})

	server.OnEvent("/", "multiplex-statechanged", func(s socketio.Conn, data RevealMultiplexData) {
		if err := bcrypt.CompareHashAndPassword([]byte(data.SocketId), []byte(data.Secret)); err != nil {
			return
		}

		data.Secret = ""
		server.BroadcastToNamespace("/", data.SocketId, data)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("Socket error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	return server
}

func contentHandler(params ServerParam, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mimeType := detectContentType(r.URL.Path); mimeType != "" {
			w.Header().Set("Content-Type", mimeType)
		}

		if r.URL.Path != "/" {
			h.ServeHTTP(w, r)
			return
		}

		tmpl, err := template.New("slide template").Parse(slideTemplate)
		if err != nil {
			log.Printf("error:%v", err)
			http.NotFound(w, r)
			return
		}

		if err := tmpl.Execute(w, params); err != nil {
			log.Fatalf("error:%v", err)
		}
	})
}

func assetsHandler(prefix string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mimeType := detectContentType(r.URL.Path); mimeType != "" {
			w.Header().Set("Content-Type", mimeType)
		}

		if prefix != "" {
			r.URL.Path = filepath.Join(prefix, r.URL.Path)
			r.URL.RawPath = filepath.Join(prefix, r.URL.RawPath)
		}

		h.ServeHTTP(w, r)
	})
}

func detectContentType(path string) string {
	switch {
	case strings.HasSuffix(path, ".css"):
		return "text/css"
	case strings.HasSuffix(path, ".js"):
		return "application/javascript"
	case strings.HasSuffix(path, ".svg"):
		return "image/svg+xml"
	}
	return ""
}
