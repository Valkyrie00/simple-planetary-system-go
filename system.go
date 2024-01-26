package sps

import (
	"errors"
	"github.com/google/uuid"
	"math/rand"
)

type PlanetSystemInterface interface {
	Generate() (system *PlanetarySystem, err error)
}

type PlanetarySystem struct {
	ID              uuid.UUID `json:"id"`
	SystemStructure `json:"system_structure"`
	Planets         []Planet `json:"planets"`
}

type SystemStructure struct {
	MinDensity      int32   `json:"min_density"`
	DensityFactor   int32   `json:"density"`
	DeviationFactor float64 `json:"deviation"`
	Size            float64 `json:"size"`
	Seed            *int64  `json:"seed"`
}

func NewPlanetarySystem(structure *SystemStructure) (PlanetSystemInterface, error) {
	if err := structure.VerifyStructure(); err != nil {
		return nil, err
	}

	return &PlanetarySystem{
		ID:              uuid.New(),
		SystemStructure: *structure,
	}, nil
}

func (s *PlanetarySystem) Generate() (system *PlanetarySystem, err error) {
	r := rand.New(rand.NewSource(*s.SystemStructure.Seed))
	chunkDensity := r.Int31n(
		s.SystemStructure.MinDensity*s.SystemStructure.DensityFactor-s.SystemStructure.MinDensity+1) +
		s.SystemStructure.MinDensity

	for i := 0; i < int(chunkDensity); i++ {
		var planet *Planet
		if planet, err = NewPlanet().Generate(s, i); err != nil {
			return nil, err
		}

		s.Planets = append(s.Planets, *planet)
	}

	return s, nil
}

func (s *SystemStructure) VerifyStructure() error {
	if s.MinDensity < 0 {
		return errors.New("invalid min density")
	}

	if s.DensityFactor < 0 {
		return errors.New("invalid density factor")
	}

	if s.DeviationFactor < 0 {
		return errors.New("invalid deviation factor")
	}

	if s.Size < 0 {
		return errors.New("invalid size")
	}

	if s.Seed == nil {
		return errors.New("seed is required to generate the system")
	}

	return nil
}
