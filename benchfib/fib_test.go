package benchfib

import (
	"testing"

	"github.com/loicalbertin/go-tests-training/benchfib/fibchan"
	"github.com/loicalbertin/go-tests-training/benchfib/fibloop"
	"github.com/loicalbertin/go-tests-training/benchfib/fibrec"
)

func BenchmarkFibChan(b *testing.B) {
	benchmarks := []struct {
		name   string
		rec    uint
		buffer uint
	}{
		{"Unbuff10Rec", 10, 0},
		{"Unbuff50Rec", 50, 0},
		{"Unbuff100Rec", 100, 0},
		{"MidBuff10Rec", 10, 5},
		{"MidBuff50Rec", 50, 25},
		{"MidBuff100Rec", 100, 50},
		{"FullBuff10Rec", 10, 10},
		{"FullBuff50Rec", 50, 50},
		{"FullBuff100Rec", 100, 100},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(bi *testing.B) {
			for i := 0; i < bi.N; i++ {
				fibchan.PrintFib(bm.rec, bm.buffer)
			}
		})
	}
}

func BenchmarkFibLoop(b *testing.B) {
	benchmarks := []struct {
		name string
		rec  int
	}{
		{"10Rec", 10},
		{"50Rec", 50},
		{"100Rec", 100},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(bi *testing.B) {
			for i := 0; i < bi.N; i++ {
				fibloop.PrintFib(bm.rec)
			}
		})
	}
}

func BenchmarkFibRec(b *testing.B) {
	benchmarks := []struct {
		name string
		rec  int
	}{
		{"10Rec", 10},
		{"50Rec", 50},
		// {"100Rec", 100},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(bi *testing.B) {
			for i := 0; i < bi.N; i++ {
				fibrec.PrintFib(bm.rec)
			}
		})
	}
}
