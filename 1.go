package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func follow_directions(d []string) (int, int, int, int) {
	var repeated_x, repeated_y int = 0, 0
	var x_pos, y_pos int = 0, 0
	var orientation string = "N"
	var found bool = false

	position_tracker := map[int]map[int]bool{
		0: {
			0: true,
		},
	}

	// make a simple state mapping
	orientation_map := map[string]map[string]string{
		"N": {
			"L": "W",
			"R": "E",
		},
		"E": {
			"L": "N",
			"R": "S",
		},
		"S": {
			"L": "E",
			"R": "W",
		},
		"W": {
			"L": "S",
			"R": "N",
		},
	}

	for _, raw_op := range d {
		raw_op = strings.Trim(raw_op, " ")
		raw_count := string(raw_op[1:])

		op := string(raw_op[0])
		count, err := strconv.Atoi(raw_count)
		check(err)

		// now let's find out which way to turn/proceed
		orientation = orientation_map[orientation][op]

		// while loop while moving to track visited coordinates
		for iter := 0; iter < count; iter++ {
			switch orientation {
			case "N":
				y_pos = y_pos + 1
			case "E":
				x_pos = x_pos + 1
			case "S":
				y_pos = y_pos - 1
			case "W":
				x_pos = x_pos - 1
			}

			if found != false {
				continue
			}
			// haven't found first repeated coordinate - check if we've been here
			if val, ok := position_tracker[x_pos][y_pos]; ok && val == true {
				repeated_x = x_pos
				repeated_y = y_pos
				found = true
			} else if val, ok := position_tracker[x_pos]; ok && val != nil {
				position_tracker[x_pos][y_pos] = true
			} else {
				position_tracker[x_pos] = map[int]bool{y_pos: true}
			}
		}
	}
	return x_pos, y_pos, repeated_x, repeated_y
}

func main() {
	// get the filename from stdin
	args := os.Args[1:]
	var dat string

	if len(args) > 0 {
		if _, err := os.Stat(args[0]); err == nil {
			// read in the input
			raw_dat, err := ioutil.ReadFile(args[0])
			check(err)
			dat = string(raw_dat)
		} else {
			// treat the input as string
			dat = strings.Join(args, ",")
		}

	} else {
		fmt.Println("Please specify an input file or string")
		return
	}

	// get a slice of just each direction
	directions := strings.Split(dat, ",")

	// get the final (x, y) coord based on directions
	x, y, repeated_x, repeated_y := follow_directions(directions)

	// manhattan distance on abs(x) and abs(y) since we assume (0,0) start
	distance := math.Abs(float64(x)) + math.Abs(float64(y))
	repeated_distance := math.Abs(float64(repeated_x)) + math.Abs(float64(repeated_y))
	fmt.Println(distance)
	fmt.Println(repeated_distance)

}
