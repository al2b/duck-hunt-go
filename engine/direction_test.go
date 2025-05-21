package engine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirection(t *testing.T) {
	tests := []struct {
		index     int
		direction Direction
		want      string
	}{
		{0, DirectionRight, "right"},
		{1, DirectionTopRight, "top right"},
		{2, DirectionTop, "top"},
		{3, DirectionTopLeft, "top left"},
		{4, DirectionLeft, "left"},
		{5, DirectionBottomLeft, "bottom left"},
		{6, DirectionBottom, "bottom"},
		{7, DirectionBottomRight, "bottom right"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.index), func(t *testing.T) {
			assert.Equal(t, test.want, test.direction.String())
		})
	}
}

func TestVerticalSemicircleDirectioner(t *testing.T) {
	tests := []struct {
		index  int
		vector Vector2D
		want   Direction
	}{
		{0, Vec2D(1.5, 1.5), DirectionRight},
		{1, Vec2D(1.5, 1), DirectionRight},
		{2, Vec2D(1.5, 0.5), DirectionRight},
		{3, Vec2D(1.5, 0), DirectionRight},
		{4, Vec2D(1.5, -0.5), DirectionRight},
		{5, Vec2D(1.5, -1), DirectionRight},
		{6, Vec2D(1.5, -1.5), DirectionRight},
		{7, Vec2D(1, 1.5), DirectionRight},
		{8, Vec2D(1, 1), DirectionRight},
		{9, Vec2D(1, 0.5), DirectionRight},
		{10, Vec2D(1, 0), DirectionRight},
		{11, Vec2D(1, -0.5), DirectionRight},
		{12, Vec2D(1, -1), DirectionRight},
		{13, Vec2D(1, -1.5), DirectionRight},
		{14, Vec2D(0.5, 1), DirectionRight},
		{15, Vec2D(0.5, 0.5), DirectionRight},
		{16, Vec2D(0.5, 0), DirectionRight},
		{17, Vec2D(0.5, -0.5), DirectionRight},
		{18, Vec2D(0.5, -1), DirectionRight},
		{19, Vec2D(0.5, -1.5), DirectionRight},
		{20, Vec2D(0, 1), DirectionRight},
		{21, Vec2D(0, 0.5), DirectionRight},
		{22, Vec2D(0, 0), DirectionRight},
		{23, Vec2D(0, -0.5), DirectionRight},
		{24, Vec2D(0, -1), DirectionRight},
		{25, Vec2D(0, -1.5), DirectionRight},
		{26, Vec2D(-0.5, 1), DirectionLeft},
		{27, Vec2D(-0.5, 0.5), DirectionLeft},
		{28, Vec2D(-0.5, 0), DirectionLeft},
		{29, Vec2D(-0.5, -0.5), DirectionLeft},
		{30, Vec2D(-0.5, -1), DirectionLeft},
		{31, Vec2D(-0.5, -1.5), DirectionLeft},
		{32, Vec2D(-1, 1.5), DirectionLeft},
		{33, Vec2D(-1, 1), DirectionLeft},
		{34, Vec2D(-1, 0.5), DirectionLeft},
		{35, Vec2D(-1, 0), DirectionLeft},
		{36, Vec2D(-1, -0.5), DirectionLeft},
		{37, Vec2D(-1, -1), DirectionLeft},
		{38, Vec2D(-1, -1.5), DirectionLeft},
		{39, Vec2D(-1.5, 1.5), DirectionLeft},
		{40, Vec2D(-1.5, 1), DirectionLeft},
		{41, Vec2D(-1.5, 0.5), DirectionLeft},
		{42, Vec2D(-1.5, 0), DirectionLeft},
		{43, Vec2D(-1.5, -0.5), DirectionLeft},
		{44, Vec2D(-1.5, -1), DirectionLeft},
		{45, Vec2D(-1.5, -1.5), DirectionLeft},
	}
	for _, test := range tests {
		directioner := VerticalSemicircleDirectioner{test.vector}
		t.Run(fmt.Sprintf("%d", test.index), func(t *testing.T) {
			assert.Equal(t, test.want, directioner.Direction())
		})
	}
}

