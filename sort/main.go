package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{13, 456, 33, 421, 789}
	sort.Ints(numbers)
	fmt.Println(numbers)

	names := []string{"BBB", "CCC", "DDD", "AAA"}
	sort.Strings(names)
	fmt.Println(names)

	users := []struct {
		name  string
		level int
	}{
		{name: "aaa", level: 2},
		{name: "zzz", level: 40},
		{name: "ccc", level: 10},
		{name: "eee", level: 25},
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].name < users[j].name
	})
	fmt.Println(users)
}
