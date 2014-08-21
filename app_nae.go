// +build !appengine

package beego

import (
	"fmt"
	"net"
	"net/http/fcgi"
	"time"
)

// Run beego application.
func (app *App) Run() {
	addr := HttpAddr

	if HttpPort != 0 {
		addr = fmt.Sprintf("%s:%d", HttpAddr, HttpPort)
	}

	BeeLogger.Info("Running on %s", addr)

	var (
		err error
		l   net.Listener
	)
	endRunning := make(chan bool, 1)

	if UseFcgi {
		if HttpPort == 0 {
			l, err = net.Listen("unix", addr)
		} else {
			l, err = net.Listen("tcp", addr)
		}
		if err != nil {
			BeeLogger.Critical("Listen: ", err)
		}
		err = fcgi.Serve(l, app.Handlers)
	} else {
		app.Server.Addr = addr
		app.Server.Handler = app.Handlers
		app.Server.ReadTimeout = time.Duration(HttpServerTimeOut) * time.Second
		app.Server.WriteTimeout = time.Duration(HttpServerTimeOut) * time.Second

		if EnableHttpTLS {
			go func() {
				time.Sleep(20 * time.Microsecond)
				if HttpsPort != 0 {
					app.Server.Addr = fmt.Sprintf("%s:%d", HttpAddr, HttpsPort)
				}
				err := app.Server.ListenAndServeTLS(HttpCertFile, HttpKeyFile)
				if err != nil {
					BeeLogger.Critical("ListenAndServeTLS: ", err)
					time.Sleep(100 * time.Microsecond)
					endRunning <- true
				}
			}()
		}

		if EnableHttpListen {
			go func() {
				app.Server.Addr = addr
				err := app.Server.ListenAndServe()
				if err != nil {
					BeeLogger.Critical("ListenAndServe: ", err)
					time.Sleep(100 * time.Microsecond)
					endRunning <- true
				}
			}()
		}
	}

	<-endRunning
}