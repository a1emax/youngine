package audio

// Track represents playable audio resource.
type Track interface {

	// Rewind requires track to be rewound after the next update.
	Rewind()

	// Play requires track to be played after the next update.
	Play()

	// Pause requires track to be paused after the next update.
	Pause()

	// Update updates track.
	Update() error

	// Close closes track.
	Close() error
}
