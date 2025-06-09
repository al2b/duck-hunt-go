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
		{0, Vector2D{1.5, 1.5}, DirectionRight},
		{1, Vector2D{1.5, 1}, DirectionRight},
		{2, Vector2D{1.5, 0.5}, DirectionRight},
		{3, Vector2D{1.5, 0}, DirectionRight},
		{4, Vector2D{1.5, -0.5}, DirectionRight},
		{5, Vector2D{1.5, -1}, DirectionRight},
		{6, Vector2D{1.5, -1.5}, DirectionRight},
		{7, Vector2D{1, 1.5}, DirectionRight},
		{8, Vector2D{1, 1}, DirectionRight},
		{9, Vector2D{1, 0.5}, DirectionRight},
		{10, Vector2D{1, 0}, DirectionRight},
		{11, Vector2D{1, -0.5}, DirectionRight},
		{12, Vector2D{1, -1}, DirectionRight},
		{13, Vector2D{1, -1.5}, DirectionRight},
		{14, Vector2D{0.5, 1}, DirectionRight},
		{15, Vector2D{0.5, 0.5}, DirectionRight},
		{16, Vector2D{0.5, 0}, DirectionRight},
		{17, Vector2D{0.5, -0.5}, DirectionRight},
		{18, Vector2D{0.5, -1}, DirectionRight},
		{19, Vector2D{0.5, -1.5}, DirectionRight},
		{20, Vector2D{0, 1}, DirectionRight},
		{21, Vector2D{0, 0.5}, DirectionRight},
		{22, Vector2D{0, 0}, DirectionRight},
		{23, Vector2D{0, -0.5}, DirectionRight},
		{24, Vector2D{0, -1}, DirectionRight},
		{25, Vector2D{0, -1.5}, DirectionRight},
		{26, Vector2D{-0.5, 1}, DirectionLeft},
		{27, Vector2D{-0.5, 0.5}, DirectionLeft},
		{28, Vector2D{-0.5, 0}, DirectionLeft},
		{29, Vector2D{-0.5, -0.5}, DirectionLeft},
		{30, Vector2D{-0.5, -1}, DirectionLeft},
		{31, Vector2D{-0.5, -1.5}, DirectionLeft},
		{32, Vector2D{-1, 1.5}, DirectionLeft},
		{33, Vector2D{-1, 1}, DirectionLeft},
		{34, Vector2D{-1, 0.5}, DirectionLeft},
		{35, Vector2D{-1, 0}, DirectionLeft},
		{36, Vector2D{-1, -0.5}, DirectionLeft},
		{37, Vector2D{-1, -1}, DirectionLeft},
		{38, Vector2D{-1, -1.5}, DirectionLeft},
		{39, Vector2D{-1.5, 1.5}, DirectionLeft},
		{40, Vector2D{-1.5, 1}, DirectionLeft},
		{41, Vector2D{-1.5, 0.5}, DirectionLeft},
		{42, Vector2D{-1.5, 0}, DirectionLeft},
		{43, Vector2D{-1.5, -0.5}, DirectionLeft},
		{44, Vector2D{-1.5, -1}, DirectionLeft},
		{45, Vector2D{-1.5, -1.5}, DirectionLeft},
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
		{0, Vector2D{1.5, 1.5}, DirectionTop},
		{1, Vector2D{1.5, 1}, DirectionTop},
		{2, Vector2D{1.5, 0.5}, DirectionTop},
		{3, Vector2D{1.5, 0}, DirectionTop},
		{4, Vector2D{1.5, -0.5}, DirectionBottom},
		{5, Vector2D{1.5, -1}, DirectionBottom},
		{6, Vector2D{1.5, -1.5}, DirectionBottom},
		{7, Vector2D{1, 1.5}, DirectionTop},
		{8, Vector2D{1, 1}, DirectionTop},
		{9, Vector2D{1, 0.5}, DirectionTop},
		{10, Vector2D{1, 0}, DirectionTop},
		{11, Vector2D{1, -0.5}, DirectionBottom},
		{12, Vector2D{1, -1}, DirectionBottom},
		{13, Vector2D{1, -1.5}, DirectionBottom},
		{14, Vector2D{0.5, 1}, DirectionTop},
		{15, Vector2D{0.5, 0.5}, DirectionTop},
		{16, Vector2D{0.5, 0}, DirectionTop},
		{17, Vector2D{0.5, -0.5}, DirectionBottom},
		{18, Vector2D{0.5, -1}, DirectionBottom},
		{19, Vector2D{0.5, -1.5}, DirectionBottom},
		{20, Vector2D{0, 1}, DirectionTop},
		{21, Vector2D{0, 0.5}, DirectionTop},
		{22, Vector2D{0, 0}, DirectionTop},
		{23, Vector2D{0, -0.5}, DirectionBottom},
		{24, Vector2D{0, -1}, DirectionBottom},
		{25, Vector2D{0, -1.5}, DirectionBottom},
		{26, Vector2D{-0.5, 1}, DirectionTop},
		{27, Vector2D{-0.5, 0.5}, DirectionTop},
		{28, Vector2D{-0.5, 0}, DirectionTop},
		{29, Vector2D{-0.5, -0.5}, DirectionBottom},
		{30, Vector2D{-0.5, -1}, DirectionBottom},
		{31, Vector2D{-0.5, -1.5}, DirectionBottom},
		{32, Vector2D{-1, 1.5}, DirectionTop},
		{33, Vector2D{-1, 1}, DirectionTop},
		{34, Vector2D{-1, 0.5}, DirectionTop},
		{35, Vector2D{-1, 0}, DirectionTop},
		{36, Vector2D{-1, -0.5}, DirectionBottom},
		{37, Vector2D{-1, -1}, DirectionBottom},
		{38, Vector2D{-1, -1.5}, DirectionBottom},
		{39, Vector2D{-1.5, 1.5}, DirectionTop},
		{40, Vector2D{-1.5, 1}, DirectionTop},
		{41, Vector2D{-1.5, 0.5}, DirectionTop},
		{42, Vector2D{-1.5, 0}, DirectionTop},
		{43, Vector2D{-1.5, -0.5}, DirectionBottom},
		{44, Vector2D{-1.5, -1}, DirectionBottom},
		{45, Vector2D{-1.5, -1.5}, DirectionBottom},
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
		{0, Vector2D{1.5, 1.5}, DirectionBottom},
		{1, Vector2D{1.5, 1}, DirectionRight},
		{2, Vector2D{1.5, 0.5}, DirectionRight},
		{3, Vector2D{1.5, 0}, DirectionRight},
		{4, Vector2D{1.5, -0.5}, DirectionRight},
		{5, Vector2D{1.5, -1}, DirectionRight},
		{6, Vector2D{1.5, -1.5}, DirectionTop},
		{7, Vector2D{1, 1.5}, DirectionBottom},
		{8, Vector2D{1, 1}, DirectionBottom},
		{9, Vector2D{1, 0.5}, DirectionRight},
		{10, Vector2D{1, 0}, DirectionRight},
		{11, Vector2D{1, -0.5}, DirectionRight},
		{12, Vector2D{1, -1}, DirectionTop},
		{13, Vector2D{1, -1.5}, DirectionTop},
		{14, Vector2D{0.5, 1}, DirectionBottom},
		{15, Vector2D{0.5, 0.5}, DirectionBottom},
		{16, Vector2D{0.5, 0}, DirectionRight},
		{17, Vector2D{0.5, -0.5}, DirectionTop},
		{18, Vector2D{0.5, -1}, DirectionTop},
		{19, Vector2D{0.5, -1.5}, DirectionTop},
		{20, Vector2D{0, 1}, DirectionBottom},
		{21, Vector2D{0, 0.5}, DirectionBottom},
		{22, Vector2D{0, 0}, DirectionTop},
		{23, Vector2D{0, -0.5}, DirectionTop},
		{24, Vector2D{0, -1}, DirectionTop},
		{25, Vector2D{0, -1.5}, DirectionTop},
		{26, Vector2D{-0.5, 1}, DirectionBottom},
		{27, Vector2D{-0.5, 0.5}, DirectionBottom},
		{28, Vector2D{-0.5, 0}, DirectionLeft},
		{29, Vector2D{-0.5, -0.5}, DirectionTop},
		{30, Vector2D{-0.5, -1}, DirectionTop},
		{31, Vector2D{-0.5, -1.5}, DirectionTop},
		{32, Vector2D{-1, 1.5}, DirectionBottom},
		{33, Vector2D{-1, 1}, DirectionBottom},
		{34, Vector2D{-1, 0.5}, DirectionLeft},
		{35, Vector2D{-1, 0}, DirectionLeft},
		{36, Vector2D{-1, -0.5}, DirectionLeft},
		{37, Vector2D{-1, -1}, DirectionTop},
		{38, Vector2D{-1, -1.5}, DirectionTop},
		{39, Vector2D{-1.5, 1.5}, DirectionBottom},
		{40, Vector2D{-1.5, 1}, DirectionLeft},
		{41, Vector2D{-1.5, 0.5}, DirectionLeft},
		{42, Vector2D{-1.5, 0}, DirectionLeft},
		{43, Vector2D{-1.5, -0.5}, DirectionLeft},
		{44, Vector2D{-1.5, -1}, DirectionLeft},
		{45, Vector2D{-1.5, -1.5}, DirectionTop},
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
		{0, Vector2D{1.5, 1.5}, DirectionTopRight},
		{1, Vector2D{1.5, 1}, DirectionTopRight},
		{2, Vector2D{1.5, 0.5}, DirectionRight},
		{3, Vector2D{1.5, 0}, DirectionRight},
		{4, Vector2D{1.5, -0.5}, DirectionRight},
		{5, Vector2D{1.5, -1}, DirectionBottomRight},
		{6, Vector2D{1.5, -1.5}, DirectionBottomRight},
		{7, Vector2D{1, 1.5}, DirectionTopRight},
		{8, Vector2D{1, 1}, DirectionTopRight},
		{9, Vector2D{1, 0.5}, DirectionRight},
		{10, Vector2D{1, 0}, DirectionRight},
		{11, Vector2D{1, -0.5}, DirectionRight},
		{12, Vector2D{1, -1}, DirectionBottomRight},
		{13, Vector2D{1, -1.5}, DirectionBottomRight},
		{14, Vector2D{0.5, 1}, DirectionTopRight},
		{15, Vector2D{0.5, 0.5}, DirectionTopRight},
		{16, Vector2D{0.5, 0}, DirectionRight},
		{17, Vector2D{0.5, -0.5}, DirectionBottomRight},
		{18, Vector2D{0.5, -1}, DirectionBottomRight},
		{19, Vector2D{0.5, -1.5}, DirectionBottomRight},
		{20, Vector2D{0, 1}, DirectionTopLeft},
		{21, Vector2D{0, 0.5}, DirectionTopLeft},
		{22, Vector2D{0, 0}, DirectionRight},
		{23, Vector2D{0, -0.5}, DirectionBottomRight},
		{24, Vector2D{0, -1}, DirectionBottomRight},
		{25, Vector2D{0, -1.5}, DirectionBottomRight},
		{26, Vector2D{-0.5, 1}, DirectionTopLeft},
		{27, Vector2D{-0.5, 0.5}, DirectionTopLeft},
		{28, Vector2D{-0.5, 0}, DirectionLeft},
		{29, Vector2D{-0.5, -0.5}, DirectionBottomLeft},
		{30, Vector2D{-0.5, -1}, DirectionBottomLeft},
		{31, Vector2D{-0.5, -1.5}, DirectionBottomLeft},
		{32, Vector2D{-1, 1.5}, DirectionTopLeft},
		{33, Vector2D{-1, 1}, DirectionTopLeft},
		{34, Vector2D{-1, 0.5}, DirectionLeft},
		{35, Vector2D{-1, 0}, DirectionLeft},
		{36, Vector2D{-1, -0.5}, DirectionLeft},
		{37, Vector2D{-1, -1}, DirectionBottomLeft},
		{38, Vector2D{-1, -1.5}, DirectionBottomLeft},
		{39, Vector2D{-1.5, 1.5}, DirectionTopLeft},
		{40, Vector2D{-1.5, 1}, DirectionTopLeft},
		{41, Vector2D{-1.5, 0.5}, DirectionLeft},
		{42, Vector2D{-1.5, 0}, DirectionLeft},
		{43, Vector2D{-1.5, -0.5}, DirectionLeft},
		{44, Vector2D{-1.5, -1}, DirectionBottomLeft},
		{45, Vector2D{-1.5, -1.5}, DirectionBottomLeft},
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
		{0, Vector2D{1.5, 1.5}, DirectionTopRight},
		{1, Vector2D{1.5, 1}, DirectionTopRight},
		{2, Vector2D{1.5, 0.5}, DirectionRight},
		{3, Vector2D{1.5, 0}, DirectionRight},
		{4, Vector2D{1.5, -0.5}, DirectionRight},
		{5, Vector2D{1.5, -1}, DirectionBottomRight},
		{6, Vector2D{1.5, -1.5}, DirectionBottomRight},
		{7, Vector2D{1, 1.5}, DirectionTopRight},
		{8, Vector2D{1, 1}, DirectionTopRight},
		{9, Vector2D{1, 0.5}, DirectionTopRight},
		{10, Vector2D{1, 0}, DirectionRight},
		{11, Vector2D{1, -0.5}, DirectionBottomRight},
		{12, Vector2D{1, -1}, DirectionBottomRight},
		{13, Vector2D{1, -1.5}, DirectionBottomRight},
		{14, Vector2D{0.5, 1}, DirectionTopRight},
		{15, Vector2D{0.5, 0.5}, DirectionTopRight},
		{16, Vector2D{0.5, 0}, DirectionRight},
		{17, Vector2D{0.5, -0.5}, DirectionBottomRight},
		{18, Vector2D{0.5, -1}, DirectionBottomRight},
		{19, Vector2D{0.5, -1.5}, DirectionBottom},
		{20, Vector2D{0, 1}, DirectionTop},
		{21, Vector2D{0, 0.5}, DirectionTop},
		{22, Vector2D{0, 0}, DirectionRight},
		{23, Vector2D{0, -0.5}, DirectionBottom},
		{24, Vector2D{0, -1}, DirectionBottom},
		{25, Vector2D{0, -1.5}, DirectionBottom},
		{26, Vector2D{-0.5, 1}, DirectionTopLeft},
		{27, Vector2D{-0.5, 0.5}, DirectionTopLeft},
		{28, Vector2D{-0.5, 0}, DirectionLeft},
		{29, Vector2D{-0.5, -0.5}, DirectionBottomLeft},
		{30, Vector2D{-0.5, -1}, DirectionBottomLeft},
		{31, Vector2D{-0.5, -1.5}, DirectionBottom},
		{32, Vector2D{-1, 1.5}, DirectionTopLeft},
		{33, Vector2D{-1, 1}, DirectionTopLeft},
		{34, Vector2D{-1, 0.5}, DirectionTopLeft},
		{35, Vector2D{-1, 0}, DirectionLeft},
		{36, Vector2D{-1, -0.5}, DirectionBottomLeft},
		{37, Vector2D{-1, -1}, DirectionBottomLeft},
		{38, Vector2D{-1, -1.5}, DirectionBottomLeft},
		{39, Vector2D{-1.5, 1.5}, DirectionTopLeft},
		{40, Vector2D{-1.5, 1}, DirectionTopLeft},
		{41, Vector2D{-1.5, 0.5}, DirectionLeft},
		{42, Vector2D{-1.5, 0}, DirectionLeft},
		{43, Vector2D{-1.5, -0.5}, DirectionLeft},
		{44, Vector2D{-1.5, -1}, DirectionBottomLeft},
		{45, Vector2D{-1.5, -1.5}, DirectionBottomLeft},
	}
	for _, test := range tests {
		directioner := OctantDirectioner{test.vector}
		t.Run(fmt.Sprintf("%d", test.index), func(t *testing.T) {
			assert.Equal(t, test.want, directioner.Direction())
		})
	}
}
