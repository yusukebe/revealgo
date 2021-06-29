package revealgo

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
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
}

func (server *Server) Serve(param ServerParam) {
	port := 3000
	if server.port > 0 {
		port = server.port
	}
	fmt.Printf("accepting connections at http://*:%d/\n", port)
	http.Handle("/", contentHandler(param, http.FileServer(http.Dir("."))))
	http.Handle("/revealjs/", assetsHandler("/assets/", http.FileServer(http.FS(revealjs))))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
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
