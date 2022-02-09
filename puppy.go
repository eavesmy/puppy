/************************************************
		Welcome to use puppy service framework.
/************************************************/

package puppy

type Option struct {
	Name     string // App name. [option]
	Group    int    // App belongs gourp id. [option]
	RootNode string // Recent node location. [option]
	id       uint8  // App id.Generate automic.
	location string // App location.
}

// Handler interface.
type Entry interface {
	Init(*App)
}

// Service instance.
type App struct {
	opt Option
}

// Create new app.
func New(options ...Option) (app *App) {

	app = new(App)
	option := *new(Option)

	if len(options) > 1 {
		option = options[0]
	}

	app.opt = option

	return
}

func (app *App) UseHandler(i Entry) {
	i.Init(app)
	// load route.
}

func (app *App) UseRemote(i Entry) {
	i.Init(app)
	// load route.
}

// listen socket.
func (app *App) Listen(location string) {
	// listen socket
	// hybridg protocal.
}
