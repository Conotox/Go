package main

import (
	"bufio"
	"fmt"
	"os"
)

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func otherGender(genderx string) {
	fmt.Println("get riggedy rekt " + genderx)
	fmt.Println("See we are all welcome here. We love you" + genderx)
}

func getRekt(gender string) {
	if gender == "boi" {
		fmt.Println("get riggedy rekt " + gender)
	} else if gender == "gurl" {
		fmt.Println("get riggedy rekt " + gender)
	} else {
		fmt.Println("It's 2019 ffs we accept all genders so, get rekt " + gender)
	}
}

func main() {
	getRekt("boi / gurl / whatever")
}
