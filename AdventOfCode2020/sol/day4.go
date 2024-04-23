package sol

import (
	"fmt"
	"strconv"
	"strings"
)

func Day4(data []string) int {
	count_valid_passport := 0
	values := []string{}
	for i := range data {
		if len(data[i]) == 0 {
			if isValidPassport(values) {
				count_valid_passport++
			}
			values = []string{}
		}
		values = append(values, data[i])
	}
	if isValidPassport(values) {
		count_valid_passport++
	}

	return count_valid_passport
}

func isValidPassport(data []string) bool {
	fmt.Println(data)
	key_check := map[string]string{
		"byr": "",
		"iyr": "",
		"eyr": "",
		"hgt": "",
		"hcl": "",
		"ecl": "",
		"pid": "",
		// "cid": false,
	}
	for i := range data {
		key_values := strings.Split(data[i], " ")
		fmt.Println(key_values)
		for j := range key_values {
			key_arr := strings.Split(key_values[j], ":")
			key_check[key_arr[0]] = key_arr[1]
		}
	}
	return isValidFormat(key_check)
}

func isValidFormat(data map[string]string) bool {
	valid_items := 0
	for k, v := range data {
		if k == "byr" {
			val, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			if (val > 1920 && val < 2002) || val == 1920 || val == 2002 {
				valid_items++
			}
		}

		if k == "iyr" {
			val, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			if (val > 2010 && val < 2020) || val == 2010 || val == 2020 {
				valid_items++
			}
		}

		if k == "eyr" {
			val, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			if (val > 2020 && val < 2030) || val == 2020 || val == 2030 {
				valid_items++
			}
		}

		if k == "hgt" {
			if v[len(v)-2:] == "in" {
				val, err := strconv.Atoi(v[0 : len(v)-2])
				if err != nil {
					fmt.Printf("cannot convert to string height %s", v)
					return false
				}
				if (val > 59 && val < 76) || val == 59 || val == 76 {
					valid_items++
				}

			} else if v[len(v)-2:] == "cm" {
				val, err := strconv.Atoi(v[0 : len(v)-2])
				if err != nil {
					fmt.Printf("cannot convert to string height %s", v)
					return false
				}
				if (val > 150 && val < 193) || val == 150 || val == 193 {
					valid_items++
				}

			} else {
				return false
			}
		}

		if k == "hcl" {
			if v[0] == '#' {
				if len(v[1:]) == 6 {
					for _, val := range v[1:] {
						if (val > '0' && val < '9') || val == '0' || val == '9' || (val > 'a' && val < 'f') || val == 'a' || val == 'f' {
							continue
						} else {
							return false
						}
					}
					valid_items++
				}
			} else {
				return false
			}
		}
		if k == "ecl" {
			if v == "amb" || v == "blu" || v == "brn" || v == "gry" || v == "hzl" || v == "oth" || v == "grn" {
				valid_items++
			} else {
				return false
			}
		}

		if k == "pid" {
			if len(v) == 9 {
				_, err := strconv.Atoi(v)
				if err != nil {
					return false
				}
				valid_items++
			} else {
				return false
			}
		}
	}
	return valid_items == 7
}
