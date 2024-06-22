package fyne_state

import (
	"image/color"
	"sync"
	"testing"
	"time"
)

func TestSetAndGet(t *testing.T) {
	Create(func(set func(string, interface{}), get func(string) interface{}) StateActions {
		return StateActions{
			State: map[string]interface{}{
				"count": 0,
			},
		}
	})

	Set("count", 10)
	count := Get[int]("count")
	if count != 10 {
		t.Errorf("expected count to be 10, but got %d", count)
	}
}

func TestSubscribe(t *testing.T) {
	Create(func(set func(string, interface{}), get func(string) interface{}) StateActions {
		return StateActions{
			State: map[string]interface{}{
				"count": 0,
			},
		}
	})

	var wg sync.WaitGroup
	wg.Add(1)

	ch := Subscribe("count")

	go func() {
		<-ch
		count := Get[int]("count")
		if count != 20 {
			t.Errorf("expected count to be 20, but got %d", count)
		}
		wg.Done()
	}()

	Set("count", 20)
	wg.Wait()
}

func TestUnsubscribe(t *testing.T) {
	Create(func(set func(string, interface{}), get func(string) interface{}) StateActions {
		return StateActions{
			State: map[string]interface{}{
				"count": 0,
			},
		}
	})

	ch := Subscribe("count")
	Unsubscribe("count", ch)
	time.Sleep(10 * time.Millisecond)
	Set("count", 30)
	select {
	case <-ch:
		t.Errorf("received an update after unsubscribe")
	default:
		// no update as expected
	}
}

func TestStateActions(t *testing.T) {
	Create(func(set func(string, interface{}), get func(string) interface{}) StateActions {
		return StateActions{
			State: map[string]interface{}{
				"count": 0,
			},
			Actions: map[string]func(){
				"inc": func() {
					count := get("count").(int)
					set("count", count+1)
				},
			},
		}
	})

	inc := Get[func()]("inc")
	inc()
	count := Get[int]("count")
	if count != 1 {
		t.Errorf("expected count to be 1, but got %d", count)
	}
}

func TestComplexState(t *testing.T) {
	Create(func(set func(string, interface{}), get func(string) interface{}) StateActions {
		return StateActions{
			State: map[string]interface{}{
				"color": color.RGBA{R: 255, G: 0, B: 0, A: 255},
			},
		}
	})

	Set("color", color.RGBA{R: 0, G: 255, B: 0, A: 255})
	c := Get[color.RGBA]("color")
	if c != (color.RGBA{R: 0, G: 255, B: 0, A: 255}) {
		t.Errorf("expected color to be {0, 255, 0, 255}, but got %v", c)
	}
}
