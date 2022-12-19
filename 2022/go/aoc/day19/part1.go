package day19

import (
	"fmt"
)

type Resources struct {
	Ore      uint8
	Clay     uint8
	Obsidian uint8
	Geode    uint8
}

func (r Resources) FingerPrint() uint64 {
	return uint64(r.Ore) | uint64(r.Clay)<<8 | uint64(r.Obsidian)<<16 | uint64(r.Geode)<<24
}

func (r1 Resources) LessOrEqual(r2 Resources) bool {
	return r1.Ore <= r2.Ore && r1.Clay <= r2.Clay && r1.Obsidian <= r2.Obsidian && r1.Geode <= r2.Geode
}

type Robots struct {
	Ore      uint8
	Clay     uint8
	Obsidian uint8
	Geode    uint8
}

func (r Robots) FingerPrint() uint64 {
	return uint64(r.Ore) | uint64(r.Clay)<<8 | uint64(r.Obsidian)<<16 | uint64(r.Geode)<<24
}

func (r1 Robots) LessOrEqual(r2 Robots) bool {
	return r1.Ore <= r2.Ore && r1.Clay <= r2.Clay && r1.Obsidian <= r2.Obsidian && r1.Geode <= r2.Geode
}

type Robot int

const (
	NoRobot Robot = iota
	Ore
	Clay
	Obsidian
	Geode
)

type Factory struct {
	// MaxRobots is used to minimize the search space. The factory
	// can produce 1 robot per minute. We don't need to produce more
	// ore than any robot costs. Equally we can limit the amount of
	// clay we need. When the factory would run longer the number of
	// obsidian robots could be limited too.
	MaxRobots Robots
	Robots    Robots
	Resources Resources
}

func (f Factory) FingerPrint() uint64 {
	return f.Robots.FingerPrint() | f.Resources.FingerPrint()<<32
}

func (f Factory) BuildableOptions(bluePrint BluePrint) []Robot {
	result := []Robot{}
	if f.Resources.Ore >= bluePrint.OreRobotCost && f.Robots.Ore <= f.MaxRobots.Ore {
		result = append(result, Ore)
	}
	if f.Resources.Ore >= bluePrint.ClayRobotCost && f.Robots.Clay <= f.MaxRobots.Clay {
		result = append(result, Clay)
	}
	if f.Resources.Ore >= bluePrint.ObsidianRobotCost.v1 && f.Resources.Clay >= bluePrint.ObsidianRobotCost.v2 {
		result = append(result, Obsidian)
	}
	if f.Resources.Ore >= bluePrint.GeodeRobotCost.v1 && f.Resources.Obsidian >= bluePrint.GeodeRobotCost.v2 {
		result = append(result, Geode)
	}

	return append(result, NoRobot)
}

func (f *Factory) Produce() {
	f.Resources.Ore += f.Robots.Ore
	f.Resources.Clay += f.Robots.Clay
	f.Resources.Obsidian += f.Robots.Obsidian
	f.Resources.Geode += f.Robots.Geode
}

func (f *Factory) Build(robot Robot, bluePrint BluePrint) {
	switch robot {
	case NoRobot:
	case Ore:
		f.Resources.Ore -= bluePrint.OreRobotCost
		f.Robots.Ore++
	case Clay:
		f.Resources.Ore -= bluePrint.ClayRobotCost
		f.Robots.Clay++
	case Obsidian:
		f.Resources.Ore -= bluePrint.ObsidianRobotCost.v1
		f.Resources.Clay -= bluePrint.ObsidianRobotCost.v2
		f.Robots.Obsidian++
	case Geode:
		f.Resources.Ore -= bluePrint.GeodeRobotCost.v1
		f.Resources.Obsidian -= bluePrint.GeodeRobotCost.v2
		f.Robots.Geode++
	}
}

// //     robots  - resources
//
//	0: 1 0 0 0 - 0 0 0 0
//	1: 1 0 0 0 - 1 0 0 0
//	2: 1 0 0 0 - 2 0 0 0
//	3: 1 1 0 0 - 1 0 0 0 | 1 0 0 0 - 3 0 0 0
//	4: 1 1 0 0 - 2 1 0 0 | 1 1 0 0 - 2 0 0 0 | 1 0 0 0 - 4 0 0 0 (discard bad options: same robots, all resources <= )
func maximumGeodeAfter(minutes int, bluePrint BluePrint) uint8 {
	var maxGeode uint8
	startingFactory := Factory{
		MaxRobots: Robots{
			Ore:  4,
			Clay: 1 + bluePrint.ObsidianRobotCost.v2/2,
		},
		Robots: Robots{Ore: 1}}
	currentFactories := []Factory{startingFactory}
	for minute := 1; minute <= minutes; minute++ {
		// fmt.Println("handle minute: ", minute)
		// fmt.Println("\t", len(currentFactories))
		nextFactories := []Factory{}
		resourcesRobots := map[Robots]Resources{}
		for _, current := range currentFactories {
			options := current.BuildableOptions(bluePrint)
			for _, option := range options {
				newFactory := current
				newFactory.Produce()
				newFactory.Build(option, bluePrint)
				if newFactory.Resources.Geode > maxGeode {
					maxGeode = newFactory.Resources.Geode
				}
				if resBefore, ok := resourcesRobots[newFactory.Robots]; ok {
					if newFactory.Resources.LessOrEqual(resBefore) {
						continue
					}
				}
				nextFactories = append(nextFactories, newFactory)
				resourcesRobots[newFactory.Robots] = newFactory.Resources
			}
		}
		currentFactories = nextFactories
	}
	return maxGeode
}

func Part1(input string) {
	bluePrints := Parse(input)

	total := 0

	for _, bp := range bluePrints {
		qualityLevel := int(bp.ID) * int(maximumGeodeAfter(24, bp))
		total += qualityLevel
	}

	fmt.Println(total)
}

func Part2(input string) {
	bluePrints := Parse(input)

	total := 1
	for _, bp := range bluePrints[:3] {
		maxGeode := int(maximumGeodeAfter(32, bp))
		total *= maxGeode
	}
	fmt.Println(total)
}
