package models

type Book struct {
	ID            string   `uri:"id" json:"_id,omitempty" bson:"_id,omitempty"`
	Title         string   `json:"title,omitempty" bison:"title,omitempty"`
	AuthorName    string   `json:"author_name,omitempty" bison:"author_name,omitempty"`
	AuthorSurname string   `json:"author_surname,omitempty" bison:"author_surname,omitempty"`
	Description   string   `json:"description,omitempty" bison:"description,omitempty"`
	Categories    []string `json:"categories,omitempty" bison:"categories,omitempty"`
	Views         int      `json:"views,omitempty" bison:"views,omitempty"`
	Likes         int      `json:"likes,omitempty" bison:"likes,omitempty"`
	AudioLink     string   `json:"link,omitempty" bson:"link,omitempty"`
}
