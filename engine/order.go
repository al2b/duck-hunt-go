package engine

import (
	"slices"
)

type Order int

func (o Order) Order() Order {
	return o
}

type Orderer interface {
	Order() Order
}

type OrderDrawer interface {
	Drawer
	Orderer
}

type OrderDrawers []OrderDrawer

func (d OrderDrawers) Draw(img *Image) {
	slices.SortStableFunc(d, func(a, b OrderDrawer) int {
		return int(a.Order()) - int(b.Order())
	})
	slices.Reverse(d)
	for _, drawer := range d {
		drawer.Draw(img)
	}
}

type OrderedDrawer struct {
	Drawer
	Orderer
}

func (pd OrderedDrawer) Draw(img *Image) {
	if pd.Drawer == nil {
		return
	}
	pd.Drawer.Draw(img)
}

func (pd OrderedDrawer) Order() Order {
	if pd.Orderer == nil {
		return Order(0)
	}
	return pd.Orderer.Order()
}
