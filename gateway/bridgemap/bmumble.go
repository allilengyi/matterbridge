// +build !nomumble

package bridgemap

import (
	bmumble "github.com/allilengyi/matterbridge/bridge/mumble"
)

func init() {
	FullMap["mumble"] = bmumble.New
}
