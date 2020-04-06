using System;
using SMPTETimecode;

namespace addsub
{
    class Program
    {
        static void Main(string[] args)
        {
            var t1 = TimeCode.FromMinutes(10.0, SmpteFrameRate.Smpte2997NonDrop);
            var t2 = TimeCode.FromMinutes(10.0, SmpteFrameRate.Smpte30);
            Console.WriteLine("TimeCode: {0}, {1}", t1.ToString(), t2.ToString());
        }
    }
}
