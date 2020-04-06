package timecode

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const _defaultPrecision = 106

func newBF(v float64) *big.Float {
	ret := new(big.Float).SetPrec(_defaultPrecision)
	if v != 0 {
		// s := fmt.Sprintf("%g", v)
		// ret.SetString(s)
		ret.SetFloat64(v)
	}
	return ret
}

func newBFInt64(v int64) *big.Float {
	return newBF(0).SetInt64(v)
}

func newBFString(v string) *big.Float {
	ret, _ := newBF(0).SetString(v)
	return ret
}

func addBF(v1, v2 *big.Float) *big.Float {
	return newBF(0).Add(v1, v2)
}
func subBF(v1, v2 *big.Float) *big.Float {
	return newBF(0).Sub(v1, v2)
}

func mulBF(v1, v2 *big.Float) *big.Float {
	return newBF(0).Mul(v1, v2)
}
func divBF(v1, v2 *big.Float) *big.Float {
	return newBF(0).Quo(v1, v2)
}

func roundBF(val *big.Float, decimals int) *big.Float {
	delta, pow := newBF(0), newBF(0).SetInt64(int64(math.Pow10(decimals)))
	if val.Sign() >= 0 {
		delta.SetString("0.5")
	} else {
		delta.SetString("-0.5")
	}
	ret := addBF(mulBF(val, pow), delta)
	intVal, _ := ret.Int(nil)
	return ret.SetInt(intVal).Quo(ret, pow)
}

// decimal ..
type decimal struct {
	value *big.Float
}

// newDecimal create a new decimal with float64
func newDecimal(value *big.Float) *decimal {
	return &decimal{
		value: value,
	}
}

func newDecimalFloat64(value float64) *decimal {
	return &decimal{
		value: newBF(value),
	}
}

// newDecimalString create a new decimal with float64
func newDecimalString(value string) *decimal {
	return &decimal{
		value: newBFString(value),
	}
}

// newDecimalInt64 create a new decimal with float64
func newDecimalInt64(value int64) *decimal {
	return &decimal{
		value: newBFInt64(value),
	}
}

/*
func addDecimal(v1, v2 *decimal) *decimal {
	return &decimal{
		value: newBF(0).Add(v1.value, v2.value),
	}
}

func subDecimal(v1, v2 *decimal) *decimal {
	return &decimal{
		value: newBF(0).Sub(v1.value, v2.value),
	}
}

func mulDecimal(v1, v2 *decimal) *decimal {
	return &decimal{
		value: newBF(0).Mul(v1.value, v2.value),
	}
}

func divDecimal(v1, v2 *decimal) *decimal {
	return &decimal{
		value: newBF(0).Quo(v1.value, v2.value),
	}
}
*/
// addFloat add a float value
func (m *decimal) addFloat(v *big.Float) *decimal {
	return newDecimal(addBF(m.value, v))
}

// subFloat sub a float value
func (m *decimal) subFloat(v *big.Float) *decimal {
	return newDecimal(subBF(m.value, v))
}

// mulFloat mul a float value
func (m *decimal) mulFloat(v *big.Float) *decimal {
	return newDecimal(mulBF(m.value, v))
}

// divFloat div a float value
func (m *decimal) divFloat(v *big.Float) *decimal {
	return newDecimal(divBF(m.value, v))
}

// Copy ..
// func (m *decimal) Copy() *decimal {
// 	ret := newBF(0)
// 	return &decimal{
// 		value: ret.Add(ret, m.value),
// 	}
// }

// Float ..
func (m *decimal) Float() *big.Float {
	return m.value
}

// Add ..
func (m *decimal) Add(v *decimal) *decimal {
	return newDecimal(addBF(m.value, v.value))
}

// AddFloat64 ..
func (m *decimal) AddFloat64(v float64) *decimal {
	return newDecimal(addBF(m.value, newBF(v)))
}

// Sub ..
func (m *decimal) Sub(v *decimal) *decimal {
	return newDecimal(subBF(m.value, v.value))
}

// Mul multiply a decimal value
func (m *decimal) Mul(v *decimal) *decimal {
	return newDecimal(mulBF(m.value, v.value))
}

// MulFloat64 ..
func (m *decimal) MulInt64(v int64) *decimal {
	return newDecimal(mulBF(m.value, newBFInt64(v)))
}

// MulFloat64 ..
func (m *decimal) MulFloat64(v float64) *decimal {
	return newDecimal(mulBF(m.value, newBF(v)))
}

// MulFloat32 multiply a float32 value
func (m *decimal) MulFloat32(v float32) *decimal {
	return newDecimal(mulBF(m.value, newBF(float64(v))))
}

// Div ..
func (m *decimal) Div(v *decimal) *decimal {
	return newDecimal(divBF(m.value, v.value))
}

// DivFloat64 ..
func (m *decimal) DivFloat64(v float64) *decimal {
	return newDecimal(divBF(m.value, newBF(v)))
}

// DivInt64 ..
func (m *decimal) DivInt64(v int64) *decimal {
	return newDecimal(divBF(m.value, newBFInt64(v)))
}

// Floor ..
func (m *decimal) Floor() *decimal {
	intVal, _ := m.value.Int(nil)
	return newDecimal(newBF(0).SetInt(intVal))
}

func (m *decimal) Round(decimals int) *decimal {
	val, delta := newBF(0), newBF(0)
	pow := newBF(0).SetInt64(int64(math.Pow10(decimals)))
	if m.value.Sign() >= 0 {
		delta.SetString("0.5")
	} else {
		delta.SetString("-0.5")
	}
	intVal, _ := val.Mul(m.value, pow).Add(val, delta).Int(nil)
	ret := val.SetInt(intVal)
	ret = ret.Quo(ret, pow)
	// return
	return newDecimal(ret)
}

// Int64 ..
func (m *decimal) Int64() int64 {
	ret, _ := m.value.Int64()
	return ret
}

// Float32 ..
func (m *decimal) Float32() float32 {
	ret, _ := m.value.Float32()
	return ret
}

// Float64 ..
func (m *decimal) Float64() float64 {
	ret, _ := m.value.Float64()
	return ret
}

// String ..
func (m *decimal) String() string {
	return m.value.String()
}

// Debug ..
func (m *decimal) Debug() {
	s := fmt.Sprintf("value: %g; %s", m.value, m.value.String())
	fmt.Println(s)
}

// SmpteFrameRate enum type
type SmpteFrameRate int

const (
	// Smpte2398 SMPTE 23.98 frame rate. Also known as Film Sync.
	Smpte2398 SmpteFrameRate = 0

	// Smpte24 SMPTE 24 fps frame rate.
	Smpte24 SmpteFrameRate = 1

	// Smpte25 SMPTE 25 fps frame rate. Also known as PAL.
	Smpte25 SmpteFrameRate = 2

	// Smpte2997Drop SMPTE 29.97 fps Drop Frame timecode. Used in the NTSC television system.
	Smpte2997Drop SmpteFrameRate = 3

	// Smpte2997NonDrop SMPTE 29.97 fps Non Drop Fram timecode. Used in the NTSC television system.
	Smpte2997NonDrop SmpteFrameRate = 4

	// Smpte30 SMPTE 30 fps frame rate.
	Smpte30 SmpteFrameRate = 5

	// Smpte50 50 fps frame rate.
	Smpte50 SmpteFrameRate = 6

	// Smpte5994Drop 59.94 fps Drop Frame timecode. Used in the NTSC television system.
	Smpte5994Drop SmpteFrameRate = 7

	// Smpte5994NonDrop 59.94 fps Non Drop Fram timecode. Used in the NTSC television system.
	Smpte5994NonDrop SmpteFrameRate = 8

	// Smpte60 60 fps frame rate.
	Smpte60 SmpteFrameRate = 9

	// Smpte96 96 fps frame rate.
	Smpte96 SmpteFrameRate = 10

	// Smpte100 100 fps frame rate.
	Smpte100 SmpteFrameRate = 11

	// Smpte120 120 fps frame rate.
	Smpte120 SmpteFrameRate = 12

	// Unknown Value.
	Unknown SmpteFrameRate = -1
)

