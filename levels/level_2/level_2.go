// Level 2: Time to Jump

/* INTRO
"OK, all set," you say impatiently.

The pilot rechecks the agent manifest and hands you your full tactical gear. Boarding the plane, you get a bad feeling about the mission.  Those manifests are triple checked and were still wrong...

On board, you keep to yourself and wait for the briefing Rand mentioned, but it seems like no one is in charge. Halfway in, you decide to keep busy by checking the source code on the deployment system for your parachute. As soon as you open your laptop the other agents start to chuckle.

"What? Afraid it won't open?" one shouts sarcasticly.

"It's something to do," you reply quietly. Undeterred, you start checking the code.

"That's odd," you whisper to yourself. Looking at this drop sequence, it doesn't look like your chute would open safely.
*/

// HINT: Variables can be modified from inside a loop

// Objective: Open your chute safely at 450 meters after 80 seconds of freefall

package main

func main() {
	seconds, meters := dropSequence(4000)
	if seconds != 80 || meters != 450 {
		println("Chute error at", seconds, "seconds and", meters, "meters")
	} else {
		println("Chute opened safely")
	}
}

func dropSequence(meters int) (int, int) {
	seconds := 0
	for meters > 450 {
		/* EDIT START */

		// Your code

		/* EDIT END */
		seconds++
		meters -= 44
		if seconds == 80 && meters == 450 {
			break
		}
	}

	return seconds, meters
}

/* OUTPUT
Chute opened safely
*/
