package layout

import "duck-hunt-go/engine"

var (
	bodyShape = engine.BodyShape{
		{0, 0},
		{width - 1, 0},
		{width - 1, ground - 1},
		{0, ground - 1},
	}
	// Tree
	treeBodyShapes = []engine.BodyShape{
		{
			{0, 0},
			{68, 0},
			{68, 150},
			{0, 150},
		},
		{
			{22, 0},
			{26, 0},
			{33, 6},
			{36, 11},
			{35, 17},
			{33, 20},
			{29, 22},
			{15, 22},
			{11, 20},
			{7, 15},
			{7, 10},
			{8, 7},
			{13, 2},
		},
		{
			{42, 24},
			{51, 25},
			{57, 30},
			{57, 34},
			{56, 39},
			{52, 40},
			{51, 42},
			{46, 43},
			{39, 41},
			{35, 39},
			{32, 36},
			{32, 31},
			{33, 28},
			{37, 25},
		},
		{
			{31, 42},
			{35, 42},
			{38, 43},
			{42, 46},
			{43, 51},
			{40, 55},
			{32, 56},
			{25, 54},
			{24, 50},
			{26, 45},
		},
		{
			{8, 48},
			{14, 48},
			{18, 49},
			{22, 55},
			{22, 60},
			{19, 65},
			{11, 67},
			{7, 65},
			{3, 63},
			{0, 60},
			{0, 55},
			{1, 52},
			{4, 50},
		},
		{
			{54, 56},
			{58, 56},
			{64, 61},
			{67, 66},
			{68, 68},
			{67, 73},
			{65, 76},
			{59, 79},
			{47, 78},
			{43, 76},
			{39, 71},
			{39, 66},
			{40, 64},
			{42, 61},
			{45, 58},
		},
	}
	// Shrub
	shrubBodyShape = engine.BodyShape{
		{0, 60},
		{0, 29},
		{1, 22},
		{3, 16},
		{7, 11},
		{8, 7},
		{9, 4},
		{16, 0},
		{23, 2},
		{25, 4},
		{29, 15},
		{30, 25},
		{30, 60},
	}
)
