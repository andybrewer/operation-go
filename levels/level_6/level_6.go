// Level 6: Laser Beams

/* INTRO
Feeling confident from cracking the passcode you enter the compound.

Things finally look to be going your way. With unrestricted access you quickly make your way to the top level where you assume Epoch resides. As you approach the top stair, you can vaguely make out a large desk with a computer and several USBs through the fading evening light.

You're tempted to charge in and grab the USBs and make your way back to the boat, but your cautious side gets the better of you and you test the air with some smoke from your lighter.

"Lasers," you mutter to yourself dejectedly.

Not wanting to leave any trace of your intrusion, you quickly open your laptop to see if there's a way to disarm them. It appears there are 7 lasers blocking your path.

The trick will be to leave the laser grid on, while disabling the lasers themselves.
*/

// HINT: If the lasers were passed by reference instead of value, they could be modified

// Objective: Turn off the lasers while keeping the LaserGrid operational

package main

func main() {
	lasers := setupLasers()

	if len(lasers) != 7 {
		println("ALERT! Wrong number of lasers.")
		return
	}

	for i := 0; i < 7; i++ {
		if !lasers[i].isRunning {
			println("ALERT! Laser not running.")
			return
		}
	}

	laserGrid := LaserGrid{"Operational", lasers}

	passedTest := testGrid(laserGrid)
	if !passedTest {
		println("ALERT! Grid test failed.")
		return
	}

	println("Grid operational")

	running := 0
	for _, laser := range laserGrid.lasers {
		if laser.isRunning {
			running++
		}
	}

	println(running, "lasers running")
}

/* EDIT START */

// LaserGrid represents a collection of 7 lasers
type LaserGrid struct {
	status string
	lasers [7]Laser
}

// Laser respresents an individual laser beam
type Laser struct {
	id        int
	isRunning bool
}

func setupLasers() [7]Laser {
	var lasers [7]Laser
	lasers[0] = Laser{1, true}
	lasers[1] = Laser{2, true}
	lasers[2] = Laser{3, true}
	lasers[3] = Laser{4, true}
	lasers[4] = Laser{5, true}
	lasers[5] = Laser{6, true}
	lasers[6] = Laser{7, true}
	return lasers
}

func testGrid(laserGrid LaserGrid) bool {
	return laserGrid.status == "Operational"
}

/* EDIT END */

/* OUTPUT
Grid operational
0 lasers running
*/
