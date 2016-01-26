package revealgo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

type Server struct {
	port   int
	mdpath string
}

type SlideParam struct {
	Path string
}

func (server *Server) Serve() {
	port := 3000
	if server.port > 0 {
		port = server.port
	}
	fmt.Printf("accepting connections at http://0:%d/\n", port)
	http.HandleFunc("/", server.HandleRoot)
	http.HandleFunc("/revealjs/", server.HandleStatic)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func (server *Server) HandleRoot(w http.ResponseWriter, r *http.Request) {

	urlPath := r.URL.Path
	path, err := filepath.Rel("./", "."+urlPath)
	if err != nil {
		log.Printf("error:%v\n", err)
	}
	_, err = os.Stat(path)
	if err == nil {
		data, err := ioutil.ReadFile(path)
		if err == nil {
			mimeType := server.DetectContentType(path, data)
			w.Header().Set("Content-Type", mimeType)
			w.Header().Set("Content-Length", strconv.Itoa(len(data)))
			if _, err := w.Write(data); err != nil {
				log.Println("unable to write data.")
			}
			return
		}
	}

	data, err := Asset("assets/templates/slide.html")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		http.NotFound(w, r)
		return
	}
	tmpl := template.New("slide template")
	tmpl.Parse(string(data))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		http.NotFound(w, r)
		return
	}
	param := SlideParam{
		Path: server.mdpath,
	}
	err = tmpl.Execute(w, param)
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
