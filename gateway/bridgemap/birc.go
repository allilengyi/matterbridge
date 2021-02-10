// +build !noirc

package bridgemap

import (
	birc "github.com/allilengyi/matterbridge/bridge/irc"
)

func init() {
	FullMap["irc"] = birc.New
}
