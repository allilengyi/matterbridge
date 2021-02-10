// +build !nodiscord

package bridgemap

import (
	bdiscord "github.com/allilengyi/matterbridge/bridge/discord"
)

func init() {
	FullMap["discord"] = bdiscord.New
	UserTypingSupport["discord"] = struct{}{}
}
