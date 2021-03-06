package main

import (
	"strconv"
	"strings"
	"testing"
)

var getRollCases = []struct {
	name    string
	ceiling int
}{
	{
		name:    "d4",
		ceiling: 4,
	},
	{
		name:    "d6",
		ceiling: 6,
	},
	{
		name:    "d8",
		ceiling: 8,
	},
	{
		name:    "d10",
		ceiling: 10,
	},
	{
		name:    "d12",
		ceiling: 12,
	},
	{
		name:    "d20",
		ceiling: 20,
	},
	{
		name:    "d100",
		ceiling: 100,
	},
}

func Test_getRoll(t *testing.T) {
	for _, tt := range getRollCases {
		t.Run(tt.name, func(t *testing.T) {
			out := getRoll(tt.ceiling)
			if out < 1 || out > tt.ceiling {
				t.Errorf("Roll out of range: %d of ceiling %d", out, tt.ceiling)
			}
		})
	}
}

var parseDiceCases = []struct {
	raw      string
	ceiling  int
	quantity int
	wantErr  bool
}{
	{
		raw:      "1d6",
		ceiling:  6,
		quantity: 1,
		wantErr:  false,
	},
	{
		raw:      "4d6",
		ceiling:  6,
		quantity: 4,
		wantErr:  false,
	},
	{
		raw:      "8d8",
		ceiling:  8,
		quantity: 8,
		wantErr:  false,
	},
	{
		raw:      "10d10",
		ceiling:  10,
		quantity: 10,
		wantErr:  false,
	},
	{
		raw:      "12d12",
		ceiling:  12,
		quantity: 12,
		wantErr:  false,
	},
	{
		raw:      "2d30",
		ceiling:  30,
		quantity: 2,
		wantErr:  true,
	},
	{
		raw:      "4d20",
		ceiling:  20,
		quantity: 4,
		wantErr:  false,
	},
	{
		raw:      "6d100",
		ceiling:  100,
		quantity: 6,
		wantErr:  false,
	},
	{
		raw:      "1d20+4",
		ceiling:  20,
		quantity: 1,
		wantErr:  false,
	},
	{
		raw:      "1d8+20",
		ceiling:  8,
		quantity: 1,
		wantErr:  false,
	},
	{
		raw:      "4d4+1",
		ceiling:  4,
		quantity: 4,
		wantErr:  false,
	},
	{
		raw:      "1d69",
		ceiling:  69,
		quantity: 1,
		wantErr:  false,
	},
	{
		raw:      "101d10",
		ceiling:  10,
		quantity: 101,
		wantErr:  true,
	},
}

func Test_parseDice(t *testing.T) {
	for _, tt := range parseDiceCases {
		t.Run(tt.raw, func(t *testing.T) {
			out, err := parseDice(tt.raw)
			if tt.raw == "1d69" && out != "n i c e" {
				t.Errorf("69 didn't return n i c e")
			}
			if tt.raw == "1d69" && out == "n i c e" {
				return
			}
			if err != nil && !tt.wantErr {
				t.Errorf("Got unexpected error: %s", err.Error())
			}
			if err == nil && tt.wantErr {
				t.Errorf("Expected error, got nil: %s", tt.raw)
			}
			if err != nil && tt.wantErr {
				return
			}

			split := strings.Split(out, ",")
			split = strings.Split(split[0], "  ")
			for i := 0; i < len(split); i++ {
				die := split[i]
				num, _ := strconv.Atoi(die)
				if num > tt.ceiling || num < 1 {
					t.Errorf("Roll out of range: %d", num)
				}
			}
		})
	}
}

func Benchmark_parseDice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range parseDiceCases {
			parseDice(tt.raw)
		}
	}
}
