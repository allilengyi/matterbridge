// +build !nowhatsapp

package bridgemap

import (
	bwhatsapp "github.com/allilengyi/matterbridge/bridge/whatsapp"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
