package models

type Post struct {
	Id 			string
	Title 		string
	Content 	string
}
//создаем конструктор
func NewNode(id, title, content string) *Post { //возвращаем указатель на Post
	return &Post{id, title, content}
}
