package think_dsp_go

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	cosSig := CosSignal{freq: 440, amp: 1.0, offset: 0}
	sinSig := SinSignal{freq: 880, amp: 0.5, offset: 0}

	/*
		You can’t do much with a Signal until you evaluate it.
		In this context, “evaluate” means taking a sequence of points in time, ts, and computing the corresponding values of the signal, ys.
		I represent ts and ys using NumPy arrays and encapsulate them in an object called a Wave.
	*/
	mix := sinSig.add(cosSig)

	/*
	      A Wave represents a signal evaluated at a sequence of points in time.
	   Each point in time is called a frame (a term borrowed from movies and video).
	   The measurement itself is called a sample, although “frame” and “sample” are sometimes used interchangeably.
	*/

	/*
		duration is the length of the Wave in seconds. start is the start time, also in seconds.
		framerate is the (integer) number of frames per second,
		which is also the number of samples per second.
	*/
	wave := mix.makeWave(0.5, 0, 11025)
	fmt.Println(wave)
	/*
		wave.plot()
		pyplot.show()
	*/
	/*
		At freq=440 there are 220 periods in 0.5 seconds,
		so this plot would look like a solid block of color.

		To zoom in on a small number of periods, we can use segment, which copies a segment of a Wave and returns a new wave:
		period = mix.period
		segment = wave.segment(start=0, duration=period*3)
	*/

	/*
		thinkdsp provides read_wave, which reads a WAV file and returns a Wave:

		violin_wave = thinkdsp.read_wave('input.wav')
	*/

	/*
		And Wave provides write, which writes a WAV file:

		wave.write(filename='output.wav')
	*/

	/*
		thinkdsp also provides play_wave, which runs the media player as a subprocess:
		thinkdsp.play_wave(filename='output.wav', player='aplay')
	*/
}
