/*
Viết function tìm ra số lớn thứ nhì trong mảng các số.
Ví dụ: max2Numbers([2, 1, 3, 4]) => 3
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(s)
	fmt.Println(s)

	s = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // unsorted
	sort.Float64s(s)
	fmt.Println(s)

}

// func max2Numbers(numbers float64) (result float64, err error) {
// 	sort.Float64s(numbers)

// }
