package userinterface

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	yearPattern  = "^[0-9]{4}$"
	monthPattern = "^[0-1][0-9]$"
	dayPattern   = "^[0-3][0-9]$"
)

func getInput(prompt string) string {
	var input string
	fmt.Print(prompt + " ")
	fmt.Scanln(&input)

	return input
}

type verify func(string) bool

func verifyYear(year string) bool {
	num, _ := strconv.Atoi(year)
	return num >= 1
}
func verifyMonth(month string) bool {
	num, _ := strconv.Atoi(month)
	return num >= 1 && num <= 12
}
func verifyDay(day string) bool {
	num, _ := strconv.Atoi(day)
	return num >= 1 && num <= 31
}

//gets input for specified time unit, such as year or month
func getTimeUnit(timeUnit, pattern string, fn verify) string {

check1:
	prompt := fmt.Sprintf("Please enter the %s:", timeUnit)
	input := getInput(prompt)

	r, _ := regexp.Compile(pattern)
	for r.MatchString(input) == false || fn(input) == false {
		fmt.Println()
		fmt.Println("Invalid input")
		goto check1
	}
	return input
}

func getDateInput() (string, string, string) {

	year := getTimeUnit("year [xxxx]", yearPattern, verifyYear)
	if year == "reset" {
		return "", "", year
	}
	month := getTimeUnit("month [xx]", monthPattern, verifyMonth)
	if month == "reset" {
		return "", "", month
	}
	day := getTimeUnit("day [xx]", dayPattern, verifyDay)
	if day == "reset" {
		return "", "", day
	}

	return year, month, day
}
func GetDate() string {
	fmt.Println()
	fmt.Println("Please enter date units below")
	format := "2006-01-02 15:04:05"
	var date time.Time
	now := time.Now()
Date:
	year, month, day := getDateInput()

	dateString := fmt.Sprintf("%s-%s-%s 00:00:00", year, month, day)

	date, err := time.Parse(format, dateString)
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		goto Date
	}
	if now.Before(date) == false {
		fmt.Println("INVALID! The date must be in the future.")
		goto Date
	}
	return fmt.Sprintf("%v%v%v", month, day, year)
}

func check(err error) bool {
	errorMessage := errors.New("you did not put in valid input. Try again")
	if err != nil {
		fmt.Println(errorMessage)
		return true
	}
	return false
}
