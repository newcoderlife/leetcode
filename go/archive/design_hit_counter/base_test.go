package design_hit_counter

import (
	"fmt"
	"testing"
)

func TestHitCounter(t *testing.T) {
	counter := Constructor()
	counter.Hit(2)
	counter.Hit(3)
	counter.Hit(4)
	fmt.Println(counter.GetHits(300))
	fmt.Println(counter.GetHits(301))
	fmt.Println(counter.GetHits(302))
	fmt.Println(counter.GetHits(303))
	fmt.Println(counter.GetHits(304))
	counter.Hit(501)
	fmt.Println(counter.GetHits(600))
}
