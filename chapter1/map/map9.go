package main

import (
	"fmt"
	"sort"
)

func main() {
	counts := map[string]int{"Barry": 96, "Aaron": 98, "Clan": 97}

	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return counts[keys[i]] > counts[keys[j]] })

	for _, key := range keys {
		fmt.Printf("%s, %d\n", key, counts[key])
	}
}

//Aaron, 98
//Clan, 97
//Barry, 96
