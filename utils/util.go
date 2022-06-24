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
func GenerateRandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

//Takes interger n as input and returns random string of length n
func GenerateRandString(n int) string{
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	length := GenerateRandInt(0, len(alphabet))
	var sb strings.Builder
	for i:=0; i<=n; i++{
		sb.WriteByte(alphabet[GenerateRandInt(0, length)])
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
	firstName.WriteString(GenerateRandString(GenerateRandInt(0,26)))
	lastName.WriteString(GenerateRandString(GenerateRandInt(0,26)))

	//Concatinate the names
	fullName.WriteString(firstName.String())
	fullName.WriteString(whiteSpace)
	fullName.WriteString(lastName.String())

	return fullName.String()
}

//Generate random password
func GenerateRandomPassword (min, max int) string {
	const characters ="`1234567890qwertyuiopasdfghjklzxcvbnm,./!@#$%^&*()-=_+?><ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sb strings.Builder
	length := GenerateRandInt(min, max);
	for i:=0; i<=int(length); i++{
		sb.WriteByte(characters[GenerateRandInt(min, length)])
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

func RandomSentenceGenerator() string{
	longString := "ou do not need to be a content creator or an influencer to land a job in tech.Back-end web development is a load of CRUD.80% of the time, the other 20% is: micro services, infrastructure, data transportation, predictive models, design patterns, architecture, data modeling, graph algorithms, …"
	words := strings.Split(longString," ")

	len := len(words)
	sentence := []string{}
	cond :=int(GenerateRandInt(4,len))
	for i := 0; i< cond; i++ {
		r := GenerateRandInt(0, len-1)
		sentence=append(sentence,words[r])
	}
	sent := strings.Join(sentence, "");
	return sent
}