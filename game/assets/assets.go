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
	gun       = engine.MustLoad(engine.ImageLoader{assets, "files/gun.png"})
	GunNormal = engine.ImageSlicer{gun, engine.Point{0, 0}, engine.Size{12, 12}}.Image()

	// Duck
	duck                    = engine.MustLoad(engine.ImageLoader{assets, "files/duck.png"})
	DuckFlyHorizontalRight1 = engine.ImageSlicer{duck, engine.Point{0, 0}, engine.Size{39, 39}}.Image()
	DuckFlyHorizontalRight2 = engine.ImageSlicer{duck, engine.Point{0, 39}, engine.Size{39, 39}}.Image()
	DuckFlyHorizontalRight3 = engine.ImageSlicer{duck, engine.Point{0, 78}, engine.Size{39, 39}}.Image()
	DuckFlyHorizontalLeft1  = engine.ImageHorizontalFlipper{DuckFlyHorizontalRight1}.Image()
	DuckFlyHorizontalLeft2  = engine.ImageHorizontalFlipper{DuckFlyHorizontalRight2}.Image()
	DuckFlyHorizontalLeft3  = engine.ImageHorizontalFlipper{DuckFlyHorizontalRight3}.Image()
	DuckFlyAngledRight1     = engine.ImageSlicer{duck, engine.Point{39, 0}, engine.Size{39, 39}}.Image()
	DuckFlyAngledRight2     = engine.ImageSlicer{duck, engine.Point{39, 39}, engine.Size{39, 39}}.Image()
	DuckFlyAngledRight3     = engine.ImageSlicer{duck, engine.Point{39, 78}, engine.Size{39, 39}}.Image()
	DuckFlyAngledLeft1      = engine.ImageHorizontalFlipper{DuckFlyAngledRight1}.Image()
	DuckFlyAngledLeft2      = engine.ImageHorizontalFlipper{DuckFlyAngledRight2}.Image()
	DuckFlyAngledLeft3      = engine.ImageHorizontalFlipper{DuckFlyAngledRight3}.Image()
	DuckShotRight           = engine.ImageSlicer{duck, engine.Point{117, 0}, engine.Size{39, 39}}.Image()
	DuckShotLeft            = engine.ImageHorizontalFlipper{DuckShotRight}.Image()
	DuckFallRight           = engine.ImageSlicer{duck, engine.Point{117, 39}, engine.Size{39, 39}}.Image()
	DuckFallLeft            = engine.ImageHorizontalFlipper{DuckFallRight}.Image()

	// Dog
	dog          = engine.MustLoad(engine.ImageLoader{assets, "files/dog.png"})
	DogTrack1    = engine.ImageSlicer{dog, engine.Point{0, 0}, engine.Size{55, 48}}.Image()
	DogTrack2    = engine.ImageSlicer{dog, engine.Point{0, 48}, engine.Size{55, 48}}.Image()
	DogTrack3    = engine.ImageSlicer{dog, engine.Point{0, 96}, engine.Size{55, 48}}.Image()
	DogTrack4    = engine.ImageSlicer{dog, engine.Point{0, 144}, engine.Size{55, 48}}.Image()
	DogSniff     = engine.ImageSlicer{dog, engine.Point{0, 192}, engine.Size{55, 48}}.Image()
	DogPant      = engine.ImageSlicer{dog, engine.Point{0, 240}, engine.Size{55, 48}}.Image()
	DogJump1     = engine.ImageSlicer{dog, engine.Point{55, 0}, engine.Size{55, 48}}.Image()
	DogJump2     = engine.ImageSlicer{dog, engine.Point{55, 48}, engine.Size{55, 48}}.Image()
	DogMock1     = engine.ImageSlicer{dog, engine.Point{55, 96}, engine.Size{29, 39}}.Image()
	DogMock2     = engine.ImageSlicer{dog, engine.Point{55, 135}, engine.Size{29, 39}}.Image()
	DogRetrieve1 = engine.ImageSlicer{dog, engine.Point{55, 174}, engine.Size{56, 39}}.Image()
	DogRetrieve2 = engine.ImageSlicer{dog, engine.Point{55, 213}, engine.Size{56, 39}}.Image()
)
