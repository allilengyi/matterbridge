// +build !nogitter

package bridgemap

import (
	bgitter "github.com/allilengyi/matterbridge/bridge/gitter"
)

func init() {
	FullMap["gitter"] = bgitter.New
}
