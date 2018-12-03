/**
*	Mailer script for secret santa.
*	Replace YOUR_EMAIL and YOUR_PASSWORD with the host gmail.
*	Load everything properly in email.json first.
*	Written in go for fun.
*	@author xlanor
**/
package main

import(
	"encoding/json"
	"math/rand"
	"time"
	"fmt"
	"os"
	"strconv"
	"gopkg.in/gomail.v2"
)

type email struct{
	Name string
	Email string
}

type jsonloader struct{
	Emails	[]email
}


func main(){
	rand.Seed(time.Now().UnixNano())
	var email_struct *jsonloader
	email_struct = loadJson()
	loaded_length := len(email_struct.Emails)
	fmt.Println(strconv.Itoa(loaded_length) + " santees loaded")
	
	printDeck(email_struct)

	/* Shuffles using fy shuffle */
	fmt.Printf("\n")
	shuffler(email_struct)
	fmt.Printf("Shuffled\n\n")
	fmt.Println("Santa - Santee")
	for i := 0; i < len(email_struct.Emails); i++{
		cur := email_struct.Emails[i]
		var next_mail email;
		if(i == len(email_struct.Emails) - 1){
			next_mail = email_struct.Emails[0]
		}else{
			next_mail = email_struct.Emails[i+1]
		}
		sendMail(cur.Name,cur.Email,next_mail.Name)
	}
}

func loadJson() *jsonloader{
	var loaded *jsonloader
	emails, err := os.Open("email.json")
	if err != nil{
		panic(err)
	}
	jsonParser := json.NewDecoder(emails)
	if err = jsonParser.Decode(&loaded); err != nil{
		panic("Could not pass config file")
	}
	return loaded
}

//Fisher-yates shuffle.
func shuffler(deck *jsonloader){
	for j := len(deck.Emails) - 1; j > 0; j-- {
		k := rand.Intn(j + 1)
		deck.Emails[j], deck.Emails[k] = deck.Emails[k], deck.Emails[j]
	  }
}

func printDeck(email_list* jsonloader){
	for j:= 0; j < len(email_list.Emails); j++{
		fmt.Println(email_list.Emails[j].Name)
	}
}

func sendMail(senderName string , senderEmail string , recipientName string ){

	msg := "Hello <b>"+senderName+"</b>\n" +"Your Secret Santa recipient for 2018 is <b>"+recipientName+"</b>! Have fun picking a gift!\n"
	m:= gomail.NewMessage()
	m.SetHeader("From", "YOUR_EMAIL")
	m.SetHeader("To", senderEmail)
	m.SetHeader("Subject", "Secret Santa 2018!")
	m.SetBody("text/html", msg)

	d := gomail.NewDialer("smtp.gmail.com", 587, "YOUR_EMAIL", "YOUR_PASSWORD")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}