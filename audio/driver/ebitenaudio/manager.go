package ebitenaudio

import (
	ebiten "github.com/hajimehoshi/ebiten/v2/audio"

	"github.com/a1emax/youngine/audio"
	"github.com/a1emax/youngine/fault"
)

// Manager of audio resources based on Ebitengine.
type Manager interface {
	audio.Manager
}

// managerImpl is the implementation of the [Manager] interface.
type managerImpl struct {
	sampleRate audio.SampleRate
	context    *ebiten.Context
}

// NewManager initializes and returns new [Manager].
//
// NewManager panics if called more than once or if the [ebiten.Context] is already initialized.
func NewManager(sampleRate audio.SampleRate) Manager {
	if ebiten.CurrentContext() != nil {
		panic(fault.Trace(fault.ErrInvalidUse))
	}

	return &managerImpl{
		sampleRate: sampleRate,
		context:    ebiten.NewContext(int(sampleRate)),
	}
}

// SampleRate implements the [audio.Manager] interface.
func (m *managerImpl) SampleRate() audio.SampleRate {
	return m.sampleRate
}

// NewTrack implements the [audio.Manager] interface.
func (m *managerImpl) NewTrack(source audio.Source, volumer audio.Volumer) audio.Track {
	if source == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if volumer == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &trackImpl{
		context: m.context,
		source:  source,
		volumer: volumer,
	}
}
