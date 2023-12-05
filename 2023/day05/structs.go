package main

/**
Step is a single step in the mapping process. Initially, the value is the seed.
*/
type Stage int

const (
	Seed Stage = iota
	Soil
	Fertilizer
	Water
	Light
	Temperature
	Humidity
	Location
)

var (
	StageMap = map[string]Stage{
		"seed":        Seed,
		"soil":        Soil,
		"fertilizer":  Fertilizer,
		"water":       Water,
		"light":       Light,
		"temperature": Temperature,
		"humidity":    Humidity,
		"location":    Location,
	}
)

type Step struct {
	Value int
	Type  Stage
}

/**
Corresponds to a single line (50 98 2) where
DestStart is 50
SourceStart is 98 (these are the seeds)
Length is 2
Destinations are 50, 51. Source is 98, 99.
This means that destination 50 is connected to source 98, and destination 51 is connected to source 99...
Any destination/source that is not explicitly listed is connected to itself (60 is connected to 60)
*/
type Route struct {
	DestStart   int
	SourceStart int
	Length      int
	Type        Stage
}

/**
List of routes, one per line
*/
type Routemap struct {
	Entries []Route
	Type    Stage
}

/**
Return if the step can be resolved in the route
*/
func (r Route) CanSolve(step Step) bool {
	for i := 0; i < r.Length; i++ {
		if step.Value == r.SourceStart+i {
			return true
		}
	}
	return false
}

/**
process a step through a route. The step will be updated to the destination of the route.
If the step is not in the Route, it will return step with the same Value, but Type will be the same as the Routes
*/
func (r Route) Solve(step Step) Step {
	if r.CanSolve(step) {
		return Step{Value: r.DestStart + (step.Value - r.SourceStart), Type: r.Type}
	}
	return Step{Value: step.Value, Type: r.Type}
}

func (m Routemap) Solve(step Step) Step {
	for _, route := range m.Entries {
		if route.CanSolve(step) {
			return route.Solve(step)
		}
	}
	return Step{Value: step.Value, Type: m.Type}
}
