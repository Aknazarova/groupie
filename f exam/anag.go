package main

import "fmt"

func main() {
	test1 := IsAnagram("listen", "silent")
	fmt.Println(test1)

	test2 := IsAnagram("alem", "school")
	fmt.Println(test2)

	test3 := IsAnagram("ne    at", "a net")
	fmt.Println(test3)

	test4 := IsAnagram("annaa madrigal", "a man and a girl")
	fmt.Println(test4)
}

func IsAnagram(str1, str2 string) bool {
	var c [1000]bool
	var sum1 [1000]int
	var sum2 [1000]int

	for i := 0; i < len(str1); i++ {
		if i != ' ' {
			sum1[str1[i]]++
			c[str1[i]] = true
		}
	}

	for i := 0; i < len(str2); i++ {
		if i != ' ' {
			sum2[str2[i]]++
		}
	}

	for i := 0; i < len(str2); i++ {
		if (c[str2[i]] == true && sum1[str2[i]] == sum2[str2[i]]) || str2[i] == ' ' {
			continue
		} else {
			return false
		}

	}
	return true

}
