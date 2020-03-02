// Level 3: Where Am I?

/* INTRO
As soon as you finish editing your code you feel a presence looking over your shoulder.

"Epoch is down there on that island, but we have intelligence that he'll be leaving tomorrow. Stay low tonight and the island will be empty in the morning," says a tough looking agent.

"Find his USBs and bring them back to the agency. A boat will signal for you tomorrow at 2300," he says, barely concealing a smile.

As you get up to ask him a question the cargo bay door drops down and a blast of frigid air rushes into the plane. Before you can brace yourself from the wind a firm shove to your back sends you hurdling into the sky. You franticly steady yourself and dive towards the island.  Luckily, your parachute deploys safely, just as planned.

When you touch down, you notice that the other agents have all landed far north of your location. You calmly unpack a satellite GPS to find your location.

Suspiciously, GPS tracking is active but there's a cryptic function that's restricting your access. There's a lot going on but if you can follow the code you should be able to bypass the restrictions.
*/

// HINT: Read the GoDoc for "strings": https://godoc.org/strings

// Objective: Make a successful GPS request

package main

import "strings"

func main() {
	println("Location:", gpsRequest(newAgent()))
}

func newAgent() string {
	/* EDIT START */

	return "D. Buggs"

	/* EDIT END */
}

func gpsRequest(agent string) string {
	a := strings.Split(agent, " ")

	denied := len(a) != 3 ||
		!strings.EqualFold(a[0], a[1]) ||
		!(strings.Count(a[2], a[1]) > 0)

	if denied {
		return "ACCESS DENIED"
	}

	return "\u0031\u0036\u002e\u0037\u0033\u0033\u0033\u002c\u002d\u0031\u0036\u0039\u002e\u0035\u0032\u0037\u0034"
}

/* OUTPUT
Location: 16.7333,-169.5274
*/
