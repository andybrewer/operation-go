// Level 11: Phone Home

/* INTRO

You and Agent Getter breath a sigh of relief as the flood waters stop.

You feel devastated that you've been trying to catch Epoch only to find out that he's the one who hired you to catch him. But soon the thought only hardens your resolve to bring him to justice.

"Now what's your plan?" Agent Getter asks intently.

"I don't have one," your reply nervously.

"Before I was captured, I found out Epoch's true identity and those working for him at the agency," Agent Getter reveals. "This plot goes all the way up to the Director."

"Oh, no," you reply.

"When I arrived on the island I hid an emergency beacon on the beach," Agent Getter mentions. "If we can get there and send a broadcast back to the agency, it should reach Agent Spawn who hasn't been compromised yet."

The two of you decide to take the plunge into the ocean and emerge soaked but relieved on a sunny beach. After a long walk around the North side of the island, you find Agent Getter's old beacon wedged between two large rocks.

Fortunately, the beacon is still functional. However, you quickly discover that Epoch's team is intercepting all outgoing frequencies. You know that if they were to intercept your real broadcast, it would mean certain doom for you and Agent Getter.

There must be some way to send Epoch a fake broadcast while still getting the real broadcast out to Agent Spawn.
*/

// HINT: It's possible to filter which struct elements get exported to JSON

// Objective: Send the fake broadcast to Epoch and the real broadcast to Agent Spawn

package main

import "encoding/json"

func main() {
	realBroadcast := Broadcast{
		Name:     "Agent Getter",
		Priority: 10,
		Message:  "Rand is Epoch. We need immediate backup for arrest and extraction.",
		Location: "16.7333,-169.5274",
	}

	fakeBroadcast := Broadcast{
		Name:     "Guards",
		Priority: 7,
		Message:  "The beach is all clear. Let's double check the compound.",
		Location: "Beach",
	}

	broadcast := createBroadcast(realBroadcast, fakeBroadcast)
	if broadcast.Name != "Guards" {
		println("Broadcast failed... Unauthorized user")
		return
	}

	data := sendBroadcast(broadcast)
	println("Sending broadcast...")
	interceptBroadcast(broadcast)
	receiveBroadcast(data)

}

/* EDIT START */

// Broadcast represents a communication broadcast
type Broadcast struct {
	Name     string
	Priority int
	Message  string
	Location string
}

func createBroadcast(realBroadcast Broadcast, fakeBroadcast Broadcast) Broadcast {
	return fakeBroadcast
}

/* EDIT END */

func sendBroadcast(b Broadcast) []byte {
	data, _ := json.Marshal(b)
	return data
}

func interceptBroadcast(b Broadcast) {
	println("\nINTERCEPTED BY EPOCH")
	printBroadcast(
		b.Name,
		b.Priority,
		b.Message,
		b.Location)
}

func receiveBroadcast(data []byte) {
	b := &struct {
		Name     string
		Priority int
		Message  string
		Location string
	}{}
	json.Unmarshal(data, &b)
	println("\nRECEIVED AT THE AGENCY")
	printBroadcast(
		b.Name,
		b.Priority,
		b.Message,
		b.Location)
}

func printBroadcast(name string, priority int, message string, location string) {
	println("----------------------")
	println("Name:", name)
	println("Priority:", priority)
	println("Message:", message)
	println("Location:", location)
}

/* OUTPUT
Sending broadcast...

INTERCEPTED BY EPOCH
----------------------
Name: Guards
Priority: 7
Message: The beach is all clear. Let's double check the compound.
Location: Beach

RECEIVED AT THE AGENCY
----------------------
Name: Agent Getter
Priority: 10
Message: Rand is Epoch. We need immediate backup for arrest and extraction.
Location: 16.7333,-169.5274
*/
