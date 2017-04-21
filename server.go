//go:generate go-bindata -o assets.go -pkg revealgo -ignore ".scss" assets/templates/ assets/revealjs/css/... assets/revealjs/js/... assets/revealjs/plugin/... assets/revealjs/lib/...
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

	"github.com/jaschaephraim/lrserver"

	fsnotify "gopkg.in/fsnotify.v1"
	yaml "gopkg.in/yaml.v2"
)

type Server struct {
	port int
}

type ServerParam struct {
	Path          string
	Theme         string
	OriginalTheme bool
	Transition    string
	Host          string
	Watch         bool
	RevealOptions map[string]interface{}
	Slides        [][]string
}

func (server *Server) Serve(param ServerParam) {
	if param.Watch == true {
		// Create file watcher
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatalln(err)
		}
		defer watcher.Close()

		// Add dir to watcher
		err = watcher.Add(".")
		err = watcher.Add("../../assets/templates/slide.html")
		if err != nil {
			log.Fatalln(err)
		}

		// Create and start LiveReload server
		lr := lrserver.New(lrserver.DefaultName, lrserver.DefaultPort)
		go lr.ListenAndServe()

		// Start goroutine that requests reload upon watcher event
		go func() {
			for {
				select {
				case event := <-watcher.Events:
					lr.Reload(event.Name)
				case err := <-watcher.Errors:
					log.Println(err)
				}
			}
		}()
	}

	port := 3000
	if server.port > 0 {
		port = server.port
	}
	fmt.Printf("accepting connections at http://*:%d/\n", port)
	http.Handle("/", &rootHandler{param: param})
	http.Handle("/revealjs/", &assetHandler{assetPath: "assets"})
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

type assetHandler struct {
	assetPath string
}

func (h *assetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filepath := h.assetPath + r.URL.Path
	data, err := Asset(filepath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	w = setResponse(w, filepath, data)
}

type rootHandler struct {
	param ServerParam
}

func (h *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	path, err := filepath.Rel("./", "."+urlPath)
	if err != nil {
		log.Fatalf("error:%v", err)
	}
	_, err = os.Stat(path)
	if err == nil {
		data, err := ioutil.ReadFile(path)
		if err == nil {
			w = setResponse(w, path, data)
			return
		}
	}

	data, err := Asset("assets/templates/slide.html")
	if err != nil {
		log.Printf("error:%v", err)
		http.NotFound(w, r)
		return
	}
	tmpl := template.New("slide template")
	tmpl.Parse(string(data))
	if err != nil {
		log.Printf("error:%v", err)
		http.NotFound(w, r)
		return
	}

	h.param = loadParams(h.param)

	err = tmpl.Execute(w, h.param)
	if err != nil {
		log.Fatalf("error:%v", err)
	}
}

func loadParams(p ServerParam) ServerParam {
	dat, _ := ioutil.ReadFile(p.Path)
	re := strings.SplitN(string(dat), "+++", 2)
	MDSlides := []string{}
	if len(re) == 2 {
		t := map[string]interface{}{}
		err := yaml.Unmarshal([]byte(re[0]), t)
		if err != nil {

		} else {
			p.Theme = t["theme"].(string)
			if strings.HasSuffix(p.Theme, ".css") {
				p.OriginalTheme = true
			} else {
				p.Theme += ".css"
				p.OriginalTheme = false
			}

			p.RevealOptions = map[string]interface{}{}
			for k, v := range t["revealOptions"].(map[interface{}]interface{}) {
				p.RevealOptions[k.(string)] = v
			}
		}
		MDSlides = strings.Split(re[1], "---")
	} else {
		MDSlides = strings.Split(string(dat), "---")
	}

	c := [][]string{}
	for _, n := range MDSlides {
		cv := []string{}
		for _, s := range strings.Split(n, "___") {
			cv = append(cv, s)
		}
		c = append(c, cv)
	}
	p.Slides = c

	return p
}

func detectContentType(path string, data []byte) string {
	switch {
	case strings.HasSuffix(path, ".css"):
		return "text/css"
	case strings.HasSuffix(path, ".js"):
		return "application/javascript"
	case strings.HasSuffix(path, ".svg"):
		return "image/svg+xml"
	}
	return http.DetectContentType(data)
}

func setResponse(w http.ResponseWriter, path string, data []byte) http.ResponseWriter {
	mimeType := detectContentType(path, data)
	w.Header().Set("Content-Type", mimeType)
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	if _, err := w.Write(data); err != nil {
		log.Fatal("unable to write data.")
	}
	return w
}
