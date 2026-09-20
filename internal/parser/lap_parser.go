package parser
import (
	"encoding/json"
	)
type rawLapPayload struct {
	LapNumber int `json:"lap_number"`
	DriverNumber int    `json:"driver_number"`
	DriverCode   string `json:"driver_code"`
	TeamName     string `json:"team_name"`
	LapTime      string `json:"lap_time"`
	Sector1Time  string `json:"sector1_time"`
	Sector2Time  string `json:"sector2_time"`
	Sector3Time  string `json:"sector3_time"`
	IsDeleted    bool   `json:"is_deleted"`
	Compound     string `json:"compound"`
	TyreAge      int    `json:"tyre_age"`
}
func ParseLapJSON(data []byte) (*LapRecord , error){
	var raw rawLapPayload 
	err := json.Unmarshal(data, &raw) 
	if err!= nil {return nil , err }
	totalTime , err := ParseLapTimeString(raw.LapTime)
	if err!= nil {return nil , err }
	s1, err:= ParseLapTimeString(raw.Sector1Time)
	if err!= nil {return nil , err }
	s2,  err := ParseLapTimeString(raw.Sector2Time)
	if err!= nil {return nil , err }
	s3,  err := ParseLapTimeString(raw.Sector3Time)
	if err!= nil {return nil , err }
	record := &LapRecord{
        LapNumber: raw.LapNumber,
        Driver: DriverInfo{
        DriverNumber: raw.DriverNumber,
        DriverCode:   raw.DriverCode,
        TeamName:     raw.TeamName,
    },
    Tyres: TyreState{
        Compound: raw.Compound,
        TyreAge:  raw.TyreAge,
    },
    Timing: TimingData{
        TotalTime: totalTime,
        Sector1:   s1,
        Sector2:   s2,
        Sector3:   s3,
        IsDeleted: raw.IsDeleted,
        Corners:   make(map[int]CornerMetrics),
    },
}

return record, nil

}
