// <copyright file="TimeCode.cs" company="Microsoft Corporation">
// ===============================================================================
//
//
// Project: SMPTE 12M Timecode
// FILES: TimeCode.cs                     
//
// ===============================================================================
// Copyright 2010 Microsoft Corporation.  All rights reserved.
// THIS CODE AND INFORMATION IS PROVIDED "AS IS" WITHOUT WARRANTY
// OF ANY KIND, EITHER EXPRESSED OR IMPLIED, INCLUDING BUT NOT
// LIMITED TO THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
// FITNESS FOR A PARTICULAR PURPOSE.
// ===============================================================================
// </copyright>

namespace SMPTETimecode
{
    public partial struct TimeCode
    {
        private const string ArgumentException = "Expected an instance of TimeCode.";
        private const string MaxValueSmpte12MOverflowException = "The resulting timecode {0} is out of the expected range of MaxValue {1}.";
        private const string MinValueSmpte12MOverflowException = "The resulting timecode is out of the expected range of MinValue.";
        private const string Smpte12MBadFormat = "The timecode provided is not in the correct format.";
        private const string Smpte12MOutOfRange = "The timecode provided is out of the expected range.";
        private const string Smpte12M_2398_BadFormat = "Timecode frame value is not in the expected range for SMPTE 23.98 IVTC.";
        private const string Smpte12M_24_BadFormat = "Timecode frame value is not in the expected range for SMPTE 24fps Film Sync.";
        private const string Smpte12M_25_BadFormat = "Timecode frame value is not in the expected range for SMPTE 25fps PAL.";
        private const string Smpte12M_2997_Drop_BadFormat = "Timecode frame value is not in the expected range for SMPTE 29.97 DropFrame.";
        private const string Smpte12M_2997_NonDrop_BadFormat = "Timecode frame value is not in the expected range for SMTPE 29.97 NonDrop.";
        private const string Smpte12M_30_BadFormat = "Timecode frame value is not in the expected range for SMPTE 30fps.";
    }
}