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

// Decimal ..
type Decimal struct {
	value *big.Float
}

// NewInt create new big.Int object with default settings
func NewInt(value int64) *big.Int {
	return big.NewInt(value)
}

// NewFloat create new big.Float object with default settings
func NewFloat(value float64) *big.Float {
	return big.NewFloat(value).SetPrec(96) // to match c#'s decimal
}

// NewDecimal create a new decimal with float64
func NewDecimal(value float64) *Decimal {
	return &Decimal{
		value: NewFloat(value),
	}
}

// NewDecimalString create a new decimal with float64
func NewDecimalString(value string) *Decimal {
	ret, _, _ := big.ParseFloat(value, 10, 128, big.ToNearestEven)
	return &Decimal{
		value: ret,
	}
}

// NewDecimalInt64 create a new decimal with float64
func NewDecimalInt64(value int64) *Decimal {
	return NewDecimal(float64(value))
}

// Copy ..
func (m *Decimal) Copy() *Decimal {
	ret := NewFloat(0)
	return &Decimal{
		value: ret.Add(ret, m.value),
	}
}

// Float ..
func (m *Decimal) Float() *big.Float {
	return m.value
}

// AddFloat64 ..
func (m *Decimal) AddFloat64(value float64) *Decimal {
	ret := NewDecimal(0)
	ret.value.Add(m.value, NewFloat(value))
	return ret
}

// Add ..
func (m *Decimal) Add(value *Decimal) *Decimal {
	ret := NewDecimal(0)
	ret.value.Add(m.value, value.value)
	return ret
}

// Sub ..
func (m *Decimal) Sub(value *Decimal) *Decimal {
	ret := NewDecimal(0)
	ret.value.Sub(m.value, value.value)
	return ret
}

// MulFloat64 ..
func (m *Decimal) MulFloat64(value float64) *Decimal {
	ret := NewDecimal(0)
	ret.value.Mul(m.value, NewFloat((value)))
	return ret
}

// MulFloat32 multiply a float32 value
func (m *Decimal) MulFloat32(value float32) *Decimal {
	ret := NewDecimal(0)
	ret.value.Mul(m.value, NewFloat(float64(value)))
	return ret
}

// Mul multiply a decimal value
func (m *Decimal) Mul(value *Decimal) *Decimal {
	ret := NewDecimal(0)
	ret.value.Mul(m.value, value.value)
	return ret
}

// DivFloat64 ..
func (m *Decimal) DivFloat64(value float64) *Decimal {
	ret := NewDecimal(0)
	ret.value.Quo(m.value, NewFloat(value))
	return ret
}

// Div ..
func (m *Decimal) Div(value *Decimal) *Decimal {
	ret := NewDecimal(0)
	ret.value.Quo(m.value, value.value)
	return ret
}

// Floor ..
func (m *Decimal) Floor() *Decimal {
	ret, intVal := NewDecimal(0), NewInt(0)
	m.value.Int(intVal)
	ret.value.SetInt(intVal)
	return ret
}

// Round ..
func (m *Decimal) Round(decimals int) *Decimal {
	ret := NewDecimal(0)
	pow10, intVal := NewFloat(math.Pow10(decimals)), NewInt(0)
	ret.value.Mul(m.value, pow10).Int(intVal)
	ret.value.SetInt(intVal).Quo(ret.value, pow10)
	return ret
}

// Int64 ..
func (m *Decimal) Int64() int64 {
	ret, _ := m.value.Int64()
	return ret
}

// Float32 ..
func (m *Decimal) Float32() float32 {
	ret, _ := m.value.Float32()
	return ret
}

// Float64 ..
func (m *Decimal) Float64() float64 {
	ret, _ := m.value.Float64()
	return ret
}

// String ..
func (m *Decimal) String() string {
	return m.value.String()
}

// Debug ..
func (m *Decimal) Debug() {
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

	// Unknown Value.
	Unknown SmpteFrameRate = -1
)

// type convert2 struct {
// }

// func (m *convert2) ToDecimal(value Double) Decimal {
// 	return Decimal(value)
// }

// func (m *convert2) ToInt64(value *big.Float) int64 {
// 	return int64(value)
// }

// func (m *convert2) ToInt32(value int64) int32 {
// 	return int32(value)
// }

// Convert ...
// var convert = &convert2{}

const (
	Smpte12MBadFormat                 = "The timecode provided is not in the correct format."
	Smpte12MOutOfRange                = "The timecode provided is out of the expected range."
	Smpte12M_2398_BadFormat           = "Timecode frame value is not in the expected range for SMPTE 23.98 IVTC."
	Smpte12M_24_BadFormat             = "Timecode frame value is not in the expected range for SMPTE 24fps Film Sync."
	Smpte12M_25_BadFormat             = "Timecode frame value is not in the expected range for SMPTE 25fps PAL."
	Smpte12M_2997_Drop_BadFormat      = "Timecode frame value is not in the expected range for SMPTE 29.97 DropFrame."
	Smpte12M_2997_NonDrop_BadFormat   = "Timecode frame value is not in the expected range for SMTPE 29.97 NonDrop."
	Smpte12M_30_BadFormat             = "Timecode frame value is not in the expected range for SMPTE 30fps."
	MaxValueSmpte12MOverflowException = "The resulting timecode %v is out of the expected range of MaxValue %v."
	MinValueSmpte12MOverflowException = "The resulting timecode is out of the expected range of MinValue."
)

const (
	// SMPTEREGEXSTRING Regular expression string used for parsing out the timecode.
	SMPTEREGEXSTRING string = "(\\d{2}):(\\d{2}):(\\d{2})(?::|;)(\\d{2})"
	// EPSILON Epsilon value to deal with rounding precision issues with decimal and double values.
)

var (
	/// Regular expression object used for validating timecode.
	validateTimecode *regexp.Regexp = regexp.MustCompile(SMPTEREGEXSTRING)
	EPSILON          *Decimal       = NewDecimalString("0.00000000000000000000001")
)

