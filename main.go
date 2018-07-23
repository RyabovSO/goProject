package main

import (
	"fmt"
	"net/http"

	"crypto/rand"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/RyabovSO/goProject/models"
	
)

var nodes map[string]*models.Node 
var counter int

func indexHandler(rnd render.Render) {
	fmt.Println(counter)

	rnd.HTML(200, "index", nodes)
}

func writeHandler(rnd render.Render) {
	rnd.HTML(200, "write", nil)
}

func editHandler(rnd render.Render, r *http.Request) {
	
	id := r.FormValue("id")
	// ищем ноду в нашей мапе в ключом id
	node, found := nodes[id]
	//если не нашел, то редиректим на главную
	if !found {	
		rnd.Redirect("/")
		return
	}

	rnd.HTML(200, "write", node)	
}

func saveNodeHandler(rnd render.Render, r *http.Request) {
	//id := GenerateId()
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var node *models.Node
	//если id не пустая строка (значит мы редактировали)
	if id != "" {
		node = nodes[id]
		node.Title = title
		node.Content = content
	} else {
		id = GenerateId()
		node := models.NewNode(id, title, content)
		//создали ноду и добавляем ее в наш map
		nodes[node.Id] = node
	}	

	rnd.Redirect("/")
}

func deleteHandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		rnd.Redirect("/")
		return
	}

	delete(nodes, id)

	rnd.Redirect("/")
}

func GenerateId() string {
	b := make([]byte, 16) //генерируем массив байтов
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func main() {
	fmt.Println("Listening on port :3000")

	nodes = make(map[string]*models.Node, 0)

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
  		Directory: "templates", 					// Specify what path to load the templates from.
  		Layout: "layout", 							// Specify a layout template. Layouts can call {{ yield }} to render the current template.
  		Extensions: []string{".tmpl", ".html"}, 	// Specify extensions to load for templates.
  		//Funcs: []template.FuncMap{AppHelpers}, 	// Specify helper function maps for templates to access.
  		//Delims: render.Delims{"{[{", "}]}"}, 		// Sets delimiters to the specified strings.
  		Charset: "UTF-8", 							// Sets encoding for json and html content-types. Default is "UTF-8".
  		IndentJSON: true, 							// Output human readable JSON
	}))

	staticOptions := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOptions))
	m.Get("/", indexHandler)
	m.Get("/write", writeHandler)
	m.Get("/edit", editHandler)
	m.Get("/delete", deleteHandler)
	m.Post("/saveNode", saveNodeHandler)


	//тестовая функция для счетчика
	counter = 0
	m.Use(func(r *http.Request) {
		//если метод write то counter++
		if r.URL.Path == "/write" {
			counter++
		}
	})
	m.Get("/test", func() string{
		return "test"
	})

	m.Run();
}