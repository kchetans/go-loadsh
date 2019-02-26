
//Creates an array of elements split into groups the length of size. If array can't be split evenly, the final chunk will be the remaining elements.
/* Arguments
array (Array): The array to process.
[size=1] (number): The length of each chunk
*/
/*
Returns
(Array): Returns the new array of chunks.
Example :-
   array := []string{"a", "b", "c", "d", "e"}
	fmt.Println(stringChunks(array, 3))
Output :-
    [[a b c] [d e]]

*/

package go_lodash

func StringChunks(array []string, size int) (newArray [][]string) {
	if len(array) <= 0 || size == 0 {
		return
	}

	var iarray []string
	currentIndex := 0
	totalItrate := 0
	mod := len(array)%size
	if mod == 0{
		totalItrate = len(array)/size
	}else {
		totalItrate = len(array)/size + 1
	}
	for a := 0;a < totalItrate; a++{
		total := currentIndex + size
		if total > len(array) {
			total = len(array)
		}
		iarray = []string{}
		for i := currentIndex; i < total; i++ {
			iarray = append(iarray, array[i])
			currentIndex = size
		}
		if len(iarray) > 0 {
			newArray = append(newArray, iarray)

		}
		currentIndex = total
	}

	return
}

func IntChunks(array []int, size int) (newArray [][]int) {
	if len(array) <= 0 || size == 0 {
		return
	}

	var iarray []int
	currentIndex := 0
	totalItrate := 0
	mod := len(array)%size
	if mod == 0{
		totalItrate = len(array)/size
	}else {
		totalItrate = len(array)/size + 1
	}
	for a := 0;a < totalItrate; a++{
		total := currentIndex + size
		if total > len(array) {
			total = len(array)
		}
		iarray = []int{}
		for i := currentIndex; i < total; i++ {
			iarray = append(iarray, array[i])
			currentIndex = size
		}
		if len(iarray) > 0 {
			newArray = append(newArray, iarray)

		}
		currentIndex = total
	}

	return
}
