package sps

import (
	"crypto/md5"
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/google/uuid"
	"math"
	"math/rand"
)

type PlanetInterface interface {
	Generate(system *PlanetarySystem, planetIndex int) (planet *Planet, err error)
}

type Planet struct {
	ID           uuid.UUID `json:"id"`
	HashPosition string    `json:"hash_position"`
	X            float64   `json:"x"`
	Y            float64   `json:"y"`
	Z            float64   `json:"z"`
}

func NewPlanet() *Planet {
	return &Planet{}
}

func (p *Planet) Generate(system *PlanetarySystem, planetIndex int) (planet *Planet, err error) {
	// Create a new seed from the system seed and the planet index
	planetSeed := (*system.SystemStructure.Seed + int64(planetIndex)) % int64(math.MaxInt64)
	r := rand.New(rand.NewSource(planetSeed))

	// Generate a random position distributed by a Gaussian distribution
	positionDistributed := mgl64.Vec3{
		r.NormFloat64() * system.DeviationFactor * system.Size,
		r.NormFloat64() * system.DeviationFactor * system.Size,
		r.NormFloat64() * system.DeviationFactor * system.Size,
	}

	hashPosition := fmt.Sprintf("%x", md5.Sum(
		[]byte(fmt.Sprintf("%v,%v,%v", positionDistributed[0], positionDistributed[1], positionDistributed[2])),
	))

	return &Planet{
		ID:           uuid.New(),
		HashPosition: hashPosition,
		X:            positionDistributed[0],
		Y:            positionDistributed[1],
		Z:            positionDistributed[2],
	}, nil
}
