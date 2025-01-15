package ebiteninput

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/clock/driver/ebitenclock"
)

// touchscreenHelper represents touchscreen state for internal purposes.
type touchscreenHelper struct {
	clock clock.Clock

	touchIDs      []ebiten.TouchID
	touches       []touchscreenHelperTouch
	prevTouches   []touchscreenHelperTouch
	maxTouchCount int

	// For testing purposes.
	appendTouchIDs func(touches []ebiten.TouchID) []ebiten.TouchID
	touchPosition  func(id ebiten.TouchID) (int, int)
}

// touchscreenHelperTouch represents touch state for internal purposes.
type touchscreenHelperTouch struct {
	id        ebiten.TouchID
	x, y      int
	startedAt clock.Time
	changedAt clock.Time
	exposedAt clock.Time
}

// newTouchscreenHelper initializes and returns new touchscreenHelper.
func newTouchscreenHelper(clk clock.Clock) touchscreenHelper {
	return touchscreenHelper{
		clock:          clk,
		appendTouchIDs: ebiten.AppendTouchIDs,
		touchPosition:  ebiten.TouchPosition,
	}
}

// Ebitengine (or gomobile) contains bug (?) due to which more than safeTouchscreenTouchCount simultaneous touches can
// "hang". Here it is assumed that if coordinates of touch did not change for more than safeTouchscreenTouchDuration,
// then it "hung" and should not be exposed until its coordinates are changed (if they are ever changed).
const (
	safeTouchscreenTouchCount    = 2
	safeTouchscreenTouchDuration = 2 * ebitenclock.Second
)

// update updates state.
func (t *touchscreenHelper) update() {
	t.touchIDs = t.appendTouchIDs(t.touchIDs[:0])
	sort.Slice(t.touchIDs, func(i, j int) bool {
		return t.touchIDs[i] < t.touchIDs[j]
	})

	if len(t.touchIDs) == 0 {
		t.maxTouchCount = 0
	} else {
		t.maxTouchCount = max(len(t.touchIDs), t.maxTouchCount)
	}

	now := t.clock.Now()

	t.prevTouches, t.touches = t.touches, t.prevTouches[:0]
	prevTouchIndex := 0
	for _, touchID := range t.touchIDs {
		touch := touchscreenHelperTouch{id: touchID}
		touch.x, touch.y = t.touchPosition(touchID)

		for prevTouchIndex < len(t.prevTouches) {
			prevTouch := &t.prevTouches[prevTouchIndex]

			if prevTouch.id > touch.id {
				break
			}

			prevTouchIndex++

			if prevTouch.id < touch.id {
				continue
			}

			touch.startedAt = prevTouch.startedAt

			if touch.x != prevTouch.x || touch.y != prevTouch.y {
				touch.changedAt = now
			} else {
				touch.changedAt = prevTouch.changedAt
			}

			if t.maxTouchCount > safeTouchscreenTouchCount &&
				now.Sub(touch.changedAt)+1 > safeTouchscreenTouchDuration {

				touch.exposedAt = clock.Time{}
			} else {
				touch.exposedAt = prevTouch.exposedAt
				if touch.exposedAt.IsZero() {
					touch.exposedAt = now
				}
			}

			break
		}
		if touch.startedAt.IsZero() {
			touch.startedAt = now
			touch.changedAt = now
			touch.exposedAt = now
		}

		t.touches = append(t.touches, touch)
	}
}
