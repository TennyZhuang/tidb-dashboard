// +build ui_server

package uiserver

import (
	"strings"
)

func InitAssetFS(prefix string) {
	a, err := _bindata["build/index.html"]()
	if err != nil {
		panic("Asset index.html not found.")
	}
	tmplText := string(a.bytes)
	updated := strings.ReplaceAll(tmplText, "$DashboardPrefix$", prefix)
	a.bytes = []byte(updated)
	_bindata["build/index.html"] = func() (*asset, error) {
		return a, nil
	}
}