package audio

// Manager of audio resources.
type Manager interface {

	// SampleRate returns used sample rate.
	SampleRate() SampleRate

	// NewTrack initializes and returns new [Track].
	//
	// NOTE that source should not be shared with other tracks.
	NewTrack(source Source, volumer Volumer) Track
}
