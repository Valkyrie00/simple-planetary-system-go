# Simple Planetary System Generator

This is a simple planetary system generator written in Go.
The generated planetary systems are not realistic or based on scientific data,
but they are useful for testing the game mechanics of a space game for example.
The entire generator is based on the Gaussian distribution and the generate planets cluster around the center of the system.

Here is an example of a generated planetary system:
![example-1.gif](assets%2Fexample-1.gif)

## Installation
```bash
go get github.com/Valkyrie00/simple-planetary-system-go
```

## Usage

You can use the generator in your project by importing the package:

```go
import sps "github.com/Valkyrie00/simple-planetary-system-go"
```
    
Then you can generate a planetary system with the following and easy steps.

First you need to define the structure of the planetary system. The structure is defined by the following parameters:
```go
seed := int64(123456) // Alternatively you can use time.Now().UnixNano()
structure := &sps.SystemStructure{
    MinDensity:      29,    // Minimum density of planets
    DensityFactor:   3,     // Density factor of planets
    DeviationFactor: 0.4,   // Deviation Gaussian factor - Increase this value to increase the deviation of the planets
    Size:            50,    // Size of the system
    Seed:            &seed, // System Seed
}
```

Then you can initialize the planetary system. This step verifies that the structure is valid and generates the planets.
```go
// Generate the system
base, err := sps.NewPlanetarySystem(structure)
if err != nil {
    fmt.Printf("Error: %v\n", err)
    os.Exit(1)
}
```

You can able now to generate a planetary system with a given seed and iterate over the planets:
```go
system, _ := base.Generate()
fmt.Println("System generated with", len(system.Planets), "planets")
for _, planet := range system.Planets {
    fmt.Println(planet)
}
```
You can find a full example [here](example/generator/generator.go).
Remember to use a valid seed to generate the same planetary system, to build or recreate the same planetary system.

## Visualize the planetary system
You can visualize the planetary system using the small web server included in the example.
This visualizer is based on [three.js](https://threejs.org/) and [vite](https://vitejs.dev/).

To generate the system you can use the generator example:
```bash
cd example/generator
go run generator.go
```
This will generate a planetary system in a new file called `system.json` under the `example/viewer` folder.

Then you can start the web server:
```bash
cd example/viewer
npm install
npx vite
```
Then open your browser at http://localhost:5173/ and you should see the planetary system.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Licenza
[MIT](https://choosealicense.com/licenses/mit/)