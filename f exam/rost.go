package main
import "os"
import "fmt"

func main() {

    args := os.Args[1:]
    fmt.Println(FirstWord(args[0]))
}

func FirstWord(str string) string {

    index1 := -1
    index2 := 0
    slice := ""
    str1 := ""
    for i := 0; i < len(str); i++ {
        if i != len(str)-1 && (str[i] == ' ' && str[i+1] != ' ') {
            str1 = str1 + string(str[i])
        } else if str[i] != ' ' {
            str1 = str1 + string(str[i])
        }
    }

    for i := 0; i < len(str1); i++ {
        if index1 == -1 && str1[i] != ' ' {
			index1 = i
			// fmt.Print(index1)
        } else if index1 != -1 && str1[i] == ' ' {
			index2 = i
			// fmt.Print(index2)
            break
        }
	}
	fmt.Println(index1)
	fmt.Println(index2)
    slice = str1[index1:index2]
    s := str1[index2+1:] + " " + slice
    return s
}