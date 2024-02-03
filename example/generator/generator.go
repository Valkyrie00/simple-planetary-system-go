package main

import (
	"encoding/json"
	"fmt"
	sps "github.com/Valkyrie00/simple-planetary-system-go"
	"os"
)

const verbose = false

func main() {
	// Define seed
	seed := int64(123456)

	// Define the system structure
	structure := &sps.SystemStructure{
		MinDensity:      500,   // Minimum density of planets
		DensityFactor:   3,     // Density factor of planets
		DeviationFactor: 0.1,   // Deviation Gaussian factor - Increase this value to increase the deviation of the planets
		Size:            600,   // Size of the system
		Seed:            &seed, // System Seed
	}

	// Generate the system
	base, err := sps.NewPlanetarySystem(structure)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	system, _ := base.Generate()
	fmt.Println("System generated with", len(system.Planets), "planets")

	// Debug
	if verbose {
		// Print System info
		fmt.Println(system.ID.String(), system.MinDensity, system.DensityFactor, system.DeviationFactor, system.Size)

		// Print the planet position
		for _, planet := range system.Planets {
			fmt.Println(planet.X, planet.Y, planet.Z)
		}
	}

	// Export the system to json
	exportToJson(system)
	os.Exit(0)
}

// Export result to json and save it to a file
func exportToJson(system *sps.PlanetarySystem) {
	jsonData, _ := json.Marshal(system)
	_ = os.WriteFile("../viewer/system.json", jsonData, 0644)
}
