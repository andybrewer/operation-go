// Level 10: Epoch's Trap

/* INTRO
As the door unlocks, Agent Getter rushes out to meet you.

"Agent Getter! I'm so relieved to see you. I'm Agent Buggs. I was on a mission to secure these USB drives but some agents have flipped and they're helping Epoch," you say desperately searching for answers.

"I'm afraid you're in serious danger. Epoch is not who you think he is," Agent Getter replies.

Just then, a shadowy figure enters the far end of the lair. It looks like your boss Rand!

"Agent Buggs, Agent Getter, it looks like you're in over your heads this time," shouts Rand.

"We sure are!" you reply. "And we sure are glad to see you! How did you find us down here?"

"This is my island," Rand replies in a low booming voice. "I'm Epoch. You were getting too close to finding my true identity, so I sent you on this dead-end mission. Despite your heroics so far, this is the end for you."

You're stunned silence fills the room as your impending fate becomes clear. As quickly as he appeared, Epoch disappears behind a corner and the door slams shut. Suddenly, valves in the floor open up and seawater starts flooding the lair...

This could be it if you can't stop the flow of water quickly.
*/

// HINT: an init() function can be used to run code before main() is called and overwrite global variables

// Objective: Pump 0 gallons without setting off the alarm

package main

var pump = func(rate int, capacity int) int {
	return fill(rate, 0, capacity)
}

func fill(rate int, filled int, capacity int) int {
	if filled < capacity {
		return fill(rate, filled+rate, capacity)
	}
	return filled
}

func main() {
	capacity := 20000
	rate := 800

	testRun := pump(rate, capacity)
	if testRun < capacity {
		println("ALARM: TestRun failed")
		return
	}

	filled := pump(rate, capacity)
	println("Gallons pumped:", filled)
}

/* EDIT START */

// Your code

/* EDIT END */

/* OUTPUT
Gallons pumped: 0
*/
