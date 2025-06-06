package config

import "time"

const TickInterval = time.Second / 60

const Ground = 183

var Debug bool

const (
	ShapeCategoryLayout = 1 << 0 // 0001
	ShapeCategoryDuck   = 1 << 1 // 0010
)
