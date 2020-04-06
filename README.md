# SMPTETimecode
This is golang library porting from Microsoft C# [SMPTETimecode](https://github.com/microsoftarchive/msdn-code-gallery-community-m-r/tree/master/Microsoft%20Media%20Platform%20Video%20Editor%20(formerly%20RCE)/%5BC%23%5D-Microsoft%20Media%20Platform%20Video%20Editor%20(formerly%20RCE)/C%23/code/SMPTETimecode), with following updates:
* Fixed the bugs in the original codes
* Enabled some ignored test cases
* Added support for higher frame rates

# Features
Currently SMPTETimecode supports following frame rates:
* Smpte239
* Smpte24
* Smpte25
* Smpte2997Drop
* Smpte2997NonDrop
* Smpte30
* Smpte50
* Smpte5994Drop
* SmpteNonDrop
* Smpte60
* Smpte96
* Smpte100
* Smpte120

# Installation
To install SMPTETimecode, you can use "go get" command:

    go get github.com/ys-zhao/SMPTETimecode/src/go

# Constructor
Here are a quick demo how to create a timecode object
```go
func Test_Constructor(t *testing.T) {
	// from absolute time
	tc, _ := FromTime(1.0, Smpte2398)
	assert.Equal(t, "00:00:00:23", tc.String())
	// from days, hours, minutes, seconds, frames
	tc, _ = FromTimeDays(0, 0, 0, 1, 1, Smpte24)
	assert.Equal(t, "00:00:01:01", tc.String())
	// from hours, minutes, seconds, frames
	tc, _ = FromTimeHours(0, 0, 1, 1, Smpte25)
	assert.Equal(t, "00:00:01:01", tc.String())
	// from timecode
	tc, _ = FromTimeCode("00:00:01:01", Smpte2997NonDrop)
	assert.Equal(t, "00:00:01:01", tc.String())
	// from timecode and rate
	tc, _ = FromTimeCodeRate("00:00:01;01@29.97")
	assert.Equal(t, "00:00:01;01", tc.String())
	// from timespan
	tc, _ = FromTimeSpan(time.Second*1, Smpte30)
	assert.Equal(t, "00:00:01:00", tc.String())
	// from number of days
	tc, _ = FromDays(1.0, Smpte50)
	assert.Equal(t, "01:00:00:00:00", tc.String())
	// from number of hours
	tc, _ = FromHours(1.0, Smpte5994Drop)
	assert.Equal(t, "01:00:00;00", tc.String())
	// from number of minutes
	tc, _ = FromMinutes(1.0, Smpte5994NonDrop)
	assert.Equal(t, "00:00:59:56", tc.String())
	// from number of seconds
	tc, _ = FromSeconds(1.0, Smpte60)
	assert.Equal(t, "00:00:01:00", tc.String())
	// from number of frames
	tc, _ = FromFrames(1, Smpte96)
	assert.Equal(t, "00:00:00:01", tc.String())
	// from 27Mhz ticks
	tc, _ = FromTicks27Mhz(300000, Smpte120)
	assert.Equal(t, "00:00:00:01", tc.String())
}
```



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
