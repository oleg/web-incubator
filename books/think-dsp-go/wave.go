package think_dsp_go

type Wave struct {
	Ys        []int     //wave array
	Ts        []float64 //array of times
	FrameRate int       //framerate: samples per second

	/*
	   def __init__(self, ys, ts=None, framerate=None):
	       """Initializes the wave.

	       ys: wave array
	       ts: array of times
	       framerate: samples per second
	       """
	       self.ys = np.asanyarray(ys)
	       self.framerate = framerate if framerate is not None else 11025

	       if ts is None:
	           self.ts = np.arange(len(ys)) / self.framerate
	       else:
	           self.ts = np.asanyarray(ts)
	*/
}

func MakeWave(ys []int, framerate int) Wave {
	max := len(ys)

	ts := make([]float64, max)
	for i := 0; i < max; i++ {
		ts[i] = float64(i) / float64(framerate)
	}
	return Wave{Ys: ys, Ts: ts, FrameRate: framerate}
}
