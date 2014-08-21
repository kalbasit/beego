// +build appengine

package beego

import "net/http"

// Run beego application.
func (app *App) Run() {
	http.Handle("/", app.Handlers)
}
