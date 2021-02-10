// +build !nonctalk

package bridgemap

import (
	btalk "github.com/allilengyi/matterbridge/bridge/nctalk"
)

func init() {
	FullMap["nctalk"] = btalk.New
}
