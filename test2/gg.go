package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println("here")
		return
	}

	str := os.Args[1]

	var result []string
	res := ""


	for i:=0; i<len(str); i++ {
		if str[i] == ' ' {
			if res != ""{
				result =append(result, res)
				res = ""
			}	
		}  else {
			
			res += string(str[i])
		}
	}

	if res != "" {
		result = append(result, res)
	}


	end := ""

	for i, char := range result {
		if i == len(result)-1 {
			end += char
		} else {
			end += char + "   "
		}
		
	}
	fmt.Println(end)
}