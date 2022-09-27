package main

import (
	"fmt"
	"math"
)

const COUNT, FIRST, CENTER = 5, 0, 2
const RED, GREEN, WHITE, YELLOW, BLUE = 0, 1, 2, 3, 4
const BRIT, SWEDE, DANE, NORWEGIAN, GERMAN = 0, 1, 2, 3, 4
const TEA, COFFEE, MILK, JUICE, WATER = 0, 1, 2, 3, 4
const PALL_MALL, DUNHILL, BLENDS, BLUE_MASTER, PRINCE = 0, 1, 2, 3, 4
const DOG, BIRD, CAT, HORSE = 0, 1, 2, 3
const COLOR, NATIONALITY, DRINK, TOBACCO, PET = 0, 1, 2, 3, 4

var rowColors = [COUNT]int {0, 1, 2, 3, 4}
var rowNationalities = [COUNT]int {0, 1, 2, 3, 4}
var rowDrinks = [COUNT]int {0, 1, 2, 3, 4}
var rowTobaccos = [COUNT]int {0, 1, 2, 3, 4}
var rowPets = [COUNT]int {0, 1, 2, 3, 4}

var colors = [COUNT]string {"red", "green", "white", "yellow", "blue"}
var nationalities = [COUNT]string {"Brit", "Swede", "Dane", "Norwegian", "German"}
var drinks = [COUNT]string {"tea", "coffee", "milk", "juice", "water"}
var tobaccos = [COUNT]string {"Pall_Mall", "Dunhill", "Blends", "Blue_Master", "Prince"}
var pets = [COUNT]string {"dog", "bird", "cat", "horse", "fish"}

func main() {
	setColors(0)
}

func setColors(index int) {
	if COUNT == index {
		if 1 == rowColors[WHITE] - rowColors[GREEN] {
			setNationalities(0)
		}
		return
	}

	setColors(index + 1)
	for i := index + 1; i < COUNT; i++ {
		rowColors[index], rowColors[i] = rowColors[i], rowColors[index]
		setColors(index + 1)
		rowColors[index], rowColors[i] = rowColors[i], rowColors[index]
	}
}

func setNationalities(index int) {
	if COUNT == index {
		if rowNationalities[BRIT] == rowColors[RED] &&
				FIRST == rowNationalities[NORWEGIAN] &&
				1 == math.Abs(float64(rowNationalities[NORWEGIAN]-rowColors[BLUE])) {
			setDrinks(0)
		}
		return
	}

	setNationalities(index + 1)
	for i := index + 1; i < COUNT; i++ {
		rowNationalities[index], rowNationalities[i] = rowNationalities[i], rowNationalities[index]
		setNationalities(index + 1)
		rowNationalities[index], rowNationalities[i] = rowNationalities[i], rowNationalities[index]
	}
}

func setDrinks(index int) {
	if COUNT == index {
		if rowNationalities[DANE] == rowDrinks[TEA] &&
				rowColors[GREEN] == rowDrinks[COFFEE] &&
				CENTER == rowDrinks[MILK] {
			setTobaccos(0)
		}
		return
	}

	setDrinks(index + 1)
	for i := index + 1; i < COUNT; i++ {
		rowDrinks[index], rowDrinks[i] = rowDrinks[i], rowDrinks[index]
		setDrinks(index + 1)
		rowDrinks[index], rowDrinks[i] = rowDrinks[i], rowDrinks[index]
	}
}

func setTobaccos(index int) {
	if COUNT == index {
		if rowColors[YELLOW] == rowTobaccos[DUNHILL] &&
				rowTobaccos[BLUE_MASTER] == rowDrinks[JUICE] &&
				rowNationalities[GERMAN] == rowTobaccos[PRINCE] &&
				1 == math.Abs(float64(rowTobaccos[BLENDS]-rowDrinks[WATER])) {
			setPets(0)
		}
		return
	}

	setTobaccos(index + 1)
	for i := index + 1; i < COUNT; i++ {
		rowTobaccos[index], rowTobaccos[i] = rowTobaccos[i], rowTobaccos[index]
		setTobaccos(index + 1)
		rowTobaccos[index], rowTobaccos[i] = rowTobaccos[i], rowTobaccos[index]
	}
}

func setPets(index int) {
	if COUNT == index {
		if rowNationalities[SWEDE] == rowPets[DOG] &&
				rowTobaccos[PALL_MALL] == rowPets[BIRD] &&
				1 == math.Abs(float64(rowTobaccos[BLENDS]-rowPets[CAT])) &&
				1 == math.Abs(float64(rowPets[HORSE] - rowTobaccos[DUNHILL])) {
			printResult()
		}
		return
	}

	setPets(index + 1)
	for i := index + 1; i < COUNT; i++ {
		rowPets[index], rowPets[i] = rowPets[i], rowPets[index]
		setPets(index + 1)
		rowPets[index], rowPets[i] = rowPets[i], rowPets[index]
	}
}

func printResult() {
	owers := [COUNT][COUNT]int{}
	for i := 0; i < COUNT; i++ {
		owers[COLOR][rowColors[i]] = i
		owers[NATIONALITY][rowNationalities[i]] = i
		owers[DRINK][rowDrinks[i]] = i
		owers[TOBACCO][rowTobaccos[i]] = i
		owers[PET][rowPets[i]] = i
	}

	for i := 0; i < COUNT; i++ {
		fmt.Println(colors[owers[COLOR][i]] + " " +
				nationalities[owers[NATIONALITY][i]] + " " +
				drinks[owers[DRINK][i]] + " " +
				tobaccos[owers[TOBACCO][i]] + " " +
				pets[owers[PET][i]])
	}
}
