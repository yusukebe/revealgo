package revealgo

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
	port int
}

func (server *Server) Serve() {
	fmt.Printf("Accepting connections at http://0:%d/\n", server.port)
	http.HandleFunc("/revealjs/", server.HandleStatic)
	http.ListenAndServe(fmt.Sprintf(":%d", server.port), nil)
}

func (server *Server) HandleExample(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "revealgo!")
}

func (server *Server) HandleStatic(w http.ResponseWriter, r *http.Request) {
	filepath := "assets" + r.URL.Path
	data, err := Asset(filepath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	mimeType := server.DetectContentType(filepath, data)
	w.Header().Set("Content-Type", mimeType)
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	if _, err := w.Write(data); err != nil {
		log.Println("unable to write data.")
	}
}

func (server *Server) DetectContentType(path string, data []byte) string {
	if strings.HasSuffix(path, ".css") {
		return "text/css"
	} else if strings.HasSuffix(path, ".js") {
		return "application/javascript"
	}
	return http.DetectContentType(data)
}
