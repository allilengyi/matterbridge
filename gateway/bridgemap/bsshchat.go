// +build !nosshchat

package bridgemap

import (
	bsshchat "github.com/allilengyi/matterbridge/bridge/sshchat"
)

func init() {
	FullMap["sshchat"] = bsshchat.New
}
