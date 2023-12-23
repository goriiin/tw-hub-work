package domain

import "time"

//id serial primary key,
//nickname varchar(50) not null unique,
//reg_data timestamp not null,
//email    text        not null
//constraint client_email_check
//check (email ~~ '%@%.%'::text),
//alive bool not null

type User struct {
	Id      int       `json:"id: int"`
	Nick    string    `json:"nick"`
	RegDate time.Time `json:"reg_data"`
	Email   string    `json:"email"`
	Alive   bool      `json:"alive"`
}
