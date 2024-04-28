package caltrain

var stops = []CalTrainStop{
	{
		Name:         "San Francisco",
		NorthboundId: 70011,
		SouthboundId: 70012,
	},
	{
		Name:         "22nd Street",
		NorthboundId: 70021,
		SouthboundId: 70022,
	},
	{
		Name:         "Bayshore",
		NorthboundId: 70031,
		SouthboundId: 70032,
	},
	{
		Name:         "South San Francisco",
		NorthboundId: 70041,
		SouthboundId: 70042,
	},
	{
		Name:         "San Bruno",
		NorthboundId: 70051,
		SouthboundId: 70052,
	},
	{
		Name:         "Millbrae",
		NorthboundId: 70061,
		SouthboundId: 70062,
	},
	{
		Name:         "Broadway",
		NorthboundId: 70071,
		SouthboundId: 70072,
	},
	{
		Name:         "Burlingame",
		NorthboundId: 70081,
		SouthboundId: 70082,
	},
	{
		Name:         "San Mateo",
		NorthboundId: 70091,
		SouthboundId: 70092,
	},
	{
		Name:         "Hayward Park",
		NorthboundId: 70101,
		SouthboundId: 70102,
	},
	{
		Name:         "Hillsdale",
		NorthboundId: 70111,
		SouthboundId: 70112,
	},
	{
		Name:         "Belmont",
		NorthboundId: 70121,
		SouthboundId: 70122,
	},
	{
		Name:         "San Carlos",
		NorthboundId: 70131,
		SouthboundId: 70132,
	},
	{
		Name:         "Redwood City",
		NorthboundId: 70141,
		SouthboundId: 70142,
	},
	{
		Name:         "Menlo Park",
		NorthboundId: 70161,
		SouthboundId: 70162,
	},
	{
		Name:         "Palo Alto",
		NorthboundId: 70171,
		SouthboundId: 70172,
	},
	{
		Name:         "California Avenue",
		NorthboundId: 70191,
		SouthboundId: 70192,
	},
	{
		Name:         "San Antonio",
		NorthboundId: 70201,
		SouthboundId: 70202,
	},
	{
		Name:         "Mountain View",
		NorthboundId: 70211,
		SouthboundId: 70212,
	},
	{
		Name:         "Sunnyvale",
		NorthboundId: 70221,
		SouthboundId: 70222,
	},
	{
		Name:         "Lawrence",
		NorthboundId: 70231,
		SouthboundId: 70232,
	},
	{
		Name:         "Santa Clara",
		NorthboundId: 70241,
		SouthboundId: 70242,
	},
	{
		Name:         "College Park",
		NorthboundId: 70251,
		SouthboundId: 70252,
	},
	{
		Name:         "San Jose Diridon",
		NorthboundId: 70261,
		SouthboundId: 70262,
	},
	{
		Name:         "Tamien",
		NorthboundId: 70271,
		SouthboundId: 70272,
	},
	{
		Name:         "Capitol",
		NorthboundId: 70281,
		SouthboundId: 70282,
	},
	{
		Name:         "Blossom Hill",
		NorthboundId: 70291,
		SouthboundId: 70292,
	},
	{
		Name:         "Morgan Hill",
		NorthboundId: 77777,
		SouthboundId: 77777,
	},
	{
		Name:         "San Martin",
		NorthboundId: 70301,
		SouthboundId: 70302,
	},
	{
		Name:         "Gilroy",
		NorthboundId: 70311,
		SouthboundId: 70312,
	},
}

func GetStopById(id int) CalTrainStop {
	var caltrainStop = CalTrainStop{}
	for _, stop := range stops {
		if stop.NorthboundId == id {
			caltrainStop = stop
			break
		} else if stop.SouthboundId == id {
			caltrainStop = stop
			break
		}
	}

	return caltrainStop
}

func GetStopByName(name string) CalTrainStop {
	var caltrainStop = CalTrainStop{}
	for _, stop := range stops {
		if stop.Name == name {
			caltrainStop = stop
			break
		}
	}
	return caltrainStop
}
