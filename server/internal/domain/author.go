package domain

type Author struct {
	Id   int    `json:"id"`
	Nick string `json:"nick,omitempty"`
	//RegDate time.Time `json:"reg_data,omitempty"`
	//Email   string    `json:"email,omitempty"`
	//Photo   string    `json:"photo,omitempty"`
}
