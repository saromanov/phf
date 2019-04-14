package phf

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
)

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

// Get returns index by the key
func (p *PHF) Get(s string) int32 {
	size := uint64(len(p.values))
	hash := p.hash(s)
	i := hash & (size - 1)
	seed := p.seeds[i]
	if seed < 0 {
		return p.values[-seed-1]
	}

	return p.values[xorshift64Star(uint64(seed)+hash)&(size-1)]
}

func (p *PHF) hashKeys(keys []string) [][]entry {
	result := make([][]entry, len(keys))
	for i, k := range keys {
		h := p.hash(k)
		result[i] = entry{key: k, value: h}
	}
	return result
}

// hash is a method for hashing string
func (p *PHF) hash(s string) string {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

// nextPower2 returns next power of 2 from input
func nextPower2(n int) int {
	i := 1
	for i < n {
		i *= 2
	}
	return i
}

// implementation of xor shift 64 star algorithm
// https://en.wikipedia.org/wiki/Xorshift#xorshift*
func xorShift64Star(x uint64) uint64 {
	x ^= x >> 12
	x ^= x << 25
	x ^= x >> 27
	return x * 2685821657736338717
}
