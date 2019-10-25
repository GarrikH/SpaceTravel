package main

import "fmt"
import "strings"

const FUEL_COST_PER_G float32 = 1.65

func fillUp(dollars float32) float32 {
	var gallons float32 = dollars / FUEL_COST_PER_G
	if dollars >= 1000000 && dollars < 1000000000 {
		dollars /= 1000000
		fmt.Printf("Your budget of $%.2fM buys %.3f gallons of rocket fuel.\n", dollars, gallons)
	} else if dollars >= 1000000000 {
		dollars /= 1000000000
		fmt.Printf("Your budget of $%.2fB buys %.3f gallons of rocket fuel.\n", dollars, gallons)
	} else {
		fmt.Printf("Your budget of $%.2f buys %.3f gallons of rocket fuel.\n", dollars, gallons)
	}
	return gallons
}

func fuelGauge(fuel float32) {
	fmt.Printf("%.3f gallons of fuel left.", fuel)
}

func calculateFuel(planet string) float32 {
	var fuel float32
	planet = strings.ToLower(planet)
	switch planet {
	case "venus":
		fuel = 300000
	case "mercury":
		fuel = 500000
	case "mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel
}

func greetPlanet(planet string) {
	fmt.Printf("Welcome to the planet %v!\n", planet)
}

func cantFly(planet string) {
	fmt.Printf("We do not have enough fuel to fly to %v.\n", planet)
}

func flyToPlanet(planet string, fuel float32) float32 {
	var fuelRemaining, fuelCost float32
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly(planet)
	}
	return fuelRemaining
}

func main() {

	// var budget float32 = 1380000000.00
	var budget float32 = 1380000.00
	// var budget float32 = 1380.00
	fuel := fillUp(budget)
	planetChoice := "Venus"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)

}