// Level 12: Secrets Revealed

/* INTRO

"Nice work!" Agent Getter says triumphantly. "The field can be kind of fun, huh?"

"I can't wait to get back to the office," you reply exhaustedly.

As the sun begins to set, you hear the sound of a half dozen helicopters flying towards Epoch's compund. You watch from a bluff as the raid goes flawlessly. Rand is arrested along with the corrupt agents. You and Agent Getter are greeted as heros by the rescue squad.

"I'm just happy to do my job," you reply to a seemlingly endless line of congradulatory remarks.

"Speaking of which, I can't wait to get my hands on those USBs," Agent Getter says breaking into a grin.

"Absolutely! Right after me." you smile back.

No sooner does the flight land back at headquarters than you're at your desk ready to get back to work.

"OK, Epoch, time to see what you were really up to," you say as you settle into your plush work chair.

As you sit down to dive in, your initial enthusiasm turns to concern when you see that as soon as the files are accessed the app panics. What's worse is that even if you could access the files, they're locked in the wrong format.

As you rub your eyes and take a second look, you smile as a small glimmer of hope reveals intself in the code.

Everyone is counting on you to deliver as many readable files as possible.
*/

// HINT: recover() can rescue a program from a panic(); the data passed to enc() is a pointer, meaning you can manipulate the original slice: http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#slice_data_corruption

// Objective: Open all 3 files from the USB using their true file extension

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
)

func init() {
}

func main() {
	println("checking secrets...")
	secrets := loadSecrets(secrets())
	files := loadFiles([]File{
		{[]byte("masterPlan.lck"), 8011},
		{[]byte("financials.lck"), 7005},
		{[]byte("doubleAgents.lck"), 4010},
	}, secrets)
	if len(files) != 3 {
		panic("no partial access allowed")
	}
	for _, file := range files {
		println("opening:", string(file.path))
	}
}

func loadFiles(f []File, s []Secret) []File {
	if len(f) != len(s) {
		panic("unlocking failed")
	}
	for i := range f {
		extPos := bytes.IndexByte(f[i].path, '.')
		if s[i].fileHash != enc(f[i].path[:extPos]) || !unlock(f[i].size, enc(f[i].path[extPos:])) {
			println("Unauthorized access")
			return nil
		}
	}
	return f
}

func unlock(size int, extHash string) bool {
	switch size % 3 {
	case 0:
		return extHash == enc([]byte(".xls"))
	case 1:
		return extHash == enc([]byte(".pdf"))
	default:
		return extHash == enc([]byte(".txt"))
	}
}

// File represents a data file
type File struct {
	path []byte
	size int
}

// Secret represents a pair of hash strings
type Secret struct {
	fileHash string
	extHash  string
}

/* EDIT START */

func enc(b []byte) string {
	sha := sha256.Sum256(b)
	return base64.StdEncoding.EncodeToString(sha[:])
}

func secrets() func(*[]Secret) {
	return nil
}

/* EDIT END */

func loadSecrets(sf func(*[]Secret)) (s []Secret) {
	defer sf(&s)
	panic("files locked")
}

/* OUTPUT
checking secrets...
opening: masterPlan.pdf
opening: financials.xls
opening: doubleAgents.txt
*/
