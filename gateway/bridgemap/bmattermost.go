// +build !nomattermost

package bridgemap

import (
	bmattermost "github.com/allilengyi/matterbridge/bridge/mattermost"
)

func init() {
	FullMap["mattermost"] = bmattermost.New
}
