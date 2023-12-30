package domain

type App struct {
	Secret string
}

func NewApp() *App {
	return &App{
		Secret: "MY_SECRET",
	}
}

type UserData struct {
	Id    uint32 `json:"id"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type RegData struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenUser struct {
	Id    uint32 `json:"id"`
	Email string `json:"email"`
}
