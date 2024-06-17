package internal

/*
[Pointer Receiver](https://sites-flow-5hu.craft.me/uN0Y07rZXR8kdZ) function
that accepts a Providers pointer and adds a method, `Routes` for instantiating
all routes
*/
func (app *Providers) Routes() {
	// Add Individual Routes below

	app.Router.NoRoute(app.noRouteProvider())

	//views
	v1 := app.Router.Group("/v1")
	{
		v1.GET("/", app.indexPageHandler())
	}

	//api
}
