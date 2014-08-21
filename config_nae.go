// +build !appengine

package beego

import (
	"os"
	"path/filepath"

	"github.com/eMxyzptlk/beego/utils"
)

func registerAppPathAndAppConfigPath() {
	workPath, _ = os.Getwd()
	workPath, _ = filepath.Abs(workPath)
	// initialize default configurations
	AppPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	AppConfigPath = filepath.Join(AppPath, "conf", "app.conf")

	if workPath != AppPath {
		if utils.FileExists(AppConfigPath) {
			os.Chdir(AppPath)
		} else {
			AppConfigPath = filepath.Join(workPath, "conf", "app.conf")
		}
	}

}
