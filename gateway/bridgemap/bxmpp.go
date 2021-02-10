// +build !noxmpp

package bridgemap

import (
	bxmpp "github.com/allilengyi/matterbridge/bridge/xmpp"
)

func init() {
	FullMap["xmpp"] = bxmpp.New
}
