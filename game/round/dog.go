package round

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/config"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewDog() *Dog {
	return &Dog{}
}

type Dog struct {
	cinematic engine.Cinematic3DPlayer
	engine.OrderedDrawer
}

func (m *Dog) Init() tea.Cmd {
	// Cinematic
	m.cinematic.Cinematic = engine.SequenceCinematic3D{
		dogCinematicTrack,
		dogCinematicMock,
		dogCinematicRetrieve1,
		dogCinematicRetrieve2,
	}
	m.cinematic.OnEnd = engine.PlayerOnEndLoop
	m.cinematic.Play()

	// Drawer
	m.OrderedDrawer.Drawer = engine.ImageDrawer{
		engine.Position2DPointer{
			engine.Position3DProjector{&m.cinematic, engine.OrthographicProjector{}},
		},
		&m.cinematic,
	}
	m.OrderedDrawer.Orderer = engine.Position3DOrderer{&m.cinematic}

	return nil
}

func (m *Dog) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		// Cinematic
		m.cinematic.Step(msg.Interval)
	}

	return nil
}

/**************/
/* Cinematics */
/**************/

var dogCinematicTrack = engine.Cinematic3D{
	{engine.Vector3D{2, 141, -10}, assets.DogTrack1, 7 * config.TickInterval},
	{engine.Vector3D{4, 141, -10}, assets.DogTrack2, 7 * config.TickInterval},
	{engine.Vector3D{6, 141, -10}, assets.DogTrack3, 7 * config.TickInterval},
	{engine.Vector3D{8, 141, -10}, assets.DogTrack4, 7 * config.TickInterval},

	{engine.Vector3D{10, 141, -10}, assets.DogTrack1, 7 * config.TickInterval},
	{engine.Vector3D{12, 141, -10}, assets.DogTrack2, 7 * config.TickInterval},
	{engine.Vector3D{14, 141, -10}, assets.DogTrack3, 7 * config.TickInterval},
	{engine.Vector3D{16, 141, -10}, assets.DogTrack4, 7 * config.TickInterval},

	{engine.Vector3D{18, 141, -10}, assets.DogTrack1, 7 * config.TickInterval},
	{engine.Vector3D{20, 141, -10}, assets.DogTrack2, 7 * config.TickInterval},
	{engine.Vector3D{22, 141, -10}, assets.DogTrack3, 7 * config.TickInterval},
	{engine.Vector3D{24, 141, -10}, assets.DogTrack4, 7 * config.TickInterval},

	{engine.Vector3D{26, 141, -10}, assets.DogTrack1, 7 * config.TickInterval},
	{engine.Vector3D{28, 141, -10}, assets.DogTrack2, 7 * config.TickInterval},
	{engine.Vector3D{30, 141, -10}, assets.DogTrack3, 7 * config.TickInterval},
	{engine.Vector3D{32, 141, -10}, assets.DogTrack4, 7 * config.TickInterval},

	{engine.Vector3D{34, 141, -10}, assets.DogTrack1, 14 * config.TickInterval},

	{engine.Vector3D{34, 141, -10}, assets.DogSniff, 9 * config.TickInterval},
	{engine.Vector3D{34, 141, -10}, assets.DogTrack1, 9 * config.TickInterval},

	{engine.Vector3D{34, 141, -10}, assets.DogSniff, 9 * config.TickInterval},
	{engine.Vector3D{34, 141, -10}, assets.DogTrack1, 9 * config.TickInterval},

	{engine.Vector3D{34, 141, -10}, assets.DogSniff, 10 * config.TickInterval},

	{engine.Vector3D{36, 141, -10}, assets.DogTrack1, 7 * config.TickInterval},
	{engine.Vector3D{38, 141, -10}, assets.DogTrack2, 7 * config.TickInterval},
	{engine.Vector3D{40, 141, -10}, assets.DogTrack3, 7 * config.TickInterval},
	{engine.Vector3D{42, 141, -10}, assets.DogTrack4, 7 * config.TickInterval},

	{engine.Vector3D{44, 141, -10}, assets.DogTrack1, 7 * config.TickInterval},
	{engine.Vector3D{46, 141, -10}, assets.DogTrack2, 7 * config.TickInterval},
	{engine.Vector3D{48, 141, -10}, assets.DogTrack3, 7 * config.TickInterval},
	{engine.Vector3D{50, 141, -10}, assets.DogTrack4, 7 * config.TickInterval},

	{engine.Vector3D{52, 141, -10}, assets.DogTrack1, 7 * config.TickInterval},
	{engine.Vector3D{54, 141, -10}, assets.DogTrack2, 7 * config.TickInterval},
	{engine.Vector3D{56, 141, -10}, assets.DogTrack3, 7 * config.TickInterval},
	{engine.Vector3D{58, 141, -10}, assets.DogTrack4, 7 * config.TickInterval},

	{engine.Vector3D{60, 141, -10}, assets.DogTrack1, 7 * config.TickInterval},
	{engine.Vector3D{62, 141, -10}, assets.DogTrack2, 7 * config.TickInterval},
	{engine.Vector3D{64, 141, -10}, assets.DogTrack3, 7 * config.TickInterval},
	{engine.Vector3D{66, 141, -10}, assets.DogTrack4, 7 * config.TickInterval},

	{engine.Vector3D{68, 141, -10}, assets.DogTrack1, 14 * config.TickInterval},

	{engine.Vector3D{68, 141, -10}, assets.DogSniff, 9 * config.TickInterval},
	{engine.Vector3D{68, 141, -10}, assets.DogTrack1, 9 * config.TickInterval},

	{engine.Vector3D{68, 141, -10}, assets.DogSniff, 9 * config.TickInterval},
	{engine.Vector3D{68, 141, -10}, assets.DogTrack1, 9 * config.TickInterval},

	{engine.Vector3D{68, 141, -10}, assets.DogSniff, 10 * config.TickInterval},

	{engine.Vector3D{70, 141, -10}, assets.DogTrack1, 7 * config.TickInterval},
	{engine.Vector3D{72, 141, -10}, assets.DogTrack2, 1 * config.TickInterval},
	{engine.Vector3D{72, 141, -10}, assets.DogPant, 18 * config.TickInterval},

	{engine.Vector3D{85, 125, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{86, 122, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{87, 119, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{88, 116, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{89, 113, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{90, 111, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{91, 109, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{92, 107, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{93, 105, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{94, 104, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{95, 102, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{96, 101, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{97, 100, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{98, 99, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{99, 98, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{100, 97, -10}, assets.DogJump1, 1 * config.TickInterval},
	{engine.Vector3D{101, 96, -10}, assets.DogJump1, 1 * config.TickInterval},

	{engine.Vector3D{102, 92, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{103, 92, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{104, 91, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{105, 91, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{106, 90, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{107, 90, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{108, 90, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{109, 90, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{110, 89, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{111, 89, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{112, 89, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{113, 89, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{114, 89, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{115, 90, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{116, 90, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{117, 91, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{118, 91, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{119, 92, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{120, 93, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{121, 94, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{122, 96, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{123, 98, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{124, 100, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{125, 103, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{126, 106, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{126, 110, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{126, 114, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{126, 119, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{126, 124, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{126, 129, 10}, assets.DogJump2, 1 * config.TickInterval},
	{engine.Vector3D{126, 134, 10}, assets.DogJump2, 1 * config.TickInterval},
}

var dogCinematicMock = engine.Cinematic3D{
	{engine.Vector3D{114, 154, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 153, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 153, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 152, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 151, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 150, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 149, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 149, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 148, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 147, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 146, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 145, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 145, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 144, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 143, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 142, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 141, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 141, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 140, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 139, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 138, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 137, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 137, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 136, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 135, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 134, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 133, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 133, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 132, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 131, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 130, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 129, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 129, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 128, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 127, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 126, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 125, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 125, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 124, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 123, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 122, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 121, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 121, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 120, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 119, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 118, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 117, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 117, 10}, assets.DogMock1, 5 * config.TickInterval},
	{engine.Vector3D{114, 117, 10}, assets.DogMock2, 5 * config.TickInterval},
	{engine.Vector3D{114, 117, 10}, assets.DogMock1, 5 * config.TickInterval},
	{engine.Vector3D{114, 117, 10}, assets.DogMock2, 5 * config.TickInterval},
	{engine.Vector3D{114, 117, 10}, assets.DogMock1, 5 * config.TickInterval},
	{engine.Vector3D{114, 117, 10}, assets.DogMock2, 5 * config.TickInterval},
	{engine.Vector3D{114, 117, 10}, assets.DogMock1, 5 * config.TickInterval},
	{engine.Vector3D{114, 117, 10}, assets.DogMock2, 5 * config.TickInterval},
	{engine.Vector3D{114, 117, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 119, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 121, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 123, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 125, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 125, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 127, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 129, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 131, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 133, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 133, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 135, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 137, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 139, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 141, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 141, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 143, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 145, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 147, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 149, 10}, assets.DogMock2, 1 * config.TickInterval},
	{engine.Vector3D{114, 149, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 151, 10}, assets.DogMock1, 1 * config.TickInterval},
	{engine.Vector3D{114, 153, 10}, assets.DogMock1, 1 * config.TickInterval},
}

var dogCinematicRetrieve1 = engine.Cinematic3D{
	{engine.Vector3D{92, 155, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 153, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 151, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 149, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 147, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 145, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 143, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 141, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 139, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 137, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 135, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 133, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 131, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 129, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 127, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 125, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 123, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 121, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 119, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 117, 10}, assets.DogRetrieve1, 19 * config.TickInterval},
	{engine.Vector3D{92, 119, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 121, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 123, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 125, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 127, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 129, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 131, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 133, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 135, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 137, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 139, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 141, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 143, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 145, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 147, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 149, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 151, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 153, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
	{engine.Vector3D{92, 155, 10}, assets.DogRetrieve1, 1 * config.TickInterval},
}

var dogCinematicRetrieve2 = engine.Cinematic3D{
	{engine.Vector3D{92, 155, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 153, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 151, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 149, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 147, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 145, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 143, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 141, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 139, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 137, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 135, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 133, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 131, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 129, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 127, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 125, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 123, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 121, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 119, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 117, 10}, assets.DogRetrieve2, 19 * config.TickInterval},
	{engine.Vector3D{92, 119, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 121, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 123, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 125, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 127, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 129, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 131, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 133, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 135, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 137, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 139, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 141, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 143, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 145, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 147, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 149, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 151, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 153, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
	{engine.Vector3D{92, 155, 10}, assets.DogRetrieve2, 1 * config.TickInterval},
}
