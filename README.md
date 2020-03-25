# SMPTETimecode
This is golang library porting from Microsoft C# [SMPTETimecode](https://github.com/microsoftarchive/msdn-code-gallery-community-m-r/tree/master/Microsoft%20Media%20Platform%20Video%20Editor%20(formerly%20RCE)/%5BC%23%5D-Microsoft%20Media%20Platform%20Video%20Editor%20(formerly%20RCE)/C%23/code/SMPTETimecode)

# Features
- Support rates: Smpte239, Smpte24, Smpte25, Smpte2997Drop, Smpte2997NonDrop, Smpte30

# Installation
To install SMPTETimecode, you can use "go get" command:

    go get github.com/ys-zhao/SMPTETimecode/src/go



# Examples
```go
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
```
