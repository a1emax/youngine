package ebitenaudio

import (
	ebiten "github.com/hajimehoshi/ebiten/v2/audio"

	"github.com/a1emax/youngine/audio"
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// trackImpl is the implementation of the [audio.Track] interface.
type trackImpl struct {
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

// Rewind implements the [audio.Track] interface.
func (t *trackImpl) Rewind() {
	t.rewind = true
}

// Play implements the [audio.Track] interface.
func (t *trackImpl) Play() {
	t.play = true
}

// Pause implements the [audio.Track] interface.
func (t *trackImpl) Pause() {
	t.play = false
}

// Update implements the [audio.Track] interface.
func (t *trackImpl) Update() error {
	rewind := t.rewind
	t.rewind = false

	if t.player == nil {
		if !t.context.IsReady() {
			return nil
		}

		player, err := t.context.NewPlayer(t.source)
		if err != nil {
			return fault.Trace(err)
		}

		t.player = player
		t.playerState.maybePlaying = false
		t.playerState.volume = 1

		rewind = false
	}

	if !t.play && t.playerState.maybePlaying {
		t.player.Pause()
		t.playerState.maybePlaying = false
	}

	volume := t.volumer.Volume()
	if t.playerState.volume != volume {
		t.player.SetVolume(volume)
		t.playerState.volume = volume
	}

	if rewind {
		err := t.player.Rewind()
		if err != nil {
			return fault.Trace(err)
		}

		t.playerState.maybePlaying = false
	}

	if t.play && !t.playerState.maybePlaying {
		t.player.Play()
		t.playerState.maybePlaying = true
	}

	return nil
}

// Close implements the [audio.Track] interface.
func (t *trackImpl) Close() error {
	player := t.player
	*t = trackImpl{}

	if player == nil {
		return nil
	}

	err := player.Close()
	if err != nil {
		return fault.Trace(err)
	}

	return nil
}
