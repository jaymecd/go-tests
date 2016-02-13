// Go's `sort` package implements sorting for builtins
// and user-defined types. We'll look at sorting for
// builtins first.

package main

import "fmt"
import "sort"

func main() {

	// Sort methods are specific to the builtin type;
	// here's an example for strings. Note that sorting is
	// in-place, so it changes the given slice and doesn't
	// return a new one.
	strs := []string{"c", "a", "b"}

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	fmt.Println("Strings:", strs)
	fmt.Println("Sorted: ", sort.StringsAreSorted(strs))

	sort.Strings(strs)

	fmt.Println("Strings:", strs)
	fmt.Println("Sorted: ", sort.StringsAreSorted(strs))

	fmt.Println();

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}

	fmt.Println("Ints:   ", ints)
	fmt.Println("Sorted: ", sort.IntsAreSorted(ints))

	sort.Ints(ints)

	fmt.Println("Ints:   ", ints)
	fmt.Println("Sorted: ", sort.IntsAreSorted(ints))
}
