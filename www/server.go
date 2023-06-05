package www

import (
	"embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/jcbbb/monkey/evaluator"
	"github.com/jcbbb/monkey/lexer"
	"github.com/jcbbb/monkey/object"
	"github.com/jcbbb/monkey/parser"
)

//go:embed index.html
var tmpl embed.FS

type handlerFunc func(r *http.Request, w http.ResponseWriter) error

func handleHomeView(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(tmpl, "index.html")
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func handleRun(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	l := lexer.New(string(b))
	p := parser.New(l, false)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		w.WriteHeader(400)
		w.Write([]byte(strings.Join(p.Errors(), ", ")))
		return
	}

	env := object.NewEnvironment()
	result := evaluator.Eval(program, env)

	if result != nil {
		_, ok := result.(*object.Error)
		if ok {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(200)
		}

		w.Write([]byte(result.Inspect()))
		return
	}
}

func Start(addr string) {
	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] -> %s\n", r.Method, r.URL)
		switch r.Method {
		case "GET":
			handleHomeView(w, r)
			break
		case "POST":
			handleRun(w, r)
			break
		default:
			w.WriteHeader(405)
			return
		}
	})

	fmt.Println("Server starting at: ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
