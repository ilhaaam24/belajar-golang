package main

import (
	"fmt"
	"regexp"
)



func main(){
	var regexp  *regexp.Regexp = regexp.MustCompile(`e([a-z]+[0-9]).o`)


	fmt.Println(regexp.MatchString("eko"))
	fmt.Println(regexp.MatchString("edo"))
	fmt.Println(regexp.MatchString("eKo"))
	fmt.Println(regexp.MatchString("ey0.o"))


	fmt.Println(regexp.FindAllString("ek1.o eKo edo e1o",6 ))
}