// +build !notelegram

package bridgemap

import (
	btelegram "github.com/allilengyi/matterbridge/bridge/telegram"
)

func init() {
	FullMap["telegram"] = btelegram.New
}