const (
	_smpte12MBadFormat        = "The timecode provided is not in the correct format."
	_smpte12MOutOfRange       = "The timecode provided is out of the expected range."
	_smpte12MMaxValueOverflow = "The resulting timecode %v is out of the expected range of MaxValue %v."
	_smpte12MMinValueOverflow = "The resulting timecode is out of the expected range of MinValue."
)

var (
	/// Regular expression object used for validating timecode.
	// SMPTEREGEXSTRING Regular expression string used for parsing out the timecode.
	_validateTimecode *regexp.Regexp = regexp.MustCompile("(\\d{2}):(\\d{2}):(\\d{2})(?::|;)(\\d{2})")
	_epsilon          *decimal       = newDecimalString("0.00000000000000000000001")
	_1000div1001      *decimal       = newDecimal(divBF(newBF(1000), newBF(1001)))
	_1001div1000      *decimal       = newDecimal(divBF(newBF(1001), newBF(1000)))
)

type rateRec struct {
	frames  int64
	hours   int64
	minutes int64
	rate    float64
	drop    bool
	pcrTb   float64
	max     *decimal
}

var _rateRecords = map[SmpteFrameRate]*rateRec{
	Smpte2398:        &rateRec{frames: 24, hours: 86400, minutes: 1440, rate: 23.98, drop: false, pcrTb: 3753.75, max: newDecimalString("86486.358291666700000")},
	Smpte24:          &rateRec{frames: 24, hours: 86400, minutes: 1440, rate: 24, drop: false, pcrTb: 3750, max: newDecimalString("86399.958333333300000")},
	Smpte25:          &rateRec{frames: 25, hours: 90000, minutes: 1500, rate: 25, drop: false, pcrTb: 3600, max: newDecimalString("86399.960000000000000")},
	Smpte2997Drop:    &rateRec{frames: 30, hours: 107892, minutes: 1798, rate: 29.97, drop: true, pcrTb: 3003, max: newDecimalString("86399.880233333300000")},
	Smpte2997NonDrop: &rateRec{frames: 30, hours: 108000, minutes: 1800, rate: 29.97, drop: false, pcrTb: 3003, max: newDecimalString("86486.366633333300000")},
	Smpte30:          &rateRec{frames: 30, hours: 108000, minutes: 1800, rate: 30, drop: false, pcrTb: 3000, max: newDecimalString("86399.966666666700000")},
	Smpte50:          &rateRec{frames: 50, hours: 180000, minutes: 3000, rate: 50, drop: false, max: newDecimalString("86399.980000000000000")},
	Smpte5994Drop:    &rateRec{frames: 60, hours: 215784, minutes: 3596, rate: 59.94, drop: true, max: newDecimalString("86399.896916666700000")},
	Smpte5994NonDrop: &rateRec{frames: 60, hours: 216000, minutes: 3600, rate: 59.94, drop: false, max: newDecimalString("86486.383316670000000")},
	Smpte60:          &rateRec{frames: 60, hours: 216000, minutes: 3600, rate: 60, drop: false, max: newDecimalString("86399.983333333300000")},
	Smpte96:          &rateRec{frames: 96, hours: 345600, minutes: 5760, rate: 96, drop: false, max: newDecimalString("86399.983333333300000")},
	Smpte100:         &rateRec{frames: 100, hours: 360000, minutes: 6000, rate: 100, drop: false, max: newDecimalString("86399.983333333300000")},
	Smpte120:         &rateRec{frames: 120, hours: 432000, minutes: 7200, rate: 120, drop: false, max: newDecimalString("86399.983333333300000")},
}

// TimeCode ...
type TimeCode struct {
	/// The private Timespan used to track absolute time for this instance.
	absoluteTime *decimal

	/// The frame rate for this instance.
	frameRate SmpteFrameRate
}

func fromTimeDecimal(time *decimal, rate SmpteFrameRate) (*TimeCode, error) {
	return &TimeCode{
		frameRate:    rate,
		absoluteTime: time,
	}, nil
}

//FromTimeHours  Initializes a new instance of the TimeCode struct to a specified number of hours, minutes, and seconds.
func FromTimeHours(hours, minutes, seconds, frames int, rate SmpteFrameRate) (*TimeCode, error) {
	timeCode := fmt.Sprintf("%02d:%02d:%02d:%02d", hours, minutes, seconds, frames)
	time, err := smpte12mToAbsoluteTime(timeCode, rate)
	if err != nil {
		return nil, err
	}
	return fromTimeDecimal(time, rate)
}

// FromTimeDays ..
/// <summary>
///  Initializes a new instance of the TimeCode struct to a specified number of hours, minutes, and seconds.
/// </summary>
/// <param name="days">Number of days.</param>
/// <param name="hours">Number of hours.</param>
/// <param name="minutes">Number of minutes.</param>
/// <param name="seconds">Number of seconds.</param>
/// <param name="frames">Number of frames.</param>
/// <param name="rate">The SMPTE frame rate.</param>
/// <exception cref="System.FormatException">
/// The parameters specify a TimeCode value less than TimeCode.MinValue
/// or greater than TimeCode.MaxValue, or the values of time code components are not valid for the SMPTE framerate.
/// </exception>
/// <code source="..\Documentation\SdkDocSamples\TimecodeSamples.cs" region="CreateTimeCode_2398FromIntegers" lang="CSharp" title="Create TimeCode from Integers"/>
func FromTimeDays(days, hours, minutes, seconds, frames int, rate SmpteFrameRate) (*TimeCode, error) {
	timeCode := fmt.Sprintf("%02d:%02d:%02d:%02d:%02d", days, hours, minutes, seconds, frames)
	time, err := smpte12mToAbsoluteTime(timeCode, rate)
	if err != nil {
		return nil, err
	}
	return fromTimeDecimal(time, rate)
}

/*
   /// <summary>
   /// Initializes a new instance of the TimeCode struct using an Int32 in hex format containing the time code value compatible with the Windows Media Format SDK.
   /// Time code is stored so that the hexadecimal value is read as if it were a decimal value. That is, the time code value 0x01133512 does not represent decimal 18035986, rather it specifies 1 hour, 13 minutes, 35 seconds, and 12 frames.
   /// </summary>
   /// <param name="windowsMediaTimeCode">The integer value of the timecode.</param>
   /// <param name="rate">The SMPTE frame rate.</param>
   public TimeCode(int windowsMediaTimeCode, SmpteFrameRate rate)
   {
       // Timecode is provided back formatted as hexadecimal bytes read in single bytes from left to right.
       byte[] timeCodeBytes = BitConverter.GetBytes(windowsMediaTimeCode);
       string timeCode = string.Format("{3:x2}:{2:x2}:{1:x2}:{0:x2}", timeCodeBytes[0], timeCodeBytes[1], timeCodeBytes[2], timeCodeBytes[3]);

       this.frameRate = rate;
       this.absoluteTime = smpte12mToAbsoluteTime(timeCode, this.frameRate);
   }

   /// <summary>
   /// Initializes a new instance of the TimeCode struct using the TotalSeconds in the supplied TimeSpan.
   /// </summary>
   /// <param name="timeSpan">The <see cref="TimeSpan"/> to be used for the new timecode.</param>
   /// <param name="rate">The SMPTE frame rate.</param>
   public TimeCode(TimeSpan timeSpan, SmpteFrameRate rate)
   {
       this.frameRate = rate;
       this.absoluteTime = TimeCode.FromTimeSpan(timeSpan, rate).absoluteTime;
   }
*/

