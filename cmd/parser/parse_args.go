package parser

import (
	"fmt"
	"slices"
	"strconv"
)

var (
	valid = []string{
		"-c",
		"--count",
		"-w",
		"--width",
		"-h",
		"--height",
	}
)

const (
	defCount  = 2000
	defWidth  = 1920
	defHeight = 1080
)

func ParseArgs(args []string) (int, int32, int32, error) {
	countS := ""
	widthS := ""
	heightS := ""

	if len(args) == 0 {
		return defCount, defWidth, defHeight, nil
	}

	for i := range len(args) {
		arg := args[i]
		if !slices.Contains(valid, arg) {
			fmt.Printf("Invalid argument: %s\n", arg)
			continue
		}

		if i+1 >= len(args) {
			fmt.Printf("Missing value for argument: %s\n", arg)
			continue
		}

		value := args[i+1]
		switch arg {
		case "-c", "--count":
			countS = value
		case "-w", "--width":
			widthS = value
		case "-h", "--height":
			heightS = value
		}

		if heightS == "" && widthS == "" && countS == "" {
			error := fmt.Errorf("No valid arguments provided.")
			return 0, 0, 0, error
		}
	}

	count, err := strconv.ParseInt(countS, 10, 32)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("Invalid count value: %s", countS)
	}

	width, err := strconv.ParseInt(widthS, 10, 32)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("Invalid width value: %s", widthS)
	}

	height, err := strconv.ParseInt(heightS, 10, 32)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("Invalid height value: %s", heightS)
	}

	return int(count), int32(width), int32(height), nil
}
