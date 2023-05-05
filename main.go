package main

import (
	"errors"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type DaysRemaining int

func main() {
	lambda.Start(Handler)
}

func Handler() {
	VACATION_DATE := "2023/06/29 23:00:00 -03"

	config := oauth1.NewConfig(
		os.Getenv("API_KEY"),
		os.Getenv("API_KEY_SECRET"),
	)
	token := oauth1.NewToken(
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_TOKEN_SECRET"),
	)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	vacationDateFormatted, err := getVacationDateFormatted(VACATION_DATE)
	if err != nil {
		log.Fatal(err)
	}

	daysToVacation := getDaysRemaining(vacationDateFormatted)
	messageToTweet, err := generateMessageToTweet(daysToVacation)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("tweet", messageToTweet)
	client.Statuses.Update(messageToTweet, nil)
}

func getVacationDateFormatted(date string) (time.Time, error) {
	const shortFormat = "2006/01/02 15:04:05 -07"
	dateParsed, err := time.Parse(shortFormat, date)

	if err != nil {
		return dateParsed, err
	}

	return dateParsed, nil
}

func getDaysRemaining(date time.Time) DaysRemaining {
	today := time.Now()
	return DaysRemaining(math.Ceil(date.Sub(today).Hours() / 24))
}

func generateMessageToTweet(d DaysRemaining) (string, error) {
	var message string

	if alreadyOnVacation(d) {
		return "", errors.New("Already on vacation")
	}

	if d == 1 {
		message = "Hoje é o ultimo dia dessa porra"
	} else {
		message = "Faltam " + strconv.Itoa(int(d)) + " dias para as férias"
	}

	return message, nil
}

func alreadyOnVacation(d DaysRemaining) bool {
	return d < 0
}
