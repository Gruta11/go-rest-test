package main
import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
	"fmt"
	"github.com/go-chi/docgen"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	flag.Parse()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("root."))
	})

	r.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	})

	r.Get("/panic", func(writer http.ResponseWriter, request *http.Request) {
			panic("test")
	})
	if *routes {
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/fravia190/go-rest-test",
			Intro: "Welcome to the go-rest-test",
		}))
	}
	http.ListenAndServe(":3333", r)
}
