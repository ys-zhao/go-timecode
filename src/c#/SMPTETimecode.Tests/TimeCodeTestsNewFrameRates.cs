using Microsoft.VisualStudio.TestTools.UnitTesting;
using System;
using System.Globalization;

namespace SMPTETimecode.Tests
{
    [TestClass]
    public class TimeCodeTestsNewFrameRates
    {
        [TestMethod]
        public void Test_Add_Frames_To_Existing_TimeCode_Smpte30_Ensure_Correct_Time()
        {
            var currentTimeCode = TimeCode.FromSeconds(0M, SmpteFrameRate.Smpte30);
            var oneFrame = TimeCode.FromFrames(1, SmpteFrameRate.Smpte30);
            var expectedFrameCount = currentTimeCode.TotalFrames;

            Assert.AreEqual(1, oneFrame.TotalFrames);

            for (var i = 1; i < 50000; i++)
            {
                expectedFrameCount++;
                currentTimeCode = currentTimeCode.Add(oneFrame);

                Assert.AreEqual(
                    expectedFrameCount,
                    currentTimeCode.TotalFrames,
                    $"Expected frames {expectedFrameCount} did not match TotalFrames {currentTimeCode.TotalFrames}"
                );
            }
        }

        [TestMethod]
        public void Test_Add_Frames_To_Existing_TimeCode_Smpte50_Ensure_Correct_Time()
        {
            var currentTimeCode = TimeCode.FromSeconds(0M, SmpteFrameRate.Smpte50);
            var oneFrame = TimeCode.FromFrames(1, SmpteFrameRate.Smpte50);
            var expectedFrameCount = currentTimeCode.TotalFrames;

            Assert.AreEqual(1, oneFrame.TotalFrames);

            for (var i = 1; i < 50000; i++)
            {
                expectedFrameCount++;
                currentTimeCode = currentTimeCode.Add(oneFrame);

                Assert.AreEqual(
                    expectedFrameCount,
                    currentTimeCode.TotalFrames,
                    $"Expected frames {expectedFrameCount} did not match TotalFrames {currentTimeCode.TotalFrames}"
                );
            }
        }

        [TestMethod]
        public void Test_Add_Frames_To_Existing_TimeCode_Smpte5994Drop_Ensure_Correct_Time()
        {
            var currentTimeCode = TimeCode.FromSeconds(0M, SmpteFrameRate.Smpte5994Drop);
            var oneFrame = TimeCode.FromFrames(1, SmpteFrameRate.Smpte5994Drop);
            var expectedFrameCount = currentTimeCode.TotalFrames;

            Assert.AreEqual(1, oneFrame.TotalFrames);

            for (var i = 1; i < 50000; i++)
            {
                expectedFrameCount++;
                currentTimeCode = currentTimeCode.Add(oneFrame);

                Assert.AreEqual(
                    expectedFrameCount,
                    currentTimeCode.TotalFrames,
                    $"Expected frames {expectedFrameCount} did not match TotalFrames {currentTimeCode.TotalFrames}"
                );
            }
        }

        [TestMethod]
        public void Test_Add_Frames_To_Existing_TimeCode_Smpte5994NonDrop_Ensure_Correct_Time()
        {
            var currentTimeCode = TimeCode.FromSeconds(0M, SmpteFrameRate.Smpte5994NonDrop);
            var oneFrame = TimeCode.FromFrames(1, SmpteFrameRate.Smpte5994NonDrop);
            var expectedFrameCount = currentTimeCode.TotalFrames;

            Assert.AreEqual(1, oneFrame.TotalFrames);

            for (var i = 1; i < 50000; i++)
            {
                expectedFrameCount++;
                currentTimeCode = currentTimeCode.Add(oneFrame);

                Assert.AreEqual(
                    expectedFrameCount,
                    currentTimeCode.TotalFrames,
                    $"Expected frames {expectedFrameCount} did not match TotalFrames {currentTimeCode.TotalFrames}"
                );
            }
        }

        [TestMethod]
        public void Test_Add_Frames_To_Existing_TimeCode_Smpte60_Ensure_Correct_Time()
        {
            var currentTimeCode = TimeCode.FromSeconds(0M, SmpteFrameRate.Smpte60);
            var oneFrame = TimeCode.FromFrames(1, SmpteFrameRate.Smpte60);
            var expectedFrameCount = currentTimeCode.TotalFrames;

            Assert.AreEqual(1, oneFrame.TotalFrames);

            for (var i = 1; i < 50000; i++)
            {
                expectedFrameCount++;
                currentTimeCode = currentTimeCode.Add(oneFrame);

                Assert.AreEqual(
                    expectedFrameCount,
                    currentTimeCode.TotalFrames,
                    $"Expected frames {expectedFrameCount} did not match TotalFrames {currentTimeCode.TotalFrames}"
                );
            }
        }

        [TestMethod]
        public void Test_Add_Frames_To_Existing_TimeCode_Smpte96_Ensure_Correct_Time()
        {
            var currentTimeCode = TimeCode.FromSeconds(0M, SmpteFrameRate.Smpte96);
            var oneFrame = TimeCode.FromFrames(1, SmpteFrameRate.Smpte96);
            var expectedFrameCount = currentTimeCode.TotalFrames;

            Assert.AreEqual(1, oneFrame.TotalFrames);

            for (var i = 1; i < 50000; i++)
            {
                expectedFrameCount++;
                currentTimeCode = currentTimeCode.Add(oneFrame);

                Assert.AreEqual(
                    expectedFrameCount,
                    currentTimeCode.TotalFrames,
                    $"Expected frames {expectedFrameCount} did not match TotalFrames {currentTimeCode.TotalFrames}"
                );
            }
        }

        [TestMethod]
        public void Test_Add_Frames_To_Existing_TimeCode_Smpte100_Ensure_Correct_Time()
        {
            var currentTimeCode = TimeCode.FromSeconds(0M, SmpteFrameRate.Smpte100);
            var oneFrame = TimeCode.FromFrames(1, SmpteFrameRate.Smpte100);
            var expectedFrameCount = currentTimeCode.TotalFrames;

            Assert.AreEqual(1, oneFrame.TotalFrames);

            for (var i = 1; i < 50000; i++)
            {
                expectedFrameCount++;
                currentTimeCode = currentTimeCode.Add(oneFrame);

                Assert.AreEqual(
                    expectedFrameCount,
                    currentTimeCode.TotalFrames,
                    $"Expected frames {expectedFrameCount} did not match TotalFrames {currentTimeCode.TotalFrames}"
                );
            }
        }

        [TestMethod]
        public void Test_Add_Frames_To_Existing_TimeCode_Smpte120_Ensure_Correct_Time()
        {
            var currentTimeCode = TimeCode.FromSeconds(0M, SmpteFrameRate.Smpte120);
            var oneFrame = TimeCode.FromFrames(1, SmpteFrameRate.Smpte120);
            var expectedFrameCount = currentTimeCode.TotalFrames;

            Assert.AreEqual(1, oneFrame.TotalFrames);

            for (var i = 1; i < 50000; i++)
            {
                expectedFrameCount++;
                currentTimeCode = currentTimeCode = currentTimeCode.Add(oneFrame);

                Assert.AreEqual(
                    expectedFrameCount,
                    currentTimeCode.TotalFrames,
                    $"Expected frames {expectedFrameCount} did not match TotalFrames {currentTimeCode.TotalFrames}"
                );
            }
        }

    } // class

} // namespace