// FromTimeCodeRate ..
/// <summary>
/// Initializes a new instance of the TimeCode struct using a time code string that contains the framerate at the end of the string.
/// </summary>
/// <remarks>
/// Pass in a timecode in the format "timecode@framrate".
/// Supported rates include @23.98, @24, @25, @29.97, @30
/// </remarks>
/// <example>
/// "00:01:00:00@29.97" is equivalent to 29.97 non drop frame.
/// "00:01:00;00@29.97" is equivalent to 29.97 drop frame.
/// </example>
/// <param name="timeCodeAndRate">The SMPTE 12m time code string.</param>
func FromTimeCodeRate(timeCodeAndRate string) (*TimeCode, error) {
	timeAndRate := strings.Split(timeCodeAndRate, "@")
	timeCode, frameRate := "", 29.07

	if len(timeAndRate) == 1 {
		timeCode = timeAndRate[0]
	} else if len(timeAndRate) == 2 {
		timeCode = timeAndRate[0]
		frameRate, _ = strconv.ParseFloat(timeAndRate[1], 10)
	}
	drop := strings.Index(timeCode, ";") >= 0
	// try to find match record
	for rate, rec := range _rateRecords {
		if rec.rate == frameRate && rec.drop == drop {
			return FromTimeCode(timeCode, rate)
		}
	}
	// now return error
	return nil, fmt.Errorf("timecode: invalid timecode with rate: '%s'", timeCodeAndRate)
}

//FromTimeCode Initializes a new instance of the TimeCode struct using a time code string and a SMPTE framerate.
func FromTimeCode(timeCode string, rate SmpteFrameRate) (*TimeCode, error) {
	time, err := smpte12mToAbsoluteTime(timeCode, rate)
	if err != nil {
		return nil, err
	}
	return fromTimeDecimal(time, rate)
}

// FromTime Initializes a new instance of the TimeCode struct using an absolute time value, and the SMPTE framerate.
func FromTime(absoluteTime float64, rate SmpteFrameRate) (*TimeCode, error) {
	return fromTimeDecimal(newDecimalFloat64(absoluteTime), rate)
}

// TicksPerDay ..
/// <summary>
///  Gets the number of ticks in 1 day.
///  This field is constant.
/// </summary>
/// <value>The number of ticks in 1 day.</value>
const TicksPerDay int64 = 864000000000

// TicksPerDayAbsoluteTime ..
/// <summary>
///  Gets the number of absolute time ticks in 1 day.
///  This field is constant.
/// </summary>
/// <value>The number of absolute time ticks in 1 day.</value>
const TicksPerDayAbsoluteTime int64 = 86400

// TicksPerHour ..
/// <summary>
///  Gets the number of ticks in 1 hour. This field is constant.
/// </summary>
/// <value>The number of ticks in 1 hour.</value>
const TicksPerHour int64 = 36000000000

// TicksPerHourAbsoluteTime ..
/// <summary>
///  Gets the number of absolute time ticks in 1 hour. This field is constant.
/// </summary>
/// <value>The number of absolute time ticks in 1 hour.</value>
const TicksPerHourAbsoluteTime int64 = 3600

// TicksPerMillisecond ..
/// <summary>
/// Gets the number of ticks in 1 millisecond. This field is constant.
/// </summary>
/// <value>The number of ticks in 1 millisecond.</value>
const TicksPerMillisecond int64 = 10000

// TicksPerMillisecondAbsoluteTime ..
/// <summary>
/// Gets the number of ticks in 1 millisecond. This field is constant.
/// </summary>
/// <value>The number of ticks in 1 millisecond.</value>
const TicksPerMillisecondAbsoluteTime float64 = 0.0010000

// TicksPerMinute ..
/// <summary>
/// Gets the number of ticks in 1 minute. This field is constant.
/// </summary>
/// <value>The number of ticks in 1 minute.</value>
const TicksPerMinute int64 = 600000000

// TicksPerMinuteAbsoluteTime ..
/// <summary>
/// Gets the number of absolute time ticks in 1 minute. This field is constant.
/// </summary>
/// <value>The number of absolute time ticks in 1 minute.</value>
const TicksPerMinuteAbsoluteTime float64 = 60

// TicksPerSecond ..
/// <summary>
/// Gets the number of ticks in 1 second.
/// </summary>
/// <value>The number of ticks in 1 second.</value>
const TicksPerSecond int64 = 10000000

// TicksPerSecondAbsoluteTime ..
/// <summary>
/// Gets the number of ticks in 1 second.
/// </summary>
/// <value>The number of ticks in 1 second.</value>
const TicksPerSecondAbsoluteTime float64 = 1.0000000

// MinValue ..
/// <summary>
/// Gets the minimum TimeCode value. This field is read-only.
/// </summary>
/// <value>The minimum TimeCode value.</value>
const MinValue float64 = 0

// Duration ..
/// <summary>
/// Gets the absolute time in seconds of the current TimeCode object.
/// </summary>
/// <returns>
///  A double that is the absolute time in seconds duration of the current TimeCode object.
/// </returns>
/// <value>The absolute time in seconds of the current TimeCode.</value>
func (m *TimeCode) Duration() float64 {
	return m.absoluteTime.Float64()
}

// FrameRate ..
/// <summary>
/// Gets or sets the current SMPTE framerate for this TimeCode instance.
/// </summary>
/// <value>The frame rate of the TimeCode.</value>
func (m *TimeCode) FrameRate() SmpteFrameRate {
	return m.frameRate
}

// DaysSegment ..
/// <summary>
///  Gets the number of whole days represented by the current TimeCode
///  structure.
/// </summary>
/// <returns>
///  The hour component of the current TimeCode structure.
/// </returns>
/// <value>The number of whole days of the TimeCode.</value>
func (m *TimeCode) DaysSegment() int64 {
	timeCode := absoluteTimeToSmpte12M(m.absoluteTime, m.frameRate)

	days := "0"

	if len(timeCode) > 11 {
		index := strings.Index(timeCode, ":")
		days = timeCode[0:index]
	}

	ret, _ := strconv.ParseInt(days, 10, 64)
	return ret
}

// HoursSegment ..
/// <summary>
///  Gets the number of whole hours represented by the current TimeCode
///  structure.
/// </summary>
/// <returns>
///  The hour component of the current TimeCode structure. The return value
///     ranges from 0 through 23.
/// </returns>
/// <value>The number of whole hours of the TimeCode.</value>
func (m *TimeCode) HoursSegment() int64 {
	timeCode := absoluteTimeToSmpte12M(m.absoluteTime, m.frameRate)

	if len(timeCode) > 11 {
		index := strings.Index(timeCode, ":") + 1
		timeCode = timeCode[index:]
	}

	hours := timeCode[0:2]

	ret, _ := strconv.ParseInt(hours, 10, 64)
	return ret
}

// MinutesSegment ..
/// <summary>
/// Gets the number of whole minutes represented by the current TimeCode structure.
/// </summary>
/// <returns>
/// The minute component of the current TimeCode structure. The return
/// value ranges from 0 through 59.
/// </returns>
/// <value>The number of whole minutes of the current TimeCode.</value>
func (m *TimeCode) MinutesSegment() int64 {
	timeCode := absoluteTimeToSmpte12M(m.absoluteTime, m.frameRate)

	if len(timeCode) > 11 {
		index := strings.Index(timeCode, ":") + 1
		timeCode = timeCode[index:]
	}

	minutes := timeCode[3:5]

	ret, _ := strconv.ParseInt(minutes, 10, 64)
	return ret
}

// SecondsSegment ..
/// <summary>
/// Gets the number of whole seconds represented by the current TimeCode structure.
/// </summary>
/// <returns>
///  The second component of the current TimeCode structure. The return
///    value ranges from 0 through 59.
/// </returns>
/// <value>The number of whole seconds of the current TimeCode.</value>
func (m *TimeCode) SecondsSegment() int64 {
	timeCode := absoluteTimeToSmpte12M(m.absoluteTime, m.frameRate)

	if len(timeCode) > 11 {
		index := strings.Index(timeCode, ":") + 1
		timeCode = timeCode[index:]
	}

	seconds := timeCode[6:8]

	ret, _ := strconv.ParseInt(seconds, 10, 64)
	return ret
}

