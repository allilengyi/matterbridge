// +build !novk

package bridgemap

import (
	bvk "github.com/allilengyi/matterbridge/bridge/vk"
)

func init() {
	FullMap["vk"] = bvk.New
}