func TestHorizontalSemicircleDirectioner(t *testing.T) {
	tests := []struct {
		index  int
		vector Vector2D
		want   Direction
	}{
		{0, Vec2D(1.5, 1.5), DirectionTop},
		{1, Vec2D(1.5, 1), DirectionTop},
		{2, Vec2D(1.5, 0.5), DirectionTop},
		{3, Vec2D(1.5, 0), DirectionTop},
		{4, Vec2D(1.5, -0.5), DirectionBottom},
		{5, Vec2D(1.5, -1), DirectionBottom},
		{6, Vec2D(1.5, -1.5), DirectionBottom},
		{7, Vec2D(1, 1.5), DirectionTop},
		{8, Vec2D(1, 1), DirectionTop},
		{9, Vec2D(1, 0.5), DirectionTop},
		{10, Vec2D(1, 0), DirectionTop},
		{11, Vec2D(1, -0.5), DirectionBottom},
		{12, Vec2D(1, -1), DirectionBottom},
		{13, Vec2D(1, -1.5), DirectionBottom},
		{14, Vec2D(0.5, 1), DirectionTop},
		{15, Vec2D(0.5, 0.5), DirectionTop},
		{16, Vec2D(0.5, 0), DirectionTop},
		{17, Vec2D(0.5, -0.5), DirectionBottom},
		{18, Vec2D(0.5, -1), DirectionBottom},
		{19, Vec2D(0.5, -1.5), DirectionBottom},
		{20, Vec2D(0, 1), DirectionTop},
		{21, Vec2D(0, 0.5), DirectionTop},
		{22, Vec2D(0, 0), DirectionTop},
		{23, Vec2D(0, -0.5), DirectionBottom},
		{24, Vec2D(0, -1), DirectionBottom},
		{25, Vec2D(0, -1.5), DirectionBottom},
		{26, Vec2D(-0.5, 1), DirectionTop},
		{27, Vec2D(-0.5, 0.5), DirectionTop},
		{28, Vec2D(-0.5, 0), DirectionTop},
		{29, Vec2D(-0.5, -0.5), DirectionBottom},
		{30, Vec2D(-0.5, -1), DirectionBottom},
		{31, Vec2D(-0.5, -1.5), DirectionBottom},
		{32, Vec2D(-1, 1.5), DirectionTop},
		{33, Vec2D(-1, 1), DirectionTop},
		{34, Vec2D(-1, 0.5), DirectionTop},
		{35, Vec2D(-1, 0), DirectionTop},
		{36, Vec2D(-1, -0.5), DirectionBottom},
		{37, Vec2D(-1, -1), DirectionBottom},
		{38, Vec2D(-1, -1.5), DirectionBottom},
		{39, Vec2D(-1.5, 1.5), DirectionTop},
		{40, Vec2D(-1.5, 1), DirectionTop},
		{41, Vec2D(-1.5, 0.5), DirectionTop},
		{42, Vec2D(-1.5, 0), DirectionTop},
		{43, Vec2D(-1.5, -0.5), DirectionBottom},
		{44, Vec2D(-1.5, -1), DirectionBottom},
		{45, Vec2D(-1.5, -1.5), DirectionBottom},
	}
	for _, test := range tests {
		directioner := HorizontalSemicircleDirectioner{test.vector}
		t.Run(fmt.Sprintf("%d", test.index), func(t *testing.T) {
			assert.Equal(t, test.want, directioner.Direction())
		})
	}
}

