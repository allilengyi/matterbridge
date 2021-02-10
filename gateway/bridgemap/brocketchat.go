// +build !norocketchat

package bridgemap

import (
	brocketchat "github.com/allilengyi/matterbridge/bridge/rocketchat"
)

func init() {
	FullMap["rocketchat"] = brocketchat.New
}
