package internal

func NewApplication() *Application {
	app := &Application{}

	app.Bootstrap()

	return app
}