func TestQuadrantDirectioner(t *testing.T) {
	tests := []struct {
		index  int
		vector Vector2D
		want   Direction
	}{
		{0, Vec2D(1.5, 1.5), DirectionBottom},
		{1, Vec2D(1.5, 1), DirectionRight},
		{2, Vec2D(1.5, 0.5), DirectionRight},
		{3, Vec2D(1.5, 0), DirectionRight},
		{4, Vec2D(1.5, -0.5), DirectionRight},
		{5, Vec2D(1.5, -1), DirectionRight},
		{6, Vec2D(1.5, -1.5), DirectionTop},
		{7, Vec2D(1, 1.5), DirectionBottom},
		{8, Vec2D(1, 1), DirectionBottom},
		{9, Vec2D(1, 0.5), DirectionRight},
		{10, Vec2D(1, 0), DirectionRight},
		{11, Vec2D(1, -0.5), DirectionRight},
		{12, Vec2D(1, -1), DirectionTop},
		{13, Vec2D(1, -1.5), DirectionTop},
		{14, Vec2D(0.5, 1), DirectionBottom},
		{15, Vec2D(0.5, 0.5), DirectionBottom},
		{16, Vec2D(0.5, 0), DirectionRight},
		{17, Vec2D(0.5, -0.5), DirectionTop},
		{18, Vec2D(0.5, -1), DirectionTop},
		{19, Vec2D(0.5, -1.5), DirectionTop},
		{20, Vec2D(0, 1), DirectionBottom},
		{21, Vec2D(0, 0.5), DirectionBottom},
		{22, Vec2D(0, 0), DirectionTop},
		{23, Vec2D(0, -0.5), DirectionTop},
		{24, Vec2D(0, -1), DirectionTop},
		{25, Vec2D(0, -1.5), DirectionTop},
		{26, Vec2D(-0.5, 1), DirectionBottom},
		{27, Vec2D(-0.5, 0.5), DirectionBottom},
		{28, Vec2D(-0.5, 0), DirectionLeft},
		{29, Vec2D(-0.5, -0.5), DirectionTop},
		{30, Vec2D(-0.5, -1), DirectionTop},
		{31, Vec2D(-0.5, -1.5), DirectionTop},
		{32, Vec2D(-1, 1.5), DirectionBottom},
		{33, Vec2D(-1, 1), DirectionBottom},
		{34, Vec2D(-1, 0.5), DirectionLeft},
		{35, Vec2D(-1, 0), DirectionLeft},
		{36, Vec2D(-1, -0.5), DirectionLeft},
		{37, Vec2D(-1, -1), DirectionTop},
		{38, Vec2D(-1, -1.5), DirectionTop},
		{39, Vec2D(-1.5, 1.5), DirectionBottom},
		{40, Vec2D(-1.5, 1), DirectionLeft},
		{41, Vec2D(-1.5, 0.5), DirectionLeft},
		{42, Vec2D(-1.5, 0), DirectionLeft},
		{43, Vec2D(-1.5, -0.5), DirectionLeft},
		{44, Vec2D(-1.5, -1), DirectionLeft},
		{45, Vec2D(-1.5, -1.5), DirectionTop},
	}
	for _, test := range tests {
		directioner := QuadrantDirectioner{test.vector}
		t.Run(fmt.Sprintf("%d", test.index), func(t *testing.T) {
			assert.Equal(t, test.want, directioner.Direction())
		})
	}
}

func TestVerticalSextantDirectioner(t *testing.T) {
	tests := []struct {
		index  int
		vector Vector2D
		want   Direction
	}{
		{0, Vec2D(1.5, 1.5), DirectionTopRight},
		{1, Vec2D(1.5, 1), DirectionTopRight},
		{2, Vec2D(1.5, 0.5), DirectionRight},
		{3, Vec2D(1.5, 0), DirectionRight},
		{4, Vec2D(1.5, -0.5), DirectionRight},
		{5, Vec2D(1.5, -1), DirectionBottomRight},
		{6, Vec2D(1.5, -1.5), DirectionBottomRight},
		{7, Vec2D(1, 1.5), DirectionTopRight},
		{8, Vec2D(1, 1), DirectionTopRight},
		{9, Vec2D(1, 0.5), DirectionRight},
		{10, Vec2D(1, 0), DirectionRight},
		{11, Vec2D(1, -0.5), DirectionRight},
		{12, Vec2D(1, -1), DirectionBottomRight},
		{13, Vec2D(1, -1.5), DirectionBottomRight},
		{14, Vec2D(0.5, 1), DirectionTopRight},
		{15, Vec2D(0.5, 0.5), DirectionTopRight},
		{16, Vec2D(0.5, 0), DirectionRight},
		{17, Vec2D(0.5, -0.5), DirectionBottomRight},
		{18, Vec2D(0.5, -1), DirectionBottomRight},
		{19, Vec2D(0.5, -1.5), DirectionBottomRight},
		{20, Vec2D(0, 1), DirectionTopLeft},
		{21, Vec2D(0, 0.5), DirectionTopLeft},
		{22, Vec2D(0, 0), DirectionRight},
		{23, Vec2D(0, -0.5), DirectionBottomRight},
		{24, Vec2D(0, -1), DirectionBottomRight},
		{25, Vec2D(0, -1.5), DirectionBottomRight},
		{26, Vec2D(-0.5, 1), DirectionTopLeft},
		{27, Vec2D(-0.5, 0.5), DirectionTopLeft},
		{28, Vec2D(-0.5, 0), DirectionLeft},
		{29, Vec2D(-0.5, -0.5), DirectionBottomLeft},
		{30, Vec2D(-0.5, -1), DirectionBottomLeft},
		{31, Vec2D(-0.5, -1.5), DirectionBottomLeft},
		{32, Vec2D(-1, 1.5), DirectionTopLeft},
		{33, Vec2D(-1, 1), DirectionTopLeft},
		{34, Vec2D(-1, 0.5), DirectionLeft},
		{35, Vec2D(-1, 0), DirectionLeft},
		{36, Vec2D(-1, -0.5), DirectionLeft},
		{37, Vec2D(-1, -1), DirectionBottomLeft},
		{38, Vec2D(-1, -1.5), DirectionBottomLeft},
		{39, Vec2D(-1.5, 1.5), DirectionTopLeft},
		{40, Vec2D(-1.5, 1), DirectionTopLeft},
		{41, Vec2D(-1.5, 0.5), DirectionLeft},
		{42, Vec2D(-1.5, 0), DirectionLeft},
		{43, Vec2D(-1.5, -0.5), DirectionLeft},
		{44, Vec2D(-1.5, -1), DirectionBottomLeft},
		{45, Vec2D(-1.5, -1.5), DirectionBottomLeft},
	}
	for _, test := range tests {
		directioner := VerticalSextantDirectioner{test.vector}
		t.Run(fmt.Sprintf("%d", test.index), func(t *testing.T) {
			assert.Equal(t, test.want, directioner.Direction())
		})
	}
}