// FramesSegment ..
/// <summary>
/// Gets the number of whole frames represented by the current TimeCode
///     structure.
/// </summary>
/// <returns>
/// The frame component of the current TimeCode structure. The return
///     value depends on the framerate selected for this instance. All frame counts start at zero.
/// </returns>
/// <value>The number of whole frames of the TimeCode.</value>
func (m *TimeCode) FramesSegment() int64 {
	timeCode := absoluteTimeToSmpte12M(m.absoluteTime, m.frameRate)

	if len(timeCode) > 11 {
		index := strings.Index(timeCode, ":") + 1
		timeCode = timeCode[index:]
	}

	frames := timeCode[9:11]

	ret, _ := strconv.ParseInt(frames, 10, 64)
	return ret
}

// TotalDays ..
/// <summary>
/// Gets the value of the current TimeCode structure expressed in whole
///     and fractional days.
/// </summary>
/// <returns>
///  The total number of days represented by this instance.
/// </returns>
/// <value>The number of days of the TimeCode.</value>
func (m *TimeCode) TotalDays() float64 {
	framecount := absoluteTimeToFrames(m.absoluteTime, m.frameRate)
	return (float64(framecount) / 108000) / 24
}

// TotalHours ...
/// <summary>
/// Gets the value of the current TimeCode structure expressed in whole
///     and fractional hours.
/// </summary>
/// <returns>
///  The total number of hours represented by this instance.
/// </returns>
/// <value>The number of hours of the TimeCode.</value>
func (m *TimeCode) TotalHours() float64 {
	framecount := absoluteTimeToFrames(m.absoluteTime, m.frameRate)
	rec := _rateRecords[m.FrameRate()]

	return float64(framecount) / float64(rec.hours)
}

// TotalMinutes ..
/// <summary>
/// Gets the value of the current TimeCode structure expressed in whole
/// and fractional minutes.
/// </summary>
/// <returns>
///  The total number of minutes represented by this instance.
/// </returns>
/// <value>The number of minutes of the TimeCode.</value>
func (m *TimeCode) TotalMinutes() float64 {
	framecount := absoluteTimeToFrames(m.absoluteTime, m.frameRate)
	rec := _rateRecords[m.FrameRate()]

	return float64(framecount) / float64(rec.minutes)
}

//TotalSeconds Gets the value of the current TimeCode structure expressed in whole
/// and fractional seconds. Not as Precise as the TotalSecondsPrecision.
func (m *TimeCode) TotalSeconds() float64 {
	return m.absoluteTime.Round(7).Float64()
}

//TotalSecondsPrecision Gets the value of the current TimeCode structure expressed in whole
/// and fractional seconds. This is returned as a <see cref="decimal"/> for greater precision.
func (m *TimeCode) TotalSecondsPrecision() float64 {
	return m.absoluteTime.Float64()
}

//TotalFrames Gets the value of the current TimeCode structure expressed in frames.
func (m *TimeCode) TotalFrames() int64 {
	return absoluteTimeToFrames(m.absoluteTime, m.frameRate)
}

//MaxValue Gets the maximum TimeCode value of a known frame rate. The Max value for Timecode.
func maxValue(frameRate SmpteFrameRate) *decimal {
	return _rateRecords[frameRate].max
}

//Sub Subtracts a specified TimeCode from another specified TimeCode.
func Sub(t1, t2 *TimeCode) (*TimeCode, error) {
	t3 := &TimeCode{
		frameRate:    t1.frameRate,
		absoluteTime: t1.absoluteTime.Sub(t2.absoluteTime),
	}

	if t3.TotalSeconds() < MinValue {
		return nil, errors.New(_smpte12MMinValueOverflow)
	}

	return t3, nil
}

// NotEqual Indicates whether two TimeCode instances are not equal.
func NotEqual(t1, t2 *TimeCode) bool {
	var timeCode1, _ = fromTimeDecimal(t1.absoluteTime, Smpte30)
	var timeCode2, _ = fromTimeDecimal(t2.absoluteTime, Smpte30)

	if timeCode1.TotalSeconds() != timeCode2.TotalSeconds() {
		return true
	}

	return false
}

// Add two specified TimeCode instances.
func Add(t1, t2 *TimeCode) (*TimeCode, error) {
	t3 := &TimeCode{
		frameRate:    t1.frameRate,
		absoluteTime: t1.absoluteTime.Add(t2.absoluteTime),
	}
	// check overflow
	maxValue := maxValue(t1.frameRate).Float64()
	if t3.TotalSecondsPrecision() > maxValue && t3.TotalSeconds() > maxValue {
		return nil, fmt.Errorf(_smpte12MMaxValueOverflow, t3.TotalSecondsPrecision(), maxValue)
	}
	// return
	return t3, nil
}

// LessThan ..
///  Indicates whether a specified TimeCode is less than another
///  specified TimeCode.
func LessThan(t1, t2 *TimeCode) bool {
	var timeCode1, _ = fromTimeDecimal(t1.absoluteTime, Smpte30)
	var timeCode2, _ = fromTimeDecimal(t2.absoluteTime, Smpte30)

	if timeCode1.TotalSeconds() < timeCode2.TotalSeconds() {
		return true
	}

	return false
}

// LessEqual ..
/// <summary>
///  Indicates whether a specified TimeCode is less than or equal to another
///  specified TimeCode.
/// </summary>
/// <param name="t1">The first TimeCode.</param>
/// <param name="t2">The second TimeCode.</param>
/// <returns>true if the value of t1 is less than or equal to the value of t2; otherwise, false.</returns>
func LessEqual(t1, t2 *TimeCode) bool {
	var timeCode1, _ = fromTimeDecimal(t1.absoluteTime, Smpte30)
	var timeCode2, _ = fromTimeDecimal(t2.absoluteTime, Smpte30)

	if timeCode1.TotalSeconds() < timeCode2.TotalSeconds() || (timeCode1.TotalSeconds() == timeCode2.TotalSeconds()) {
		return true
	}

	return false
}

//Equal  Indicates whether two TimeCode instances are equal.
func Equal(t1, t2 *TimeCode) bool {
	var timeCode1, _ = fromTimeDecimal(t1.absoluteTime, Smpte30)
	var timeCode2, _ = fromTimeDecimal(t2.absoluteTime, Smpte30)

	if timeCode1.TotalSeconds() == timeCode2.TotalSeconds() {
		return true
	}

	return false
}

// GreatThan Indicates whether a specified TimeCode is greater than another specified
func GreatThan(t1, t2 *TimeCode) bool {
	var timeCode1, _ = fromTimeDecimal(t1.absoluteTime, Smpte30)
	var timeCode2, _ = fromTimeDecimal(t2.absoluteTime, Smpte30)

	if timeCode1.TotalSeconds() > timeCode2.TotalSeconds() {
		return true
	}

	return false
}

// GreatEqual Indicates whether a specified TimeCode is greater than or equal to
///     another specified TimeCode.
func GreatEqual(t1, t2 *TimeCode) bool {

	var timeCode1, _ = fromTimeDecimal(t1.absoluteTime, Smpte30)
	var timeCode2, _ = fromTimeDecimal(t2.absoluteTime, Smpte30)

	if timeCode1.TotalSeconds() > timeCode2.TotalSeconds() || (timeCode1.TotalSeconds() == timeCode2.TotalSeconds()) {
		return true
	}

	return false
}

