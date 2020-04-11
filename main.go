// Hangman. Coded in March 2020.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

// Now the game coding begins below
func main() {
	// Moving images and wordlist into main.go to avid using PACKR library
	introImage :=
		`
               .--------.
              /           \
            _|_            |
           /   \           | 
           '==='           |
           . ' .           |
           . : ' .         |
           . ' .           |
           . : ' .         |
           '.              |
           . '    .        |
          .-"""-.          |
          /  \___ \        |
          |/    \|         |
          (  a  a )        |
          |   _\ |         |
          )\  =  |         |
      .--'  '----;         |
     /            '-.      |
    |                \     |
    |     |   .  & .   \   |
     \    /      &   |  ;  |
     |   |           |  ;  |
     |   /\          /  |  |
     \   \ )   -:-  /\  \  |
      \    -.  -:-  |  \ \_|
       '.--. \      (   \/\ \
          / \  \    |    \/_/
         |    \  |  |      |
         |    /'-\  /      |
          \   \   | |      |
           \   )_/\ |      |
            \      \|      |
             \      \      |
              '.     |     |
                /   /      |
               /  .';      |
              /  /  |      |
            /   /   |      |
           |  .' \  |      |
           /  \  )  |      |
           \   \ /  '-.._  |
           '.ooO\__._.Ooo  | 
`
	wordList :=
		`
abate
aberration
abhor
abhorrence
abstruse
accost
acrimony
acumen
adamant
adept
adroit
affected
alacrity
allocate
altruistic
altruism
amenable
amiable
amicable
antediluvian
anthropology
antipathy
apathetic
apathy
apt
arcane
ascendancy
ascetic
asceticism
aspire
assail
assiduous
assuage
atrophy
attenuate
august
aura
auspicious
autocrat
autocratic
automaton
avarice
banal
barrage
belie
belligerent
benevolent
bequeath
berate
bipartisanship
blighted
bog
bolster
bombastic
boorish
boorishness
buoyant
burgeon
buttress
byzantine
cacophonous
cacophony
cajole
callous
cantankerous
cantankerousness
capricious
castigate
caustic
censorious
censure
cerebral
chagrin
charlatan
chastise
chide
churlish
circuitous
circumscribe
circumvent
clandestine
coalesce
compendious
complacency
complacent
compliant
compliance
conciliate
conciliatory
concur
conflagration
confluence
congenial
conscientious
consternation
contempt
contemptuous
contemptible
contentious
convivial
copious
corroborate
cosmopolitan
credulity
credulous
culpable
cursory
dauntless
dearth
debacle
debilitate
debilitated
debunk
decimate
decorum
decorous
deference
deferential
degradation
deleterious
delineate
demonstrative
demure
demystify
denigrate
depose
depravity
deprecate
depreciation
depreciatory
deride
derivative
derogatory
derogate
desecration
despondent
despot
destitute
deterrent
devoid
didactic
diffident
diffidence
diffuse
digress
digression
dilatory
diminutive
diminution
dire
discern
discerning
discomfited
discount
disheartening
disillusionment
disingenuous
disparage
dispassionate
dispel
disputatious
disquieting
disseminate
distaste
divergent
divisive
divulge
doctrine
dogmatic
dormant
dupe
duplicitous
duplicity
ebullient
eclectic
effacement
effervesce
egalitarian
elated
elation
elicit
elucidate
elude
elusive
embittered
embroiled
embroil
empathetic
empathy
empirical
encompass
encroaching
encumbrance
enigmatic
enumerate
ephemeral
epiphany
epitome
epitomize
equanimity
equitable
equivocal
erudite
erudition
esoteric
estrange
estrangement
eulogy
eulogize
evoke
exacting
excavate
exemplar
exhibitionist
exhort
exorbitant
expedient
exposé
extol
extricate
facile
faction
fallacious
fallacy
fanaticism
fastidious
fathom
felicitous
finesse
flagrant
flippant
flippancy
florid
flummox
folly
foolhardy
forlorn
fortitude
fortuitous
fraudulent
frugal
furor
furtive
futile
futility
gait
gallant
gargantuan
garish
genial
germinate
glutton
grandiose
hackneyed
hamper
hardy
hasten
heresy
heretic
histrionic
hubris
idiosyncracy
idiosyncratic
idyllic
ignominy
ignominious
illicit
impasse
imperious
impetuous
impudence
impudent
inane
incongruity
incongruous
incredulous
incredulity
incriminate
incubate
indeterminate
indict
indictment
indigenous
indignant
indiscriminate
indolent
indomitable
induce
indulgent
ineffable
ineptitude
inert
inertia
ingenuous
inherent
inhibit
innate
innocuous
innuendo
inscrutable
insipid
insolence
insolent
instigate
insular
intrepid
inundate
invoke
irate
irony
ironic
irreverent
irreverence
jaded
jocular
jovial
judicious
lackadaisical
laconic
laggard
languid
latent
latency
laud
laudatory
listless
lithe
lucid
lucrative
lull
lurid
luxuriant
magnanimity
magnanimous
malleable
marred
maudlin
melancholy
mercenary
mercurial
miserly
mitigate
mitigator
modicum
morose
motley
multifarious
nebulous
nefarious
neophyte
notoriety
notorious
noxious
nuance
obdurate
obstinate
officious
onerous
opportunist
opportunistic
oracle
orthodox
ostensible
oversight
pacifist
pacify
painstaking
palliate
palliative
paradigm
parch
parody
partisan
patronize
paucity
pedant
pedantic
pedantry
peevish
penchant
penurious
peremptory
perfunctory
peripheral
perquisite
petulant
philanthropist
philanthropic
piety
pious
placate
placid
plasticity
plausible
plausibility
plethora
plethoric
pliable
pliant
polemical
prattle
precarious
precipitate
preclude
precocious
presumptuous
pretext
prevaricator
procure
prodigious
profound
profuse
prohibitive
prohibition
proliferate
proliferation
prolific
pronouncement
propensity
proponent
prosaic
prospective
provident
provincial
punctilious
pundit
quell
quixotic
rampant
ramshackle
rancorous
rancor
rapport
ratify
raucous
ravenous
raze
reap
rebuttal
recalcitrant
recant
recessive
recluse
reclusive
rectify
rectitude
redolent
refutation
refute
regressive
relegate
relinquish
renounce
repertory
reprehensible
reprimand
reproach
repudiate
repugnant
rescind
reticent
reticence
reverent
rhetorical
rouse
rousing
sage
sanctimonious
sanction
sanctity
sanguine
satiate
satire
satirical
satirize
saturate
scanty
scathing
scintillating
scope
scrupulous
scrutinize
scrutiny
self-righteous
self-serving
serendipity
servile
shrewd
shroud
simile
slight
slipshod
solace
solicitous
somber
sophistry
spartan
sporadic
spurious
spurn
squander
stagnant
stagnation
stark
static
staunch
steadfast
stock
strident
stupefy
stupefaction
subservient
substantiate
subversive
succulent
supercilious
superfluous
supplant
surfeit
susceptible
sycophant
tangential
teem
teeming
temperamental
temporize
tenacious
tenacity
tenuous
tirade
toady
torpor
totalitarian
tout
tractable
transient
treatise
trepidation
tribulation
trifling
trite
truculent
truculence
ubiquitous
unabashed
uncanny
uncouth
unfathomable
ungainly
unruly
unwitting
urbane
usurp
vacuous
vacuity
vanquish
vapid
venality
venerable
verbose
vicarious
vigilant
vindicate
vindication
vindictive
virtuoso
virtuosity
virulent
viscous
vocation
vying
waning
wayward
wrath
wry
zealot
`
	introName :=
		`
         ___   ___   ________   ___   __    _______    ___ __ __   ________   ___   __      
        /__/\ /__/\ /_______/\ /__/\ /__/\ /______/\  /__//_//_/\ /_______/\ /__/\ /__/\    
        \::\ \\  \ \\::: _  \ \\::\_\\  \ \\::::__\/__\::\| \| \ \\::: _  \ \\::\_\\  \ \   
         \::\/_\ .\ \\::(_)  \ \\:.  -\  \ \\:\ /____/\\:.      \ \\::(_)  \ \\:.  -\  \ \  
          \:: ___::\ \\:: __  \ \\:. _    \ \\:\\_  _\/ \:.\-/\  \ \\:: __  \ \\:. _    \ \ 
           \: \ \\::\ \\:.\ \  \ \\. \ -\  \ \\:\_\ \ \  \. \  \  \ \\:.\ \  \ \\. \ -\  \ \
            \__\/ \::\/ \__\/\__\/ \__\/ \__\/ \_____\/   \__\/ \__\/ \__\/\__\/ \__\/ \__\/
                                                                                            
                                BY MATTHEW BRANDEBURG                                       
                                    MARCH 2020                                             
																							
																							
`

	outroWin :=
		`
	                                                                          
               CONGRADULATIONS !!!!!                                               
                                                                           
          YOU'VE WON !!!!                                       
															  
                               .                                  
      . .                     -:-               .  .  .
    .'.:,'.        .  .  .     '.              . \ | / .
    .'.;.'.       ._. ! ._.      \             .__\:/__.
     ',:.'         ._\!/_.                       .';'.      . ' .
     .'             . ! .              ,.,      ..======..       .:.
    .                 .               ._!_.     ||::: : | .        ',
    .====.,                  .           ;  .~.===: : : :|   ..===.
    |.:: ||      .=====.,    ..=======.~,   |"|: :|::::::|   ||:::|=====|
 ___| :::|!__.,  |:::::|!_,   |: :: ::|"|l_l|"|:: |:;;:::|___!| ::|: : :|
|: :|::: |:: |!__|; :: |: |===::: :: :|"||_||"| : |: :: :|: : |:: |:::::|
|:::| _::|: :|:::|:===:|::|:::|:===F=:|"!/|\!"|::F|:====:|::_:|: :|::__:|
!_[]![_]_!_[]![]_!_[__]![]![_]![_][I_]!//_:_\\![]I![_][_]!_[_]![]_!_[__]!
 -----------------------------------"---'''''''---"-----------------------
 _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ |= _ _:_ _ =| _ _ _ _ _ _ _ _ _ _ _ _
                                     |=    :    =|                        
_____________________________________L___________J________________________
--------------------------------------------------------------------------
															   
                        GAME OVER                                             
															  
`

	outroLose :=
		`
	                       
	YOU'VE LOST      
                       
	| .__________))______|
	| | / /      ||
	| |/ /       ||
	| | /        ||.-''.
	| |/         |/  _  \
	| |          ||  '/,|
	| |          (\\'_.'
	| |         .-'--'.
	| |        /Y . . Y\
	| |       // |   | \\
	| |      //  | . |  \\
	| |     ')   |   |   ('
	| |          ||'||
	| |          || ||
	| |          || ||
	| |          || ||
	| |         / | | \
	""""""""""|_'-' '-' |"""|
	|"|"""""""\ \       '"|"|
	| |        \ \        | |
	: :         \ \       : : 
	. .          '.       . .
							 
		   GAME OVER         
							 
	`

	// // Use Packr to ingest our ascii files used in game below
	// box := packr.New("My Box", ".")
	// wordList, _ := box.FindString("wordlist.txt")
	// introImage, _ := box.FindString("ascii-image.txt")
	// introName, _ := box.FindString("ascii-name.txt")
	// outroWin, _ := box.FindString("ascii-win.txt")
	// outroLose, _ := box.FindString("ascii-lose.txt")

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
							os.Exit(0) // need more than break here because in round == 0 nested loop
						}
					}
				} else {
					for _, p := range players {
						if len(incorrectlyGuessedLetters) < chancesCount {
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
								os.Exit(0) // need more than break here because in round == 0 nested loop
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
							os.Exit(0) // need more than break here because in round == 0 nested loop
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
				os.Exit(0) // need more than break here because in round == 0 nested loop
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
			os.Exit(0) // need more than break here because in round == 0 nested loop
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
