// +build !nomatrix

package bridgemap

import (
	bmatrix "github.com/allilengyi/matterbridge/bridge/matrix"
)

func init() {
	FullMap["matrix"] = bmatrix.New
}
