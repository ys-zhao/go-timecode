package timecode

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/// <summary>
/// Creates the time code_2398 from string.
/// </summary>
func Test_CreateTimeCode_2398FromString(t *testing.T) {
	tc, err := FromTimeCode("01:24:12:13", Smpte2398)
	assert.Nil(t, err)
	assert.Equal(t, "01:24:12:13", tc.String())
}

/// <summary>
/// Creates the time code_2398 from integers.
/// </summary>
func Test_CreateTimeCode_2398FromIntegers(t *testing.T) {
	tc, err := FromTimeHours(1, 24, 12, 13, Smpte2398)
	assert.Nil(t, err)
	assert.Equal(t, "01:24:12:13", tc.String())
}

/// <summary>
/// Creates the timecode 24fps from string.
/// </summary>
func Test_CreateTimeCode_24FromString(t *testing.T) {
	tc, err := FromTimeCode("01:24:12:13", Smpte24)
	assert.Nil(t, err)
	assert.Equal(t, "01:24:12:13", tc.String())
}

/// <summary>
/// Creates the timecode 24fps from integers.
/// </summary>

func Test_CreateTimeCode_24FromIntegers(t *testing.T) {
	tc, err := FromTimeHours(1, 24, 12, 13, Smpte24)

	assert.Nil(t, err)
	assert.Equal(t, "01:24:12:13", tc.String())
}

/// <summary>
/// Creates the timecode 25fps from string.
/// </summary>

func Test_CreateTimeCode_25FromString(t *testing.T) {
	tc, err := FromTimeCode("01:24:12:13", Smpte25)

	assert.Nil(t, err)
	assert.Equal(t, "01:24:12:13", tc.String())
}

/// <summary>
/// Creates the timecode 25fps from integers.
/// </summary>

func Test_CreateTimeCode_25FromIntegers(t *testing.T) {
	tc, err := FromTimeHours(1, 24, 12, 13, Smpte25)

	assert.Nil(t, err)
	assert.Equal(t, "01:24:12:13", tc.String())
}

/// <summary>
/// Creates the timecode  2997 non drop from string.
/// </summary>

func Test_CreateTimeCode_2997NonDropFromString(t *testing.T) {
	tc, err := FromTimeCode("01:42:12:20", Smpte2997NonDrop)

	assert.Nil(t, err)
	assert.Equal(t, "01:42:12:20", tc.String())
}

/// <summary>
/// Creates the timecode 2997 non drop from integers.
/// </summary>

func Test_CreateTimeCode_2997NonDropFromIntegers(t *testing.T) {
	tc, err := FromTimeHours(1, 24, 12, 13, Smpte2997NonDrop)

	assert.Nil(t, err)
	assert.Equal(t, "01:24:12:13", tc.String())
}

/// <summary>
/// Creates the timecode 2997 drop from string.
/// </summary>

func Test_CreateTimeCode_2997DropFromString(t *testing.T) {
	tc, err := FromTimeCode("01:42:12;22", Smpte2997Drop)

	assert.Nil(t, err)
	assert.Equal(t, "01:42:12;22", tc.String())
}

/// <summary>
/// Creates the timecode 2997 drop from integers.
/// </summary>

func Test_CreateTimeCode_2997DropFromIntegers(t *testing.T) {
	tc, err := FromTimeHours(1, 24, 12, 13, Smpte2997Drop)

	assert.Nil(t, err)
	assert.Equal(t, "01:24:12;13", tc.String())
}

/// <summary>
/// Creates the timecode 30 from string.
/// </summary>

func Test_CreateTimeCode_30FromString(t *testing.T) {
	tc, err := FromTimeCode("01:42:12:22", Smpte30)

	assert.Nil(t, err)
	assert.Equal(t, "01:42:12:22", tc.String())
}

/// <summary>
/// Creates the time code_30 from integers.
/// </summary>

func Test_CreateTimeCode_30FromIntegers(t *testing.T) {
	tc, err := FromTimeHours(1, 24, 12, 13, Smpte30)

	assert.Nil(t, err)
	assert.Equal(t, "01:24:12:13", tc.String())
}

/// <summary>
/// Adds two  timecodes @ 30fps.
/// </summary>

func Test_Add2TimeCodes30Fps(t *testing.T) {
	var t1, _ = FromTimeHours(0, 01, 00, 00, Smpte30)
	var t2, _ = FromTimeHours(0, 01, 00, 22, Smpte30)

	t3, err := Add(t1, t2)
	assert.Nil(t, err)
	assert.Equal(t, "00:02:00:22", t3.String())
}

/// <summary>
/// Adds two timecodes @29 nondrop.
/// </summary>

func Test_Add2TimeCodes_29Nondrop(t *testing.T) {
	var t1, _ = FromTimeHours(0, 01, 00, 29, Smpte2997NonDrop)
	var t2, _ = FromTimeHours(0, 00, 00, 02, Smpte2997NonDrop)

	t3, err := Add(t1, t2)
	assert.Nil(t, err)
	assert.Equal(t, "00:01:01:01", t3.String())
}

func Test_Add2TimeCodes_29Drop1(t *testing.T) {
	var t1, _ = FromTimeHours(0, 01, 00, 29, Smpte2997Drop)
	var t2, _ = FromTimeHours(0, 01, 00, 02, Smpte2997Drop)

	t3, err := Add(t1, t2)
	assert.Nil(t, err)
	assert.Equal(t, "00:02:01;01", t3.String())
}

func Test_Add2TimeCodes_29Drop2(t *testing.T) {
	var t1, _ = FromTimeHours(0, 01, 00, 29, Smpte2997Drop)
	var t2, _ = FromTimeHours(0, 01, 00, 02, Smpte2997Drop)

	t3, err := Add(t1, t2)
	assert.Nil(t, err)
	assert.Equal(t, "00:02:01;01", t3.String())
}

/// <summary>
/// Adds two timecodes @ 25 fps PAL/EBU
/// </summary>

func Test_Add2TimeCodes25Fps(t *testing.T) {
	var t1, _ = FromTimeHours(0, 00, 10, 13, Smpte25)
	var t2, _ = FromTimeHours(0, 00, 12, 22, Smpte25)

	t3, err := Add(t1, t2)
	assert.Nil(t, err)
	assert.Equal(t, "00:00:23:10", t3.String())
}

/// <summary>
/// Adds two time codes @ 23.98fps.
/// </summary>

func Test_Add2TimeCodes2398Fps(t *testing.T) {
	var t1, _ = FromTimeHours(0, 00, 10, 13, Smpte2398)
	var t2, _ = FromTimeHours(0, 00, 12, 22, Smpte2398)

	t3, err := Add(t1, t2)
	assert.Nil(t, err)
	assert.Equal(t, "00:00:23:11", t3.String())
}

