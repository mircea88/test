// pack
package pack

func SumArr(a [4]int) int {
	su := 0
	for _, v := range a {
		su += v
	}
	return su
}
