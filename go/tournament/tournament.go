package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type stats struct {
	matchedPlayed int
	win           int
	draw          int
	lose          int
	points        int
}

type teamStats struct {
	AllegoricAlaskians     stats
	DevastatingDonkeys     stats
	BlitheringBadgers      stats
	CourageousCalifornians stats
}

type (
	m map[string]stats
)

//Tally tallies the matches
func Tally(r io.Reader, w io.Writer) error {

	s, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.New("error reading input data")
	}

	t := teamStats{
		AllegoricAlaskians:     stats{},
		DevastatingDonkeys:     stats{},
		BlitheringBadgers:      stats{},
		CourageousCalifornians: stats{},
	}

	c := strings.Split(string(s), "\n")

	for _, element := range c {
		var a []string
		a = strings.Split(string(element), ";")
		if len(a) > 1 {
			status := a[2]
			team := a[0]
			secondTeam := a[1]

			switch status {
			case "win":
				switch team {
				case "Allegoric Alaskians":
					t.AllegoricAlaskians.win++
				case "Devastating Donkeys":
					t.DevastatingDonkeys.win++
				case "Courageous Californians":
					t.CourageousCalifornians.win++
				case "Blithering Badgers":
					t.BlitheringBadgers.win++
				}
			case "loss":
				switch team {
				case "Allegoric Alaskians":
					t.AllegoricAlaskians.lose++
				case "Devastating Donkeys":
					t.DevastatingDonkeys.lose++
				case "Courageous Californians":
					t.CourageousCalifornians.lose++
				case "Blithering Badgers":
					t.BlitheringBadgers.lose++
				}
			case "draw":
				switch team {
				case "Allegoric Alaskians":
					t.AllegoricAlaskians.draw++
				case "Devastating Donkeys":
					t.DevastatingDonkeys.draw++
				case "Courageous Californians":
					t.CourageousCalifornians.draw++
				case "Blithering Badgers":
					t.BlitheringBadgers.draw++
				}

				switch secondTeam {
				case "Allegoric Alaskians":
					t.AllegoricAlaskians.draw++
				case "Devastating Donkeys":
					t.DevastatingDonkeys.draw++
				case "Courageous Californians":
					t.CourageousCalifornians.draw++
				case "Blithering Badgers":
					t.BlitheringBadgers.draw++
				}
			}
			switch secondTeam {
			case "Allegoric Alaskians":
				t.AllegoricAlaskians.matchedPlayed++
			case "Devastating Donkeys":
				t.DevastatingDonkeys.matchedPlayed++
			case "Courageous Californians":
				t.CourageousCalifornians.matchedPlayed++
			case "Blithering Badgers":
				t.BlitheringBadgers.matchedPlayed++
			}

			switch team {
			case "Allegoric Alaskians":
				t.AllegoricAlaskians.matchedPlayed++
			case "Devastating Donkeys":
				t.DevastatingDonkeys.matchedPlayed++
			case "Courageous Californians":
				t.CourageousCalifornians.matchedPlayed++
			case "Blithering Badgers":
				t.BlitheringBadgers.matchedPlayed++
			}
		}
	}
	fmt.Printf("%+v", t)
	return nil
}