/// <summary>
/// Adds two time codes @ 23.98fps.
/// </summary>

func Test_Add2TimeCodes2398Fps2(t *testing.T) {
	var t1, _ = FromTimeHours(15, 54, 25, 12, Smpte2398)
	var t2, _ = FromTimeHours(1, 42, 35, 15, Smpte2398)

	t3, err := Add(t1, t2)
	assert.Nil(t, err)
	assert.Equal(t, "17:37:01:03", t3.String())
}

/// <summary>
/// Adds two time codes @ 23.98fps that are both equal.
/// </summary>

func Test_TimeCodesAreEqualBoth2398(t *testing.T) {
	var t1, _ = FromTimeHours(0, 12, 10, 13, Smpte2398)
	var t2, _ = FromTimeHours(0, 12, 10, 13, Smpte2398)

	assert.Equal(t, t1, t2)
}

func Test_Add2TimeCodes_PastMaxTimeCode(t *testing.T) {
	var t1, _ = FromTimeHours(12, 01, 00, 00, Smpte30)
	var t2, _ = FromTimeHours(12, 01, 00, 22, Smpte30)

	_, err := Add(t1, t2)
	assert.NotNil(t, err, "OverflowException")
}

func Test_TimeCodes_AreEqual_DifferentFrameRates_DropAndNon(t *testing.T) {
	var t1, _ = FromTimeCode("00:12:33:26", Smpte2997Drop)
	var t2, _ = FromTimeCode("00:12:33:04", Smpte2997NonDrop)

	assert.True(t, Equal(t1, t2))
}

/// <summary>
/// Checks to see which is less using different rates drop and PAL.
/// </summary>

func Test_TimeCodesLessThanOrEqualDifferentFrameRatesDropAndPAL(t *testing.T) {
	t1, _ := FromTimeCode("00:12:33:26", Smpte2997Drop)
	t2, _ := FromTimeCode("00:12:33:23", Smpte25)

	assert.True(t, LessEqual(t1, t2))
}

func Test_TimeCodesLessThanOrEqual2DifferentFrameRatesDropAndPAL(t *testing.T) {
	t1, _ := FromTimeCode("00:12:31:26", Smpte2997Drop)
	t2, _ := FromTimeCode("00:12:33:22", Smpte25)

	assert.True(t, LessEqual(t1, t2))
}

func Test_TimeCodesLessThanOrEqual3DifferentFrameRates_DropAndPAL(t *testing.T) {
	t1, _ := FromTimeCode("00:12:35:26", Smpte2997Drop)
	t2, _ := FromTimeCode("00:12:33:22", Smpte25)

	assert.False(t, LessEqual(t1, t2))
}

func Test_TimeCodesGreaterThanOrEqualDifferentFrameRatesDropAndPAL(t *testing.T) {
	t1, _ := FromTimeCode("00:12:33:26", Smpte2997Drop)
	t2, _ := FromTimeCode("00:12:33:22", Smpte25)

	assert.True(t, GreatEqual(t1, t2))
}

func Test_TimeCodesGreaterThanOrEqual2DifferentFrameRatesDropAndPAL(t *testing.T) {
	t1, _ := FromTimeCode("00:12:35:26", Smpte2997Drop)
	t2, _ := FromTimeCode("00:12:33:22", Smpte25)

	assert.True(t, GreatEqual(t1, t2))
}

func Test_TimeCodesGreaterThanOrEqual3DifferentFrameRatesDropAndPAL(t *testing.T) {
	t1, _ := FromTimeCode("00:12:31:26", Smpte2997Drop)
	t2, _ := FromTimeCode("00:12:33:22", Smpte25)

	assert.False(t, GreatEqual(t1, t2))
}

func Test_Ticks27MhzToSMPTEDrop(t *testing.T) {
	const Ticks27 int64 = 156523374000
	timecode := Ticks27MhzToSmpte12M(Ticks27, Smpte2997Drop)

	assert.Equal(t, "01:36:37;05", timecode)
}

func Test_Ticks27MhzToSMPTENonDrop(t *testing.T) {
	const Ticks27 int64 = 156523374000
	timecode := Ticks27MhzToSmpte12M(Ticks27, Smpte2997NonDrop)

	assert.Equal(t, "01:36:31:11", timecode)
}

func Test_Ticks27MhzToSMPTE2398(t *testing.T) {
	const Ticks27 int64 = 156523374000
	timecode := Ticks27MhzToSmpte12M(Ticks27, Smpte2398)

	assert.Equal(t, "01:36:31:08", timecode)
}

func Test_Ticks27MhzToSMPTE24(t *testing.T) {
	const Ticks27 int64 = 156523374000
	timecode := Ticks27MhzToSmpte12M(Ticks27, Smpte24)

	assert.Equal(t, "01:36:37:03", timecode)
}

func Test_Ticks27MhzToSMPTE25(t *testing.T) {
	const Ticks27 int64 = 156523374000
	timecode := Ticks27MhzToSmpte12M(Ticks27, Smpte25)

	assert.Equal(t, "01:36:37:04", timecode)
}

func Test_Ticks27MhzToSMPTE30(t *testing.T) {
	const Ticks27 int64 = 156522600000
	timecode := Ticks27MhzToSmpte12M(Ticks27, Smpte30)

	assert.Equal(t, "01:36:37:04", timecode)
}

func Test_SMPTE30ToTicks27Mhz(t *testing.T) {
	const Timecode string = "01:36:37:04"
	ticks27Mhz := smpte12MToTicks27Mhz(Timecode, Smpte30)

	assert.Equal(t, int64(156522600000), ticks27Mhz)
}

func Test_TimecodeFromTicks27BackToString(t *testing.T) {
	const Timecode string = "01:36:37:04"
	ticks27Mhz := smpte12MToTicks27Mhz(Timecode, Smpte30)

	t1, _ := FromTicks27Mhz(ticks27Mhz, Smpte30)

	assert.Equal(t, "01:36:37:04", t1.String())
}

func Test_TimecodeFromTicks27BackToString2398(t *testing.T) {
	const Timecode string = "01:36:31:08"
	ticks27Mhz := smpte12MToTicks27Mhz(Timecode, Smpte2398)

	t1, _ := FromTicks27Mhz(ticks27Mhz, Smpte2398)

	assert.Equal(t, "01:36:31:08", t1.String())
}

