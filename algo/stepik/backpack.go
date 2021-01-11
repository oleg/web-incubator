package stepik

import (
	"fmt"
	"io"
	"sort"
)

type item struct {
	price, volume, ppi float64
}

type backpack struct {
	volume float64
	items  []item
}

func solveBackpackProblem(r io.Reader, w io.Writer) {
	writePrice(w, calculatePrice(readBackpack(r)))
}

func readBackpack(r io.Reader) backpack {
	var n int
	var v float64
	_, _ = fmt.Fscan(r, &n, &v)
	items := make([]item, n)
	for i := 0; i < n; i++ {
		var price, volume float64
		_, _ = fmt.Fscan(r, &price, &volume)
		items[i] = item{price: price, volume: volume}
	}
	return backpack{
		volume: v,
		items:  items,
	}
}

func writePrice(w io.Writer, price float64) {
	_, _ = fmt.Fprintf(w, "%.3f", price)
}

func calculatePrice(bp backpack) float64 {
	for i := range bp.items {
		v := &bp.items[i]
		v.ppi = v.price / v.volume
	}

	sort.Slice(bp.items, func(i, j int) bool {
		return bp.items[i].ppi > bp.items[j].ppi
	})

	price := 0.0
	volumeLeft := bp.volume
	for _, v := range bp.items {
		if v.volume <= volumeLeft {
			price += v.price
			volumeLeft -= v.volume
		} else {
			price += volumeLeft * v.ppi
			volumeLeft -= volumeLeft
		}
	}
	return price
}
