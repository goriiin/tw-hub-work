package domain

import "time"

type User struct {
	Id      int       `json:"id: int"`
	Nick    string    `json:"nick,omitempty"`
	RegDate time.Time `json:"reg_data,omitempty"`
	Email   string    `json:"email"`
	Alive   bool      `json:"alive,omitempty"`
	Pass    string    `json:"pass,omitempty"`
	OldPass string    `json:"old_pass,omitempty"`
}
