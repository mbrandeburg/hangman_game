// Hangman. Coded in March 2020.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gobuffalo/packr/v2"
)

// Player : number and what letters are guessed
type Player struct {
	Number  int
	Guesses []string
}

// Find if a letter guessed is in a string
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// NumFind if a number is in a slice
func NumFind(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// MultiFind used to find if a letter guessed multiple times is in a string
func MultiFind(slice []string, val string) (sliceOut []int) {
	var multiFindSlice []int
	for i, item := range slice {
		if item == val {
			// fmt.Printf("%v at %v", item, i)
			multiFindSlice = append(multiFindSlice, i)
		}
	}
	return multiFindSlice
}

// Unique will make correct guesses unique values only - in case ppl repeat - NOTE: we won't apply to incorrect so we can penalize them (we do offer a chance to not mess up when they do that though)
func Unique(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

// Generate our ascii images
func textGenerator(textFileName string, timerDuration int64) {
	for _, line := range strings.Split(strings.TrimSuffix(textFileName, "\n"), "\n") {
		fmt.Println(line)
		time.Sleep(110 * time.Millisecond)
	}
}

// Correct guesses _ _ _ etc. builder
func correctGuessMapper(slice []string, val string, m map[int]string, finalString string) {
	// First do a loop to figure out where letters are in string?
	var guessWordSlice = strings.Split(val, "")
	for _, letter := range slice {
		foundSlice := MultiFind(guessWordSlice, letter)
		for _, index := range foundSlice {
			// fmt.Printf("\n%v value was found in %s at position %d", letter, val, index)
			m[index] = letter // Because otherwise it overwrites the letter when the letter is key
		}
	}
	var tempMap = m
	// fmt.Printf("\n%v", tempMap)

	keys := make([]int, 0, len(tempMap))

	for key := range tempMap {
		keys = append(keys, key)
	}
	// fmt.Printf("\n%v", keys)
	for number := range val {
		_, found := NumFind(keys, number)
		if found {
			continue
		} else {
			tempMap[number] = "_"
		}
	}
	// fmt.Printf("\n%v", tempMap)
	fmt.Println("\nThe word so far is: ")
	for keyHere := range val {
		// fmt.Printf("\n%v,%v  ", keyHere, tempMap[keyHere])
		fmt.Printf("%v  ", tempMap[keyHere])
	}
}

func main() {
	// Use Packr to ingest our ascii files used in game below
	box := packr.New("My Box", ".")
	wordList, _ := box.FindString("wordlist.txt")
	introImage, _ := box.FindString("ascii-image.txt")
	introName, _ := box.FindString("ascii-name.txt")
	outroWin, _ := box.FindString("ascii-win.txt")
	outroLose, _ := box.FindString("ascii-lose.txt")

	// For correct guess blank builder
	var guessWordBlankBuilider = make(map[int]string)
	var guessWordFinal string

	// Select word at random from file
	rand.Seed(time.Now().Unix())
	var lenghtOfWorldList int
	var guessWord string
	for index, _ := range strings.Split(strings.TrimSuffix(wordList, "\n"), "\n") {
		lenghtOfWorldList = index // will overwrite each time, which is good, just need lenght b/c len() won't work the way the newlines are formatted - len is being character based otherwise
	}
	n := rand.Int() % lenghtOfWorldList
	for index, line := range strings.Split(strings.TrimSuffix(wordList, "\n"), "\n") {
		if index == n {
			// Assign it here
			guessWord = line
		}
	}
	// guessWord = "test"

	// Introduce the game
	fmt.Println("\n")
	fmt.Println("\nWELCOME TO\n")
	time.Sleep(1 * time.Second)
	textGenerator(introImage, 110)
	fmt.Println("***********************************************************")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("\n")
	textGenerator(introName, 50)
	time.Sleep(1 * time.Second)
	fmt.Println("\n")

	//Set up the variables
	var playerCount int
	for {
		fmt.Println("First, how many players are there? (1-20)")
		fmt.Scanf("%d\n", &playerCount)
		if playerCount >= 1 && playerCount <= 20 {
			break
		}
		fmt.Println("\nPlease enter a number in the valid range.")
	}

	// Add the above variable responses to the Player struct at the top
	var players []*Player
	for i := 1; i <= playerCount; i++ {
		players = append(players, &Player{
			Number: i,
		})
	}

	// Add in number of chances to play - doesn't need to get tagged to a specific player
	var chancesCount int
	for {
		fmt.Println("Now, how many incorrect guesses can you make?")
		fmt.Scanf("%d\n", &chancesCount)
		if chancesCount > 0 {
			break
		}
		fmt.Println("\nPlease enter a number in the valid range.")
	}

	// Now stage the new variables needed for the game below
	var input string
	var correctlyGuessedLetters []string
	var incorrectlyGuessedLetters []string

	// Now begin the game
	fmt.Println("\n***********************************************************")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("***********************************************************\n")
	time.Sleep(110 * time.Millisecond)
	fmt.Println("\nNow begins the game..")
	time.Sleep(100 * time.Millisecond)
	if chancesCount == 1 {
		fmt.Printf("\nYour hangman has %d chance to live!", chancesCount)
	} else {
		fmt.Printf("\nYour hangman has %d chances to live!", chancesCount)
	}
	time.Sleep(600 * time.Millisecond)
	fmt.Printf("\nYour word is %d letters long...\n", len(guessWord))
	var blankSlate = strings.Repeat("_ ", len(guessWord))
	fmt.Printf("%s", blankSlate)
	time.Sleep(600 * time.Millisecond)

	// Rounds coded here
	for round := 0; round < 1000; round++ { // 1000 is arbitrary - we loop out when we die per the p.Chances value
		if len(incorrectlyGuessedLetters) < chancesCount {
			var guessWordSliceForLen = strings.Split(guessWord, "") // If a word has repeating letters, can't do len(word), need len(unique values in slice made of word)
			if len(correctlyGuessedLetters) < len(Unique(guessWordSliceForLen)) {
				if round == 0 {
					for _, p := range players {
						if len(correctlyGuessedLetters) < len(Unique(guessWordSliceForLen)) {
							time.Sleep(1 * time.Second)
							if p.Number == 1 {
								fmt.Printf("\n\nPlayer %d: What is your first letter guess?\n", p.Number)
							} else {
								fmt.Printf("\n\nPlayer %d: What is your letter guess? (Or guess the word!) \n", p.Number)
							}
							fmt.Scanf("%s\n", &input)
							if len(input) == 1 {
								//Check if in word
								_, found := Find(correctlyGuessedLetters, input)
								if found {
									fmt.Println("That value was previously used as a correctly guessed letter.\n")
									time.Sleep(800 * time.Millisecond)
									fmt.Printf("Remeber, you have correctly guessed previously: %s\n", strings.Trim(fmt.Sprint(correctlyGuessedLetters), "[]"))
									time.Sleep(800 * time.Millisecond)
									if len(incorrectlyGuessedLetters) > 0 {
										fmt.Printf("And, you have INCORRECTLY guessed previously: %s\n", strings.Trim(fmt.Sprint(incorrectlyGuessedLetters), "[]"))
										time.Sleep(800 * time.Millisecond)
									}
									fmt.Println("Last chance to make a correct guess!\n")
									time.Sleep(800 * time.Millisecond)
									fmt.Println("What is your letter guess?\n")
									fmt.Scanf("%s\n", &input)
								}
								p.Guesses = append(p.Guesses, input)
								fmt.Printf("You guessed: %s\n", p.Guesses[len(p.Guesses)-1])
								if strings.ContainsAny(guessWord, input) {
									fmt.Printf("You have chosen well. %s is in the magic word.\n", input)
									correctlyGuessedLetters = append(correctlyGuessedLetters, input)
									correctlyGuessedLetters = Unique(correctlyGuessedLetters)
								} else {
									fmt.Printf("You have chosen poorly. %s is not in the magic word.\n", input)
									incorrectlyGuessedLetters = append(incorrectlyGuessedLetters, input)
								}
								time.Sleep(1 * time.Second)
								if len(incorrectlyGuessedLetters) >= 1 {
									fmt.Printf("\nFor the record, everyone's incorrect guesses are: ")
									for _, item := range incorrectlyGuessedLetters {
										fmt.Printf("%s ", item)
									}
								}
								time.Sleep(1 * time.Second)
								if len(correctlyGuessedLetters) >= 1 {
									correctGuessMapper(correctlyGuessedLetters, guessWord, guessWordBlankBuilider, guessWordFinal)
								}
							} else if len(input) > 1 {
								if input == guessWord {
									time.Sleep(3 * time.Second)
									fmt.Println("\n\n\nYOUR MAN HAS BEEN SAVED!")
									time.Sleep(2 * time.Second)
									fmt.Printf("\nThe word was: %s\n", guessWord)
									time.Sleep(2 * time.Second)
									fmt.Println("\n")
									textGenerator(outroWin, 110)
									time.Sleep(2 * time.Second)
									os.Exit(0) // need more than break here because in round == 0 nested loop
								} else {
									fmt.Printf("You have chosen poorly. %s is not the magic word and will be recorded as a letter guess.\n", input)
									incorrectlyGuessedLetters = append(incorrectlyGuessedLetters, input)
									time.Sleep(1 * time.Second)
									if len(incorrectlyGuessedLetters) >= 1 {
										fmt.Printf("\nFor the record, everyone's incorrect guesses are: ")
										for _, item := range incorrectlyGuessedLetters {
											fmt.Printf("%s ", item)
										}
									}
									time.Sleep(1 * time.Second)
									if len(correctlyGuessedLetters) >= 1 {
										correctGuessMapper(correctlyGuessedLetters, guessWord, guessWordBlankBuilider, guessWordFinal)
									}
								}
							}
						} else {
							time.Sleep(3 * time.Second)
							fmt.Println("\n\n\nYOUR MAN HAS BEEN SAVED!")
							time.Sleep(2 * time.Second)
							fmt.Printf("\nThe word was: %s\n", guessWord)
							time.Sleep(2 * time.Second)
							fmt.Println("\n")
							textGenerator(outroWin, 110)
							time.Sleep(2 * time.Second)
							break // now break the 1000 loop
						}
					}
				} else {
					for _, p := range players {
						if len(correctlyGuessedLetters) < len(Unique(guessWordSliceForLen)) {
							time.Sleep(1 * time.Second)
							fmt.Printf("\n\nPlayer %d: What is your next letter guess? (Or guess the word!) \n", p.Number)
							fmt.Scanf("%s\n", &input)
							if len(input) == 1 {
								//Check if in word
								_, found := Find(correctlyGuessedLetters, input)
								if found {
									fmt.Println("That value was previously used as a correctly guessed letter.\n")
									time.Sleep(800 * time.Millisecond)
									fmt.Printf("Remeber, you have correctly guessed previously: %s\n", strings.Trim(fmt.Sprint(correctlyGuessedLetters), "[]"))
									time.Sleep(800 * time.Millisecond)
									if len(incorrectlyGuessedLetters) > 0 {
										fmt.Printf("And, you have INCORRECTLY guessed previously: %s\n", strings.Trim(fmt.Sprint(incorrectlyGuessedLetters), "[]"))
										time.Sleep(800 * time.Millisecond)
									}
									fmt.Println("Last chance to make a correct guess!\n")
									time.Sleep(800 * time.Millisecond)
									fmt.Println("What is your letter guess?\n")
									fmt.Scanf("%s\n", &input)
								}
								p.Guesses = append(p.Guesses, input)
								fmt.Printf("You guessed: %s\n", p.Guesses[len(p.Guesses)-1])
								if strings.ContainsAny(guessWord, input) {
									fmt.Printf("You have chosen well. %s is in the magic word.\n", input)
									correctlyGuessedLetters = append(correctlyGuessedLetters, input)
									correctlyGuessedLetters = Unique(correctlyGuessedLetters)
								} else {
									fmt.Printf("You have chosen poorly. %s is not in the magic word.\n", input)
									incorrectlyGuessedLetters = append(incorrectlyGuessedLetters, input)
								}
								time.Sleep(1 * time.Second)
								if len(incorrectlyGuessedLetters) >= 1 {
									fmt.Printf("\nFor the record, everyone's incorrect guesses are: ")
									for _, item := range incorrectlyGuessedLetters {
										fmt.Printf("%s ", item)
									}
								}
								time.Sleep(1 * time.Second)
								if len(correctlyGuessedLetters) >= 1 {
									correctGuessMapper(correctlyGuessedLetters, guessWord, guessWordBlankBuilider, guessWordFinal)
								}
							} else if len(input) > 1 {
								if input == guessWord {
									time.Sleep(3 * time.Second)
									fmt.Println("\n\n\nYOUR MAN HAS BEEN SAVED!")
									time.Sleep(2 * time.Second)
									fmt.Printf("\nThe word was: %s\n", guessWord)
									time.Sleep(2 * time.Second)
									fmt.Println("\n")
									textGenerator(outroWin, 110)
									time.Sleep(2 * time.Second)
									break // now break the 1000 loop
								} else {
									fmt.Printf("You have chosen poorly. %s is not the magic word and will be recorded as a letter guess.\n", input)
									incorrectlyGuessedLetters = append(incorrectlyGuessedLetters, input)
									time.Sleep(1 * time.Second)
									if len(incorrectlyGuessedLetters) >= 1 {
										fmt.Printf("\nFor the record, everyone's incorrect guesses are: ")
										for _, item := range incorrectlyGuessedLetters {
											fmt.Printf("%s ", item)
										}
									}
									time.Sleep(1 * time.Second)
									if len(correctlyGuessedLetters) >= 1 {
										correctGuessMapper(correctlyGuessedLetters, guessWord, guessWordBlankBuilider, guessWordFinal)
									}
								}
							}
						} else {
							time.Sleep(3 * time.Second)
							fmt.Println("\n\n\nYOUR MAN HAS BEEN SAVED!")
							time.Sleep(2 * time.Second)
							fmt.Printf("\nThe word was: %s\n", guessWord)
							time.Sleep(2 * time.Second)
							fmt.Println("\n")
							textGenerator(outroWin, 110)
							time.Sleep(2 * time.Second)
							break // now break the 1000 loop
						}
					}
				}
			} else {
				time.Sleep(3 * time.Second)
				fmt.Println("\n\n\nYOUR MAN HAS BEEN SAVED!")
				time.Sleep(2 * time.Second)
				fmt.Printf("\nThe word was: %s\n", guessWord)
				time.Sleep(2 * time.Second)
				fmt.Println("\n")
				textGenerator(outroWin, 110)
				time.Sleep(2 * time.Second)
				break // now break the 1000 loop
			}
		} else {
			time.Sleep(3 * time.Second)
			fmt.Println("\n\n\nOH NO. YOUR MAN HAS BEEN HUNG!")
			time.Sleep(2 * time.Second)
			fmt.Printf("\nThe word you were looking for is: %s\n", guessWord)
			time.Sleep(2 * time.Second)
			fmt.Println("\n")
			textGenerator(outroLose, 110)
			time.Sleep(2 * time.Second)
			break // now break the 1000 loop
		}
	}
	// End the game
	time.Sleep(1 * time.Second) // pause afer the player's round
	fmt.Printf("\nThanks for playing! (Hit enter to quit)\n")
	fmt.Scanf("%s\n", &input) // again, doesn't do anything, but gives us time before moving on
}

// EXPORT COMMANDS
// PACKR: $packr2
// MACOS: env GOOS=darwin GOARCH=amd64 go build -v github.com/gophercises/hangman_game
// WINDOWS: env GOOS=windows GOARCH=386 go build -v github.com/gophercises/hangman_game
// PACKR: $packr2 clean

// Fully bounded commands:
// packr2 && env GOOS=darwin GOARCH=amd64 go build -v github.com/gophercises/hangman_game && env GOOS=windows GOARCH=386 go build -v github.com/gophercises/hangman_game && packr2 clean

// ************************************************************************************************************************************************
// ARCHIVING ASCII ART FUNCTION ATTEMPTS BELOW FOR POSTERITY
// func textGenerator(textFileName string, timerDuration int64) {
// 	// // Quick way to produce ascii art
// 	// asciiart, err := ioutil.ReadFile("ascii-image.txt")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Println(string(asciiart))

// 	// Sadly, needed a more verbose option to go line by line....
// 	asciiart, err := os.Open(textFileName)
// 	if err != nil {
// 		log.Fatalf("failed opening file: %s", err)
// 	}
// 	scanner := bufio.NewScanner(asciiart)
// 	scanner.Split(bufio.ScanLines)
// 	var txtlines []string
// 	for scanner.Scan() {
// 		txtlines = append(txtlines, scanner.Text())
// 	}
// 	asciiart.Close()
// 	for _, eachline := range txtlines {
// 		fmt.Println(eachline)
// 		time.Sleep(time.Duration(timerDuration) * time.Millisecond)
// 	}
// }
