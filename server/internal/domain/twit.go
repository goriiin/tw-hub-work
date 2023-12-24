package domain

import "time"

// TODO: добавление времени

type Twit struct {
	Id       int       `json:"id,omitempty"`
	AuthorId int       `json:"author_id,omitempty"`
	Text     string    `json:"text,omitempty"`
	Photo    string    `json:"photo,omitempty"`
	Date     time.Time `json:"date,omitempty"`
}
