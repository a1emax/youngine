package ebiteninput

import (
	"reflect"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/clock"
)

func TestTouchscreenHelper(t *testing.T) {
	type touchMap map[ebiten.TouchID][2]int

	var currentTime clock.Time
	var currentTouchMap touchMap
	helper := &touchscreenHelper{
		clock: testClockFunc(func() clock.Time {
			return currentTime
		}),
		appendTouchIDs: func(touches []ebiten.TouchID) []ebiten.TouchID {
			touches = touches[:0]
			for touchID := range currentTouchMap {
				touches = append(touches, touchID)
			}

			return touches
		},
		touchPosition: func(id ebiten.TouchID) (int, int) {
			xy := currentTouchMap[id]

			return xy[0], xy[1]
		},
	}

	next := func(touchMap touchMap, expectedTouches []touchscreenHelperTouch) {
		currentTime = currentTime.Add(1)
		currentTouchMap = touchMap
		helper.update()

		if !reflect.DeepEqual(helper.touches, expectedTouches) {
			t.Fatalf("wrong touches\nexpected: %+v\ngot:      %+v", expectedTouches, helper.touches)
		}
	}

	next(touchMap{
		1: {1, 1},
		2: {2, 2},
		3: {3, 3},
	}, []touchscreenHelperTouch{
		{id: 1, x: 1, y: 1, startedAt: clock.At(1), changedAt: clock.At(1), exposedAt: clock.At(1)},
		{id: 2, x: 2, y: 2, startedAt: clock.At(1), changedAt: clock.At(1), exposedAt: clock.At(1)},
		{id: 3, x: 3, y: 3, startedAt: clock.At(1), changedAt: clock.At(1), exposedAt: clock.At(1)},
	})

	next(touchMap{
		4: {4, 4},
		2: {2, 2},
		5: {5, 5},
	}, []touchscreenHelperTouch{
		{id: 2, x: 2, y: 2, startedAt: clock.At(1), changedAt: clock.At(1), exposedAt: clock.At(1)},
		{id: 4, x: 4, y: 4, startedAt: clock.At(2), changedAt: clock.At(2), exposedAt: clock.At(2)},
		{id: 5, x: 5, y: 5, startedAt: clock.At(2), changedAt: clock.At(2), exposedAt: clock.At(2)},
	})

	next(touchMap{
		1: {1, 1},
		2: {2, 2},
		5: {5, 5},
	}, []touchscreenHelperTouch{
		{id: 1, x: 1, y: 1, startedAt: clock.At(3), changedAt: clock.At(3), exposedAt: clock.At(3)},
		{id: 2, x: 2, y: 2, startedAt: clock.At(1), changedAt: clock.At(1), exposedAt: clock.At(1)},
		{id: 5, x: 5, y: 5, startedAt: clock.At(2), changedAt: clock.At(2), exposedAt: clock.At(2)},
	})

	next(touchMap{
		7: {7, 7},
		8: {8, 8},
		9: {9, 9},
	}, []touchscreenHelperTouch{
		{id: 7, x: 7, y: 7, startedAt: clock.At(4), changedAt: clock.At(4), exposedAt: clock.At(4)},
		{id: 8, x: 8, y: 8, startedAt: clock.At(4), changedAt: clock.At(4), exposedAt: clock.At(4)},
		{id: 9, x: 9, y: 9, startedAt: clock.At(4), changedAt: clock.At(4), exposedAt: clock.At(4)},
	})

	next(touchMap{
		10: {10, 10},
		8:  {8, 8},
		9:  {9, 9},
	}, []touchscreenHelperTouch{
		{id: 8, x: 8, y: 8, startedAt: clock.At(4), changedAt: clock.At(4), exposedAt: clock.At(4)},
		{id: 9, x: 9, y: 9, startedAt: clock.At(4), changedAt: clock.At(4), exposedAt: clock.At(4)},
		{id: 10, x: 10, y: 10, startedAt: clock.At(5), changedAt: clock.At(5), exposedAt: clock.At(5)},
	})

	j := safeTouchscreenTouchDuration
	currentTime = currentTime.Add(j)

	next(touchMap{
		10: {10, 10},
		11: {11, 11},
		12: {12, 12},
	}, []touchscreenHelperTouch{
		{id: 10, x: 10, y: 10, startedAt: clock.At(5), changedAt: clock.At(5), exposedAt: clock.At(0)},
		{id: 11, x: 11, y: 11, startedAt: clock.At(j + 6), changedAt: clock.At(j + 6), exposedAt: clock.At(j + 6)},
		{id: 12, x: 12, y: 12, startedAt: clock.At(j + 6), changedAt: clock.At(j + 6), exposedAt: clock.At(j + 6)},
	})

	next(touchMap{
		10: {0, 0},
		11: {11, 11},
		12: {12, 12},
	}, []touchscreenHelperTouch{
		{id: 10, x: 0, y: 0, startedAt: clock.At(5), changedAt: clock.At(j + 7), exposedAt: clock.At(j + 7)},
		{id: 11, x: 11, y: 11, startedAt: clock.At(j + 6), changedAt: clock.At(j + 6), exposedAt: clock.At(j + 6)},
		{id: 12, x: 12, y: 12, startedAt: clock.At(j + 6), changedAt: clock.At(j + 6), exposedAt: clock.At(j + 6)},
	})
}
