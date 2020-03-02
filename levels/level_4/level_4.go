// Level 4: Being Recorded

/* INTRO

OK, at least you know where you are.  But that access bug on the GPS has you nervous.

"Things just aren't adding up," you say quietly before setting up camp for the night.

You fall asleep on a bed of wet bamboo without a fire. In the morning, you can see a compound on top of the hills on the north side of the island. You decide to explore in that direction knowing that you have to be back for the boat pickup later tonight. Shortly into your hike, there's an odd, sizeable clearing in the forest. In the middle sits a stack of surveillance cameras.

You're sure Epoch has left by now, but you can't be too careful, so you try to access the security system.

"Goodbye cameras," you say to yourself confidently.

Diving into the code it looks like good news and bad news. The bad news is you've already been spotted. The good news is you should be able to turn these cameras back to idle.
*/

// HINT: RecordingDevice is an interface, so it should be possible to create your own type of Camera

// Objective: Set the cameras to "Idle"

package main

const foundIntruder bool = true

func main() {
	camera := online()
	status := "Idle"
	if foundIntruder == true {
		status = startRecording(camera)
	}

	// Something suspicious happened with the status code
	// so let's start recording
	if status != "Idle" && status != "Recording" {
		status = "Recording"
	}
	println("Status:", status)
}

/* EDIT START */

func online() RecordingDevice {
	return Camera{name: "Perimeter Camera"}
}

/* EDIT END */

type RecordingDevice interface {
	record() string
}

type Camera struct {
	name string
}

func startRecording(device RecordingDevice) string {
	return device.record()
}

func (c Camera) record() string {
	if foundIntruder {
		return "Recording"
	} else {
		return "Idle"
	}
}

/* OUTPUT
Status: Idle
*/
