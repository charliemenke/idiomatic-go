package main

import (
	"errors"
	"io"
	"os"
	"sort"
)

type Ranker interface {
    Ranking() []string
}

type Team struct {
    TeamName    string
    PlayerNames []string
}

type League struct {
    Name    string
    Teams   []Team
    Wins    map[string]int
}
func (l *League) MatchResult(t1 string, s1 int, t2 string, s2 int) error {
    if l == nil {
       return errors.New("League is nil") 
    }
    if s1 > s2 {
        l.Wins[t1]++
    } else if s2 > s1 {
        l.Wins[t2]++
    }
    return nil
}
func (l *League) Ranking() []string {
    sortedSlice := make([]string, 0, len(l.Teams))
    for _, t := range l.Teams {
        sortedSlice = append(sortedSlice, t.TeamName)
    }
    sort.Slice(sortedSlice, func(i, j int) bool {
       return l.Wins[sortedSlice[i]] > l.Wins[sortedSlice[j]]
    })
    return sortedSlice
}

func main() {
l := League{
		Name: "Big League",
		Teams: []Team{
			{
				TeamName:    "Italy",
				PlayerNames: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			{
				TeamName:    "France",
				PlayerNames: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
            {
				TeamName:    "India",
				PlayerNames: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
            {
				TeamName:    "Nigeria",
				PlayerNames: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
    RankPrinter(&l, os.Stdout)
}

func RankPrinter(ranker Ranker, writer io.Writer) {
    for _, v := range ranker.Ranking() {
        io.WriteString(writer, v)
        writer.Write([]byte("\n"))
    }
}
