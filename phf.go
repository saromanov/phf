package phf

import "errors"

var errNoKeys = errors.New("keys is not defined")

// PHF defines main struct
type PHF struct {
	values []int32
}

type entry struct {
	key   string
	index int32
	value uint32
}

// New creates new app
func New() *PHF {
	return &PHF{
		values: []int32{},
	}
}

// Add provides adding of the keys
func (p *PHF) Add(keys []string) error {
	if len(keys) == 0 {
		return errNoKeys
	}
	return nil
}

func (p *PHF) hashKeys(keys []string) [][]entry {
	result := make([][]entry, len(keys))
	for i, k := range keys {
		result[i] = entry{key: k}
	}
	return entry
}