// ticks27MhzToSmpte12M Returns a SMPTE 12M formatted time code string from a 27Mhz ticks value.
func ticks27MhzToSmpte12M(ticks27Mhz int64, rate SmpteFrameRate) string {
	var days, hours, minutes, seconds, frames int64
	rec := _rateRecords[rate]
	pcrTb := float64(ticks27MhzToPcrTb(ticks27Mhz))
	framecount := int64(pcrTb / rec.pcrTb)
	dropFrame := false
	// calcuate
	days = (framecount / rec.hours) / 24
	hours = (framecount / rec.hours) % 24
	framecountHours := rec.hours * (hours + days*24)
	switch rate {
	case Smpte2997Drop:
		minutes = ((framecount + int64(2*int32((framecount-framecountHours)/1800)) - int64(2*int32((framecount-framecountHours)/18000)) - (framecountHours)) / 1800) % 60
		seconds = ((framecount - int64(rec.minutes*minutes) - int64(2*int32(minutes/10)) - int64(framecountHours)) / rec.frames) % 3600
		frames = (framecount - int64(rec.frames*seconds) - (rec.minutes * minutes) - int64(2*int32(minutes/10)) - (framecountHours)) % rec.frames
		dropFrame = true
		break
	case Smpte5994Drop:
		minutes = ((framecount + int64(4*int32((framecount-framecountHours)/3600)) - int64(4*int32((framecount-framecountHours)/36000)) - (framecountHours)) / 3600) % 60
		seconds = ((framecount - int64(rec.minutes*minutes) - int64(4*int32(minutes/10)) - int64(framecountHours)) / rec.frames) % 3600
		frames = (framecount - int64(rec.frames*seconds) - (rec.minutes * minutes) - int64(4*int32(minutes/10)) - (framecountHours)) % rec.frames
		dropFrame = true
		break
	default:
		minutes = ((framecount - framecountHours) / rec.minutes) % 60
		seconds = ((framecount - (rec.minutes * minutes) - (framecountHours)) / rec.frames) % 3600
		frames = (framecount - (rec.frames * seconds) - (rec.minutes * minutes) - (framecountHours)) % rec.frames
		break
	}
	return formatTimeCodeString(int32(days), int32(hours), int32(minutes), int32(seconds), int32(frames), dropFrame)
}

/*
   /// <summary>
   /// Compares two TimeCode values and returns an integer that indicates their relationship.
   /// </summary>
   /// <param name="t1">The first TimeCode.</param>
   /// <param name="t2">The second TimeCode.</param>
   /// <returns>
   /// Value Condition -1 t1 is less than t2, 0 t1 is equal to t2, 1 t1 is greater than t2.
   /// </returns>
   public static int Compare(TimeCode t1, TimeCode t2)
   {
       if (t1 < t2)
       {
           return -1;
       }

       if (t1 == t2)
       {
           return 0;
       }

       return 1;
   }

   /// <summary>
   ///  Returns a value indicating whether two specified instances of TimeCode
   ///  are equal.
   /// </summary>
   /// <param name="t1">The first TimeCode.</param>
   /// <param name="t2">The second TimeCode.</param>
   /// <returns>true if the values of t1 and t2 are equal; otherwise, false.</returns>
   public static bool Equals(TimeCode t1, TimeCode t2)
   {
       if (t1 == t2)
       {
           return true;
       }

       return false;
   }
*/

// FromDays ..
/// <summary>
///  Returns a TimeCode that represents a specified number of days, where
///  the specification is accurate to the nearest millisecond.
/// </summary>
/// <param name="value">A number of days accurate to the nearest millisecond.</param>
/// <param name="rate">The desired framerate for this instance.</param>
/// <returns> A TimeCode that represents value.</returns>
/// <exception cref="System.OverflowException">
/// value is less than TimeCode.MinValue or greater than TimeCode.MaxValue.
/// -or-value is System.Double.PositiveInfinity.-or-value is System.Double.NegativeInfinity.
/// </exception>
/// <exception cref="System.FormatException">
/// value is equal to System.Double.NaN.
/// </exception>
func FromDays(days float64, rate SmpteFrameRate) (*TimeCode, error) {
	time := newDecimalFloat64(days).MulInt64(TicksPerDayAbsoluteTime)
	return fromTimeDecimal(time, rate)
}

// FromHours ..
/// <summary>
///  Returns a TimeCode that represents a specified number of hours, where
///  the specification is accurate to the nearest millisecond.
/// </summary>
/// <param name="value">A number of hours accurate to the nearest millisecond.</param>
/// <param name="rate">The desired framerate for this instance.</param>
/// <returns> A TimeCode that represents value.</returns>
/// <exception cref="System.OverflowException">
/// value is less than TimeCode.MinValue or greater than TimeCode.MaxValue.
/// -or-value is System.Double.PositiveInfinity.-or-value is System.Double.NegativeInfinity.
/// </exception>
/// <exception cref="System.FormatException">
/// value is equal to System.Double.NaN.
/// </exception>
func FromHours(hours float64, rate SmpteFrameRate) (*TimeCode, error) {
	absoluteTime := hours * float64(TicksPerHourAbsoluteTime)
	return FromTime(absoluteTime, rate)
}

// FromMinutes ..
/// <summary>
///   Returns a TimeCode that represents a specified number of minutes,
///   where the specification is accurate to the nearest millisecond.
/// </summary>
/// <param name="value">A number of minutes, accurate to the nearest millisecond.</param>
/// <param name="rate">The <see cref="SmpteFrameRate"/> to use for the calculation.</param>
/// <returns>A TimeCode that represents value.</returns>
/// <exception cref="System.OverflowException">
/// value is less than TimeCode.MinValue or greater than TimeCode.MaxValue.-or-value
/// is System.Double.PositiveInfinity.-or-value is System.Double.NegativeInfinity.
/// </exception>
/// <exception cref="System.ArgumentException">
/// value is equal to System.Double.NaN.
/// </exception>
func FromMinutes(minutes float64, rate SmpteFrameRate) (*TimeCode, error) {
	absoluteTime := minutes * float64(TicksPerMinuteAbsoluteTime)
	return FromTime(absoluteTime, rate)
}

// FromSeconds ..
/// <summary>
/// Returns a TimeCode that represents a specified number of seconds,
/// where the specification is accurate to the nearest millisecond.
/// </summary>
/// <param name="value">A number of seconds, accurate to the nearest millisecond.</param>
/// /// <param name="rate">The framerate of the Timecode.</param>
/// <returns>A TimeCode that represents value.</returns>
/// <exception cref="System.OverflowException">
/// value is less than TimeCode.MinValue or greater than TimeCode.MaxValue.-or-value
///  is System.Double.PositiveInfinity.-or-value is System.Double.NegativeInfinity.
/// </exception>
/// <exception cref="System.ArgumentException">
/// value is equal to System.Double.NaN.
/// </exception>
func FromSeconds(seconds float64, rate SmpteFrameRate) (*TimeCode, error) {
	return FromTime(seconds, rate)
}

/// <summary>
/// Returns a TimeCode that represents a specified number of seconds,
/// where the specification is accurate to the nearest millisecond.
/// </summary>
/// <param name="value">A number of seconds in <see cref="decimal"/> precision, accurate to the nearest millisecond.</param>
/// /// <param name="rate">The framerate of the Timecode.</param>
/// <returns>A TimeCode that represents value.</returns>
/// <exception cref="System.OverflowException">
/// value is less than TimeCode.MinValue or greater than TimeCode.MaxValue.-or-value
///  is System.Double.PositiveInfinity.-or-value is System.Double.NegativeInfinity.
/// </exception>
/// <exception cref="System.ArgumentException">
/// value is equal to System.Double.NaN.
/// </exception>
//    public static TimeCode FromSeconds(decimal value, SmpteFrameRate rate)
//    {
//        return new TimeCode(value, rate);
//    }

// FromFrames ..
/// <summary>
/// Returns a TimeCode that represents a specified number of frames.
/// </summary>
/// <param name="value">A number of frames.</param>
/// <param name="rate">The framerate of the Timecode.</param>
/// <returns>A TimeCode that represents value.</returns>
/// <exception cref="System.OverflowException">
///  value is less than TimeCode.MinValue or greater than TimeCode.MaxValue.-or-value
///    is System.Double.PositiveInfinity.-or-value is System.Double.NegativeInfinity.
/// </exception>
/// <exception cref="System.ArgumentException">
/// value is equal to System.Double.NaN.
/// </exception>
func FromFrames(frames int64, rate SmpteFrameRate) (*TimeCode, error) {
	time := framesToAbsoluteTime(frames, rate)
	return fromTimeDecimal(time, rate)
}

