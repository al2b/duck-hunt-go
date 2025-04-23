package dog

import (
	"duck-hunt-go/engine"
	"embed"
	"time"
)

const frameDuration = time.Second / 60

var (
	//go:embed assets/*.png
	assets embed.FS

	// Images
	imageDogTrack1 = engine.Must(engine.LoadImage(assets, "assets/dog.track.1.png"))
	imageDogTrack2 = engine.Must(engine.LoadImage(assets, "assets/dog.track.2.png"))
	imageDogTrack3 = engine.Must(engine.LoadImage(assets, "assets/dog.track.3.png"))
	imageDogTrack4 = engine.Must(engine.LoadImage(assets, "assets/dog.track.4.png"))
	imageDogSniff  = engine.Must(engine.LoadImage(assets, "assets/dog.sniff.png"))
	imageDogPant   = engine.Must(engine.LoadImage(assets, "assets/dog.pant.png"))
	imageDogJump1  = engine.Must(engine.LoadImage(assets, "assets/dog.jump.1.png"))
	imageDogJump2  = engine.Must(engine.LoadImage(assets, "assets/dog.jump.2.png"))

	// Cinematics
	cinematicDogTrack = engine.Cinematic3D{
		{engine.Vec3D(2, 141, -10), imageDogTrack1, 7 * frameDuration},
		{engine.Vec3D(4, 141, -10), imageDogTrack2, 7 * frameDuration},
		{engine.Vec3D(6, 141, -10), imageDogTrack3, 7 * frameDuration},
		{engine.Vec3D(8, 141, -10), imageDogTrack4, 7 * frameDuration},

		{engine.Vec3D(10, 141, -10), imageDogTrack1, 7 * frameDuration},
		{engine.Vec3D(12, 141, -10), imageDogTrack2, 7 * frameDuration},
		{engine.Vec3D(14, 141, -10), imageDogTrack3, 7 * frameDuration},
		{engine.Vec3D(16, 141, -10), imageDogTrack4, 7 * frameDuration},

		{engine.Vec3D(18, 141, -10), imageDogTrack1, 7 * frameDuration},
		{engine.Vec3D(20, 141, -10), imageDogTrack2, 7 * frameDuration},
		{engine.Vec3D(22, 141, -10), imageDogTrack3, 7 * frameDuration},
		{engine.Vec3D(24, 141, -10), imageDogTrack4, 7 * frameDuration},

		{engine.Vec3D(26, 141, -10), imageDogTrack1, 7 * frameDuration},
		{engine.Vec3D(28, 141, -10), imageDogTrack2, 7 * frameDuration},
		{engine.Vec3D(30, 141, -10), imageDogTrack3, 7 * frameDuration},
		{engine.Vec3D(32, 141, -10), imageDogTrack4, 7 * frameDuration},

		{engine.Vec3D(34, 141, -10), imageDogTrack1, 14 * frameDuration},

		{engine.Vec3D(34, 141, -10), imageDogSniff, 9 * frameDuration},
		{engine.Vec3D(34, 141, -10), imageDogTrack1, 9 * frameDuration},

		{engine.Vec3D(34, 141, -10), imageDogSniff, 9 * frameDuration},
		{engine.Vec3D(34, 141, -10), imageDogTrack1, 9 * frameDuration},

		{engine.Vec3D(34, 141, -10), imageDogSniff, 10 * frameDuration},

		{engine.Vec3D(36, 141, -10), imageDogTrack1, 7 * frameDuration},
		{engine.Vec3D(38, 141, -10), imageDogTrack2, 7 * frameDuration},
		{engine.Vec3D(40, 141, -10), imageDogTrack3, 7 * frameDuration},
		{engine.Vec3D(42, 141, -10), imageDogTrack4, 7 * frameDuration},

		{engine.Vec3D(44, 141, -10), imageDogTrack1, 7 * frameDuration},
		{engine.Vec3D(46, 141, -10), imageDogTrack2, 7 * frameDuration},
		{engine.Vec3D(48, 141, -10), imageDogTrack3, 7 * frameDuration},
		{engine.Vec3D(50, 141, -10), imageDogTrack4, 7 * frameDuration},

		{engine.Vec3D(52, 141, -10), imageDogTrack1, 7 * frameDuration},
		{engine.Vec3D(54, 141, -10), imageDogTrack2, 7 * frameDuration},
		{engine.Vec3D(56, 141, -10), imageDogTrack3, 7 * frameDuration},
		{engine.Vec3D(58, 141, -10), imageDogTrack4, 7 * frameDuration},

		{engine.Vec3D(60, 141, -10), imageDogTrack1, 7 * frameDuration},
		{engine.Vec3D(62, 141, -10), imageDogTrack2, 7 * frameDuration},
		{engine.Vec3D(64, 141, -10), imageDogTrack3, 7 * frameDuration},
		{engine.Vec3D(66, 141, -10), imageDogTrack4, 7 * frameDuration},

		{engine.Vec3D(68, 141, -10), imageDogTrack1, 14 * frameDuration},

		{engine.Vec3D(68, 141, -10), imageDogSniff, 9 * frameDuration},
		{engine.Vec3D(68, 141, -10), imageDogTrack1, 9 * frameDuration},

		{engine.Vec3D(68, 141, -10), imageDogSniff, 9 * frameDuration},
		{engine.Vec3D(68, 141, -10), imageDogTrack1, 9 * frameDuration},

		{engine.Vec3D(68, 141, -10), imageDogSniff, 10 * frameDuration},

		{engine.Vec3D(70, 141, -10), imageDogTrack1, 7 * frameDuration},
		{engine.Vec3D(72, 141, -10), imageDogTrack2, 1 * frameDuration},
		{engine.Vec3D(72, 141, -10), imageDogPant, 18 * frameDuration},

		{engine.Vec3D(85, 125, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(86, 122, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(87, 119, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(88, 116, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(89, 113, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(90, 111, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(91, 109, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(92, 107, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(93, 105, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(94, 104, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(95, 102, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(96, 101, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(97, 100, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(98, 99, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(99, 98, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(100, 97, -10), imageDogJump1, 1 * frameDuration},
		{engine.Vec3D(101, 96, -10), imageDogJump1, 1 * frameDuration},

		{engine.Vec3D(102, 92, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(103, 92, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(104, 91, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(105, 91, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(106, 90, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(107, 90, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(108, 90, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(109, 90, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(110, 89, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(111, 89, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(112, 89, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(113, 89, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(114, 89, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(115, 90, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(116, 90, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(117, 91, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(118, 91, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(119, 92, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(120, 93, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(121, 94, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(122, 96, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(123, 98, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(124, 100, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(125, 103, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(126, 106, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(126, 110, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(126, 114, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(126, 119, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(126, 124, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(126, 129, 10), imageDogJump2, 1 * frameDuration},
		{engine.Vec3D(126, 134, 10), imageDogJump2, 1 * frameDuration},
	}
)
