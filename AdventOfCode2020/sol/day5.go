package sol

import "fmt"

func Day5(data []string) int {
	max_seat_id := 0
	for _, val := range data {
		seat := GetSeatNumber(val)
		seatId := getSeatId(seat)
		if seatId > max_seat_id {
			max_seat_id = seatId
		}
	}

	return max_seat_id
}

// returns a array witj row number, col number
func GetSeatNumber(data string) []int {
	lr, hr := 0, 127
	for _, val := range data {
		fmt.Println(lr, hr, string(val))
		if val == 'F' {
			hr = hr - (hr-lr)/2
		} else if val == 'B' {
			lr = lr + (hr-lr)/2
		}
	}
	return []int{lr, hr}
}

func getSeatId(data []int) int {
	return 0
}
