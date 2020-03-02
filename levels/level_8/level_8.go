// Level 8: Need to Escape

/* INTRO
"They're heading south," you say with a decent sigh of relief.

You're heading west and can see a clearing leading down to the water's edge. You quickly make your way to a set of palm trees lining the beach. You decide to catch your breath and wait for the progressively distant sound of voices to fade away completely.

As you survey your surroundings, you notice a small dock to your north. The realization that your boat ride home might never be coming has already sunk in. You're devastated to find your fellow agents helping Epoch, but you decide to put them out of your mind and focus on your escape.

The dock hosts a small boat and a Jet Ski. You try your luck with the boat, but the gas tank is empty. The Jet Ski electronics start up and a digital display prompts you for a computer connection before the motor will engage.

You dig out your laptop and plug into the dashboard for the Jet Ski. It looks like this Jet Ski would make a better rock than an escape plan as it's been hardwired not to start. Good thing you still have a trick or two up your sleeve.

You need this engine to start, and fast.
*/

// HINT: startup() accepts an interface, so login() can return a struct that matches what validSequence() is expecting

// Objective: Start the Jet Ski

package main

import "reflect"

func main() {
	println("Logging in...")
	authorized := startup(login())
	if reflect.ValueOf(authorized).Bool() {
		println("Starting the engine")
		return
	}
	println("Startup failed")
}

func validSequence(i int, el interface{}) bool {
	return reflect.TypeOf(el).String() == "*main.Sequence" &&
		!reflect.ValueOf(el).IsNil() &&
		reflect.ValueOf(el).Elem().NumField() == 2 &&
		reflect.TypeOf(reflect.ValueOf(el).Elem().Field(0).Interface()).String() == "int" &&
		int(reflect.ValueOf(el).Elem().Field(0).Int()) == i*i-i &&
		!reflect.ValueOf(reflect.ValueOf(el).Elem().Field(1).Interface()).IsNil()
}

func startup(seq interface{}) bool {
	for i := 0; i < 5; i++ {
		if !validSequence(i, seq) {
			return false
		}
		seq = reflect.ValueOf(seq).Elem().Field(1).Interface()
	}

	return true
}

/* EDIT START */

func login() bool {
	return false
}

/* EDIT END */

/* OUTPUT
Logging in...
Starting the engine
*/
