package domain

type App struct {
	Secret string
}

func NewApp() *App {
	return &App{
		Secret: "MY_SECRET",
	}
}
