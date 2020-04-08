package models

import (
	"blog-server/config"
)

var Config config.config
func init() {
	Config, _ = config.Init()
}
