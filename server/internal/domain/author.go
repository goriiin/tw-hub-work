package domain

import "time"

type Author struct {
	Id      int       `json:"id: int"`
	Nick    string    `json:"nick,omitempty"`
	RegDate time.Time `json:"reg_data,omitempty"`
	Email   string    `json:"email"`
}
