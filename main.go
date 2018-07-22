package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/RyabovSO/goProject/models"
)

var posts map[string]*models.Post 

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Println(posts)

	t.ExecuteTemplate(w, "index", nil)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "write", nil)
}

func saveNodeHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	node := models.NewNode(id, title, content)
	node[node.Id] = node

	http.Redirect(w, r, "/", 302)
}


func main() {
	fmt.Println("Listening on port :3000")

	post = make(map[string]*models.Post, 0)

	// /css/app.css
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/saveNode", saveNodeHandler)

	http.ListenAndServe(":3000", nil)
}
