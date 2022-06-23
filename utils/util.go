package utils

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
//Generates random int between min and max
func GenerateRandInt(min,max int64) int64{
	dif := max-min
	randomNumber := min + rand.Int63n(dif)
   return randomNumber
}

//Takes interger n as input and returns random string of length n
func GenerateRandString(n int) string{
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	length := GenerateRandInt(0,int64(len(alphabet)))
	var sb strings.Builder
	for i:=0; i<=n; i++{
		sb.WriteByte(alphabet[GenerateRandInt(0,int64(length))])
	}

	return sb.String()
}

//Function to generate random username
func GenerateRandomUserName() string{
	const whiteSpace = " "

	var firstName strings.Builder
	var lastName strings.Builder
	var fullName strings.Builder

	//Generate random first and last names
	firstName.WriteString(GenerateRandString(int(GenerateRandInt(0,50))))
	lastName.WriteString(GenerateRandString(int(GenerateRandInt(0,50))))

	//Concatinate the names
	fullName.WriteString(firstName.String())
	fullName.WriteString(whiteSpace)
	fullName.WriteString(lastName.String())

	return fullName.String()
}

//Generate random password
func GenerateRandomPassword (min, max int64) string {
	const characters ="`1234567890qwertyuiopasdfghjklzxcvbnm,./!@#$%^&*()-=_+?><ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sb strings.Builder
	length := GenerateRandInt(min, max);
	for i:=0; i<=int(length); i++{
		sb.WriteByte(characters[GenerateRandInt(min,length)])
	}
	return sb.String()
}


//Generates random email
func GenerateRandomEmail () string{
	const at = "@"
	var sb strings.Builder
	domains := []string{".com",".net",".tech"}
	providers:=[]string{"gmail","outlook","yahoo","icloud"}
	name := GenerateRandString(int(GenerateRandInt(0,64)));
	domain:=domains[rand.Intn(len(domains))]
	provider := providers[rand.Intn(len(domains))]
	sb.WriteString(name)
	sb.WriteString(at)
	sb.WriteString(provider)
	sb.WriteString(domain)
	return sb.String()
}