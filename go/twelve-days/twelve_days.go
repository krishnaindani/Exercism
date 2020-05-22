package twelve

import (
	"fmt"
	"strings"
)

type dayToSong struct {
	day    string
	lyrics string
}

var d = []dayToSong{
	{
		day:    "first",
		lyrics: "a Partridge in a Pear Tree.",
	}, {
		day:    "second",
		lyrics: "two Turtle Doves",
	}, {
		day:    "third",
		lyrics: "three French Hens",
	}, {
		day:    "fourth",
		lyrics: "four Calling Birds",
	}, {
		day:    "fifth",
		lyrics: "five Gold Rings",
	}, {
		day:    "sixth",
		lyrics: "six Geese-a-Laying",
	}, {
		day:    "seventh",
		lyrics: "seven Swans-a-Swimming",
	}, {
		day:    "eighth",
		lyrics: "eight Maids-a-Milking",
	}, {
		day:    "ninth",
		lyrics: "nine Ladies Dancing",
	}, {
		day:    "tenth",
		lyrics: "ten Lords-a-Leaping",
	}, {
		day:    "eleventh",
		lyrics: "eleven Pipers Piping",
	}, {
		day:    "twelfth",
		lyrics: "twelve Drummers Drumming",
	},
}

const startLine = "On the %s day of Christmas my true love gave to me:"

//Song ...
func Song() string {
	return ""
}

//Verse ...
func Verse(day int) string {
	lines := []string{fmt.Sprintf(startLine, d[day-1].day)}
	for i := day - 1; i >= 0; i-- {
		if i == 0 && day != 1 {
			lines = append(lines, "and "+d[i].lyrics)
		} else {
			lines = append(lines, " "+d[i].lyrics)
		}
	}
	fmt.Println("lines:", lines)
	return strings.Join(lines, "")
}