func Test_TimecodeFromTicks27BackToString24(t *testing.T) {
	const Timecode string = "01:36:37:05"
	ticks27Mhz := smpte12MToTicks27Mhz(Timecode, Smpte24)

	t1, _ := FromTicks27Mhz(ticks27Mhz, Smpte24)

	assert.Equal(t, "01:36:37:05", t1.String())
}

func Test_TimecodeFromTicks27BackToString25(t *testing.T) {
	const Timecode string = "01:36:37:05"
	ticks27Mhz := smpte12MToTicks27Mhz(Timecode, Smpte25)

	t1, _ := FromTicks27Mhz(ticks27Mhz, Smpte25)

	assert.Equal(t, "01:36:37:05", t1.String())
}

func Test_TimecodeFromTicks27BackToString2997Drop(t *testing.T) {
	const Timecode string = "01:36:37;05"
	ticks27Mhz := smpte12MToTicks27Mhz(Timecode, Smpte2997Drop)

	t1, _ := FromTicks27Mhz(ticks27Mhz, Smpte2997Drop)

	assert.Equal(t, "01:36:37;05", t1.String())
}

func Test_TimecodeFromTicks27BackToString2997NonDrop(t *testing.T) {
	const Timecode string = "01:36:37:05"
	ticks27Mhz := smpte12MToTicks27Mhz(Timecode, Smpte2997NonDrop)

	t1, _ := FromTicks27Mhz(ticks27Mhz, Smpte2997NonDrop)

	assert.Equal(t, "01:36:37:05", t1.String())
}

func Test_ValidateBadTimecode1(t *testing.T) {
	valid := validateSmpte12MTimecode("24:00:00:12")

	assert.False(t, valid)
}

func Test_ValidateBadTimecode2(t *testing.T) {
	valid := validateSmpte12MTimecode("01:60:10:10")

	assert.False(t, valid)
}

func Test_ValidateGoodTimecode(t *testing.T) {
	valid := validateSmpte12MTimecode("23:38:10:10")

	assert.True(t, valid)
}

func Test_ValidateAreEqualDropFrame(t *testing.T) {
	t1, _ := FromTimeCode("01:01:43;02", Smpte2997Drop)
	t2, _ := FromTimeCode("01:01:43;02", Smpte2997Drop)

	assert.Equal(t, t1, t2)
}

func Test_ValidateAreNotEqualDropFrame(t *testing.T) {
	t1, _ := FromTimeCode("01:01:43;01", Smpte2997Drop)
	t2, _ := FromTimeCode("01:01:43;02", Smpte2997Drop)

	assert.NotEqual(t, t1, t2)
}

func Test_ValidateAreEqualAddOneFrame(t *testing.T) {
	t1, _ := FromTimeCode("01:29:45:15", Smpte2997Drop)
	t2, _ := FromTimeCode("01:29:45:16", Smpte2997Drop)
	t3, _ := FromTimeHours(0, 0, 0, 1, Smpte2997Drop)

	expected, _ := Add(t1, t3)

	assert.Equal(t, expected.TotalSecondsPrecision(), t2.TotalSecondsPrecision())
}

func Test_AddAndSubtractTimecodesDrop(t *testing.T) {
	t1, _ := FromTimeCode("00:58:12:15", Smpte2997Drop)
	t2, _ := FromTimeCode("00:02:00:00", Smpte2997Drop)
	t3, _ := FromTimeCode("01:22:12:15", Smpte2997Drop)

	t4, _ := Add(t1, t3)
	t4, _ = Sub(t4, t2)

	assert.Equal(t, "02:18:25;00", t4.String())
}

func Test_Add3TimecodesDrop(t *testing.T) {
	t1, _ := FromTimeCode("00:01:00:02", Smpte2997Drop)
	t2, _ := FromTimeCode("00:10:00:00", Smpte2997Drop)
	t3, _ := FromTimeCode("01:00:00:00", Smpte2997Drop)

	t4, _ := Add(t1, t2)
	t4, _ = Add(t4, t3)

	assert.Equal(t, "01:11:00;02", t4.String())
}

func Test_AbsoluteTimeToDropFrame(t *testing.T) {
	const AbsoluteTime float64 = 8304.963333333335

	t1, _ := FromTime(AbsoluteTime, Smpte2997Drop)

	assert.Equal(t, "02:18:25;00", t1.String())
}

func Test_GetTotalHours(t *testing.T) {
	t1, _ := FromTimeCode("01:30:00:00", Smpte2997NonDrop)

	assert.Equal(t, float64(1.5), t1.TotalHours())
}

func Test_GetTotalMinutes(t *testing.T) {
	t1, _ := FromTimeCode("01:30:00:00", Smpte2997NonDrop)

	assert.Equal(t, float64(90), t1.TotalMinutes())
}

func Test_FromHours(t *testing.T) {
	t1, _ := FromHours(1.5, Smpte30)

	assert.Equal(t, float64(1.5), t1.TotalHours())
}

func Test_AbsoluteToSmpte2997Drop(t *testing.T) {
	const AbsoluteTime float64 = 37.837800000000000000000

	tc, err := FromTime(AbsoluteTime, Smpte2997NonDrop)
	assert.Nil(t, err)
	assert.Equal(t, "00:00:37:24", tc.String())
}

func Test_SubtractToMinValue(t *testing.T) {
	timecode1, _ := FromTimeHours(0, 0, 0, 1, Smpte2997NonDrop)
	timecode2, _ := FromTimeHours(0, 0, 0, 1, Smpte2997NonDrop)

	time3, err := Sub(timecode1, timecode2)
	assert.Nil(t, err)
	assert.Equal(t, "00:00:00:00", time3.String())
}

func Test_SubtractPastMinValue(t *testing.T) {
	timecode1, _ := FromTimeHours(0, 0, 0, 1, Smpte2997NonDrop)
	timecode2, _ := FromTimeHours(0, 0, 0, 4, Smpte2997NonDrop)

	_, err := Sub(timecode1, timecode2)
	assert.NotNil(t, err, "MinValueSmpte12MOverflowException")
}

/// <summary>
/// Checks to see that 1000/1001 is coming back as expected. This is the slowdown rate of 29.97 video.
/// </summary>

func Test_CheckDotNetMathDivide1000By1001(t *testing.T) {
	value, _ := strconv.ParseFloat("0.999000999000999000999000999", 64)
	assert.Equal(t, float64(1000)/1001, value)
}

/// <summary>
/// Checks to see that 30 / (1000/1001) is coming back as expected
/// </summary>

func Test_CheckDotNetMathDivide30By1000By1001(t *testing.T) {
	t.Skip()
	v, _ := strconv.ParseFloat("30.03", 64)
	assert.Equal(t, (30 / (1000 / float64(1001))), v)
}

/// <summary>
/// Checks to see that 29 frames is converting to absolute time as expected
/// </summary>

