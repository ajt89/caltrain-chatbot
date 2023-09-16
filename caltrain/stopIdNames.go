package caltrain

var stops = []CalTrainStop{
	CalTrainStop{
		Name:         "San Francisco",
		NorthboundId: 70011,
		SouthboundId: 70012,
	},
	CalTrainStop{
		Name:         "22nd Street",
		NorthboundId: 70021,
		SouthboundId: 70022,
	},
	CalTrainStop{
		Name:         "Bayshore",
		NorthboundId: 70031,
		SouthboundId: 70032,
	},
	CalTrainStop{
		Name:         "South San Francisco",
		NorthboundId: 70041,
		SouthboundId: 70042,
	},
	CalTrainStop{
		Name:         "San Bruno",
		NorthboundId: 70051,
		SouthboundId: 70052,
	},
	CalTrainStop{
		Name:         "Millbrae",
		NorthboundId: 70061,
		SouthboundId: 70062,
	},
	CalTrainStop{
		Name:         "Broadway",
		NorthboundId: 70071,
		SouthboundId: 70072,
	},
	CalTrainStop{
		Name:         "Burlingame",
		NorthboundId: 70081,
		SouthboundId: 70082,
	},
	CalTrainStop{
		Name:         "San Mateo",
		NorthboundId: 70091,
		SouthboundId: 70092,
	},
	CalTrainStop{
		Name:         "Hayward Park",
		NorthboundId: 70101,
		SouthboundId: 70102,
	},
	CalTrainStop{
		Name:         "Hillsdale",
		NorthboundId: 70111,
		SouthboundId: 70112,
	},
	CalTrainStop{
		Name:         "Belmont",
		NorthboundId: 70121,
		SouthboundId: 70122,
	},
	CalTrainStop{
		Name:         "San Carlos",
		NorthboundId: 70131,
		SouthboundId: 70132,
	},
	CalTrainStop{
		Name:         "Redwood City",
		NorthboundId: 70141,
		SouthboundId: 70142,
	},
	CalTrainStop{
		Name:         "Menlo Park",
		NorthboundId: 70161,
		SouthboundId: 70162,
	},
	CalTrainStop{
		Name:         "Palo Alto",
		NorthboundId: 70171,
		SouthboundId: 70172,
	},
	CalTrainStop{
		Name:         "California Avenue",
		NorthboundId: 70191,
		SouthboundId: 70192,
	},
	CalTrainStop{
		Name:         "San Antonio",
		NorthboundId: 70201,
		SouthboundId: 70202,
	},
	CalTrainStop{
		Name:         "Mountain View",
		NorthboundId: 70211,
		SouthboundId: 70212,
	},
	CalTrainStop{
		Name:         "Sunnyvale",
		NorthboundId: 70221,
		SouthboundId: 70222,
	},
	CalTrainStop{
		Name:         "Lawrence",
		NorthboundId: 70231,
		SouthboundId: 70232,
	},
	CalTrainStop{
		Name:         "Santa Clara",
		NorthboundId: 70241,
		SouthboundId: 70242,
	},
	CalTrainStop{
		Name:         "College Park",
		NorthboundId: 70251,
		SouthboundId: 70252,
	},
	CalTrainStop{
		Name:         "San Jose Diridon",
		NorthboundId: 70261,
		SouthboundId: 70262,
	},
	CalTrainStop{
		Name:         "Tamien",
		NorthboundId: 70271,
		SouthboundId: 70272,
	},
	CalTrainStop{
		Name:         "Capitol",
		NorthboundId: 70281,
		SouthboundId: 70282,
	},
	CalTrainStop{
		Name:         "Blossom Hill",
		NorthboundId: 70291,
		SouthboundId: 70292,
	},
	CalTrainStop{
		Name:         "Morgan Hill",
		NorthboundId: 77777,
		SouthboundId: 77777,
	},
	CalTrainStop{
		Name:         "San Martin",
		NorthboundId: 70301,
		SouthboundId: 70302,
	},
	CalTrainStop{
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
