package domain

//id serial primary key,
//text text,
//photo text

type Content struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Photo   string `json:"photo"`
}
