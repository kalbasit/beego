// +build appengine

package beego

import "path/filepath"

func registerAppPathAndAppConfigPath() {
	AppPath = ""
	AppConfigPath = filepath.Join(AppPath, "conf", "app.conf")
}
