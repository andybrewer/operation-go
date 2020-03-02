// Level 7: Familiar Security

/* INTRO
"Not too bad at all," you say and smile as the lasers power down.

Without hesitation, you grab the USBs on the desk and run back downstairs. As you reach the lower level of the compound, you notice some commotion on an ocean-side terrace. The noise startles you, but luckily it's your fellow agents.

"I got 'em!" you cry out as you wave the USBs high in the air.

Several of the agents turn around and start pointing in your direction and shouting, but the words are drowned out by the deafening roar of an approaching helicopter. Your stomach turns to knots as you recognize the helicopter as Epoch's and watch your team guide him in for a landing. No arrest, no questioning, just more pointing in your direction.

You run faster than you ever have before and take cover outside the compound in some dense brush to pause and think. This is bad. Really bad... You need to get off this island, but how? You need more time. Desperate for any advantage you decide to try to intercept their communications.

You can easily capture their communications, but for this to work you need to be able to intercept their communication while keeping their channels up and running.
*/

// HINT: Try using a go routine when reading from the channel

// Objective: Send messages to your interceptComms channel while preserving the messages in epochComms

package main

func main() {
	epochComms := make(chan string)
	go func() {
		epochComms <- messageQueue(0)
		epochComms <- messageQueue(1)
		epochComms <- messageQueue(2)
		epochComms <- messageQueue(3)
		close(epochComms)
	}()
	/* EDIT START */

	interceptComms := make(chan string)
	close(interceptComms)

	/* EDIT END */
	println("Intercepted")
	println("---------------------------------")
	for message := range interceptComms {
		println("->", message)
	}

	println()
	println("Sent to Epoch")
	println("---------------------------------")
	for message := range epochComms {
		println(message)
	}
}

func messageQueue(i int) string {
	messages := make(map[int]string)
	messages[0] = "[Len] All agents head south."
	messages[1] = "[Epoch] Get those USBs!"
	messages[2] = "[Val] Move out team."
	messages[3] = "[Epoch] Faster!"
	return messages[i]
}

/* OUTPUT
Intercepted
---------------------------------
-> [Len] All agents head south.
-> [Epoch] Get those USBs!
-> [Val] Move out team.
-> [Epoch] Faster!

Sent to Epoch
---------------------------------
[Len] All agents head south.
[Epoch] Get those USBs!
[Val] Move out team.
[Epoch] Faster!
*/
