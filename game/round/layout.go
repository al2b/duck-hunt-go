package round

import (
	"duck-hunt-go/engine"
	enginecp "duck-hunt-go/engine-cp"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/config"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/jakecoffman/cp/v2"
)

/**********/
/* Layout */
/**********/

func NewLayout(space *cp.Space) *Layout {
	return &Layout{
		space: space,
	}
}

type Layout struct {
	space *cp.Space
	engine.OrderedDrawer
}

func (m *Layout) Init() tea.Cmd {
	// Space
	borders := []cp.Vector{
		{0, 0}, {255, 0},
		{255, 0}, {255, config.Ground},
		{255, config.Ground}, {0, config.Ground},
		{0, config.Ground}, {0, 0},
	}

	for i := 0; i < len(borders)-1; i += 2 {
		shape := m.space.AddShape(cp.NewSegment(m.space.StaticBody, borders[i], borders[i+1], 0))
		shape.SetElasticity(1)
		shape.SetFriction(0)
	}

	// Drawer
	m.OrderedDrawer.Drawer = engine.ImageDrawer{
		engine.Pt(0, 0),
		assets.Layout,
	}
	m.OrderedDrawer.Orderer = engine.Order(0)

	return nil
}

func (m *Layout) Update(_ tea.Msg) tea.Cmd {
	return nil
}

/*****************/
/* Layout - Tree */
/*****************/

func NewLayoutTree(space *cp.Space) *LayoutTree {
	return &LayoutTree{
		space: space,
	}
}

type LayoutTree struct {
	space *cp.Space
	engine.ImageDrawer
}

func (m *LayoutTree) Init() tea.Cmd {
	// Space
	body := m.space.AddBody(cp.NewKinematicBody())
	body.SetPosition(cp.Vector{6, 32})

	vertices := [][]cp.Vector{
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

	for _, vertices := range vertices {
		shape := m.space.AddShape(
			cp.NewPolyShapeRaw(body, len(vertices), vertices, 0),
		)
		shape.SetFilter(cp.NewShapeFilter(cp.NO_GROUP, config.ShapeCategoryLayout, config.ShapeCategoryDuck))
		shape.SetElasticity(1)
		shape.SetFriction(0)
	}

	// Drawer
	m.ImageDrawer.Pointer = enginecp.PositionPointer{body}
	m.ImageDrawer.Imager = assets.LayoutTree

	return nil
}

func (m *LayoutTree) Update(_ tea.Msg) tea.Cmd {
	return nil
}

/*****************/
/* Layout - Bush */
/*****************/

func NewLayoutBush(space *cp.Space) *LayoutBush {
	return &LayoutBush{
		space: space,
	}
}

type LayoutBush struct {
	space *cp.Space
	engine.ImageDrawer
}

func (m *LayoutBush) Init() tea.Cmd {
	// Space
	body := m.space.AddBody(cp.NewKinematicBody())
	body.SetPosition(cp.Vector{193, 122})

	vertices := []cp.Vector{
		{0, 61},
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
		{30, 61},
	}

	shape := m.space.AddShape(
		cp.NewPolyShapeRaw(body, len(vertices), vertices, 0),
	)
	shape.SetFilter(cp.NewShapeFilter(cp.NO_GROUP, config.ShapeCategoryLayout, config.ShapeCategoryDuck))
	shape.SetElasticity(1)
	shape.SetFriction(0)

	// Drawer
	m.ImageDrawer.Pointer = enginecp.PositionPointer{body}
	m.ImageDrawer.Imager = assets.LayoutBush

	return nil
}

func (m *LayoutBush) Update(_ tea.Msg) tea.Cmd {
	return nil
}
