package ebitenaudio

import (
	ebiten "github.com/hajimehoshi/ebiten/v2/audio"

	"github.com/a1emax/youngine/audio"
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// instanceImpl is the implementation of the [audio.Instance] interface.
type instanceImpl struct {
	context *ebiten.Context
	source  audio.Source
	volumer audio.Volumer

	rewind bool
	play   bool

	player      *ebiten.Player
	playerState struct {
		volume       basic.Float
		maybePlaying bool
	}
}

// Rewind implements the [audio.Instance] interface.
func (i *instanceImpl) Rewind() {
	i.rewind = true
}

// Play implements the [audio.Instance] interface.
func (i *instanceImpl) Play() {
	i.play = true
}

// Pause implements the [audio.Instance] interface.
func (i *instanceImpl) Pause() {
	i.play = false
}

// Update implements the [audio.Instance] interface.
func (i *instanceImpl) Update() error {
	rewind := i.rewind
	i.rewind = false

	if i.player == nil {
		if !i.context.IsReady() {
			return nil
		}

		player, err := i.context.NewPlayer(i.source)
		if err != nil {
			return fault.Trace(err)
		}

		i.player = player
		i.playerState.maybePlaying = false
		i.playerState.volume = 1

		rewind = false
	}

	if !i.play && i.playerState.maybePlaying {
		i.player.Pause()
		i.playerState.maybePlaying = false
	}

	volume := i.volumer.Volume()
	if i.playerState.volume != volume {
		i.player.SetVolume(volume)
		i.playerState.volume = volume
	}

	if rewind {
		err := i.player.Rewind()
		if err != nil {
			return fault.Trace(err)
		}

		i.playerState.maybePlaying = false
	}

	if i.play && !i.playerState.maybePlaying {
		i.player.Play()
		i.playerState.maybePlaying = true
	}

	return nil
}

// Close implements the [audio.Instance] interface.
func (i *instanceImpl) Close() error {
	player := i.player
	*i = instanceImpl{}

	if player == nil {
		return nil
	}

	err := player.Close()
	if err != nil {
		return fault.Trace(err)
	}

	return nil
}