func Test_CheckDotNetMathConvert29Frames2997ToAbsoluteTime(t *testing.T) {
	const Frames int = 29
	AbsoluteTime := float64(Frames) / float64(30) / (1000 / float64(1001))
	value, _ := strconv.ParseFloat("0.9676333333333333333333333334", 64)

	assert.Equal(t, AbsoluteTime, value)
}

/// <summary>
/// Checks to see that 31 frames is converting to absolute time as expected
/// </summary>

func Test_CheckDotNetMathConvert31Frames2997ToAbsoluteTime(t *testing.T) {
	const Frames int = 31
	AbsoluteTime := float64(Frames) / float64(30) / (1000 / float64(1001))
	v, _ := strconv.ParseFloat("1.0343666666666666666666666666", 64)

	assert.Equal(t, AbsoluteTime, v)
}

/// <summary>
/// Checks to see that 29 frames is actually 00:00:00:29 in Smpte 29.97 timecode.
/// </summary>

func Test_CheckFromFrames29IsEqualToString(t *testing.T) {
	const Frames int64 = 29
	t1, _ := FromFrames(Frames, Smpte2997NonDrop)

	assert.Equal(t, "00:00:00:29", t1.String())
}

/// <summary>
/// Checks to see that 30 frames is actually 00:00:01:00 in Smpte 29.97 timecode.
/// </summary>

func Test_CheckFromFrames30IsEqualToString(t *testing.T) {
	const Frames int64 = 30
	t1, _ := FromFrames(Frames, Smpte2997NonDrop)

	assert.Equal(t, "00:00:01:00", t1.String())
}

/// <summary>
/// Checks to see that 31 frames is actually 00:00:01:01 in Smpte 29.97 timecode.
/// </summary>

func Test_CheckFromFrames31IsEqualToString(t *testing.T) {
	const Frames int64 = 31
	t1, _ := FromFrames(Frames, Smpte2997NonDrop)

	assert.Equal(t, "00:00:01:01", t1.String())
}

/// <summary>
/// This test adds frames to an existing timecode and makes sure it is correct.
/// </summary>

func Test_AddFramesToExistingTimeCodeEnsureCorrectTime(t *testing.T) {
	timeSpanInitial := time.Second * 1

	startingTimeCode, _ := FromTimeSpan(timeSpanInitial, Smpte2997NonDrop)

	assert.Equal(t, int64(29), startingTimeCode.TotalFrames())

	totalFrameCount := startingTimeCode.TotalFrames()
	assert.Equal(t, int64(29), totalFrameCount)

	newFrameCount := totalFrameCount + 1
	assert.Equal(t, int64(30), newFrameCount)

	frameAddedTimeCode, _ := FromFrames(newFrameCount, Smpte2997NonDrop)
	assert.Equal(t, int64(30), frameAddedTimeCode.TotalFrames())

	newTime := frameAddedTimeCode.TotalSeconds()

	newTimeCode, _ := FromSeconds(newTime, Smpte2997NonDrop)

	assert.NotEqual(t, startingTimeCode.TotalFrames(), newTimeCode.TotalFrames())
	assert.Equal(t, int64(30), newTimeCode.TotalFrames())
}

/// <summary>
/// This test adds frames to an existing 2398 format timecode and makes sure it is correct.
/// </summary>

func Test_Add_Frames_To_Existing_TimeCode_2398_Ensure_Correct_Time(t *testing.T) {
	currentTimeCode, _ := FromSeconds(float64(1000), Smpte2398)

	expectedFrameCount := currentTimeCode.TotalFrames()

	for i := 1; i < 30000; i++ {
		expectedFrameCount++
		// Console.WriteLine("----------------------------------------");
		// Console.WriteLine("Iteration # {0} ExpectedFrames: {1}", i, expectedFrameCount);

		detal, _ := FromFrames(1, Smpte2398)
		currentTimeCode, _ = currentTimeCode.Add(detal)
		// Console.WriteLine("currentTimeCode: " + currentTimeCode);
		// Console.WriteLine("currentTimeCode.TotalFrames: " + currentTimeCode.TotalFrames);
		// Console.WriteLine("currentTimeCode.TotalSeconds: " + currentTimeCode.TotalSeconds);
		// Console.WriteLine("currentTimeCode.TotalSecondsPrecision: " + currentTimeCode.TotalSecondsPrecision);
		// Console.WriteLine("currentTimeCode.ToTicks27Mhz: " + currentTimeCode.ToTicks27Mhz());
		assert.Equal(t, expectedFrameCount, currentTimeCode.TotalFrames(), "Expected frames {0} did not match TotalFrames {1}", expectedFrameCount, currentTimeCode.TotalFrames)
		// Console.WriteLine("----------------------------------------");
	}
}

/// <summary>
/// This test adds frames to an existing 2997 dropframe format timecode and makes sure it is correct.
/// </summary>

func Test_AddFramesToExistingTimeCode2997DfEnsureCorrectTime(t *testing.T) {
	currentTimeCode, _ := FromSeconds(float64(1000), Smpte2997Drop)

	expectedFrameCount := currentTimeCode.TotalFrames()

	for i := 1; i < 30000; i++ {
		expectedFrameCount++
		// Console.WriteLine("----------------------------------------");
		// Console.WriteLine("Iteration # {0} ExpectedFrames: {1}", i, expectedFrameCount);

		detal, _ := FromFrames(1, Smpte2997Drop)
		currentTimeCode, _ = currentTimeCode.Add(detal)
		// Console.WriteLine("currentTimeCode: " + currentTimeCode);
		// Console.WriteLine("currentTimeCode.TotalFrames: " + currentTimeCode.TotalFrames);
		// Console.WriteLine("currentTimeCode.TotalSeconds: " + currentTimeCode.TotalSeconds);
		// Console.WriteLine("currentTimeCode.TotalSecondsPrecision: " + currentTimeCode.TotalSecondsPrecision);
		// Console.WriteLine("currentTimeCode.ToTicks27Mhz: " + currentTimeCode.ToTicks27Mhz());
		assert.Equal(t, expectedFrameCount, currentTimeCode.TotalFrames(), "Expected frames %v did not match TotalFrames %v", expectedFrameCount, currentTimeCode.TotalFrames())
		// Console.WriteLine("----------------------------------------");
	}
}

/// <summary>
/// This test adds frames to an existing 2997 Non Drop format timecode and makes sure it is correct.
/// </summary>

