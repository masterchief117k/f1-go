package parser

import ()

type WheelState struct {
	SurfaceTemp   float64
	CarcassTemp   float64
	InnerPressure float64
}
type CornerMetrics struct {
	EntrySpeed float64
	ApexSpeed  float64
}
type DriverInfo struct {
	DriverNumber int
	DriverCode   string
	TeamName     string
}
type TyreState struct {
	Compound   string
	TyreAge    int
	FrontLeft  WheelState
	FrontRight WheelState
	RearLeft   WheelState
	RearRight  WheelState
}
type WeatherState struct {
	AirTemp       float64
	TrackTemp     float64
	WindSpeed     float64
	WindDirection int
	IsRaining     bool
	Rainfall      float64
	AirPressure   float64
	Humidity      float64
}
type TimingData struct {
	TotalTime     *float64
	Sector1       *float64
	Sector2       *float64
	Sector3       *float64
	IsDeleted     bool
	PitStopStatus string
	Corners       map[int]CornerMetrics
}
type LapRecord struct {
	LapNumber int
	Driver    DriverInfo
	Tyres     TyreState
	Weather   WeatherState
	Timing    TimingData
}
