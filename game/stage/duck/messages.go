package duck

type ShotMsg struct{}

type DiscriminatedShotMsg struct {
	ShotMsg
	Discriminator any
}

type FallMsg struct{}

type DownMsg struct{}
