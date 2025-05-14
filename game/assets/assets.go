package assets

import (
	"duck-hunt-go/engine"
	"embed"
)

var (
	//go:embed files/*.png
	assets embed.FS

	// Font
	Font = engine.Must(engine.LoadFont(assets, "files/font.png"))

	// Title
	TitleLayout = engine.Must(engine.LoadImage(assets, "files/title.layout.png"))
	TitleCursor = engine.Must(engine.LoadImage(assets, "files/title.cursor.png"))

	// Layout
	Layout      = engine.Must(engine.LoadImage(assets, "files/layout.png"))
	LayoutShrub = engine.Must(engine.LoadImage(assets, "files/layout.shrub.png"))
	LayoutTree  = engine.Must(engine.LoadImage(assets, "files/layout.tree.png"))

	// Gun
	Gun = engine.Must(engine.LoadImage(assets, "files/gun.png"))

	// Duck
	DuckHorizontal1 = engine.Must(engine.LoadImage(assets, "files/duck.horizontal.1.png"))
	DuckHorizontal2 = engine.Must(engine.LoadImage(assets, "files/duck.horizontal.2.png"))
	DuckHorizontal3 = engine.Must(engine.LoadImage(assets, "files/duck.horizontal.3.png"))
	DuckAngled1     = engine.Must(engine.LoadImage(assets, "files/duck.angled.1.png"))
	DuckAngled2     = engine.Must(engine.LoadImage(assets, "files/duck.angled.2.png"))
	DuckAngled3     = engine.Must(engine.LoadImage(assets, "files/duck.angled.3.png"))

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
