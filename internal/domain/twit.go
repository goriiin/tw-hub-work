package domain

//id serial primary key,
//author_id integer not null,
//text text,
//photo text,

type Twit struct {
	Id       int    `json:"id,omitempty"`
	AuthorId int    `json:"author_id,omitempty"`
	Text     string `json:"text,omitempty"`
	Photo    string `json:"photo,omitempty"`
}
