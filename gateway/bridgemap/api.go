// +build !noapi

package bridgemap

import (
	"github.com/allilengyi/matterbridge/bridge/api"
)

func init() {
	FullMap["api"] = api.New
}
