// +build !nozulip

package bridgemap

import (
	bzulip "github.com/allilengyi/matterbridge/bridge/zulip"
)

func init() {
	FullMap["zulip"] = bzulip.New
}
