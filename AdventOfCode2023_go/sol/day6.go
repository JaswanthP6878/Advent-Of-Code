package sol

import (
	"strconv"
	"strings"
)

type raceRecord struct {
	time   int
	length int
}

func Day6_1(data []string) int {
	races := readRaceRecords(data)
	total_possible := 1
	for _, race := range races {
		total_possible *= getPossibleRecordBreaks(race)
	}
	return total_possible
}

// per race, return the possible record breaks;
func getPossibleRecordBreaks(race raceRecord) int {
	possible := 0
	for i := 0; i <= race.time; i++ {
		speed := i
		time_left := race.time - i
		if speed*time_left > race.length {
			possible++
		}
	}
	return possible

}

// changed it for day6_2
// need to revert changes to get the day6_1
func readRaceRecords(data []string) []raceRecord {
	times := data[0]
	times_int := []int{}
	str_split := strings.Split(times, ": ")
	time_split := strings.Split(str_split[1], " ")
	var sb strings.Builder
	for _, val := range time_split {
		if len(val) != 0 {
			sb.WriteString(val)
		}
	}
	val, _ := strconv.Atoi(sb.String())
	times_int = append(times_int, val)

	dist := data[1]
	dist_int := []int{}
	str_split = strings.Split(dist, ": ")
	dist_split := strings.Split(str_split[1], " ")
	var pb strings.Builder
	for _, val := range dist_split {
		if len(val) != 0 {
			pb.WriteString(val)
		}
	}
	val, _ = strconv.Atoi(pb.String())
	dist_int = append(dist_int, val)

	raceRecords := []raceRecord{}
	for i := 0; i < len(times_int); i++ {
		new_record := raceRecord{}
		new_record.time = times_int[i]
		new_record.length = dist_int[i]

		raceRecords = append(raceRecords, new_record)
	}

	return raceRecords
}
