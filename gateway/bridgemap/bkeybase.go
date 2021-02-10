// +build !nokeybase

package bridgemap

import (
	bkeybase "github.com/allilengyi/matterbridge/bridge/keybase"
)

func init() {
	FullMap["keybase"] = bkeybase.New
}
