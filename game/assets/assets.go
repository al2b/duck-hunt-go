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
	Font = engine.Must(engine.BitmapFontLoader{
		assets, "files/font.png",
		engine.SquareBitmapFontMaskMapper{}, charmap.CodePage850,
	}.Load())

	// Menu
	MenuLayout = engine.Must(engine.LoadImage(assets, "files/menu.layout.png"))
	MenuCursor = engine.Must(engine.LoadImage(assets, "files/menu.cursor.png"))

	// Layout
	Layout     = engine.Must(engine.LoadImage(assets, "files/layout.png"))
	LayoutBush = engine.Must(engine.LoadImage(assets, "files/layout.bush.png"))
	LayoutTree = engine.Must(engine.LoadImage(assets, "files/layout.tree.png"))
	LayoutAmmo = engine.Must(engine.LoadImage(assets, "files/layout.ammo.png"))

	// Gun
	Gun = engine.Must(engine.LoadImage(assets, "files/gun.png"))

	// Duck
	DuckFlyHorizontalRight1 = engine.Must(engine.LoadImage(assets, "files/duck.fly.horizontal.1.png"))
	DuckFlyHorizontalRight2 = engine.Must(engine.LoadImage(assets, "files/duck.fly.horizontal.2.png"))
	DuckFlyHorizontalRight3 = engine.Must(engine.LoadImage(assets, "files/duck.fly.horizontal.3.png"))
	DuckFlyHorizontalLeft1  = DuckFlyHorizontalRight1.FlipHorizontal()
	DuckFlyHorizontalLeft2  = DuckFlyHorizontalRight2.FlipHorizontal()
	DuckFlyHorizontalLeft3  = DuckFlyHorizontalRight3.FlipHorizontal()
	DuckFlyAngledRight1     = engine.Must(engine.LoadImage(assets, "files/duck.fly.angled.1.png"))
	DuckFlyAngledRight2     = engine.Must(engine.LoadImage(assets, "files/duck.fly.angled.2.png"))
	DuckFlyAngledRight3     = engine.Must(engine.LoadImage(assets, "files/duck.fly.angled.3.png"))
	DuckFlyAngledLeft1      = DuckFlyAngledRight1.FlipHorizontal()
	DuckFlyAngledLeft2      = DuckFlyAngledRight2.FlipHorizontal()
	DuckFlyAngledLeft3      = DuckFlyAngledRight3.FlipHorizontal()
	DuckShotRight           = engine.Must(engine.LoadImage(assets, "files/duck.shot.png"))
	DuckShotLeft            = DuckShotRight.FlipHorizontal()
	DuckFallRight           = engine.Must(engine.LoadImage(assets, "files/duck.fall.png"))
	DuckFallLeft            = DuckFallRight.FlipHorizontal()

	// Dog
	DogTrack1    = engine.Must(engine.LoadImage(assets, "files/dog.track.1.png"))
	DogTrack2    = engine.Must(engine.LoadImage(assets, "files/dog.track.2.png"))
	DogTrack3    = engine.Must(engine.LoadImage(assets, "files/dog.track.3.png"))
	DogTrack4    = engine.Must(engine.LoadImage(assets, "files/dog.track.4.png"))
	DogSniff     = engine.Must(engine.LoadImage(assets, "files/dog.sniff.png"))
	DogPant      = engine.Must(engine.LoadImage(assets, "files/dog.pant.png"))
	DogJump1     = engine.Must(engine.LoadImage(assets, "files/dog.jump.1.png"))
	DogJump2     = engine.Must(engine.LoadImage(assets, "files/dog.jump.2.png"))
	DogMock1     = engine.Must(engine.LoadImage(assets, "files/dog.mock.1.png"))
	DogMock2     = engine.Must(engine.LoadImage(assets, "files/dog.mock.2.png"))
	DogRetrieve1 = engine.Must(engine.LoadImage(assets, "files/dog.retrieve.1.png"))
	DogRetrieve2 = engine.Must(engine.LoadImage(assets, "files/dog.retrieve.2.png"))
)