// TimeCode ...
type TimeCode struct {
	/// The private Timespan used to track absolute time for this instance.
	absoluteTime *Decimal

	/// The frame rate for this instance.
	frameRate SmpteFrameRate
}

//FromTimeHours  Initializes a new instance of the TimeCode struct to a specified number of hours, minutes, and seconds.
func FromTimeHours(hours, minutes, seconds, frames int, rate SmpteFrameRate) (*TimeCode, error) {
	timeCode := fmt.Sprintf("%02d:%02d:%02d:%02d", hours, minutes, seconds, frames)
	time, err := smpte12mToAbsoluteTime(timeCode, rate)
	if err != nil {
		return nil, err
	}
	return &TimeCode{
		frameRate:    rate,
		absoluteTime: time,
	}, nil
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
	return &TimeCode{
		frameRate:    rate,
		absoluteTime: time,
	}, nil
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
   public TimeCode(string timeCodeAndRate)
   {
       string[] timeAndRate = timeCodeAndRate.Split('@');

       string time = string.Empty;
       string rate = string.Empty;

       if (timeAndRate.Length == 1)
       {
           time = timeAndRate[0];
           rate = "29.97";
       }
       else if (timeAndRate.Length == 2)
       {
           time = timeAndRate[0];
           rate = timeAndRate[1];
       }

       this.frameRate = Smpte2997NonDrop;

       if (rate == "29.97" && time.IndexOf(';') > -1)
       {
           this.frameRate = Smpte2997Drop;
       }
       else if (rate == "29.97" && time.IndexOf(';') == -1)
       {
           this.frameRate = Smpte2997NonDrop;
       }
       else if (rate == "25")
       {
           this.frameRate = Smpte25;
       }
       else if (rate == "23.98")
       {
           this.frameRate = Smpte2398;
       }
       else if (rate == "24")
       {
           this.frameRate = Smpte24;
       }
       else if (rate == "30")
       {
           this.frameRate = Smpte30;
       }

       this.absoluteTime = smpte12mToAbsoluteTime(time, this.frameRate);
   }
*/

//FromTimeCode Initializes a new instance of the TimeCode struct using a time code string and a SMPTE framerate.
func FromTimeCode(timeCode string, rate SmpteFrameRate) (*TimeCode, error) {
	time, err := smpte12mToAbsoluteTime(timeCode, rate)
	if err != nil {
		return nil, err
	}
	return &TimeCode{
		frameRate:    rate,
		absoluteTime: time,
	}, nil
}

// FromTime Initializes a new instance of the TimeCode struct using an absolute time value, and the SMPTE framerate.
func FromTime(absoluteTime float64, rate SmpteFrameRate) (*TimeCode, error) {
	return FromTimeDecimal(NewDecimal(absoluteTime), rate)
}

// FromTimeDecimal Initializes a new instance of the TimeCode struct using an absolute time value in <see cref="decimal"/> precision, and the SMPTE framerate.
func FromTimeDecimal(absoluteTime *Decimal, rate SmpteFrameRate) (*TimeCode, error) {
	return &TimeCode{
		frameRate:    rate,
		absoluteTime: absoluteTime.Copy(),
	}, nil
}

/*
   /// <summary>
   /// Initializes a new instance of the TimeCode struct a long value that represents a value of a 27 Mhz clock.
   /// </summary>
   /// <param name="ticks27Mhz">The long value in 27 Mhz clock ticks.</param>
   /// <param name="rate">The SMPTE frame rate to use for this instance.</param>
   public TimeCode(long ticks27Mhz, SmpteFrameRate rate)
   {
       this.absoluteTime = ticks27MhzToAbsoluteTime(ticks27Mhz);
       this.frameRate = rate;
   }
*/
/// <summary>
///  Gets the number of ticks in 1 day.
///  This field is constant.
/// </summary>
/// <value>The number of ticks in 1 day.</value>
const TicksPerDay int64 = 864000000000

/// <summary>
///  Gets the number of absolute time ticks in 1 day.
///  This field is constant.
/// </summary>
/// <value>The number of absolute time ticks in 1 day.</value>
const TicksPerDayAbsoluteTime int64 = 86400

/// <summary>
///  Gets the number of ticks in 1 hour. This field is constant.
/// </summary>
/// <value>The number of ticks in 1 hour.</value>
const TicksPerHour int64 = 36000000000

/// <summary>
///  Gets the number of absolute time ticks in 1 hour. This field is constant.
/// </summary>
/// <value>The number of absolute time ticks in 1 hour.</value>
const TicksPerHourAbsoluteTime int64 = 3600

/// <summary>
/// Gets the number of ticks in 1 millisecond. This field is constant.
/// </summary>
/// <value>The number of ticks in 1 millisecond.</value>
const TicksPerMillisecond int64 = 10000

/// <summary>
/// Gets the number of ticks in 1 millisecond. This field is constant.
/// </summary>
/// <value>The number of ticks in 1 millisecond.</value>
const TicksPerMillisecondAbsoluteTime float64 = 0.0010000

/// <summary>
/// Gets the number of ticks in 1 minute. This field is constant.
/// </summary>
/// <value>The number of ticks in 1 minute.</value>
const TicksPerMinute int64 = 600000000

/// <summary>
/// Gets the number of absolute time ticks in 1 minute. This field is constant.
/// </summary>
/// <value>The number of absolute time ticks in 1 minute.</value>
const TicksPerMinuteAbsoluteTime float64 = 60

/// <summary>
/// Gets the number of ticks in 1 second.
/// </summary>
/// <value>The number of ticks in 1 second.</value>
const TicksPerSecond int64 = 10000000

/// <summary>
/// Gets the number of ticks in 1 second.
/// </summary>
/// <value>The number of ticks in 1 second.</value>
const TicksPerSecondAbsoluteTime float64 = 1.0000000

/// <summary>
/// Gets the minimum TimeCode value. This field is read-only.
/// </summary>
/// <value>The minimum TimeCode value.</value>
const MinValue float64 = 0

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

/// <summary>
/// Gets or sets the current SMPTE framerate for this TimeCode instance.
/// </summary>
/// <value>The frame rate of the TimeCode.</value>
func (m *TimeCode) FrameRate() SmpteFrameRate {
	return m.frameRate
}

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

	var hours float64 = 0

	switch m.frameRate {
	case Smpte2398, Smpte24:
		hours = float64(framecount) / 86400
		break
	case Smpte25:
		hours = float64(framecount) / 90000
		break
	case Smpte2997Drop:
		hours = float64(framecount) / 107892
		break
	case Smpte2997NonDrop, Smpte30:
		hours = float64(framecount) / 108000
		break
	default:
		hours = float64(framecount) / 108000
		break
	}

	return hours
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

	var minutes float64 = 0

	switch m.frameRate {
	case Smpte2398, Smpte24:
		minutes = float64(framecount) / 1400
		break
	case Smpte25:
		minutes = float64(framecount) / 1500
		break
	case Smpte2997Drop, Smpte2997NonDrop, Smpte30:
		minutes = float64(framecount) / 1800
		break
	default:
		minutes = float64(framecount) / 1800
		break
	}

	return minutes
}

