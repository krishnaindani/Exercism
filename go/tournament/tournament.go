package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type stats struct {
	name          string
	matchesPlayed int
	win           int
	draw          int
	loss          int
	points        int
}

type scoreBoard map[string]stats

//Tally tallies the matches
func Tally(r io.Reader, w io.Writer) error {

	scanner := bufio.NewScanner(r)
	sb := make(scoreBoard)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Split(line, ";")
		if len(fields) != 3 {
			return fmt.Errorf("invalid game: %s", line)
		}

		teamA := fields[0]
		teamB := fields[1]
		status := fields[2]

		teamAStats := sb[teamA]
		teamBStats := sb[teamB]

		teamAStats.name = teamA
		teamBStats.name = teamB

		teamAStats.matchesPlayed++
		teamBStats.matchesPlayed++

		switch status {
		case "win":
			teamAStats.points += 3
			teamAStats.win++
			teamBStats.loss++
		case "loss":
			teamAStats.loss++
			teamBStats.points += 3
			teamBStats.win++
		case "draw":
			teamAStats.points++
			teamBStats.points++
			teamAStats.draw++
			teamBStats.draw++
		default:
			return fmt.Errorf("unknown condition: %s", line)
		}

		sb[teamA] = teamAStats
		sb[teamB] = teamBStats
	}

	allTeams := make([]stats, 0, len(sb))

	for _, v := range sb {
		allTeams = append(allTeams, v)
	}

	sort.Slice(allTeams, func(i, j int) bool {
		if allTeams[i].points == allTeams[j].points {
			return allTeams[i].name < allTeams[j].name
		}
		return allTeams[i].points > allTeams[j].points
	})

	header := "Team                           | MP |  W |  D |  L |  P"
	fmtString := "%-31s| %2d | %2d | %2d | %2d | %2d\n"

	fmt.Fprintln(w, header)
	for _, team := range allTeams {
		fmt.Fprintf(w, fmtString,
			team.name,
			team.matchesPlayed,
			team.win,
			team.draw,
			team.loss,
			team.points,
		)
	}

	return nil
}
