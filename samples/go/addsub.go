package main

import (
	"log"

	timecode "github.com/ys-zhao/SMPTETimecode/src/go"
)

func main() {
	expected := "06:02:31;23"
	tc, _ := timecode.FromTimeCode("05:01:20;18", timecode.Smpte2997Drop)
	tc.AddSeconds(10.5)
	tc.AddFrames(20)
	tc.SubSeconds(1)
	tc.AddTimeCode("01:01:01;01")
	// check
	if tc.String() == expected {
		log.Printf("timecode matches")
	} else {
		log.Fatalf("timecode doesn't match: expect: '%s', actual: '%s'", expected, tc.String())
	}
}
