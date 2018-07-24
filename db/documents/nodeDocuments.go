package documents

type NodeDocument struct {
	Id		string `bson:"_id,omitempty"`
	Title 	string
	ContentHtml string
}