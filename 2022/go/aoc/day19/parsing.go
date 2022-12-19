package day19

import (
	"regexp"
	"strconv"
	"strings"
)

var testinput = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

type Tuple struct {
	v1, v2 uint8
}
type BluePrint struct {
	ID                uint8
	OreRobotCost      uint8
	ClayRobotCost     uint8
	ObsidianRobotCost Tuple
	GeodeRobotCost    Tuple
}

func Parse(input string) []BluePrint {
	re := regexp.MustCompile(`Blueprint (\d*): Each ore robot costs (\d*) ore. Each clay robot costs (\d*) ore. Each obsidian robot costs (\d*) ore and (\d*) clay. Each geode robot costs (\d*) ore and (\d*) obsidian.`)
	bluePrints := []BluePrint{}
	for _, line := range strings.Split(input, "\n") {
		results := re.FindStringSubmatch(line)
		vs := []uint8{}
		for _, svalue := range results[1:] {
			v, err := strconv.ParseUint(svalue, 10, 16)
			if err != nil {
				panic(err.Error())
			}
			vs = append(vs, uint8(v))
		}
		bluePrints = append(bluePrints,
			BluePrint{
				ID:                vs[0],
				OreRobotCost:      vs[1],
				ClayRobotCost:     vs[2],
				ObsidianRobotCost: Tuple{v1: vs[3], v2: vs[4]},
				GeodeRobotCost:    Tuple{v1: vs[5], v2: vs[6]},
			})
	}
	return bluePrints
}
