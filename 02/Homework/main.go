/*
Bài 1 Viết function tìm ra số lớn thứ nhì trong mảng các số.
Ví dụ: max2Numbers([2, 1, 3, 4]) => 3

Bài 2 Cho 1 mảng các chuỗi. Viết function lọc ra các phần tử có độ dài lớn nhất.
Ví dụ: findMaxLengthElement["aba", "aa", "ad", "c", "vcd"] => ["aba", "vcd"]

Bài 3 Viết function remove những phần tử bị trùng nhau trong mảng
Ví dụ: removeDuplicates([1,2,5,2,6,2,5]) => [1,2,5,6]

Bài 4 Một nhân viên trong công ty bao gồm các thuộc tính sau : Tên, Hệ số lương, Tiền trợ cấp
Tạo 1 mảng nhân viên (số lượng tuỳ ý) và thực hiện các chức năng sau:

Sắp xếp tên nhân viên tăng dần theo bảng chữ cái
Sắp xếp nhân viên theo mức lương giảm dần (lương = Hệ số lương * 1.500.000 + Tiền trợ cấp)
Lấy ra danh sách nhân viên có mức lương lớn thứ 2 trong mảng nhân viên
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Bài 1
	numbers1 := []float64{2.5, 1.7, 3.2, 4.9}
	fmt.Println("The second largest number is (function 1): ", max2Numbers1(numbers1))
	numbers2 := []float64{6.5, 11.7, 8.2, 1.9}
	fmt.Println("The second largest number is (function 2): ", max2Numbers2(numbers2))

	Bài 2
	elements := []string{"aba", "aa", "ad", "c", "vcd"}
	fmt.Println("The max length element: ", findMaxLengthElement(elements))

	fmt.Println(findLargestNumber(numbers2))
}

/*
Sắp xếp các số từ nhỏ đến lớn sử dụng hàm sẵn có và sau đó trả về số gần cuối
*/
func max2Numbers1(numbers []float64) float64 {
	sort.Float64s(numbers)
	return numbers[len(numbers)-2]
}

/*
Sắp xếp các số từ nhỏ đến lớn bằng vòng lặp sau đó trả về số gần cuối
*/
func max2Numbers2(numbers []float64) float64 {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				a := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = a
			}
		}
	}
	return numbers[len(numbers)-2]
}

/*
Hàm tìm số lớn nhất
*/
func findLargestNumber(numbers []float64) float64 {
	for i := 1; i < len(numbers); i++ {
		if numbers[0] < numbers[i] {
			numbers[0] = numbers[i]
		}
	}
	return numbers[0]
}

/*
Hàm lọc ra các phần tử có độ dài lớn nhất
*/
func findMaxLengthElement(strings []string) []string {
	array := []int{}
	for i := 0; i < len(strings); i++ {
		array = append(array, len(strings[i]))
	}
}
