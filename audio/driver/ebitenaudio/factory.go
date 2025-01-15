package ebitenaudio

import (
	ebitenaudio "github.com/hajimehoshi/ebiten/v2/audio"

	"github.com/a1emax/youngine/audio"
	"github.com/a1emax/youngine/fault"
)

// Factory of playable audio resources based on Ebitengine.
type Factory interface {
	audio.Factory
}

// factoryImpl is the implementation of the [Factory] interface.
type factoryImpl struct {
	sampleRate audio.SampleRate
	context    *ebitenaudio.Context
}

// NewFactory initializes and returns new [Factory].
//
// NewFactory panics if called more than once or if the [github.com/hajimehoshi/ebiten/v2/audio.Context]
// is already initialized.
func NewFactory(sampleRate audio.SampleRate) Factory {
	if ebitenaudio.CurrentContext() != nil {
		panic(fault.Trace(fault.ErrInvalidUse))
	}

	return &factoryImpl{
		sampleRate: sampleRate,
		context:    ebitenaudio.NewContext(int(sampleRate)),
	}
}

// SampleRate implements the [audio.Factory] interface.
func (f *factoryImpl) SampleRate() audio.SampleRate {
	return f.sampleRate
}

// NewInstance implements the [audio.Factory] interface.
func (f *factoryImpl) NewInstance(source audio.Source, volumer audio.Volumer) audio.Instance {
	if source == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if volumer == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &instanceImpl{
		context: f.context,
		source:  source,
		volumer: volumer,
	}
}
