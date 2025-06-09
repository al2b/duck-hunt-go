package assets

import (
	"duck-hunt-go/engine"
	"embed"
	"golang.org/x/text/encoding/charmap"
)

var (
	//go:embed files/*.png
	assets embed.FS

	// Fonts
	Font = engine.MustLoad(engine.BitmapFontLoader{
		assets, "files/font.png",
		engine.SquareBitmapFontMapper{}, charmap.CodePage850,
	})

	// Menu
	MenuLayout = engine.MustLoad(engine.ImageLoader{assets, "files/menu.layout.png"})
	MenuCursor = engine.MustLoad(engine.ImageLoader{assets, "files/menu.cursor.png"})

	// Layout
	Layout     = engine.MustLoad(engine.ImageLoader{assets, "files/layout.png"})
	LayoutBush = engine.MustLoad(engine.ImageLoader{assets, "files/layout.bush.png"})
	LayoutTree = engine.MustLoad(engine.ImageLoader{assets, "files/layout.tree.png"})
	LayoutAmmo = engine.MustLoad(engine.ImageLoader{assets, "files/layout.ammo.png"})

	// Gun
	Gun = engine.MustLoad(engine.ImageLoader{assets, "files/gun.png"})

	// Duck
	DuckFlyHorizontalRight1 = engine.MustLoad(engine.ImageLoader{assets, "files/duck.fly.horizontal.1.png"})
	DuckFlyHorizontalRight2 = engine.MustLoad(engine.ImageLoader{assets, "files/duck.fly.horizontal.2.png"})
	DuckFlyHorizontalRight3 = engine.MustLoad(engine.ImageLoader{assets, "files/duck.fly.horizontal.3.png"})
	DuckFlyHorizontalLeft1  = DuckFlyHorizontalRight1.FlipHorizontal()
	DuckFlyHorizontalLeft2  = DuckFlyHorizontalRight2.FlipHorizontal()
	DuckFlyHorizontalLeft3  = DuckFlyHorizontalRight3.FlipHorizontal()
	DuckFlyAngledRight1     = engine.MustLoad(engine.ImageLoader{assets, "files/duck.fly.angled.1.png"})
	DuckFlyAngledRight2     = engine.MustLoad(engine.ImageLoader{assets, "files/duck.fly.angled.2.png"})
	DuckFlyAngledRight3     = engine.MustLoad(engine.ImageLoader{assets, "files/duck.fly.angled.3.png"})
	DuckFlyAngledLeft1      = DuckFlyAngledRight1.FlipHorizontal()
	DuckFlyAngledLeft2      = DuckFlyAngledRight2.FlipHorizontal()
	DuckFlyAngledLeft3      = DuckFlyAngledRight3.FlipHorizontal()
	DuckShotRight           = engine.MustLoad(engine.ImageLoader{assets, "files/duck.shot.png"})
	DuckShotLeft            = DuckShotRight.FlipHorizontal()
	DuckFallRight           = engine.MustLoad(engine.ImageLoader{assets, "files/duck.fall.png"})
	DuckFallLeft            = DuckFallRight.FlipHorizontal()

	// Dog
	DogTrack1    = engine.MustLoad(engine.ImageLoader{assets, "files/dog.track.1.png"})
	DogTrack2    = engine.MustLoad(engine.ImageLoader{assets, "files/dog.track.2.png"})
	DogTrack3    = engine.MustLoad(engine.ImageLoader{assets, "files/dog.track.3.png"})
	DogTrack4    = engine.MustLoad(engine.ImageLoader{assets, "files/dog.track.4.png"})
	DogSniff     = engine.MustLoad(engine.ImageLoader{assets, "files/dog.sniff.png"})
	DogPant      = engine.MustLoad(engine.ImageLoader{assets, "files/dog.pant.png"})
	DogJump1     = engine.MustLoad(engine.ImageLoader{assets, "files/dog.jump.1.png"})
	DogJump2     = engine.MustLoad(engine.ImageLoader{assets, "files/dog.jump.2.png"})
	DogMock1     = engine.MustLoad(engine.ImageLoader{assets, "files/dog.mock.1.png"})
	DogMock2     = engine.MustLoad(engine.ImageLoader{assets, "files/dog.mock.2.png"})
	DogRetrieve1 = engine.MustLoad(engine.ImageLoader{assets, "files/dog.retrieve.1.png"})
	DogRetrieve2 = engine.MustLoad(engine.ImageLoader{assets, "files/dog.retrieve.2.png"})
)
