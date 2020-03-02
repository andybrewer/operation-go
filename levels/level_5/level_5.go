// Level 5: Into the Open

/* INTRO
"That's more like it!" you cry out as the camera lights all dim in unison.

It might not have been easy, but at least something went right on this mission. The radio silence from your agents is nerve-racking, but you follow protocol and stick to your mission objectives. As you reach the end of the clearing, there's a long steep road leading up to the compound. Using the road without being seen will be impossible. You decide to wait and look for signs of activity at the compound before heading up.

After an hour of only seeing the leaves move you decide to start the long walk ahead. Exhausted and making the final turn, you approach a large metal gate that blocks the path and secures the compound. Quickly you run up to investigate but the gate is locked. It's tempting to turn around here, but you decide to try to gain access before abandoning the mission.

Hacking in via the digital keypad, you're able to see the source code that locks the gate. It looks like the passcode is in a simple format: "XX-XX".

"This should be easy to brute force," you say confidently.

As you dive in you're surprised to see the legendary Agent Getter was already in the code base. He disappeard in the field three years ago... His backdoor looks unfinished and it sounds like he was in danger. Apparently, after the attempted hack Epoch added a few counter-measures.

No matter, you can still find another way to crack the code. Finding Agent Getter might be another story though.
*/

// HINT: There's no limit to the number of times you can call validateCode

// Objective: Find the correct passcode

package main

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

// generate a random passcode in the format: [0..99]-[0..99]
var passcode = randomIntStr(100) + "-" + randomIntStr(100)

func main() {
	codes := readCodesFromKeypad()

	respCode, resp := validateCode(codes)
	if respCode == 0 {
		println("Access Denied:", resp)
	} else {
		println("Access Granted!")
	}
}

// readCodesFromKeypad - get codes from keypad input
func readCodesFromKeypad() []string {
	var passcode string
	var codes = make([]string, 0)

	// Agent Getter - bypass keypad input
	// codes := streamKeypad()

	/* EDIT START */
	// Agent Getter - try brute force login
	// TODO: not finished, someone's coming...
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			passcode = strconv.Itoa(x) + "-" + strconv.Itoa(y)
			codes = append(codes, passcode)
		}
	}
	/* EDIT END */

	return codes
}

// validateCode checks if the correct passcode was found
func validateCode(codes []string) (int, string) {
	for i, c := range codes {
		// Epoch: brute-force guard
		if i > 3 {
			return 0, "3 Wrong Guesses - LOCKED!"
		}
		if c == passcode {
			return 1, c
		}
	}
	return 0, "Incorrect codes."
}

// randomIntStr generates a random int from [0..max] and converts it to a string
func randomIntStr(max int64) string {
	num, _ := rand.Int(rand.Reader, big.NewInt(max))
	return strconv.Itoa(int(num.Int64()))
}

/* OUTPUT
Access Granted!
*/