func Test_AddFramesToExistingTimeCode2997NdEnsureCorrectTime(t *testing.T) {
	currentTimeCode, _ := FromSeconds(float64(1000), Smpte2997NonDrop)

	expectedFrameCount := currentTimeCode.TotalFrames()

	for i := 1; i < 30000; i++ {
		expectedFrameCount++
		// Console.WriteLine("----------------------------------------");
		// Console.WriteLine("Iteration # {0} ExpectedFrames: {1}", i, expectedFrameCount);

		delta, _ := FromFrames(1, Smpte2997NonDrop)
		currentTimeCode, _ = currentTimeCode.Add(delta)
		// Console.WriteLine("currentTimeCode: " + currentTimeCode);
		// Console.WriteLine("currentTimeCode.TotalFrames: " + currentTimeCode.TotalFrames);
		// Console.WriteLine("currentTimeCode.TotalSeconds: " + currentTimeCode.TotalSeconds);
		// Console.WriteLine("currentTimeCode.TotalSecondsPrecision: " + currentTimeCode.TotalSecondsPrecision);
		// Console.WriteLine("currentTimeCode.ToTicks27Mhz: " + currentTimeCode.ToTicks27Mhz());
		// Console.WriteLine("currentTimeCode.PcrTB: " + currentTimeCode.ToTicksPcrTb());
		assert.Equal(t, expectedFrameCount, currentTimeCode.TotalFrames(), "Expected frames %v did not match TotalFrames %v", expectedFrameCount, currentTimeCode.TotalFrames())
		// Console.WriteLine("----------------------------------------");
	}
}

func Test_Add_Frames_To_Existing_TimeCode_Smpte2398_Ensure_Correct_Time(t *testing.T) {
	currentTimeCode, _ := FromSeconds(float64(0), Smpte2398)
	oneFrame, _ := FromFrames(1, Smpte2398)
	expectedFrameCount := currentTimeCode.TotalFrames()
	assert.Equal(t, int64(1), oneFrame.TotalFrames())
	for i := 1; i < 50000; i++ {
		expectedFrameCount++
		currentTimeCode, _ = currentTimeCode.Add(oneFrame)
		assert.Equal(t, expectedFrameCount, currentTimeCode.TotalFrames(), "Expected frames %v did not match TotalFrames %v", expectedFrameCount, currentTimeCode.TotalFrames)
	}
}
func Test_Add_Frames_To_Existing_TimeCode_Smpte24_Ensure_Correct_Time(t *testing.T) {
	currentTimeCode, _ := FromSeconds(float64(0), Smpte24)
	oneFrame, _ := FromFrames(1, Smpte24)
	expectedFrameCount := currentTimeCode.TotalFrames()
	assert.Equal(t, int64(1), oneFrame.TotalFrames())
	for i := 1; i < 50000; i++ {
		expectedFrameCount++
		currentTimeCode, _ = currentTimeCode.Add(oneFrame)
		assert.Equal(t, expectedFrameCount, currentTimeCode.TotalFrames(), "Expected frames %v did not match TotalFrames %v", expectedFrameCount, currentTimeCode.TotalFrames)
	}
}

/// <summary>
/// This test adds frames to an existing 25 PAL format timecode and makes sure it is correct.
/// </summary>

func Test_Add_Frames_To_Existing_TimeCode_Smpte25_Ensure_Correct_Time(t *testing.T) {
	t.Skip("TODO")
	currentTimeCode, _ := FromSeconds(float64(0), Smpte25)
	oneFrame, _ := FromFrames(1, Smpte25)
	expectedFrameCount := currentTimeCode.TotalFrames()
	assert.Equal(t, int64(1), oneFrame.TotalFrames())
	for i := 1; i < 50000; i++ {
		expectedFrameCount++
		currentTimeCode, _ = currentTimeCode.Add(oneFrame)
		assert.Equal(t, expectedFrameCount, currentTimeCode.TotalFrames(), "Expected frames %v did not match TotalFrames %v", expectedFrameCount, currentTimeCode.TotalFrames)
	}
}

func Test_Add_Frames_To_Existing_TimeCode_Smpte2997Drop_Ensure_Correct_Time(t *testing.T) {
	currentTimeCode, _ := FromSeconds(float64(0), Smpte2997Drop)
	oneFrame, _ := FromFrames(1, Smpte2997Drop)
	expectedFrameCount := currentTimeCode.TotalFrames()
	assert.Equal(t, int64(1), oneFrame.TotalFrames())
	for i := 1; i < 50000; i++ {
		expectedFrameCount++
		currentTimeCode, _ = currentTimeCode.Add(oneFrame)
		assert.Equal(t, expectedFrameCount, currentTimeCode.TotalFrames(), "Expected frames %v did not match TotalFrames %v", expectedFrameCount, currentTimeCode.TotalFrames)
	}
}

func Test_Add_Frames_To_Existing_TimeCode_Smpte2997NonDrop_Ensure_Correct_Time(t *testing.T) {
	currentTimeCode, _ := FromSeconds(float64(0), Smpte2997NonDrop)
	oneFrame, _ := FromFrames(1, Smpte2997NonDrop)
	expectedFrameCount := currentTimeCode.TotalFrames()
	assert.Equal(t, int64(1), oneFrame.TotalFrames())
	for i := 1; i < 50000; i++ {
		expectedFrameCount++
		currentTimeCode, _ = currentTimeCode.Add(oneFrame)
		assert.Equal(t, expectedFrameCount, currentTimeCode.TotalFrames(), "Expected frames %v did not match TotalFrames %v", expectedFrameCount, currentTimeCode.TotalFrames)
	}
}

func Test_Add_Frames_To_Existing_TimeCode_Smpte30_Ensure_Correct_Time(t *testing.T) {
	t.Skip()
	currentTimeCode, _ := FromSeconds(float64(0), Smpte30)
	oneFrame, _ := FromFrames(1, Smpte30)
	expectedFrameCount := currentTimeCode.TotalFrames()
	assert.Equal(t, int64(1), oneFrame.TotalFrames())
	for i := 1; i < 50000; i++ {
		expectedFrameCount++
		currentTimeCode, _ = currentTimeCode.Add(oneFrame)
		assert.Equal(t, expectedFrameCount, currentTimeCode.TotalFrames(), "Expected frames %v did not match TotalFrames %v", expectedFrameCount, currentTimeCode.TotalFrames)
	}
}

/// <summary>
/// Determines whether this value equals one frame.
/// </summary>

func Test_IsThisValueEqualToOneFrame(t *testing.T) {
	t.Skip("This is test case has some issue")
	const OneFrame float64 = 0.033366666666667
	var oneSecTimeCode, _ = FromTimeHours(0, 0, 0, 1, Smpte2997NonDrop)

	var newTimeCode, _ = FromTime(OneFrame, Smpte2997NonDrop)
	assert.Equal(t, oneSecTimeCode.TotalFrames(), newTimeCode.TotalFrames())
}