// FromTicks27Mhz ..
/// <summary>
/// Returns a TimeCode that represents a specified time, where the specification
///  is in units of ticks.
/// </summary>
/// <param name="ticks"> A number of ticks that represent a time.</param>
/// <param name="rate">The Smpte framerate.</param>
/// <returns>A TimeCode with a value of value.</returns>
//    public static TimeCode FromTicks(long ticks, SmpteFrameRate rate)
//    {
//        double absoluteTime = Math.Pow(10, -7) * ticks;
//        return new TimeCode(absoluteTime, rate);
//    }
/// <summary>
/// Returns a TimeCode that represents a specified time, where the specification is
/// in units of 27 Mhz clock ticks.
/// </summary>
/// <param name="value">A number of ticks in 27 Mhz clock format.</param>
/// <param name="rate">A Smpte framerate.</param>
/// <returns>A TimeCode.</returns>
func FromTicks27Mhz(ticks27Mhz int64, rate SmpteFrameRate) (*TimeCode, error) {
	absoluteTime := ticks27MhzToAbsoluteTime(ticks27Mhz)
	return fromTimeDecimal(absoluteTime, rate)
}

/*
   /// <summary>
   /// Returns a TimeCode that represents a specified time, where the specification is
   /// in units of absolute time.
   /// </summary>
   /// <param name="value">The absolute time in 100 nanosecond units.</param>
   /// <param name="rate">The SMPTE framerate.</param>
   /// <returns>A TimeCode.</returns>
   public static TimeCode FromAbsoluteTime(double value, SmpteFrameRate rate)
   {
       return new TimeCode(value, rate);
   }
*/

// FromTimeSpan ..
/// <summary>
/// Returns a TimeCode that represents a specified time, where the specification is
/// in units of absolute time.
/// </summary>
/// <param name="value">The <see cref="TimeSpan"/> object.</param>
/// <param name="rate">The SMPTE framerate.</param>
/// <returns>A TimeCode.</returns>
func FromTimeSpan(span time.Duration, rate SmpteFrameRate) (*TimeCode, error) {
	return FromTime(span.Seconds(), rate)
}

// validateSmpte12MTimecode ..
/// <summary>
/// Validates that the string provided is in the correct format for SMPTE 12M time code.
/// </summary>
/// <param name="timeCode">String that is the time code.</param>
/// <returns>True if this is a valid SMPTE 12M time code string.</returns>
func validateSmpte12MTimecode(timeCode string) bool {
	if !_validateTimecode.Match([]byte(timeCode)) {
		return false
	}

	index, times := -1, regexp.MustCompile("[\\:\\;]+").Split(timeCode, -1)

	days := 0

	if len(times) > 4 {
		index++
		days, _ = strconv.Atoi(times[index])
	}

	index++
	hours, _ := strconv.Atoi(times[index])
	index++
	minutes, _ := strconv.Atoi(times[index])
	index++
	seconds, _ := strconv.Atoi(times[index])
	index++
	frames, _ := strconv.Atoi(times[index])

	if (days < 0) || (hours >= 24) || (minutes >= 60) || (seconds >= 60) || (frames >= 30) {
		return false
	}

	return true
}

/*
   /// <summary>
   /// Validates that the hexadecimal formatted integer provided is in the correct format for SMPTE 12M time code
   /// Time code is stored so that the hexadecimal value is read as if it were an integer value.
   /// That is, the time code value 0x01133512 does not represent integer 18035986, rather it specifies 1 hour, 13 minutes, 35 seconds, and 12 frames.
   /// </summary>
   /// <param name="windowsMediaTimeCode">Integer that is the time code stored in hexadecimal format.</param>
   /// <returns>True if this is a valid SMPTE 12M time code string.</returns>
   public static bool validateSmpte12MTimecode(int windowsMediaTimeCode)
   {
       byte[] timeCodeBytes = BitConverter.GetBytes(windowsMediaTimeCode);
       string timeCode = string.Format("{3:x2}:{2:x2}:{1:x2}:{0:x2}", timeCodeBytes[0], timeCodeBytes[1], timeCodeBytes[2], timeCodeBytes[3]);
       string[] times = timeCode.Split(':', ';');

       hours := ret, _ := strconv.Atoi(times[0]);
       minutes := ret, _ := strconv.Atoi(times[1]);
       seconds := ret, _ := strconv.Atoi(times[2]);
       frames := ret, _ := strconv.Atoi(times[3]);

       if ((hours >= 24) || (minutes >= 60) || (seconds >= 60) || (frames >= 30))
       {
           return false;
       }

       return true;
   }
*/
// smpte12MToTicks27Mhz Returns the value of the provided time code string and framerate in 27Mhz ticks.
func smpte12MToTicks27Mhz(timeCode string, rate SmpteFrameRate) int64 {
	t, _ := FromTimeCode(timeCode, rate)
	hours, minutes, seconds, frames := t.HoursSegment(), t.MinutesSegment(), t.SecondsSegment(), t.FramesSegment()
	days, hours, minutes, seconds, frames := t.DaysSegment(), t.HoursSegment(), t.MinutesSegment(), t.SecondsSegment(), t.FramesSegment()
	rec := _rateRecords[rate]
	switch rate {
	case Smpte2997Drop:
		ret := (3003 * frames) + (90090 * seconds) + (5399394 * minutes) + (6006 * int64(float64(seconds)/10)) + (323999676 * hours) + (7775992224 * days)
		return int64(ret * 300)
	case Smpte5994Drop:
		ret := (6006 * frames) + (90090 * seconds) + (5399394 * minutes) + (12012 * int64(float64(seconds)/10)) + (323999676 * hours) + (7775992224 * days)
		return ret * 300
	default:
		valFrames := rec.pcrTb
		valSeconds := rec.pcrTb * float64(rec.frames)
		valMinutes := valSeconds * 60
		valHours := valMinutes * 60
		valDays := valHours * 24
		ret := (valFrames * float64(frames)) + (valSeconds * float64(seconds)) + (valMinutes * float64(minutes)) + (valHours * float64(hours)) + (valDays * float64(days))
		return int64(ret * float64(300))
	}
}

// ParseFramerate ..
/// <summary>
/// Parses a framerate value as double and converts it to a member of the SmpteFrameRate enumeration.
/// </summary>
/// <param name="rate">Double value of the framerate.</param>
/// <returns>A SmpteFrameRate enumeration value that matches the incoming rates.</returns>
func ParseFramerate(rate float64) SmpteFrameRate {
	rateRounded := int64(math.Floor(rate))
	for key, rec := range _rateRecords {
		if rec.frames == rateRounded {
			return key
		}
	}
	return Unknown
}

// Add ..
/// <summary>
/// Adds the specified TimeCode to this instance.
/// </summary>
/// <param name="ts">A TimeCode.</param>
/// <returns>A TimeCode that represents the value of this instance plus the value of ts.
/// </returns>
/// <exception cref="System.OverflowException">
/// The resulting TimeCode is less than TimeCode.MinValue or greater than TimeCode.MaxValue.
/// </exception>
func (m *TimeCode) Add(tc *TimeCode) error {
	m.absoluteTime.value = addBF(m.absoluteTime.value, tc.absoluteTime.value)
	return nil
}

// Sub substracts timecode
func (m *TimeCode) Sub(tc *TimeCode) error {
	m.absoluteTime.value = subBF(m.absoluteTime.value, tc.absoluteTime.value)
	return nil
}

// AddSeconds ..
func (m *TimeCode) AddSeconds(seconds float64) error {
	tc, err := FromSeconds(seconds, m.frameRate)
	if err != nil {
		return err
	}
	return m.Add(tc)
}

// SubSeconds ..
func (m *TimeCode) SubSeconds(seconds float64) error {
	tc, err := FromSeconds(seconds, m.frameRate)
	if err != nil {
		return err
	}
	return m.Sub(tc)
}

