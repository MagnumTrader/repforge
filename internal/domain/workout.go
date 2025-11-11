package domain

type Workout struct {
	Id       int
	Date     string
	Type     string
	Duration int
	Notes    string
}

var Workouts = []Workout{
	{
		Id:       1,
		Date:     "2025-11-10",
		Type:     "Running",
		Duration: 30,
		Notes:    "Morning jog in the park",
	},
	{
		Id:       2,
		Date:     "2025-11-09",
		Type:     "Cycling",
		Duration: 45,
		Notes:    "Evening ride with friends",
	},
	{
		Id:       3,
		Date:     "2026-11-08",
		Type:     "Yoga",
		Duration: 60,
		Notes:    "Relaxing session at home",
	},
}
