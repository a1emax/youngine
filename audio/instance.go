package audio

// Instance of audio resource for playing.
type Instance interface {

	// Rewind requires instance to be rewound after the next update.
	Rewind()

	// Play requires instance to be played after the next update.
	Play()

	// Pause requires instance to be paused after the next update.
	Pause()

	// Update updates instance.
	Update() error

	// Close closes instance.
	Close() error
}
