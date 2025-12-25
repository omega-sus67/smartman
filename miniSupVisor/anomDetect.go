package main

import(
	"math"
)

// a struct that records history of a program
type anomalyDetector struct{
	Windowsize int
	History []float64
	Threshold float64
}
// self explanatory
func addData(window int, threshold float64) *anomalyDetector{
	return &anomalyDetector{
		Windowsize: window,
		History: make([]float64, 0, window),
		Threshold: threshold,
	}
}
// updates the history and detects the anomaly, sends the error
func (ad *anomalyDetector) addDataAndCheck(latency float64) (bool, string){
	ad.History = append(ad.History, latency)

	if len(ad.History) < ad.Windowsize{
		return false, "not enough data"
	}
	if len(ad.History) > ad.Windowsize{
		ad.History = ad.History[1:]
	}

	mean := ad.mean()
	stdDev := ad.stdDev(mean)

	if stdDev == 0{
		return false, "no anomaly detected"
	}

	zScore := (latency - mean)/stdDev

	if zScore > ad.Threshold{
		return true, "High anomaly"
	}
	if zScore < -ad.Threshold{
		return true, "Low anomaly"
	}

	return false, "Normal"
}
// helper functions
func (ad *anomalyDetector) mean () float64{
	sum := 0.0

	for _, s := range ad.History{
		sum += s
	}

	return sum/ad.Threshold
}

func (ad *anomalyDetector) stdDev (mean float64) float64{
	sqsum := 0.0

	for _, s := range ad.History{
		diff := mean - s
		sqsum += diff * diff
	}

	return math.Sqrt(sqsum/float64(ad.Windowsize))
}

