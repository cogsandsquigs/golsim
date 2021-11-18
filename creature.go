package main

import (
	"math/rand"
	"time"
)

type Creature struct {
	pos    [2]uint64
	genome []uint64
}

func RandomCreature(pwidth, plen, genomelen uint) Creature {
	rand.Seed(time.Now().UnixNano())
	c := Creature{[2]uint64{uint64(rand.Int63n(int64(pwidth))), uint64(rand.Int63n(int64(plen)))}, []uint64{}}
	for genomelen > 0 {
		c.genome = append(c.genome, rand.Uint64())
		genomelen -= 1
	}

	return c
}