func TestOctantDirectioner(t *testing.T) {
	tests := []struct {
		index  int
		vector Vector2D
		want   Direction
	}{
		{0, Vec2D(1.5, 1.5), DirectionTopRight},
		{1, Vec2D(1.5, 1), DirectionTopRight},
		{2, Vec2D(1.5, 0.5), DirectionRight},
		{3, Vec2D(1.5, 0), DirectionRight},
		{4, Vec2D(1.5, -0.5), DirectionRight},
		{5, Vec2D(1.5, -1), DirectionBottomRight},
		{6, Vec2D(1.5, -1.5), DirectionBottomRight},
		{7, Vec2D(1, 1.5), DirectionTopRight},
		{8, Vec2D(1, 1), DirectionTopRight},
		{9, Vec2D(1, 0.5), DirectionTopRight},
		{10, Vec2D(1, 0), DirectionRight},
		{11, Vec2D(1, -0.5), DirectionBottomRight},
		{12, Vec2D(1, -1), DirectionBottomRight},
		{13, Vec2D(1, -1.5), DirectionBottomRight},
		{14, Vec2D(0.5, 1), DirectionTopRight},
		{15, Vec2D(0.5, 0.5), DirectionTopRight},
		{16, Vec2D(0.5, 0), DirectionRight},
		{17, Vec2D(0.5, -0.5), DirectionBottomRight},
		{18, Vec2D(0.5, -1), DirectionBottomRight},
		{19, Vec2D(0.5, -1.5), DirectionBottom},
		{20, Vec2D(0, 1), DirectionTop},
		{21, Vec2D(0, 0.5), DirectionTop},
		{22, Vec2D(0, 0), DirectionRight},
		{23, Vec2D(0, -0.5), DirectionBottom},
		{24, Vec2D(0, -1), DirectionBottom},
		{25, Vec2D(0, -1.5), DirectionBottom},
		{26, Vec2D(-0.5, 1), DirectionTopLeft},
		{27, Vec2D(-0.5, 0.5), DirectionTopLeft},
		{28, Vec2D(-0.5, 0), DirectionLeft},
		{29, Vec2D(-0.5, -0.5), DirectionBottomLeft},
		{30, Vec2D(-0.5, -1), DirectionBottomLeft},
		{31, Vec2D(-0.5, -1.5), DirectionBottom},
		{32, Vec2D(-1, 1.5), DirectionTopLeft},
		{33, Vec2D(-1, 1), DirectionTopLeft},
		{34, Vec2D(-1, 0.5), DirectionTopLeft},
		{35, Vec2D(-1, 0), DirectionLeft},
		{36, Vec2D(-1, -0.5), DirectionBottomLeft},
		{37, Vec2D(-1, -1), DirectionBottomLeft},
		{38, Vec2D(-1, -1.5), DirectionBottomLeft},
		{39, Vec2D(-1.5, 1.5), DirectionTopLeft},
		{40, Vec2D(-1.5, 1), DirectionTopLeft},
		{41, Vec2D(-1.5, 0.5), DirectionLeft},
		{42, Vec2D(-1.5, 0), DirectionLeft},
		{43, Vec2D(-1.5, -0.5), DirectionLeft},
		{44, Vec2D(-1.5, -1), DirectionBottomLeft},
		{45, Vec2D(-1.5, -1.5), DirectionBottomLeft},
	}
	for _, test := range tests {
		directioner := OctantDirectioner{test.vector}
		t.Run(fmt.Sprintf("%d", test.index), func(t *testing.T) {
			assert.Equal(t, test.want, directioner.Direction())
		})
	}
}
