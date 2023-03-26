package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

func SendEmail(toEmail string, code string, emailType string) error {
	// Set parameters for email 
	from := os.Getenv("MAIL")
	pw := os.Getenv("PW")
	to := []string{toEmail}
	var subject string
	var body string

	// Determine subject/body based on emailType
	if emailType == "signup" {
		subject = "Confirm Your Signup"
		body = "This is an email from our Grocery App for CEN3031.\n\nPlease confirm your signup by entering this code:\n" + code + "\n\nThanks!\nGroup 105"
	} else if emailType == "reset" {
		subject = "Pasword Reset Confirmation"
		body = "This is an email from our Grocery App for CEN3031.\n\nTo verify that it's you trying to reset your password, please enter this code:\n" + code + "\n\nThanks!\nGroup 105"
	} else {
		log.Fatal("invalid email type")
	}

	// Combine subject/body into formatted message
	message := []byte("Subject: " + subject + "\r\n" +
	"To: " + to[0] + "\r\n" +
	"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
	body + "\r\n")

	// Authorize with CEN3031 email login
	auth := smtp.PlainAuth("", from, pw, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)

	// Handle any email errors
	if err != nil {
		fmt.Println(err)
		return errors.New("error sending " + emailType + " email")
	}

	return nil
}

func GenerateRandomCode() string {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())
	// Generate a random 6-digit code
	return strconv.Itoa(rand.Intn(900000) + 100000)
}