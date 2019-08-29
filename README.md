# go-loadsh!

Hi!  go-loadsh is similar librieries as **lodash**. lodash is used for javascript and go-loadsh fo golang.

## Array Chunks

Creates an array of elements split into groups the length of size. If array can't be split evenly, the final chunk will be the remaining elements.
 ####  Arguments
	 >  array (Array): The array to process.  
	 >  [size=1] (number): The length of each chunk
 ####  Returns
	 >  (Array): Returns the new array of chunks. 

#### Example :-  
     ```golang
        array := []string{"a", "b", "c", "d", "e"}
	go_lodash.StringChunks(array, 3)
        fmt.Println(go_lodash.StringChunks(array, 3))
       ```
  #####  Output :-
      [[a b c] [d e]]
  2.  
   >  array := []int{1,2,3,4}
     fmt.Println(go_lodash.IntChunks(array, 3))
     Output :-
      [[1 2 3] [4]]