// AddFrames ..
func (m *TimeCode) AddFrames(frames int64) error {
	tc, err := FromFrames(frames, m.frameRate)
	if err != nil {
		return err
	}
	return m.Add(tc)
}

// SubFrames ..
func (m *TimeCode) SubFrames(frames int64) error {
	tc, err := FromFrames(frames, m.frameRate)
	if err != nil {
		return err
	}
	return m.Sub(tc)
}

// AddTimeCode ..
func (m *TimeCode) AddTimeCode(timecode string) error {
	tc, err := FromTimeCode(timecode, m.frameRate)
	if err != nil {
		return err
	}
	return m.Add(tc)
}

// SubTimeCode ..
func (m *TimeCode) SubTimeCode(timecode string) error {
	tc, err := FromTimeCode(timecode, m.frameRate)
	if err != nil {
		return err
	}
	return m.Sub(tc)
}

/*
   /// <summary>
   ///  Compares this instance to a specified object and returns an indication of
   ///   their relative values.
   /// </summary>
   /// <param name="value">An object to compare, or null.</param>
   /// <returns>
   ///  Value Condition -1 The value of this instance is less than the value of value.
   ///    0 The value of this instance is equal to the value of value. 1 The value
   ///    of this instance is greater than the value of value.-or- value is null.
   /// </returns>
   /// <exception cref="System.ArgumentException">
   ///  value is not a TimeCode.
   /// </exception>
   public int CompareTo(object value)
   {
       if (!(value is TimeCode))
       {
           throw new ArgumentException(_smpte12MOutOfRange);
       }

       TimeCode t1 = (TimeCode)value;

       if (this < t1)
       {
           return -1;
       }

       if (this == t1)
       {
           return 0;
       }

       return 1;
   }

   /// <summary>
   /// Compares this instance to a specified TimeCode object and returns
   /// an indication of their relative values.
   /// </summary>
   /// <param name="value"> A TimeCode object to compare to this instance.</param>
   /// <returns>
   /// A signed number indicating the relative values of this instance and value.Value
   /// Description A negative integer This instance is less than value. Zero This
   /// instance is equal to value. A positive integer This instance is greater than
   /// value.
   /// </returns>
   public int CompareTo(TimeCode value)
   {
       if (this < value)
       {
           return -1;
       }

       if (this == value)
       {
           return 0;
       }

       return 1;
   }

   /// <summary>
   ///  Returns a value indicating whether this instance is equal to a specified
   ///  object.
   /// </summary>
   /// <param name="value">An object to compare with this instance.</param>
   /// <returns>
   /// A true if value is a TimeCode object that represents the same time interval
   /// as the current TimeCode structure; otherwise, false.
   /// </returns>
   public override bool Equals(object value)
   {
       if (this == (TimeCode)value)
       {
           return true;
       }

       return false;
   }

   /// <summary>
   /// Returns a value indicating whether this instance is equal to a specified
   ///     TimeCode object.
   /// </summary>
   /// <param name="obj">An TimeCode object to compare with this instance.</param>
   /// <returns>true if obj represents the same time interval as this instance; otherwise, false.
   /// </returns>
   public bool Equals(TimeCode obj)
   {
       if (this == obj)
       {
           return true;
       }

       return false;
   }

   /// <summary>
   /// Returns a hash code for this instance.
   /// </summary>
   /// <returns> A 32-bit signed integer hash code.</returns>
   public override int GetHashCode()
   {
       return this.GetHashCode();
   }

   /// <summary>
   /// Subtracts the specified TimeCode from this instance.
   /// </summary>
   /// <param name="ts">A TimeCode.</param>
   /// <returns>A TimeCode whose value is the result of the value of this instance minus the value of ts.</returns>
   /// <exception cref="OverflowException">The return value is less than TimeCode.MinValue or greater than TimeCode.MaxValue.</exception>
   public TimeCode Subtract(TimeCode ts)
   {
       return this - ts;
   }
*/

//ToString Returns the SMPTE 12M string representation of the value of this instance.
func (m *TimeCode) String() string {
	return absoluteTimeToSmpte12M(m.absoluteTime, m.frameRate)
}

/*
   /// <summary>
   /// Outputs a string of the current time code in the requested framerate.
   /// </summary>
   /// <param name="rate">The SmpteFrameRate required for the string output.</param>
   /// <returns>SMPTE 12M formatted time code string converted to the requested framerate.</returns>
   public string ToString(SmpteFrameRate rate)
   {
       return absoluteTimeToSmpte12M(this.absoluteTime, rate);
   }

   /// <summary>
   /// Returns the value of this instance in 27 Mhz ticks.
   /// </summary>
   /// <returns>A long value that is in 27 Mhz ticks.</returns>
   public long ToTicks27Mhz()
   {
       return AbsoluteTimeToTicks27Mhz(this.absoluteTime);
   }

   /// <summary>
   /// Returns the value of this instance in MPEG 2 PCR time base (PcrTb) format.
   /// </summary>
   /// <returns>A long value that is in PcrTb.</returns>
   public long ToTicksPcrTb()
   {
       return AbsoluteTimeToTicksPcrTb(this.absoluteTime);
   }
*/

// smpte12mToAbsoluteTime Converts a SMPTE timecode to absolute time.
func smpte12mToAbsoluteTime(timeCode string, rate SmpteFrameRate) (*decimal, error) {
	days, hours, minutes, seconds, frames, err := parseTimecodeString(timeCode)
	if err != nil {
		return nil, err
	}
	// get record
	rec := _rateRecords[rate]
	if frames >= rec.frames {
		return nil, fmt.Errorf("Timecode frame value is not in the expected range for SMPTE %v", rec.frames)
	}
	// get time
	ret := frames + (rec.frames * seconds) + (rec.minutes * minutes) + (rec.hours * hours) + (rec.hours * 24 * days)
	switch rate {
	case Smpte2398, Smpte2997NonDrop, Smpte5994NonDrop, Smpte2997Drop, Smpte5994Drop:
		if rate == Smpte2997Drop {
			ret += 2 * (minutes / 10)
		} else if rate == Smpte5994Drop {
			ret += 4 * (minutes / 10)
		}
		return _1001div1000.DivInt64(rec.frames).MulInt64(ret), nil
	default:
		return newDecimalInt64(1).DivInt64(rec.frames).MulInt64(ret), nil
	}
}

// parseTimecodeString Parses a timecode string for the different parts of the timecode.
func parseTimecodeString(timeCode string) (days, hours, minutes, seconds, frames int64, err error) {
	if !_validateTimecode.Match([]byte(timeCode)) {
		err = errors.New(_smpte12MBadFormat)
		return
	}
	index, times := -1, regexp.MustCompile("[\\:\\;]+").Split(timeCode, -1)

	days = 0
	if len(times) > 4 {
		index++
		days, _ = strconv.ParseInt(times[index], 10, 64)
	}

	index++
	hours, _ = strconv.ParseInt(times[index], 10, 64)
	index++
	minutes, _ = strconv.ParseInt(times[index], 10, 64)
	index++
	seconds, _ = strconv.ParseInt(times[index], 10, 64)
	index++
	frames, _ = strconv.ParseInt(times[index], 10, 64)

	if (days < 0) || (hours >= 24) || (minutes >= 60) || (seconds >= 60) || (frames >= 30) {
		err = errors.New(_smpte12MOutOfRange)
		return
	}
	err = nil
	return
}

/*
   /// <summary>
   /// Parses a timecode string for the different parts of the timecode.
   /// </summary>
   /// <param name="timeCode">The source timecode to parse.</param>
   /// <param name="hours">The Hours section from the timecode.</param>
   /// <param name="minutes">The Minutes section from the timecode.</param>
   /// <param name="seconds">The Seconds section from the timecode.</param>
   /// <param name="frames">The frames section from the timecode.</param>
   private static void parseTimecodeString(string timeCode, out int hours, out int minutes, out int seconds, out int frames)
   {
       if (!validateTimecode.IsMatch(timeCode))
       {
           throw new FormatException(_smpte12MBadFormat);
       }

       string[] times = timeCode.Split(':', ';');

       hours = ret, _ := strconv.Atoi(times[0]);
       minutes = ret, _ := strconv.Atoi(times[1]);
       seconds = ret, _ := strconv.Atoi(times[2]);
       frames = ret, _ := strconv.Atoi(times[3]);

       if ((hours >= 24) || (minutes >= 60) || (seconds >= 60) || (frames >= 30))
       {
           throw new FormatException(_smpte12MOutOfRange);
       }
   }
*/

