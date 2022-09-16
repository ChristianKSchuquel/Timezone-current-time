package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	for {
		t, err := GetTimezones()
		if err != nil {
			fmt.Println(err)
			main()
		}

		fmt.Println("     ")
		a, err := CheckTimezones(t)
		if err != nil {
			fmt.Println(err)
			main()
		}

		fmt.Println(a)
		fmt.Println("     ")
	}
}

func GetTimezones() ([]byte, error) {
	var t string
	fmt.Println("Please insert a valid timezone: \n(type '--h' to get all timezones)")
	fmt.Println("↴")
	_, err := fmt.Scan(&t)
	if err != nil {
		return []byte(t), err
	}

	if t == "--h" {
		timezones, err := ioutil.ReadFile("timezones.txt")
		if err != nil {
			fmt.Println("Unable to read data: ", err)
		}
		fmt.Println("\nTimezones: ", string(timezones))
		fmt.Println("  ")
		GetTimezones()
	}
	return []byte(t), nil
}

func CheckTimezones(t1 []byte) (string, error) {

	loc1, e1 := time.LoadLocation(string(t1))
	if e1 != nil {
		return "error: ", e1
	}

	loc1Time := time.Now().UTC().In(loc1)
	return loc1Time.Format("15:04"), nil
}
