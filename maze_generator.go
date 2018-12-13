package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	var seed = generate_seed()
	var maze = generate_layout(seed)
	fmt.Println(maze)
}

func generate_seed() int {
	return rand.Int()
}

func generate_layout(seed int) string {
	var x = 0
	var layers = 0
	for layers == 0 {
		layers = get_n_digit(seed, x)
		x += 1
	}
	var y = 0
	for y < layers {
		fmt.Println(y)
		y += 1
	}
	return "hah"
}

func get_n_digit(seed int, digit int) int {
	var seed_string = strconv.Itoa(seed)[digit:]
	var new_seed, err = strconv.Atoi(seed_string)
	if err != nil {
		fmt.Println(err)
	}
	var value = int(float64(new_seed) / math.Pow(10, float64(len(seed_string)-1)))
	return value
}