//formatTimeCodeString Generates a string representation of the timecode.
func formatTimeCodeString(days, hours, minutes, seconds, frames int32, dropFrame bool) string {
	framesSeparator := ":"
	if dropFrame {
		framesSeparator = ";"
	}

	if days > 0 {
		return fmt.Sprintf("%02d:%02d:%02d:%02d%s%02d", days, hours, minutes, seconds, framesSeparator, frames)
	}
	return fmt.Sprintf("%02d:%02d:%02d%s%02d", hours, minutes, seconds, framesSeparator, frames)
}

/*
   /// <summary>
   /// Generates a string representation of the timecode.
   /// </summary>
   /// <param name="hours">The Hours section from the timecode.</param>
   /// <param name="minutes">The Minutes section from the timecode.</param>
   /// <param name="seconds">The Seconds section from the timecode.</param>
   /// <param name="frames">The frames section from the timecode.</param>
   /// <param name="dropFrame">Indicates whether the timecode is drop frame or not.</param>
   /// <returns>The timecode in string format.</returns>
   private static string formatTimeCodeString(int hours, int minutes, int seconds, int frames, bool dropFrame)
   {
       string framesSeparator = ":";
       if (dropFrame)
       {
           framesSeparator = ";";
       }

       return string.Format("{0:D2}:{1:D2}:{2:D2}{4}{3:D2}", hours, minutes, seconds, frames, framesSeparator);
   }
*/

// ticks27MhzToPcrTb Converts from 27Mhz ticks to PCRTb.
func ticks27MhzToPcrTb(ticks27Mhz int64) int64 {
	return ticks27Mhz / 300
}

/*
   /// <summary>
   ///     Converts the provided absolute time to PCRTb.
   /// </summary>
   /// <param name="absoluteTime">Absolute time to be converted.</param>
   /// <returns>The number of PCRTb ticks.</returns>
   private static long AbsoluteTimeToTicksPcrTb(decimal absoluteTime)
   {
       return int64(((absoluteTime * 90000) % convert.ToDecimal(Math.Pow(2, 33))));
   }

   /// <summary>
   ///     Converts the specified absolute time to 27 mhz ticks.
   /// </summary>
   /// <param name="absoluteTime">Absolute time to be converted.</param>
   /// <returns>THe number of 27Mhz ticks.</returns>
   private static long AbsoluteTimeToTicks27Mhz(decimal absoluteTime)
   {
       return AbsoluteTimeToTicksPcrTb(absoluteTime) * 300;
   }
*/

// ticksPcrTbToAbsoluteTime ..
/// <summary>
///     Converts the specified absolute time to absolute time.
/// </summary>
/// <param name="ticksPcrTb">Ticks PCRTb to be converted.</param>
/// <returns>The absolute time.</returns>
func ticksPcrTbToAbsoluteTime(ticksPcrTb int64) *decimal {
	return newDecimalInt64(ticksPcrTb).DivInt64(90000)
}

// ticks27MhzToAbsoluteTime ..
/// <summary>
/// Converts the specified absolute time to absolute time.
/// </summary>
/// <param name="ticks27Mhz">Ticks 27Mhz to be converted.</param>
/// <returns>The absolute time.</returns>
func ticks27MhzToAbsoluteTime(ticks27Mhz int64) *decimal {
	ticksPcrTb := ticks27MhzToPcrTb(ticks27Mhz)
	return ticksPcrTbToAbsoluteTime(ticksPcrTb)
}

// absoluteTimeToSmpte12M Converts to SMPTE 12M.
func absoluteTimeToSmpte12M(absoluteTime *decimal, rate SmpteFrameRate) string {
	framecount := absoluteTimeToFrames(absoluteTime, rate)
	// get rate record
	rec := _rateRecords[rate]
	var days, hours, minutes, seconds, frames int64
	var dropFrame = false

	days = (framecount / rec.hours) / 24
	hours = (framecount / rec.hours) % 24
	framecountHours := rec.hours * (hours + days*24)
	if rate == Smpte2997Drop {
		minutes = ((framecount + int64(2*int64(framecount-framecountHours)/1800) - (2 * int64(framecount-framecountHours) / 18000) - (framecountHours)) / 1800) % 60
		seconds = ((framecount - (rec.minutes * minutes) - (2 * int64(float64(minutes)/10)) - (framecountHours)) / rec.frames) % 3600
		frames = (framecount - (rec.frames * seconds) - (rec.minutes * minutes) - int64(2*((int)(float64(minutes)/float64(10)))) - (framecountHours)) % rec.frames
		dropFrame = true
	} else if rate == Smpte5994Drop {
		minutes = ((framecount + int64(4*int64(framecount-framecountHours)/3600) - (4 * int64(framecount-framecountHours) / 3600) - (framecountHours)) / 3600) % 60
		seconds = ((framecount - (rec.minutes * minutes) - (4 * int64(float64(minutes)/10)) - (framecountHours)) / rec.frames) % 3600
		frames = (framecount - (rec.frames * seconds) - (rec.minutes * minutes) - int64(4*((int)(float64(minutes)/float64(10)))) - (framecountHours)) % rec.frames
		dropFrame = true
	} else {
		minutes = ((framecount - framecountHours) / rec.minutes) % 60
		seconds = ((framecount - (rec.minutes * minutes) - framecountHours) / rec.frames) % 3600
		frames = (framecount - (rec.frames * seconds) - (rec.minutes * minutes) - framecountHours) % rec.frames
	}
	return formatTimeCodeString(int32(days), int32(hours), int32(minutes), int32(seconds), int32(frames), dropFrame)
}

// absoluteTimeToFrames Returns the number of frames.
func absoluteTimeToFrames(absoluteTime *decimal, rate SmpteFrameRate) int64 {
	rec := _rateRecords[rate]
	switch rate {
	case Smpte2398, Smpte2997Drop, Smpte2997NonDrop, Smpte5994Drop, Smpte5994NonDrop:
		ret := _1000div1001.MulInt64(rec.frames).Mul(absoluteTime.Add(_epsilon))
		return ret.Round(4).Int64()
		// return newDecimalFloat64(float64(ret)).Round(25).Floor().Int64()
	default:
		ret := absoluteTime.MulInt64(rec.frames)
		return ret.Round(4).Int64()
		// return absoluteTime.MulInt64(rec.frames).Round(25).Int64()
		// ret := float32(rec.frames) * absoluteTime.Float32()
		// return newDecimal(float64(ret)).Floor().Int64()
		// return Convert.ToInt64(30 * (float)absoluteTime);
	}
}

// framesToAbsoluteTime ..
/// <summary>
/// Returns the absolute time.
/// </summary>
/// <param name="frames">The number of frames.</param>
/// <param name="rate">The SMPTE frame rate to use for the conversion.</param>
/// <returns>The absolute time.</returns>
func framesToAbsoluteTime(frames int64, rate SmpteFrameRate) *decimal {
	rec := _rateRecords[rate]

	switch rate {
	case Smpte2398, Smpte2997Drop, Smpte2997NonDrop, Smpte5994Drop, Smpte5994NonDrop:
		return newDecimalInt64(frames).DivInt64(rec.frames).Div(_1000div1001)
		// return newDecimalInt64(frames).DivInt64(rec.frames).Div(newDecimal(1000).DivFloat64(1001).Round(11))
	default:
		return newDecimalInt64(frames).DivInt64(rec.frames)
	}
}
