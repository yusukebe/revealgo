package revealgo

import (
	"fmt"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
)

type Server struct {
	port int
}

func (server *Server) serve() {
	http.Handle("/revealjs/", http.StripPrefix("/revealjs/", http.FileServer(&assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "assets/revealjs",
	})))
	fmt.Printf("Accepting connections at http://0:%d/\n", server.port)
	http.ListenAndServe(fmt.Sprintf(":%d", server.port), nil)
}

func (server *Server) handleExample(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "revealgo!")
}
