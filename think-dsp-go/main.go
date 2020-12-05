package think_dsp_go

type Signal struct {
}

type CosSignal struct {
	freq   int
	amp    float64
	offset int
}
type SinSignal struct {
	freq   int
	amp    float64
	offset int
}

func (s SinSignal) add(sig CosSignal) SumSignal {
	return SumSignal{}
}

type SumSignal struct {
}

//duration is the length of the Wave in seconds.
//start is the start time, also in seconds.
//framerate is the (integer) number of frames per second, which is also the number of samples per second.
func (s SumSignal) makeWave(duration float64, start int, frameRate int) Wave {
	return Wave{}
}

/*
cos_sig = thinkdsp.CosSignal(freq=440, amp=1.0, offset=0)
sin_sig = thinkdsp.SinSignal(freq=880, amp=0.5, offset=0)
*/
