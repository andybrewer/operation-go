// Level 9: Hidden Lair

/* INTRO
"Nothing like getting some closure on this mission," you chuckle to yourself.

You turn on the gas and smile as the engine roars to life. Instantly, a message flashes on the screen: "Starting Underwater Mode".

"Oh no..." you sigh and tap the dashboard franticly. Suddenly, a glass enclosure surrounds the Jet Ski as it begins to move on it's own. The Jet Ski dives underwater as you unnecessarily hold your breath out of instinct. Through the clear ocean water you can see a rail guiding you towards a submerged bunker with a large glass window.

As you slowly enter the bunker, the Jet Ski rises up into a cavernous air chamber. Looking around, you surprisingly find yourself in Epoch's hidden lair. The underwater view from the bunker is amazing, but your awe quickly turns to fear as you see someone's reflection in the glass.

Spinning around, you realize it's someone trapped behind a thick plexiglass door and you instantly recognize them as the famous Agent Getter!

"Hold on. I'll get you out," you state confidently.

Unbelievably, the sequence of buttons to unlock the door appears to be sitting right there in the code. You try your luck, but the sequence is getting scrambled and the door still remains locked. This has to be the first time you've ever seen a race condition double as a security feature.

If you can execute the code safely, it looks like the door should unlock for you.
*/

// HINT: The sync package includes the ability to use mutexes

// Objective: Fix the race condition

package main

import (
	"math/rand"
	"sync"
	"time"
)

var buttons = []string{"red", "blue", "green", "yellow", "purple"}

func main() {
	rand.Seed(111009)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(buttons))

	sequence := make([]Button, 0)

	for x := 0; x < len(buttons); x++ {
		go addButton(x, &sequence, &waitGroup)
	}

	waitGroup.Wait()

	println("Sending Code Sequence...")
	for i, button := range sequence {
		println(i+1, ":", button.color)
	}
}

/* EDIT START */

func setButton(x int, sequence *[]Button) {

	/* EDIT END */
	newButton := Button{buttons[x]}
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	*sequence = append(*sequence, newButton)
}

func addButton(x int, sequence *[]Button, waitGroup *sync.WaitGroup) {
	setButton(x, sequence)
	waitGroup.Done()
}

// Button represents a colored button
type Button struct {
	color string
}

/* OUTPUT
Sending Code Sequence...
1 : purple
2 : red
3 : blue
4 : green
5 : yellow
*/
