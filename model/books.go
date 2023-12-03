package model

type ReqGetBooks struct {
	Subjects string
	Limit    int
	Offset   int
}

type RespLibraryApi struct {
	Works []struct {
		Key     string `json:"key"`
		Title   string `json:"title"`
		Authors []struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"authors"`
		EditionCount int `json:"edition_count"`
	} `json:"works"`
}

type RetData struct {
	Key            string `json:"key"`
	Title          string `json:"title"`
	Authors        string `json:"authors"`
	EditionNumber  int    `json:"edition_number"`
	PickupSchedule string `json:"pickup_schedule,omitempty"`
}

type ReqPickup struct {
	Key        string `json:"key"`
	PickupTime string `json:"pickup_time"`
}
