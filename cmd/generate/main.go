package main

import (
	"log"
	"math"
	"os"

	"github.com/fr3fou/gusic/gusic"
)

// Memory C# F# G#
func Memory() gusic.Melody {
	sampleRate := 48000.0
	bpm := 80
	noteLength := 4 // 4/4, therefore 4
	volume := 0.125

	m := gusic.NewMelody(
		sampleRate,
		bpm,
		noteLength,
		math.Sin,
		gusic.NewEasedADSR(func(t float64) float64 {
			return t * t
		}, gusic.NewRatios(0.25, 0.25, 0.25, 0.25), 1.35, 0.35),
	)

	m.AddNotes(
		gusic.NewChord(
			gusic.C(4, m.Crotchet(), volume),
			gusic.E(4, m.Crotchet(), volume),
			gusic.G(4, m.Crotchet(), volume),
		),
	)

	//m.AddNotes(
	//	gusic.D(5, m.Quaver(), volume),
	//	gusic.B(5, m.Quaver(), volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	gusic.E(5, m.Quaver(), volume),
	//	gusic.GS(5, m.Quaver()+m.Semiquaver(), volume),
	//	gusic.GS(5, m.Semiquaver()*3, volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	//
	//	gusic.Rest(m.Quaver()),
	//	gusic.E(5, m.Quaver(), volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	gusic.E(5, m.Quaver(), volume),
	//	gusic.GS(5, m.Quaver()+m.Semiquaver(), volume),
	//	gusic.GS(5, m.Semiquaver()*3, volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	//
	//	gusic.D(5, m.Quaver(), volume),
	//	gusic.B(5, m.Quaver(), volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	gusic.E(5, m.Quaver(), volume),
	//	gusic.GS(5, m.Quaver()+m.Semiquaver(), volume),
	//	gusic.GS(5, m.Semiquaver()*3, volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	//
	//	gusic.Rest(m.Quaver()),
	//	gusic.E(5, m.Quaver(), volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	gusic.CS(6, m.Quaver(), volume),
	//	gusic.B(5, m.Quaver()+m.Semiquaver(), volume),
	//	gusic.A(5, m.Semiquaver()*3, volume),
	//	gusic.B(5, m.Quaver(), volume),
	//	// repeat
	//	gusic.D(5, m.Quaver(), volume),
	//	gusic.B(5, m.Quaver(), volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	gusic.E(5, m.Quaver(), volume),
	//	gusic.GS(5, m.Quaver()+m.Semiquaver(), volume),
	//	gusic.GS(5, m.Semiquaver()*3, volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	//
	//	gusic.Rest(m.Quaver()),
	//	gusic.E(5, m.Quaver(), volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	gusic.E(5, m.Quaver(), volume),
	//	gusic.GS(5, m.Quaver()+m.Semiquaver(), volume),
	//	gusic.GS(5, m.Semiquaver()*3, volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	//
	//	gusic.D(5, m.Quaver(), volume),
	//	gusic.B(5, m.Quaver(), volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	gusic.E(5, m.Quaver(), volume),
	//	gusic.GS(5, m.Quaver()+m.Semiquaver(), volume),
	//	gusic.GS(5, m.Semiquaver()*3, volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	//
	//	gusic.Rest(m.Quaver()),
	//	gusic.E(5, m.Quaver(), volume),
	//	gusic.A(5, m.Quaver(), volume),
	//	gusic.CS(6, m.Quaver(), volume),
	//	gusic.B(5, m.Quaver()+m.Semiquaver(), volume),
	//	gusic.A(5, m.Semiquaver()*3, volume),
	//	gusic.B(5, m.Quaver(), volume),
	//)

	//m.AddRun(
	//	//
	//	gusic.Rest(m.Quaver()),
	//	gusic.D(4, m.Crotchet(), volume*2/3),
	//	gusic.D(4, m.Crotchet(), volume*2/3),
	//	gusic.D(4, m.Crotchet(), volume*2/3),
	//	gusic.D(4, m.Crotchet(), volume*2/3), //
	//	gusic.D(4, m.Crotchet(), volume*2/3),
	//	gusic.D(4, m.Crotchet(), volume*2/3),
	//	gusic.D(4, m.Crotchet(), volume*2/3),
	//	gusic.D(4, m.Crotchet(), volume*2/3), //
	//	gusic.D(4, m.Crotchet(), volume*2/3),
	//	gusic.D(4, m.Crotchet(), volume*2/3),
	//	gusic.D(4, m.Crotchet(), volume*2/3),
	//	gusic.D(4, m.Crotchet(), volume*2/3), //
	//	gusic.D(4, m.Crotchet(), volume*2/3),
	//	gusic.D(4, m.Crotchet(), volume*2/3),
	//	gusic.D(4, m.Crotchet(), volume*2/3),
	//	gusic.D(4, m.Crotchet(), volume*2/3), //
	//	gusic.C(4, m.Crotchet(), volume*2/3),
	//	gusic.C(4, m.Crotchet(), volume*2/3),
	//	gusic.C(4, m.Crotchet(), volume*2/3),
	//	gusic.C(4, m.Crotchet(), volume*2/3), //
	//	gusic.C(4, m.Crotchet(), volume*2/3),
	//	gusic.C(4, m.Crotchet(), volume*2/3),
	//	gusic.C(4, m.Crotchet(), volume*2/3),
	//	gusic.C(4, m.Crotchet(), volume*2/3), //
	//	gusic.C(4, m.Crotchet(), volume*2/3),
	//	gusic.C(4, m.Crotchet(), volume*2/3),
	//	gusic.C(4, m.Crotchet(), volume*2/3),
	//	gusic.C(4, m.Crotchet(), volume*2/3), //
	//	gusic.FS(4, m.Crotchet(), volume*2/3),
	//	gusic.FS(4, m.Crotchet(), volume*2/3),
	//	gusic.FS(4, m.Crotchet(), volume*2/3),
	//	gusic.FS(4, m.Quaver(), volume*2/3),
	//)

	return *m
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("gusic: File name argument required\nUsage: %s <filename>", os.Args[0])
	}

	FileName := os.Args[1]
	file, err := os.Create(FileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	m := Memory()

	if err := m.Write(file); err != nil {
		panic(err)
	}
}
