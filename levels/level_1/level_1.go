// Level 1: All Aboard

/* INTRO
"You're going into the field Agent Buggs," yells your boss Rand.

This is the moment you've been dreading since joining the agency. You're perfectly fine hunting down the notorious cyber criminal Epoch from your dual monitors at the agency.

"We've tracked Epoch to a remote island in the Pacific. You're on a plane at 0700, with a briefing on the flight.  Operation Go should be pretty routine," Rand mentions already half way down the hall.

"Thank you sir," you reply hesitantly. Nothing at this agency is routine you think to yourself as you pack up your field laptop and head to the airport. When you arrive, the pilot looks confused.

"Where's your gear Agent Buggs?" he barks.

"There must be some mistake," you reply after noticing all of the other agents wearing full tactical gear. "Let me take a look..."

Logging into the agency network, you're quickly into the operation manifest code. Sure enough, you're listed with no equipment, a typical agency snafu.

You're going to need full tactical gear for this mission. Let's go ahead and fix that...
*/

// HINT: Individual agent data can be modified using agents[X].gear

// Objective: Set your gear to "full"

// NOTES:
//  * Code on a grey background is locked and cannot be edited
//  * Output is validated, so don't use extra print statements

package main

func main() {
	agents := make([]Agent, 0)
	agents = append(agents, Agent{name: "J. Son", gear: "full"})
	agents = append(agents, Agent{name: "A. Pend", gear: "full"})
	agents = append(agents, Agent{name: "D. Buggs", gear: "none"})
	agents = append(agents, Agent{name: "X. Itwon", gear: "full"})
	agents = append(agents, Agent{name: "D. Fercloze", gear: "full"})
	/* EDIT START */

	// TODO: give "D. Buggs" full gear

	/* EDIT END */
	println("Operation Go: Agent Manifest")
	println("----------------------------")
	for _, agent := range agents {
		println(agent.name, "-> Gear:", agent.gear)
	}
}

// Agent represents an agency employee
type Agent struct {
	name string
	gear string
}

/* OUTPUT
Operation Go: Agent Manifest
----------------------------
J. Son -> Gear: full
A. Pend -> Gear: full
D. Buggs -> Gear: full
X. Itwon -> Gear: full
D. Fercloze -> Gear: full
*/
