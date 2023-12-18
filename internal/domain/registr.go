package domain

//id       serial
//primary key,
//email    text        not null
//unique
//constraint sign_up_email_check
//check (email ~~ '%@%.%'::text),
//nick     varchar(50) not null unique,
//pass     varchar(50) not null,
//old_pass varchar(50) not null

type SignUpUser struct {
	Id      int    `json:"id"`
	Email   string `json:"email"`
	Nick    string `json:"nick"`
	Pass    string `json:"pass"`
	OldPass string `json:"old_pass"`
}
