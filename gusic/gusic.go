package gusic

import "time"

// Generator is a function for producing samples
type Generator = func(x float64) float64

// Easer is an easing function meant to be used in conjuction with an EasedADSR
type Easer = func(t float64) float64

// NoteDuration is the duration of a single note
type NoteDuration = time.Duration

// Melody is a melody
type Melody struct {
	Notes      []Note
	SampleRate float64
	NoteLength int
	BPM        int
	Generator  Generator
	Envelope   ADSR

	breve          NoteDuration
	semibreve      NoteDuration
	minim          NoteDuration
	crotchet       NoteDuration
	quaver         NoteDuration
	semiquaver     NoteDuration
	demisemiquaver NoteDuration
}

// LinearADSR is a linear Envelope implementation
type LinearADSR struct {
	ratios            ADSRRatios
	attackMultiplier  float64
	decayMultiplier   float64
	releaseMultiplier float64
	currentMultiplier float64
}

// EasedADSR is an eased Envelope implementation
type EasedADSR struct {
	easing            Easer
	ratios            ADSRRatios
	attackMultiplier  float64
	decayMultiplier   float64
	releaseMultiplier float64
	currentMultiplier float64
}

// ADSRRatios is a struct for determining what duration should each state in an ADSR take.
type ADSRRatios struct {
	AttackRatio  float64
	DecayRatio   float64
	SustainRatio float64
	ReleaseRatio float64
}

// ADSR defines what an Envelope should behave like
type ADSR interface {
	Attack(notes []float64)
	Decay(notes []float64)
	Sustain(notes []float64)
	Release(notes []float64)
	GetRatios() ADSRRatios
}

// Note is a note
type Note struct {
	Frequency float64
	Duration  NoteDuration
	Volume    float64
}