/// <summary>
/// Determines if two timecode values with slightly differing absoluteTimes are still equal.
/// </summary>

func Test_TwoTimecodesAreEqualDurationInSeconds(t *testing.T) {
	var t1, _ = FromTime(2096.111665, Smpte2997Drop)
	var t2, _ = FromTime(2096.11166497009, Smpte2997Drop)

	assert.Equal(t, t1.TotalSeconds(), t2.TotalSeconds(), "The two timecodes are not evaluating to Equal as expected.")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_1(t *testing.T) {
	const Known float64 = 1300.073

	var t1, _ = FromTime(Known, Smpte25)

	assert.Equal(t, "00:21:40:01", t1.String(), "Not valid timecode for 1300.073")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_2(t *testing.T) {
	const Known float64 = 355.04

	var t1, _ = FromTime(Known, Smpte2997NonDrop)

	assert.Equal(t, "00:05:54:20", t1.String(), "Not valid timecode for 355.04D")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_3(t *testing.T) {
	const Known float64 = 1315.736

	var t1, _ = FromTime(Known, Smpte2997NonDrop)

	assert.Equal(t, "00:21:54:12", t1.String(), "Not valid timecode for 1315.736")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_4(t *testing.T) {
	const Known float64 = 1655.112

	var t1, _ = FromTime(Known, Smpte2997NonDrop)

	assert.Equal(t, "00:27:33:13", t1.String(), "Not valid timecode for 1655.112")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_6(t *testing.T) {
	const Known float64 = 1754.315

	var t1, _ = FromTime(Known, Smpte2997NonDrop)

	assert.Equal(t, "00:29:12:16", t1.String(), "Not valid timecode for 1754.315")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_7(t *testing.T) {
	const Known float64 = 1926.613

	var t1, _ = FromTime(Known, Smpte2997Drop)

	assert.Equal(t, "00:32:06;18", t1.String(), "Not valid timecode for 1926.613")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_8(t *testing.T) {
	const Known float64 = 4965.337

	var t1, _ = FromTime(Known, Smpte2997Drop)

	assert.Equal(t, "01:22:45;09", t1.String(), "Not valid timecode for 4965.337")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_9(t *testing.T) {
	const Known float64 = 12342.52342

	var t1, _ = FromTime(Known, Smpte2997Drop)

	assert.Equal(t, "03:25:42;15", t1.String(), "Not valid timecode for 12342.52342")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_10(t *testing.T) {
	const Known float64 = 4885.23489

	var t1, _ = FromTime(Known, Smpte25)

	assert.Equal(t, "01:21:25:05", t1.String(), "Not valid timecode for 4885.23489")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_11(t *testing.T) {
	const Known float64 = 8948.233667

	var t1, _ = FromTime(Known, Smpte24)

	assert.Equal(t, "02:29:08:05", t1.String(), "Not valid timecode for 8948.233667")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_12(t *testing.T) {
	const Known float64 = 5797.133333

	var t1, _ = FromTime(Known, Smpte30)

	assert.Equal(t, "01:36:37:04", t1.String(), "Not valid timecode for 5797.133333")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_13(t *testing.T) {
	const Known float64 = 56197.4333333333

	var t1, _ = FromTime(Known, Smpte30)

	assert.Equal(t, "15:36:37:13", t1.String(), "Not valid timecode for 56197.4333333333")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_14(t *testing.T) {
	t.Skip("[Ignore]")
	const Known27Mhz float64 = 43443649200

	var t1, _ = FromTime(Known27Mhz, Smpte2398)

	assert.Equal(t, "00:26:47:09", t1.String(), "Not valid timecode for 43443649200 (27Mhz)")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_15(t *testing.T) {
	const Known2398Time string = "00:26:47:09"

	var t1, _ = FromTimeCode(Known2398Time, Smpte2398)

	assert.Equal(t, int64(38577), t1.TotalFrames(), "Should be 38578 frames")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownAbsoluteTimeIsValid_16(t *testing.T) {
	const Known2398Time1 string = "00:24:50:01"
	var t1, _ = FromTimeCode(Known2398Time1, Smpte2398)
	assert.Equal(t, int64(35761), t1.TotalFrames(), "Should be 35761 frames")

	const Known2398Time2 string = "00:24:50:02"
	var t2, _ = FromTimeCode(Known2398Time2, Smpte2398)
	assert.Equal(t, int64(35762), t2.TotalFrames(), "Should be 35762 frames")
}

func Test_KnownTimecode_Smpte24_MatchesString1(t *testing.T) {
	var t1, _ = FromTimeHours(10, 10, 20, 5, Smpte24)
	result := t1.String()
	const Expected string = "10:10:20:05"

	assert.Equal(t, Expected, result)
}

func Test_KnownTimecode_Smpte24_MatchesString2(t *testing.T) {
	var t1, _ = FromTimeHours(0, 23, 20, 5, Smpte24)
	result := t1.String()
	const Expected string = "00:23:20:05"

	assert.Equal(t, Expected, result)
}

func Test_KnownTimecode_Smpte24_MatchesString3(t *testing.T) {
	t.Skip("[Ignore]")
	var t1, _ = FromTimeHours(7, 01, 20, 5, Smpte24)
	result := t1.String()
	const Expected string = "07:01:20:05"

	assert.Equal(t, Expected, result)
}

func Test_KnownTimecode_Smpte24_MatchesString4_Days(t *testing.T) {
	t.Skip("[Ignore]")
	var t1, _ = FromTimeDays(12, 22, 01, 20, 5, Smpte24)
	result := t1.String()
	const Expected string = "12:22:01:20:05"

	assert.Equal(t, Expected, result)
}

func Test_KnownTimecode_Smpte2997DF_MatchesString1_Days(t *testing.T) {
	t.Skip("[Ignore]")
	var t1, _ = FromTimeDays(12, 22, 01, 20, 5, Smpte2997Drop)
	result := t1.String()
	const Expected string = "12:22:01:20;05"

	assert.Equal(t, Expected, result)
}

func Test_KnownTimecode_Smpte299ND_MatchesString1_Days(t *testing.T) {
	t.Skip("[Ignore]")
	var t1, _ = FromTimeDays(12, 22, 01, 20, 5, Smpte2997NonDrop)
	result := t1.String()
	const Expected string = "12:22:01:20:05"

	assert.Equal(t, Expected, result)
}

/// <summary>
/// Checks the absolute time to frames algorithm for 2398.
/// </summary>

func Test_CheckSomeAbsoluteTimeToFramesAlgorithmFor2398(t *testing.T) {
	Epsilon := newDecimalString("0.00000000000000000000000001")

	absoluteTime := newDecimalString("23.481791666666666666666666648")
	frames := newDecimal(24).MulFloat64(1000).DivFloat64(1001).Mul(absoluteTime.Add(Epsilon)).Round(26).Floor().Int64()
	assert.Equal(t, int64(562), frames, "wrong # of frames for this absolute Time.")

	absoluteTime = newDecimal(23.48179166667)
	frames = newDecimal(24).MulFloat64(1000).DivFloat64(1001).Mul(absoluteTime.Add(Epsilon)).Round(26).Floor().Int64()
	assert.Equal(t, int64(563), frames, "wrong # of frames for this absolute Time")
}

func Test_ShouldHandleMax2398Value(t *testing.T) {
	const Known float64 = 86486.36

	var t1, _ = FromTime(Known, Smpte2398)

	assert.Equal(t, "23:59:59:23", t1.String(), "Not valid timecode for 86486.358291666700000")
}

func Test_ShouldHandleMax24Value(t *testing.T) {
	const Known float64 = 86399.958333333300000

	var t1, _ = FromTime(Known, Smpte24)

	assert.Equal(t, "23:59:59:23", t1.String(), "Not valid timecode for 86399.958333333300000")
}

func Test_ShouldHandleMax25Value(t *testing.T) {
	const Known float64 = 86399.960000000000000

	var t1, _ = FromTime(Known, Smpte25)

	assert.Equal(t, "23:59:59:24", t1.String(), "Not valid timecode for 86399.960000000000000")
}

func Test_ShouldHandleMax2997DFValue(t *testing.T) {
	const Known float64 = 86399.880233333300000

	var t1, _ = FromTime(Known, Smpte2997Drop)

	assert.Equal(t, "23:59:59;29", t1.String(), "Not valid timecode for 86399.880233333300000")
}

func Test_ShouldHandleMax2997NDValue(t *testing.T) {
	const Known float64 = 86486.366633333300000

	var t1, _ = FromTime(Known, Smpte2997NonDrop)

	assert.Equal(t, "23:59:59:29", t1.String(), "Not valid timecode for 86486.366633333300000")
}

func Test_ShouldHandleMax30Value(t *testing.T) {
	const Known float64 = 86399.966666666700000

	var t1, _ = FromTime(Known, Smpte30)

	assert.Equal(t, "23:59:59:29", t1.String(), "Not valid timecode for 86399.966666666700000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn2398_1(t *testing.T) {
	const Known float64 = 86486.400000000000000

	var t1, _ = FromTime(Known, Smpte2398)

	assert.Equal(t, "01:00:00:00:00", t1.String(), "Not valid timecode for 86486.400000000000000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn2398_2(t *testing.T) {
	const Known float64 = 86486.441708333300000

	var t1, _ = FromTime(Known, Smpte2398)

	assert.Equal(t, "01:00:00:00:01", t1.String(), "Not valid timecode for 86486.441708333300000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn2398_3(t *testing.T) {
	const FrameCount int64 = 4924825

	var t1, _ = FromFrames(FrameCount, Smpte2398)

	assert.Equal(t, "02:09:00:01:01", t1.String(), "Not valid timecode for 205406.242708333")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn24_1(t *testing.T) {
	const Known float64 = 86400.000000000000000

	var t1, _ = FromTime(Known, Smpte24)

	assert.Equal(t, "01:00:00:00:00", t1.String(), "Not valid timecode for 86400.000000000000000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn24_2(t *testing.T) {
	const Known float64 = 86400.041666666700000

	var t1, _ = FromTime(Known, Smpte24)

	assert.Equal(t, "01:00:00:00:01", t1.String(), "Not valid timecode for 86400.041666666700000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn24_3(t *testing.T) {
	t.Skip("[Ignore]")
	const KnownTicks27 int64 = 759108344625000

	var t1, _ = FromTicks27Mhz(KnownTicks27, Smpte24)

	assert.Equal(t, "325:09:45:23:21", t1.String(), "759108344625000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn25_1(t *testing.T) {
	const Known float64 = 86400.000000000000000

	var t1, _ = FromTime(Known, Smpte25)

	assert.Equal(t, "01:00:00:00:00", t1.String(), "Not valid timecode for 86400.000000000000000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn25_2(t *testing.T) {
	const Known float64 = 86400.040000000000000

	var t1, _ = FromTime(Known, Smpte25)

	assert.Equal(t, "01:00:00:00:01", t1.String(), "Not valid timecode for 86400.040000000000000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn25_3(t *testing.T) {
	const Known float64 = 167462.200000000000000

	var t1, _ = FromTime(Known, Smpte25)

	assert.Equal(t, "01:22:31:02:05", t1.String(), "Not valid timecode for 167462.200000000000000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn2997ND_1(t *testing.T) {
	const Known float64 = 86486.400000000000000

	var t1, _ = FromTime(Known, Smpte2997NonDrop)

	assert.Equal(t, "01:00:00:00:00", t1.String(), "Not valid timecode for 86486.400000000000000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn2997ND_2(t *testing.T) {
	const Known float64 = 86486.433366666700000

	var t1, _ = FromTime(Known, Smpte2997NonDrop)

	assert.Equal(t, "01:00:00:00:01", t1.String(), "Not valid timecode for 86486.433366666700000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn2997ND_3(t *testing.T) {
	const Known float64 = 309753.877766667000000

	var t1, _ = FromTime(Known, Smpte2997NonDrop)

	assert.Equal(t, "03:13:57:24:13", t1.String(), "Not valid timecode for 309753.877766667000000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn2997DF_1(t *testing.T) {
	t.Skip("[Ignore]")

	const Known float64 = 86399.913600000000000

	var t1, _ = FromTime(Known, Smpte2997Drop)

	assert.Equal(t, "01:00:00:00;00", t1.String(), "Not valid timecode for 86399.913600000000000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn2997DF_2(t *testing.T) {
	t.Skip("[Ignore]")
	const Known float64 = 86399.946966666700000

	var t1, _ = FromTime(Known, Smpte2997Drop)

	assert.Equal(t, "01:00:00:00;01", t1.String(), "Not valid timecode for 86399.946966666700000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn2997DF_3(t *testing.T) {
	t.Skip("[Ignore]")
	const Known float64 = 441921.313166667000000

	var t1, _ = FromTime(Known, Smpte2997Drop)

	assert.Equal(t, "05:02:45:21;23", t1.String(), "Not valid timecode for 441921.313166667000000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn30_1(t *testing.T) {
	const Known float64 = 86400.000000000000000

	var t1, _ = FromTime(Known, Smpte30)

	assert.Equal(t, "01:00:00:00:00", t1.String(), "Not valid timecode for 86400.000000000000000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn30_2(t *testing.T) {
	const Known float64 = 86400.033333333300000

	var t1, _ = FromTime(Known, Smpte30)

	assert.Equal(t, "01:00:00:00:01", t1.String(), "Not valid timecode for 86400.033333333300000")
}

/// <summary>
/// Knows the absolute time is valid.
/// </summary>

func Test_KnownLongRunningAbsoluteTimeIsValidIn30_3(t *testing.T) {
	const Known float64 = 334769.833333333000000

	var t1, _ = FromTime(Known, Smpte30)

	assert.Equal(t, "03:20:59:29:25", t1.String(), "Not valid timecode for 334769.833333333000000")
}

/// <summary>
/// Creates the time code_2398 from string.
/// </summary>

func Test_CreateLongRunningTimeCode_2398FromString(t *testing.T) {
	tc, err := FromTimeCode("2:01:24:12:13", Smpte2398)

	assert.Nil(t, err)

	assert.Equal(t, "02:01:24:12:13", tc.String())
}

/// <summary>
/// Creates the time code_2398 from integers.
/// </summary>

func Test_CreateLongRunningTimeCode_2398FromIntegers(t *testing.T) {
	tc, err := FromTimeDays(2, 1, 24, 12, 13, Smpte2398)

	assert.Nil(t, err)

	assert.Equal(t, "02:01:24:12:13", tc.String())
}

/// <summary>
/// Creates the timecode 24fps from string.
/// </summary>

func Test_CreateLongRunningTimeCode_24FromString(t *testing.T) {
	tc, err := FromTimeCode("2:01:24:12:13", Smpte24)

	assert.Nil(t, err)
	assert.Equal(t, "02:01:24:12:13", tc.String())
}

/// <summary>
/// Creates the timecode 24fps from integers.
/// </summary>

func Test_CreateLongRunningTimeCode_24FromIntegers(t *testing.T) {
	tc, err := FromTimeDays(2, 1, 24, 12, 13, Smpte24)

	assert.Nil(t, err)

	assert.Equal(t, "02:01:24:12:13", tc.String())
}

/// <summary>
/// Creates the timecode 25fps from string.
/// </summary>

func Test_CreateLongRunningTimeCode_25FromString(t *testing.T) {
	tc, err := FromTimeCode("2:01:24:12:13", Smpte25)

	assert.Nil(t, err)
	assert.Equal(t, "02:01:24:12:13", tc.String())
}

/// <summary>
/// Creates the timecode 25fps from integers.
/// </summary>

func Test_CreateLongRunningTimeCode_25FromIntegers(t *testing.T) {
	tc, err := FromTimeDays(2, 1, 24, 12, 13, Smpte25)

	assert.Nil(t, err)

	assert.Equal(t, "02:01:24:12:13", tc.String())
}

/// <summary>
/// Creates the timecode  2997 non drop from string.
/// </summary>

func Test_CreateLongRunningTimeCode_2997NonDropFromString(t *testing.T) {
	tc, err := FromTimeCode("2:01:42:12:20", Smpte2997NonDrop)

	assert.Nil(t, err)
	assert.Equal(t, "02:01:42:12:20", tc.String())
}

/// <summary>
/// Creates the timecode 2997 non drop from integers.
/// </summary>

func Test_CreateLongRunningTimeCode_2997NonDropFromIntegers(t *testing.T) {
	tc, err := FromTimeDays(2, 1, 24, 12, 13, Smpte2997NonDrop)
	assert.Nil(t, err)
	assert.Equal(t, "02:01:24:12:13", tc.String())
}

/// <summary>
/// Creates the timecode 2997 drop from string.
/// </summary>

func Test_CreateLongRunningTimeCode_2997DropFromString(t *testing.T) {
	t.Skip("ignore this test case")
	tc, err := FromTimeCode("2:01:42:12;22", Smpte2997Drop)
	assert.Nil(t, err)
	assert.Equal(t, "02:01:42:12;22", tc.String())
}

/// <summary>
/// Creates the timecode 2997 drop from integers.
/// </summary>

func Test_CreateLongRunningTimeCode_2997DropFromIntegers(t *testing.T) {
	t.Skip("ignore this test case")
	tc, err := FromTimeDays(2, 1, 24, 12, 13, Smpte2997Drop)
	assert.Nil(t, err)
	assert.Equal(t, "02:01:24:12;13", tc.String())
}

/// <summary>
/// Creates the timecode 30 from string.
/// </summary>

func Test_CreateLongRunningTimeCode_30FromString(t *testing.T) {
	tc, err := FromTimeCode("2:01:42:12:22", Smpte30)
	assert.Nil(t, err)
	assert.Equal(t, "02:01:42:12:22", tc.String())
}

/// <summary>
/// Creates the time code_30 from integers.
/// </summary>

func Test_CreateLongRunningTimeCode_30FromIntegers(t *testing.T) {
	tc, err := FromTimeDays(2, 1, 24, 12, 13, Smpte30)
	assert.Nil(t, err)
	assert.Equal(t, "02:01:24:12:13", tc.String())
}

func Test_ValidateBadLongRunningTimecode(t *testing.T) {
	valid := validateSmpte12MTimecode("-1:01:00:00:12")

	assert.False(t, valid)
}

func Test_GetTotalDays(t *testing.T) {
	t1, _ := FromTimeCode("01:30:00:00", Smpte2997NonDrop)

	assert.Equal(t, 0.0625, t1.TotalDays())
}

func Test_GetLongRunningTotalDays(t *testing.T) {
	t1, _ := FromTimeCode("01:01:30:00:00", Smpte2997NonDrop)

	assert.Equal(t, 1.0625, t1.TotalDays())
}

func Test_GetLongRunningTotalHours(t *testing.T) {
	t1, _ := FromTimeCode("2:01:30:00:00", Smpte2997NonDrop)

	assert.Equal(t, float64(49.5), t1.TotalHours())
}

func Test_GetLongRunningTotalMinutes(t *testing.T) {
	t1, _ := FromTimeCode("2:01:30:00:00", Smpte2997NonDrop)

	assert.Equal(t, float64(2970), t1.TotalMinutes())
}

func Test_FromHours_LongRunning(t *testing.T) {
	t1, _ := FromHours(25.5, Smpte30)

	assert.Equal(t, float64(25.5), t1.TotalHours())
}

func Test_FromDays(t *testing.T) {
	t1, _ := FromDays(10.5, Smpte30)

	assert.Equal(t, float64(10.5), t1.TotalDays())
}
