package main

import (
	"fmt"
	"unsafe"
)

/*prints:

824634818264
824634818216
I was 0!    
[]

*/

func main() {
	var count int
	pCount := unsafe.Pointer(&count)

	var array **byte                 // native char**
	pArray := unsafe.Pointer(&array) // native char*** (input)

	/*success, err := parseSyscall(syscall.SyscallN(pFunction, uintptr(pArray), uintptr(pCount)))
	if !success {
		err = buildError(m, err)

		return nil, err
	}*/

	// do something with pCount and pArray
	fmt.Println(uintptr(pArray)) // prints the address
	fmt.Println(uintptr(pCount)) // prints the address

	var result []string

	if count == 0 {
		fmt.Println("I was 0!")
	}

	// convert raw C pointer (array) to slice
	transitionSlice := unsafe.Slice(array, count)
	for _, pString := range transitionSlice {
		// convert each raw C string (char*) to a string
		// and append to result slice
		result = append(result, string(toString(pString)))
		fmt.Println(pString)
	}

	fmt.Println(result)
}

func toString(buffer *byte) string {
	if buffer == nil || *buffer == 0 {
		return ""
	}

	// figure out how large our string is
	strLength, iterator, charSize := 0, unsafe.Pointer(buffer), unsafe.Sizeof(*buffer)
	for *(*byte)(iterator) != 0 {
		iterator = unsafe.Add(iterator, charSize)
		strLength++
	}

	return string(unsafe.Slice(buffer, strLength))
}
