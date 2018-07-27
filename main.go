package main

import (
	"fmt"
	"net/http"
	"crypto/rand"

	"github.com/RyabovSO/goProject/db/documents"
	"github.com/RyabovSO/goProject/models"	

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"gopkg.in/mgo.v2"
)

 
var nodesCollection *mgo.Collection

func indexHandler(rnd render.Render) {
	nodeDocuments := []documents.NodeDocument{}
	nodesCollection.Find(nil).All(&nodeDocuments)

	nodes := []models.Node{}
	for _, doc :=  range nodeDocuments {
		node := models.Node{doc.Id, doc.Title, doc.ContentHtml}
		nodes = append(nodes, node)
	}

	rnd.HTML(200, "index", nodes)
}

func writeHandler(rnd render.Render) {
	node := models.Node{}
	rnd.HTML(200, "write", node)
}

func editHandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	nodeDocument := documents.NodeDocument{}
	err := nodesCollection.FindId(id).One(&nodeDocument)
	//если не нашел, то редиректим на главную
	if err != nil {	
		rnd.Redirect("/")
		return
	}
	node := models.Node{nodeDocument.Id, nodeDocument.Title, nodeDocument.ContentHtml}

	rnd.HTML(200, "write", node)	
}

func saveNodeHandler(rnd render.Render, r *http.Request) {
	//id := GenerateId()
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentHtml := r.FormValue("content")

	nodeDocument := documents.NodeDocument{id, title, contentHtml}
	//если id не пустая строка (значит мы редактировали)
	if id != "" {
		nodesCollection.UpdateId(id, nodeDocument)
	} else {
		id = GenerateId()
		nodeDocument.Id = id
		nodesCollection.Insert(nodeDocument)
	}	

	rnd.Redirect("/")
}

func deleteHandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		rnd.Redirect("/")
		return
	}

	nodesCollection.RemoveId(id)
	rnd.Redirect("/")
}

func GenerateId() string {
	b := make([]byte, 16) //генерируем массив байтов
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func main() {
	fmt.Println("Listening on port :3000")

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	nodesCollection = session.DB("blog").C("nodes")

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

	m.Run();
}