package caltrain

type CalTrainStopsStatus struct {
	ErrorMsg      string
	Status        int
	CalTrainStops []CalTrainStop
}

type CalTrainStop struct {
	Name                     string
	NorthboundId             int
	SouthboundId             int
	NorthBoundTrainsToArrive []CalTrainVehicle
	SouthBoundTrainsToArrive []CalTrainVehicle
}

type CalTrainVehicleStatus struct {
	ErrorMsg         string
	Status           int
	CalTrainVehicles []CalTrainVehicle
}

type CalTrainVehicle struct {
	Id            string
	ArrivalTime   int64
	DepartureTime int64
	StopsLeft     int
	CurrentStop   string
	TripType      string
	Direction     string
}

type CalTrainTripStatus struct {
	ErrorMsg      string
	Status        int
	CalTrainTrips []CalTrainTrip
}

type CalTrainTrip struct {
	Id          string
	RouteId     string
	DirectionId int
	Stops       []Stop
}

type Stop struct {
	Id        int
	Arrival   int64
	Departure int64
}

type RequestStatus struct {
	ErrorMsg string
	Status   int
	Data     []byte
}

type CalTrainDecodeStatus struct {
	ErrorMsg string
	Status   int
	RealTime RealTime
}

type RealTimeStatus struct {
	ErrorMsg string
	Status   int
	RealTime RealTime
}

type RealTime struct {
	Header   Header
	Entities []Entity
}

type Header struct {
	GtfsRealtimeVersion string `json:"GtfsRealtimeVersion"`
	Incrementality      int    `json:"incrementality"`
	Timestamp           int64  `json:"Timestamp"`
}

type Entity struct {
	Id         string `json:"Id"`
	TripUpdate TripUpdate
}

type TripUpdate struct {
	Trip           Trip
	Vehicle        Vehicle
	StopTimeUpdate []StopTimeUpdate
	Timestamp      int64 `json:"Timestamp"`
}

type Trip struct {
	Id          string `json:"TripId"`
	RouteId     string `json:"RouteId"`
	DirectionId int    `json:"DirectionId"`
}

type Vehicle struct {
	Id           string `json:"Id"`
	Label        string `json:"Label"`
	LicensePlate string `json:"LicensePlate"`
}

type StopTimeUpdate struct {
	Id        string `json:"StopId"`
	Arrival   ArrivalDeparture
	Departure ArrivalDeparture
}

type ArrivalDeparture struct {
	Time int64 `json:"Time"`
}
