package main

import "fmt"

func IsAnagram(str1, str2 string) bool {
	var bol [1000]bool
	var was [1000]int
	var was2 [1000]int

	for i := 0; i < len(str1); i++ {
		if str1[i] != ' ' {
			was[str1[i]]++
			bol[str1[i]] = true
		}
		// fmt.Println(was[str1[i]])

	}
	for i := 0; i < len(str2); i++ {
		if str2[i] != ' ' {
			was2[str2[i]]++
		}
		// fmt.Println(was[str2[i]])
	}
	for i := 0; i < len(str2); i++ {
		if (bol[str2[i]] == true && was[str2[i]] == was2[str2[i]]) || str2[i] == ' ' {
			continue

		} else {
			return false
		}
	}
	return true

}

func main() {
	test1 := IsAnagram("listen", "silent")
	fmt.Println(test1)

	// test2 := IsAnagram("alem", "school")
	// fmt.Println(test2)

	// test3 := IsAnagram("neat", "a net")
	// fmt.Println(test3)

	// test4 := IsAnagram("anna madrigal", "a man and a girl")
	// fmt.Println(test4)
}
