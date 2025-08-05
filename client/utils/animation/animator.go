package animation

import (
	"time"
)

// ----------- Core Types -----------

type PlayMode int

const (
	Once PlayMode = iota
	Loop
	PingPong
)

// ----------- Keyframes -----------
type Keyframe[T any] struct {
	Time  float32 // 0.0 to 1.0
	Value T
}

type Interpolator[T any] func(a, b T, t float32) T

type KeyframedProperty[T any] struct {
	Keyframes []Keyframe[T]
	Interp    Interpolator[T]
}

func (kf KeyframedProperty[T]) Sample(t float32) T {
	n := len(kf.Keyframes)
	if n == 0 {
		var zero T
		return zero
	}
	if t <= kf.Keyframes[0].Time {
		return kf.Keyframes[0].Value
	}
	if t >= kf.Keyframes[n-1].Time {
		return kf.Keyframes[n-1].Value
	}
	for i := 0; i < n-1; i++ {
		a := kf.Keyframes[i]
		b := kf.Keyframes[i+1]
		if t >= a.Time && t <= b.Time {
			localT := (t - a.Time) / (b.Time - a.Time)
			return kf.Interp(a.Value, b.Value, localT)
		}
	}
	return kf.Keyframes[n-1].Value
}

// ----------- Animator -----------

type Animator struct {
	start    time.Time
	duration time.Duration
	running  bool
	progress float32
	reversed bool
	playMode PlayMode
	OnFinish func()
	OnLoop   func()
	OnStop   func()
}

func (a *Animator) Start(duration time.Duration, mode PlayMode) {
	a.start = time.Now()
	a.duration = duration
	a.running = true
	a.reversed = false
	a.playMode = mode
	a.progress = 0
}

func (a *Animator) Stop() {
	if a.running {
		a.running = false
		if a.OnStop != nil {
			a.OnStop()
		}
	}
}

func (a *Animator) Cancel() {
	a.Stop()
	a.progress = 0
	a.reversed = false
}

func (a *Animator) Update() {
	if !a.running {
		return
	}
	elapsed := time.Since(a.start)
	t := float32(elapsed) / float32(a.duration)
	if t >= 1 {
		switch a.playMode {
		case Once:
			a.progress = 1
			a.running = false
			if a.OnFinish != nil {
				a.OnFinish()
			}
		case Loop:
			a.start = time.Now()
			a.progress = 0
			if a.OnLoop != nil {
				a.OnLoop()
			}
		case PingPong:
			a.start = time.Now()
			a.reversed = !a.reversed
			a.progress = 0
			if a.OnLoop != nil {
				a.OnLoop()
			}
		}
	} else {
		a.progress = t
	}
}

func (a *Animator) Progress() float32 {
	if a.reversed {
		return 1 - a.progress
	}
	return a.progress
}

func (a *Animator) Running() bool    { return a.running }
func (a *Animator) Reversed() bool   { return a.reversed }
func (a *Animator) IsFinished() bool { return !a.running }
