package main

import (
	"fmt"
	"html/template"
	"net/http"

	"crypto/rand"

	"github.com/RyabovSO/goProject/models"
)

var nodes map[string]*models.Node 

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Println(nodes)

	t.ExecuteTemplate(w, "index", nodes)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "write", nil)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	id := r.FormValue("id")
	// ищем ноду в нашей мапе в ключом id
	node, found := nodes[id]
	//если не нашел, то даем ему NotFound
	if !found {
		http.NotFound(w, r)
	}
	//передать ноду в write
	t.ExecuteTemplate(w, "write", node )
}

func saveNodeHandler(w http.ResponseWriter, r *http.Request) {
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

	http.Redirect(w, r, "/", 302)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	delete(nodes, id)

	http.Redirect(w, r, "/", 302)
}

func GenerateId() string {
	b := make([]byte, 16) //генерируем массив байтов
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func main() {
	fmt.Println("Listening on port :3000")

	nodes = make(map[string]*models.Node, 0)

	// /css/app.css
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/saveNode", saveNodeHandler)

	http.ListenAndServe(":3000", nil)
}

