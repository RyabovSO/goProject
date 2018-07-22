package models

type Node struct {
	Id 			string
	Title 		string
	Content 	string
}
//создаем конструктор
func NewNode(id, title, content string) *Node { //возвращаем указатель на Node
	return &Node{id, title, content}
}
