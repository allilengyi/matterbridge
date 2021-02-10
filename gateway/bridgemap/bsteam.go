// +build !nosteam

package bridgemap

import (
	bsteam "github.com/allilengyi/matterbridge/bridge/steam"
)

func init() {
	FullMap["steam"] = bsteam.New
}
