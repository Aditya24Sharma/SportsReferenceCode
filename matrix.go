package main

import (
	"fmt"
	"sort"
)

type Record struct {
	W int `json:"W"`
	L int `json:"L"`
}

type TeamData map[string]Record

type HeadToHeadData map[string]TeamData

func Matrix(data HeadToHeadData) {
	teams := make([]string, 0, len(data))
	for team := range data {
		teams = append(teams, team)
	}
	sort.Strings(teams)

	fmt.Printf("%-5s", "Tm")
	for _, team := range teams {
		fmt.Printf("%5s", team)
	}
	fmt.Println()

	for _, team := range teams {
		fmt.Printf("%-5s", team)

		for _, opponent := range teams {
			if team == opponent {
				fmt.Printf("%5s", "--")
				continue
			}

			if record, ok := data[team][opponent]; ok {
				fmt.Printf("%5d", record.W)
			} else {
				fmt.Printf("%5s", "-")
			}
		}
		fmt.Println()
	}
}

func main() {
	data := HeadToHeadData{
		"BRO": {
			"BSN": {W: 10, L: 12},
			"CHC": {W: 15, L: 7},
			"CIN": {W: 15, L: 7},
			"NYG": {W: 14, L: 8},
			"PHI": {W: 14, L: 8},
			"PIT": {W: 15, L: 7},
			"STL": {W: 11, L: 11},
		},
		"BSN": {
			"BRO": {W: 12, L: 10},
			"CHC": {W: 13, L: 9},
			"CIN": {W: 13, L: 9},
			"NYG": {W: 13, L: 9},
			"PHI": {W: 14, L: 8},
			"PIT": {W: 12, L: 10},
			"STL": {W: 9, L: 13},
		},
		"CHC": {
			"BRO": {W: 7, L: 15},
			"BSN": {W: 9, L: 13},
			"CIN": {W: 12, L: 10},
			"NYG": {W: 7, L: 15},
			"PHI": {W: 16, L: 6},
			"PIT": {W: 8, L: 14},
			"STL": {W: 10, L: 12},
		},
		"CIN": {
			"BRO": {W: 7, L: 15},
			"BSN": {W: 9, L: 13},
			"CHC": {W: 10, L: 12},
			"NYG": {W: 13, L: 9},
			"PHI": {W: 13, L: 9},
			"PIT": {W: 13, L: 9},
			"STL": {W: 8, L: 14},
		},
		"NYG": {
			"BRO": {W: 8, L: 14},
			"BSN": {W: 9, L: 13},
			"CHC": {W: 15, L: 7},
			"CIN": {W: 9, L: 13},
			"PHI": {W: 12, L: 10},
			"PIT": {W: 15, L: 7},
			"STL": {W: 13, L: 9},
		},
		"PHI": {
			"BRO": {W: 8, L: 14},
			"BSN": {W: 8, L: 14},
			"CHC": {W: 6, L: 16},
			"CIN": {W: 9, L: 13},
			"NYG": {W: 10, L: 12},
			"PIT": {W: 13, L: 9},
			"STL": {W: 8, L: 14},
		},
		"PIT": {
			"BRO": {W: 7, L: 15},
			"BSN": {W: 10, L: 12},
			"CHC": {W: 14, L: 8},
			"CIN": {W: 9, L: 13},
			"NYG": {W: 7, L: 15},
			"PHI": {W: 9, L: 13},
			"STL": {W: 6, L: 16},
		},
		"STL": {
			"BRO": {W: 11, L: 11},
			"BSN": {W: 13, L: 9},
			"CHC": {W: 12, L: 10},
			"CIN": {W: 14, L: 8},
			"NYG": {W: 9, L: 13},
			"PHI": {W: 14, L: 8},
			"PIT": {W: 16, L: 6},
		},
	}

	Matrix(data)
}
