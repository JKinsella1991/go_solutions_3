//James Kinsella - G00282261@gmit.ie
//Go Problem Sheet 3 - Nov 2017

//Sources: https://golang.org/pkg/regexp/syntax/
// https://golang.org/pkg/regexp/
// https://gobyexample.com/regular-expressions
// https://github.com/StefanSchroeder/Golang-Regex-Tutorial
package main

import (
	"fmt"       //Formatting (Obvs)
	"math/rand" //Randomizer
	"regexp"    //Regular Expressions
	"strings"   //Ease of styring manipulation
	"time"      //Import Unix modifier for random responses
)

func ElizaResponse(input string) string {

	responses := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	//Catch for "I am" expressions
	iam := regexp.MustCompile("(?i)i(?:'| a|)?m(.*)")
	if iam.MatchString(input) {
		return iam.ReplaceAllString(input, "How do I know you are $1")
	}

	//Parses input for "father"
	father, _ := regexp.MatchString(`(?i)\\bfather\\b`, input) //\b and ?i help parse "grandfather" and "Father"

	//Catches "father" input, response:
	if father {
		return ("Why don’t you tell me more about your father?")
	}

	return responses[rand.Intn(len(responses))]
}

func reflection(input string) string {

	//Pronoun exp switch
	pronouns := [][]string{
		{`am`, `are`},
		{`I`, `you`},
		{`you`, `I`},
		{`me`, `you`},
		{`your`, `my`},
		{`my`, `your`},
	}

	// Split input into values
	boundaries := regexp.MustCompile(`\b`)

	values := boundaries.Split(input, -1)

	//Parse pronouns for a match
	for i, token := range values {
		for _, reflection := range pronouns {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {

				values[i] = reflection[1]
				break
			}
		}
	}

	//Piece String back together, after pronoun
	answer := strings.Join(values, ``)

	counterResp := []string{
		"Why do ",
		"How do you know that ",
		"I find it fasinating that ",
	}

	return (counterResp[rand.Intn(len(counterResp))] + answer)
}

func main() {
	//Utilize Unix time for randomization
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("\nInput: " + "People say I look like both my mother and father.")
	fmt.Println("Output: " + ElizaResponse("People say I look like both my mother and father."))

	fmt.Println("\nInput: " + "Father was a teacher.")
	fmt.Println("Output: " + ElizaResponse("Father was a teacher."))

	fmt.Println("\nInput: " + "I was my father’s favourite.")
	fmt.Println("Output: " + ElizaResponse("I was my father’s favourite."))

	fmt.Println("\nInput: " + "I'm looking forward to the weekend.")
	fmt.Println("Output: " + ElizaResponse("I'm looking forward to the weekend."))

	fmt.Println("\nInput: " + "My grandfather was French!")
	fmt.Println("Output: " + ElizaResponse("My grandfather was French!"))

	fmt.Println("\nInput: " + "I am happy.")
	fmt.Println("Output: " + ElizaResponse("I am happy."))

	fmt.Println("\nInput: " + "I'm not happy with your responses")
	fmt.Println("Output: " + ElizaResponse("I'm not happy with your responses"))

	fmt.Println("\nInput: " + "“I AM not sure that you understand the effect that your questions are having on me.”")
	fmt.Println("Output: " + ElizaResponse("I AM not sure that you understand the effect that your questions are having on me."))

	fmt.Println("\nInput: " + "Im supposed to just take what you’re saying at face value?")
	fmt.Println("Output: " + ElizaResponse("Im supposed to just take what you’re saying at face value?"))

	fmt.Println("\nInput: I am not sure that you understand the effect your questions are having on me.")
	fmt.Println(reflection("I am not sure that you understand the effect your questions are having on me."))

	fmt.Println("\nInput: About time I hit the road")
	fmt.Println("Output: " + reflection("About time I hit the road"))

	fmt.Println("\nInput: I think we should go for pints")
	fmt.Println("Output: " + reflection("I think we should go for pints"))

	fmt.Println("\nInput: Do you smell that?")
	fmt.Println("Output: " + reflection("Do you smell that?"))
}
