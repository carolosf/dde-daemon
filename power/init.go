package power

import "dde-daemon"

func init() {
	loader.Register(&loader.Module{"power", Start, nil, true})
}