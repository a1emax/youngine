package audio

// Factory of playable audio resources.
type Factory interface {

	// SampleRate returns used sample rate.
	SampleRate() SampleRate

	// NewInstance initializes and returns new [Instance].
	//
	// Source should not be shared with other instances.
	NewInstance(source Source, volumer Volumer) Instance
}
