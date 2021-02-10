// +build !nomsteams

package bridgemap

import (
	bmsteams "github.com/allilengyi/matterbridge/bridge/msteams"
)

func init() {
	FullMap["msteams"] = bmsteams.New
}