//TotalSeconds Gets the value of the current TimeCode structure expressed in whole
/// and fractional seconds. Not as Precise as the TotalSecondsPrecision.
func (m *TimeCode) TotalSeconds() float64 {
	return m.absoluteTime.Copy().Round(7).Float64()
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
func MaxValue(frameRate SmpteFrameRate) *Decimal {
	switch frameRate {
	case Smpte2398:
		return NewDecimalString("86486.358291666700000")

	case Smpte24:
		return NewDecimalString("86399.958333333300000")

	case Smpte25:
		return NewDecimalString("86399.960000000000000")

	case Smpte2997Drop:
		return NewDecimalString("86399.880233333300000")

	case Smpte2997NonDrop:
		return NewDecimalString("86486.366633333300000")

	case Smpte30:
		return NewDecimalString("86399.966666666700000")

	default:
		return NewDecimalString("86399")
	}
}

//Sub Subtracts a specified TimeCode from another specified TimeCode.
func Sub(t1, t2 *TimeCode) (*TimeCode, error) {
	rate, time := t1.frameRate, NewDecimal(0)
	time.Float().Sub(t1.absoluteTime.Float(), t2.absoluteTime.Float())
	t3 := &TimeCode{
		frameRate:    rate,
		absoluteTime: time,
	}

	if t3.TotalSeconds() < MinValue {
		return nil, errors.New(MinValueSmpte12MOverflowException)
	}

	return t3, nil
}

// NotEqual Indicates whether two TimeCode instances are not equal.
func NotEqual(t1, t2 *TimeCode) bool {
	var timeCode1, _ = FromTimeDecimal(t1.absoluteTime, Smpte30)
	var timeCode2, _ = FromTimeDecimal(t2.absoluteTime, Smpte30)

	if timeCode1.TotalSeconds() != timeCode2.TotalSeconds() {
		return true
	}

	return false
}

// Add two specified TimeCode instances.
func Add(t1, t2 *TimeCode) (*TimeCode, error) {
	rate, time := t1.frameRate, NewDecimal(0)
	time.Float().Add(t1.absoluteTime.Float(), t2.absoluteTime.Float())
	t3 := &TimeCode{
		frameRate:    rate,
		absoluteTime: time,
	}
	// check overflow
	maxValue := MaxValue(rate).Float64()
	if t3.TotalSecondsPrecision() > maxValue && t3.TotalSeconds() > maxValue {
		return nil, fmt.Errorf(MaxValueSmpte12MOverflowException, t3.TotalSecondsPrecision(), maxValue)
	}
	// return
	return t3, nil
}

// LessThan ..
///  Indicates whether a specified TimeCode is less than another
///  specified TimeCode.
func LessThan(t1, t2 *TimeCode) bool {
	var timeCode1, _ = FromTimeDecimal(t1.absoluteTime, Smpte30)
	var timeCode2, _ = FromTimeDecimal(t2.absoluteTime, Smpte30)

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
	var timeCode1, _ = FromTimeDecimal(t1.absoluteTime, Smpte30)
	var timeCode2, _ = FromTimeDecimal(t2.absoluteTime, Smpte30)

	if timeCode1.TotalSeconds() < timeCode2.TotalSeconds() || (timeCode1.TotalSeconds() == timeCode2.TotalSeconds()) {
		return true
	}

	return false
}

//Equal  Indicates whether two TimeCode instances are equal.
func Equal(t1, t2 *TimeCode) bool {
	var timeCode1, _ = FromTimeDecimal(t1.absoluteTime, Smpte30)
	var timeCode2, _ = FromTimeDecimal(t2.absoluteTime, Smpte30)

	if timeCode1.TotalSeconds() == timeCode2.TotalSeconds() {
		return true
	}

	return false
}

// GreatThan Indicates whether a specified TimeCode is greater than another specified
func GreatThan(t1, t2 *TimeCode) bool {
	var timeCode1, _ = FromTimeDecimal(t1.absoluteTime, Smpte30)
	var timeCode2, _ = FromTimeDecimal(t2.absoluteTime, Smpte30)

	if timeCode1.TotalSeconds() > timeCode2.TotalSeconds() {
		return true
	}

	return false
}

// GreatEqual Indicates whether a specified TimeCode is greater than or equal to
///     another specified TimeCode.
func GreatEqual(t1, t2 *TimeCode) bool {

	var timeCode1, _ = FromTimeDecimal(t1.absoluteTime, Smpte30)
	var timeCode2, _ = FromTimeDecimal(t2.absoluteTime, Smpte30)

	if timeCode1.TotalSeconds() > timeCode2.TotalSeconds() || (timeCode1.TotalSeconds() == timeCode2.TotalSeconds()) {
		return true
	}

	return false
}

// Ticks27MhzToSmpte12M Returns a SMPTE 12M formatted time code string from a 27Mhz ticks value.
func Ticks27MhzToSmpte12M(ticks27Mhz int64, rate SmpteFrameRate) string {
	switch rate {
	case Smpte2398:
		return ticks27MhzToSmpte12M2398Fps(ticks27Mhz)
	case Smpte24:
		return ticks27MhzToSmpte12M24Fps(ticks27Mhz)
	case Smpte25:
		return ticks27MhzToSmpte12M25Fps(ticks27Mhz)
	case Smpte2997Drop:
		return ticks27MhzToSmpte12M_29_27_Drop(ticks27Mhz)
	case Smpte2997NonDrop:
		return ticks27MhzToSmpte12M_29_27_NonDrop(ticks27Mhz)
	case Smpte30:
		return ticks27MhzToSmpte12M_30fps(ticks27Mhz)
	default:
		return ticks27MhzToSmpte12M_30fps(ticks27Mhz)
	}
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
	absoluteTime := days * float64(TicksPerDayAbsoluteTime)
	return FromTime(absoluteTime, rate)
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
	abs := framesToAbsoluteTime(frames, rate)
	return FromTime(abs, rate)
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
	return FromTimeDecimal(absoluteTime, rate)
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
	if !validateTimecode.Match([]byte(timeCode)) {
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
	switch rate {
	case Smpte2398:
		return smpte12M2398FpsToTicks27Mhz(timeCode)
	case Smpte24:
		return smpte12M24FpsToTicks27Mhz(timeCode)
	case Smpte25:
		return smpte12M25FpsToTicks27Mhz(timeCode)
	case Smpte2997Drop:
		return smpte12M_29_27_DropToTicks27Mhz(timeCode)
	case Smpte2997NonDrop:
		return smpte12M_29_27_NonDropToTicks27Mhz(timeCode)
	case Smpte30:
		return smpte12M30FpsToTicks27Mhz(timeCode)
	default:
		return smpte12M30FpsToTicks27Mhz(timeCode)
	}
}

/*
   /// <summary>
   /// Parses a framerate value as double and converts it to a member of the SmpteFrameRate enumeration.
   /// </summary>
   /// <param name="rate">Double value of the framerate.</param>
   /// <returns>A SmpteFrameRate enumeration value that matches the incoming rates.</returns>
   public static SmpteFrameRate ParseFramerate(double rate)
   {
       int rateRounded = (int)Math.Floor(rate);

       switch (rateRounded)
       {
           case 23: return Smpte2398;
           case 24: return Smpte24;
           case 25: return Smpte25;
           case 29: return Smpte2997NonDrop;
           case 30: return Smpte30;
           case 50: return Smpte25;
           case 60: return Smpte30;
           case 59: return Smpte2997NonDrop;
       }

       return Unknown;
   }
*/
/// <summary>
/// Adds the specified TimeCode to this instance.
/// </summary>
/// <param name="ts">A TimeCode.</param>
/// <returns>A TimeCode that represents the value of this instance plus the value of ts.
/// </returns>
/// <exception cref="System.OverflowException">
/// The resulting TimeCode is less than TimeCode.MinValue or greater than TimeCode.MaxValue.
/// </exception>
func (m *TimeCode) Add(tc *TimeCode) (*TimeCode, error) {
	return Add(m, tc)
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
           throw new ArgumentException(Smpte12MOutOfRange);
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
func smpte12mToAbsoluteTime(timeCode string, rate SmpteFrameRate) (*Decimal, error) {
	switch rate {
	case Smpte2398:
		return smpte12M_23_98_ToAbsoluteTime(timeCode)
	case Smpte24:
		return smpte12M_24_ToAbsoluteTime(timeCode)
	case Smpte25:
		return smpte12M_25_ToAbsoluteTime(timeCode)
	case Smpte2997Drop:
		return smpte12M_29_97_Drop_ToAbsoluteTime(timeCode)
	case Smpte2997NonDrop:
		return smpte12M_29_97_NonDrop_ToAbsoluteTime(timeCode)
	case Smpte30:
		return smpte12M_30_ToAbsoluteTime(timeCode)
	default:
		return nil, nil
	}
}

// parseTimecodeString Parses a timecode string for the different parts of the timecode.
func parseTimecodeString(timeCode string) (days, hours, minutes, seconds, frames int64, err error) {
	if !validateTimecode.Match([]byte(timeCode)) {
		err = errors.New(Smpte12MBadFormat)
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
		err = errors.New(Smpte12MOutOfRange)
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
           throw new FormatException(Smpte12MBadFormat);
       }

       string[] times = timeCode.Split(':', ';');

       hours = ret, _ := strconv.Atoi(times[0]);
       minutes = ret, _ := strconv.Atoi(times[1]);
       seconds = ret, _ := strconv.Atoi(times[2]);
       frames = ret, _ := strconv.Atoi(times[3]);

       if ((hours >= 24) || (minutes >= 60) || (seconds >= 60) || (frames >= 30))
       {
           throw new FormatException(Smpte12MOutOfRange);
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

// smpte12M_23_98_ToAbsoluteTime Converts to Absolute time from SMPTE 12M 23.98.
func smpte12M_23_98_ToAbsoluteTime(timeCode string) (*Decimal, error) {
	days, hours, minutes, seconds, frames, err := parseTimecodeString(timeCode)
	if err != nil {
		return nil, err
	}

	if frames >= 24 {
		return nil, errors.New(Smpte12M_2398_BadFormat)
	}

	return NewDecimal(1001).DivFloat64(24000).MulFloat64(float64(frames + (24 * seconds) + (1440 * minutes) + (86400 * hours) + (2073600 * days))), nil
}

//smpte12M_24_ToAbsoluteTime Converts to Absolute time from SMPTE 12M 24.
func smpte12M_24_ToAbsoluteTime(timeCode string) (*Decimal, error) {
	days, hours, minutes, seconds, frames, err := parseTimecodeString(timeCode)
	if err != nil {
		return nil, err
	}

	if frames >= 24 {
		return nil, errors.New(Smpte12M_24_BadFormat)
	}

	return NewDecimal(1).DivFloat64(24).MulFloat64(float64(frames + (24 * seconds) + (1440 * minutes) + (86400 * hours) + (2073600 * days))), nil
}

//smpte12M_25_ToAbsoluteTime Converts to Absolute time from SMPTE 12M 25.
func smpte12M_25_ToAbsoluteTime(timeCode string) (*Decimal, error) {
	days, hours, minutes, seconds, frames, err := parseTimecodeString(timeCode)
	if err != nil {
		return nil, err
	}

	if frames >= 25 {
		return nil, errors.New(Smpte12M_25_BadFormat)
	}

	return NewDecimal(1).DivFloat64(25).MulFloat64(float64(frames + (25 * seconds) + (1500 * minutes) + (90000 * hours) + (2160000 * days))), nil
}

//smpte12M_29_97_Drop_ToAbsoluteTime Converts to Absolute time from SMPTE 12M 29.97 Drop frame.
func smpte12M_29_97_Drop_ToAbsoluteTime(timeCode string) (*Decimal, error) {
	days, hours, minutes, seconds, frames, err := parseTimecodeString(timeCode)
	if err != nil {
		return nil, err
	}

	if frames >= 30 {
		return nil, errors.New(Smpte12M_2997_Drop_BadFormat)
	}

	return NewDecimal(1001).DivFloat64(30000).MulFloat64(float64(frames + (30 * seconds) + (1798 * minutes) + ((2 * (minutes / 10)) + (107892 * hours) + (2589408 * days)))), nil
	// return (1001 / 30000M) * (frames + (30 * seconds) + (1798 * minutes) + ((2 * (minutes / 10)) + (107892 * hours) + (2589408 * days)));
}

//smpte12M_29_97_NonDrop_ToAbsoluteTime Converts to Absolute time from SMPTE 12M 29.97 Non Drop.
func smpte12M_29_97_NonDrop_ToAbsoluteTime(timeCode string) (*Decimal, error) {
	days, hours, minutes, seconds, frames, err := parseTimecodeString(timeCode)
	if err != nil {
		return nil, err
	}

	if frames >= 30 {
		return nil, errors.New(Smpte12M_2997_NonDrop_BadFormat)
	}

	// round to 8 decimals
	return NewDecimal(1001).DivFloat64(30000).MulFloat64(float64(frames + (30 * seconds) + (1800 * minutes) + (108000 * hours) + (2592000 * days))).Round(8), nil
}

//smpte12M_30_ToAbsoluteTime Converts to Absolute time from SMPTE 12M 30.
func smpte12M_30_ToAbsoluteTime(timeCode string) (*Decimal, error) {
	days, hours, minutes, seconds, frames, err := parseTimecodeString(timeCode)
	if err != nil {
		return nil, err
	}
	if frames >= 30 {
		return nil, errors.New(Smpte12M_30_BadFormat)
	}

	return NewDecimal(1).DivFloat64(30).MulFloat64(float64(frames + (30 * seconds) + (1800 * minutes) + (108000 * hours) + (2592000 * days))), nil
}

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
func ticksPcrTbToAbsoluteTime(ticksPcrTb int64) *Decimal {
	return NewDecimal(float64(ticksPcrTb)).DivFloat64(90000)
}

// ticks27MhzToAbsoluteTime ..
/// <summary>
/// Converts the specified absolute time to absolute time.
/// </summary>
/// <param name="ticks27Mhz">Ticks 27Mhz to be converted.</param>
/// <returns>The absolute time.</returns>
func ticks27MhzToAbsoluteTime(ticks27Mhz int64) *Decimal {
	ticksPcrTb := ticks27MhzToPcrTb(ticks27Mhz)
	return ticksPcrTbToAbsoluteTime(ticksPcrTb)
}

// absoluteTimeToSmpte12M Converts to SMPTE 12M.
func absoluteTimeToSmpte12M(absoluteTime *Decimal, rate SmpteFrameRate) string {
	timeCode := ""

	if rate == Smpte2398 {
		timeCode = absoluteTimeToSmpte12M_23_98fps(absoluteTime)
	} else if rate == Smpte24 {
		timeCode = absoluteTimeToSmpte12M_24fps(absoluteTime)
	} else if rate == Smpte25 {
		timeCode = absoluteTimeToSmpte12M_25fps(absoluteTime)
	} else if rate == Smpte2997Drop {
		timeCode = absoluteTimeToSmpte12M_29_97_Drop(absoluteTime)
	} else if rate == Smpte2997NonDrop {
		timeCode = absoluteTimeToSmpte12M_29_97_NonDrop(absoluteTime)
	} else if rate == Smpte30 {
		timeCode = absoluteTimeToSmpte12M30Fps(absoluteTime)
	}

	return timeCode
}

// absoluteTimeToFrames Returns the number of frames.
func absoluteTimeToFrames(absoluteTime *Decimal, rate SmpteFrameRate) int64 {
	if rate == Smpte2398 {
		ret := 24 * NewDecimal(1000).DivFloat64(1001).Float32() * absoluteTime.Add(EPSILON).Float32()
		return NewDecimal(float64(ret)).Round(25).Floor().Int64()
		// return Convert.ToInt64(decimal.Floor(decimal.Round((decimal)(24 * (float)(1000 / 1001M) * (float)(absoluteTime + EPSILON)), 25)));
	}

	if rate == Smpte24 {
		ret := 24 * absoluteTime.Float32()
		return NewDecimal(float64(ret)).Floor().Int64()
		// return Convert.ToInt64(decimal.Floor((decimal)(24 * (float)absoluteTime)));
	}

	if rate == Smpte25 {
		ret := 25 * absoluteTime.Float32()
		return NewDecimal(float64(ret)).Floor().Int64()
		// return Convert.ToInt64(decimal.Floor((decimal)(25 * (float)absoluteTime)));
	}

	if rate == Smpte2997Drop {
		ret := 30 * NewDecimal(1000).DivFloat64(1001).Float32() * absoluteTime.Add(EPSILON).Float32()
		return NewDecimal(float64(ret)).Round(25).Floor().Int64()
		// return Convert.ToInt64(decimal.Floor(decimal.Round((decimal)(30 * (float)(1000 / 1001M) * (float)(absoluteTime + EPSILON)), 25)));
	}

	if rate == Smpte2997NonDrop {
		ret := 30 * NewDecimal(1000).DivFloat64(1001).Float32() * absoluteTime.Add(EPSILON).Float32()
		return NewDecimal(float64(ret)).Round(25).Floor().Int64()
		// return Convert.ToInt64(decimal.Floor(decimal.Round((decimal)(30 * (float)(1000 / 1001M) * (float)(absoluteTime + EPSILON)), 25)));
	}

	if rate == Smpte30 {
		ret := 30 * absoluteTime.Float32()
		return int64(float64(ret))
		// return Convert.ToInt64(30 * (float)absoluteTime);
	}

	return NewDecimal(30).Mul(absoluteTime).Floor().Int64()
	// return Convert.ToInt64(decimal.Floor(30 * absoluteTime));
}

// framesToAbsoluteTime ..
/// <summary>
/// Returns the absolute time.
/// </summary>
/// <param name="frames">The number of frames.</param>
/// <param name="rate">The SMPTE frame rate to use for the conversion.</param>
/// <returns>The absolute time.</returns>
func framesToAbsoluteTime(frames int64, rate SmpteFrameRate) float64 {
	var absoluteTimeInDecimal *Decimal

	if rate == Smpte2398 {
		absoluteTimeInDecimal = NewDecimalInt64(frames).DivFloat64(24).DivFloat64(NewDecimal(1000).DivFloat64(1001).Round(11).Float64())
		//    frames / 24M / decimal.Round((1000 / 1001M), 11);
	} else if rate == Smpte24 {
		absoluteTimeInDecimal = NewDecimalInt64(frames).DivFloat64(24)
		//    absoluteTimeInDecimal = frames / 24M;
	} else if rate == Smpte25 {
		absoluteTimeInDecimal = NewDecimalInt64(frames).DivFloat64(25)
		//    absoluteTimeInDecimal = frames / 25M;
	} else if rate == Smpte2997Drop || rate == Smpte2997NonDrop {
		absoluteTimeInDecimal = NewDecimalInt64(frames).DivFloat64(30).DivFloat64(NewDecimal(1000).DivFloat64(1001).Round(11).Float64())
		//    absoluteTimeInDecimal = frames / 30M / decimal.Round((1000 / 1001M), 11);
	} else if rate == Smpte30 {
		absoluteTimeInDecimal = NewDecimalInt64(frames).DivFloat64(30)
		//    absoluteTimeInDecimal = frames / 30M;
	} else {
		absoluteTimeInDecimal = NewDecimalInt64(frames).DivFloat64(30)
		//    absoluteTimeInDecimal = frames / 30M;
	}

	return absoluteTimeInDecimal.Float64()
}

// absoluteTimeToSmpte12M_23_98fps Returns the SMPTE 12M 23.98 timecode.
func absoluteTimeToSmpte12M_23_98fps(absoluteTime *Decimal) string {
	framecount := absoluteTimeToFrames(absoluteTime, Smpte2398)

	days := int32((framecount / 86400) / 24)
	hours := int32((framecount / 86400) % 24)
	minutes := int32(((framecount - int64(86400*hours)) / 1440) % 60)
	seconds := int32(((framecount - int64(1440*minutes) - int64(86400*hours)) / 24) % 3600)
	frames := int32((framecount - int64(24*seconds) - int64(1440*minutes) - int64(86400*hours)) % 24)

	return formatTimeCodeString(days, hours, minutes, seconds, frames, false)
}

// absoluteTimeToSmpte12M_24fps Converts to SMPTE 12M 24fps.
func absoluteTimeToSmpte12M_24fps(absoluteTime *Decimal) string {
	framecount := absoluteTimeToFrames(absoluteTime, Smpte24)

	days := int32((framecount / 86400) / 24)
	hours := int32((framecount / 86400) % 24)
	minutes := int32(((framecount - int64(86400*hours)) / 1440) % 60)
	seconds := int32(((framecount - int64(1440*minutes) - int64(86400*hours)) / 24) % 3600)
	frames := int32((framecount - int64(24*seconds) - int64(1440*minutes) - int64(86400*hours)) % 24)

	return formatTimeCodeString(days, hours, minutes, seconds, frames, false)
}

// absoluteTimeToSmpte12M_25fps Converts to SMPTE 12M 25fps.
func absoluteTimeToSmpte12M_25fps(absoluteTime *Decimal) string {
	framecount := absoluteTimeToFrames(absoluteTime, Smpte25)

	days := int32((framecount / 90000) / 24)
	hours := int32((framecount / 90000) % 24)
	minutes := int32(((framecount - int64(90000*hours)) / 1500) % 60)
	seconds := int32(((framecount - int64(1500*minutes) - int64(90000*hours)) / 25) % 3600)
	frames := int32((framecount - int64(25*seconds) - int64(1500*minutes) - int64(90000*hours)) % 25)

	return formatTimeCodeString(days, hours, minutes, seconds, frames, false)
}

// absoluteTimeToSmpte12M_29_97_Drop Converts to SMPTE 12M 29.97fps Drop.
func absoluteTimeToSmpte12M_29_97_Drop(absoluteTime *Decimal) string {
	framecount := absoluteTimeToFrames(absoluteTime, Smpte2997Drop)
	days := (int32)((framecount / 107892) / 24)
	hours := (int32)((framecount / 107892) % 24)
	minutes := int32(((framecount + int64(2*((int)((framecount-int64(107892*hours))/1800))) - int64(2*((int)((framecount-int64(107892*hours))/18000))) - int64(107892*hours)) / 1800) % 60)
	seconds := int32(((framecount - int64(1798*minutes) - int64(2*((int)(float64(minutes)/float64(10)))) - int64(107892*hours)) / 30) % 3600)
	frames := int32((framecount - int64(30*seconds) - int64(1798*minutes) - int64(2*((int)(float64(minutes)/float64(10)))) - int64(107892*hours)) % 30)

	return formatTimeCodeString(days, hours, minutes, seconds, frames, true)
}

// absoluteTimeToSmpte12M_29_97_NonDrop Converts to SMPTE 12M 29.97fps Non Drop.
func absoluteTimeToSmpte12M_29_97_NonDrop(absoluteTime *Decimal) string {
	framecount := absoluteTimeToFrames(absoluteTime, Smpte2997NonDrop)

	days := int32((framecount / 108000) / 24)
	hours := int32((framecount / 108000) % 24)
	minutes := int32(((framecount - int64(108000*hours)) / 1800) % 60)
	seconds := int32(((framecount - int64(1800*minutes) - int64(108000*hours)) / 30) % 3600)
	frames := int32((framecount - int64(30*seconds) - int64(1800*minutes) - int64(108000*hours)) % 30)

	return formatTimeCodeString(days, hours, minutes, seconds, frames, false)
}

// absoluteTimeToSmpte12M30Fps Converts to SMPTE 12M 30fps.
func absoluteTimeToSmpte12M30Fps(absoluteTime *Decimal) string {
	framecount := absoluteTimeToFrames(absoluteTime, Smpte30)

	days := int32((framecount / 108000) / 24)
	hours := int32((framecount / 108000) % 24)
	minutes := int32(((framecount - int64(108000*hours)) / 1800) % 60)
	seconds := int32(((framecount - int64(1800*minutes) - int64(108000*hours)) / 30) % 3600)
	frames := int32((framecount - int64(30*seconds) - int64(1800*minutes) - int64(108000*hours)) % 30)

	return formatTimeCodeString(days, hours, minutes, seconds, frames, false)
}

// smpte12M30FpsToTicks27Mhz ..
/// <summary>
/// Converts to Ticks 27Mhz.
/// </summary>
/// <param name="timeCode">The timecode to convert from.</param>
/// <returns>The number of 27Mhz ticks.</returns>
func smpte12M30FpsToTicks27Mhz(timeCode string) int64 {
	t, _ := FromTimeCode(timeCode, Smpte30)
	days, hours, minutes, seconds, frames := t.DaysSegment(), t.HoursSegment(), t.MinutesSegment(), t.SecondsSegment(), t.FramesSegment()
	ticksPcrTb := (frames * 3000) + (90000 * seconds) + (5400000 * minutes) + (324000000 * hours) + (7776000000 * days)
	// long ticksPcrTb = (t.FramesSegment * 3000) + (90000 * t.SecondsSegment) + (5400000 * t.MinutesSegment) + (324000000 * t.HoursSegment) + (7776000000 * t.DaysSegment);
	return ticksPcrTb * 300
}

// smpte12M2398FpsToTicks27Mhz Converts to Ticks 27Mhz.
func smpte12M2398FpsToTicks27Mhz(timeCode string) int64 {
	t, _ := FromTimeCode(timeCode, Smpte2398)
	days, hours, minutes, seconds, frames := t.DaysSegment(), t.HoursSegment(), t.MinutesSegment(), t.SecondsSegment(), t.FramesSegment()
	ticksPcrTb := int64((math.Ceil(float64(1001)*(float64(15)/4)*float64(frames) + float64(90090*seconds) + float64(5405400*minutes) + float64(324324000*hours) + float64(7783776000*days))))
	// long ticksPcrTb = Convert.ToInt64((Math.Ceiling(1001 * (15 / 4D) * t.FramesSegment) + (90090 * t.SecondsSegment) + (5405400 * t.MinutesSegment) + (324324000D * t.HoursSegment) + (7783776000 * t.DaysSegment)));
	return ticksPcrTb * 300
}

// smpte12M24FpsToTicks27Mhz Converts to Ticks 27Mhz.
func smpte12M24FpsToTicks27Mhz(timeCode string) int64 {
	t, _ := FromTimeCode(timeCode, Smpte24)
	days, hours, minutes, seconds, frames := t.DaysSegment(), t.HoursSegment(), t.MinutesSegment(), t.SecondsSegment(), t.FramesSegment()
	ticksPcrTb := (frames * 3750) + (90000 * seconds) + (5400000 * minutes) + (324000000 * hours) + (7776000000 * days)
	// long ticksPcrTb = (t.FramesSegment * 3750) + (90000 * t.SecondsSegment) + (5400000 * t.MinutesSegment) + (324000000 * t.HoursSegment) + (7776000000 * t.DaysSegment);
	return ticksPcrTb * 300
}

// smpte12M25FpsToTicks27Mhz Converts to Ticks 27Mhz.
func smpte12M25FpsToTicks27Mhz(timeCode string) int64 {
	t, _ := FromTimeCode(timeCode, Smpte25)
	days, hours, minutes, seconds, frames := t.DaysSegment(), t.HoursSegment(), t.MinutesSegment(), t.SecondsSegment(), t.FramesSegment()
	ticksPcrTb := (frames * 3600) + (90000 * seconds) + (5400000 * minutes) + (324000000 * hours) + (7776000000 * days)
	// long ticksPcrTb = (t.FramesSegment * 3600) + (90000 * t.SecondsSegment) + (5400000 * t.MinutesSegment) + (324000000 * t.HoursSegment) + (7776000000 * t.DaysSegment);
	return ticksPcrTb * 300
}

// smpte12M_29_27_NonDropToTicks27Mhz Converts to Ticks 27Mhz.
func smpte12M_29_27_NonDropToTicks27Mhz(timeCode string) int64 {
	t, _ := FromTimeCode(timeCode, Smpte2997NonDrop)
	days, hours, minutes, seconds, frames := t.DaysSegment(), t.HoursSegment(), t.MinutesSegment(), t.SecondsSegment(), t.FramesSegment()
	ticksPcrTb := (frames * 3003) + (90090 * seconds) + (5405400 * minutes) + (324324000 * hours) + (7783776000 * days)
	// long ticksPcrTb = (t.FramesSegment * 3003) + (90090 * t.SecondsSegment) + (5405400 * t.MinutesSegment) + (324324000 * t.HoursSegment) + (7783776000 * t.DaysSegment);
	return ticksPcrTb * 300
}

// smpte12M_29_27_DropToTicks27Mhz Converts to Ticks 27Mhz.
func smpte12M_29_27_DropToTicks27Mhz(timeCode string) int64 {
	t, _ := FromTimeCode(timeCode, Smpte2997Drop)
	hours, minutes, seconds, frames := t.HoursSegment(), t.MinutesSegment(), t.SecondsSegment(), t.FramesSegment()
	ticksPcrTb := (3003 * frames) + (90090 * seconds) + (5399394 * minutes) + (6006 * int64(float64(seconds)/10)) + (323999676 * hours)
	// long ticksPcrTb = (3003 * t.FramesSegment) + (90090 * t.SecondsSegment) + (5399394 * t.MinutesSegment) + (6006 * (int)(t.MinutesSegment / 10D)) + (323999676 * t.HoursSegment);
	return ticksPcrTb * 300
}

// ticks27MhzToSmpte12M_29_27_NonDrop Converts to SMPTE 12M 29.27fps Non Drop.
func ticks27MhzToSmpte12M_29_27_NonDrop(ticks27Mhz int64) string {
	pcrTb := ticks27MhzToPcrTb(ticks27Mhz)
	framecount := int32(pcrTb / 3003)

	days := int32((framecount / 108000) / 24)
	hours := int32((framecount / 108000) % 24)
	minutes := int32(((framecount - (108000 * hours)) / 1800) % 60)
	seconds := int32(((framecount - (1800 * minutes) - (108000 * hours)) / 30) % 3600)
	frames := (framecount - (30 * seconds) - (1800 * minutes) - (108000 * hours)) % 30

	return formatTimeCodeString(days, hours, minutes, seconds, frames, false)
}

// ticks27MhzToSmpte12M_29_27_Drop Converts to SMPTE 12M 29.27fps Non Drop.
func ticks27MhzToSmpte12M_29_27_Drop(ticks27Mhz int64) string {
	pcrTb := ticks27MhzToPcrTb(ticks27Mhz)
	framecount := int32(pcrTb / 3003)
	hours := int32((framecount / 107892) % 24)
	minutes := int32((framecount + (2 * int32((framecount-(107892*hours))/1800)) - (2 * int32((framecount-(107892*hours))/18000)) - (107892 * hours)) / 1800)
	seconds := int32((framecount - (1798 * minutes) - (2 * int32(minutes/10)) - (107892 * hours)) / 30)
	frames := framecount - (30 * seconds) - (1798 * minutes) - (2 * int32(minutes/10)) - (107892 * hours)

	return formatTimeCodeString(0, hours, minutes, seconds, frames, true)
}

// ticks27MhzToSmpte12M2398Fps Converts to SMPTE 12M 23.98fps.
func ticks27MhzToSmpte12M2398Fps(ticks27Mhz int64) string {
	pcrTb := ticks27MhzToPcrTb(ticks27Mhz)

	framecount := (int32)((float64(4) / 15) * (float64(pcrTb) / 1001))

	days := int32((framecount / 86400) / 24)
	hours := int32((framecount / 86400) % 24)
	minutes := int32(((framecount - (86400 * hours)) / 1440) % 60)
	seconds := int32(((framecount - (1440 * minutes) - (86400 * hours)) / 24) % 3600)
	frames := (framecount - (24 * seconds) - (1440 * minutes) - (86400 * hours)) % 24

	return formatTimeCodeString(days, hours, minutes, seconds, frames, false)
}

// ticks27MhzToSmpte12M24Fps Converts to SMPTE 12M 24fps.
func ticks27MhzToSmpte12M24Fps(ticks27Mhz int64) string {
	pcrTb := ticks27MhzToPcrTb(ticks27Mhz)
	framecount := (int32)(pcrTb / 3750)

	days := int32((framecount / 86400) / 24)
	hours := int32((framecount / 86400) % 24)
	minutes := int32(((framecount - (86400 * hours)) / 1440) % 60)
	seconds := int32(((framecount - (1440 * minutes) - (86400 * hours)) / 24) % 3600)
	frames := (framecount - (24 * seconds) - (1440 * minutes) - (86400 * hours)) % 24

	return formatTimeCodeString(days, hours, minutes, seconds, frames, false)
}

// ticks27MhzToSmpte12M25Fps Converts to SMPTE 12M 25fps.
func ticks27MhzToSmpte12M25Fps(ticks27Mhz int64) string {
	pcrTb := ticks27MhzToPcrTb(ticks27Mhz)
	framecount := (int32)(pcrTb / 3600)

	days := int32((framecount / 90000) / 24)
	hours := int32((framecount / 90000) % 24)
	minutes := int32(((framecount - (90000 * hours)) / 1500) % 60)
	seconds := int32(((framecount - (1500 * minutes) - (90000 * hours)) / 25) % 3600)
	frames := (framecount - (25 * seconds) - (1500 * minutes) - (90000 * hours)) % 25

	return formatTimeCodeString(days, hours, minutes, seconds, frames, false)
}

// ticks27MhzToSmpte12M_30fps Converts to SMPTE 12M 30fps.
func ticks27MhzToSmpte12M_30fps(ticks27Mhz int64) string {
	pcrTb := ticks27MhzToPcrTb(ticks27Mhz)
	framecount := (int32)(pcrTb / 3000)

	days := int32((framecount / 108000) / 24)
	hours := int32((framecount / 108000) % 24)
	minutes := int32(((framecount - (108000 * hours)) / 1800) % 60)
	seconds := int32(((framecount - (1800 * minutes) - (108000 * hours)) / 30) % 3600)
	frames := (framecount - (30 * seconds) - (1800 * minutes) - (108000 * hours)) % 30

	return formatTimeCodeString(days, hours, minutes, seconds, frames, false)
}
