package data

type VoiceData struct {
	Country             string  `json:"Country"`
	Bandwidth           int     `json:"Bandwidth"`
	AverageResponse     int     `json:"AverageResponse"`
	Provider            string  `json:"Provider"`
	ConnectionStability float32 `json:"ConnectionStability"`
	TTFB                int     `json:"TTFB"`
	VoicePurity         int     `json:"VoicePurity"`
	MedianCallDuration  int     `json:"MedianCallDuration"`
}
