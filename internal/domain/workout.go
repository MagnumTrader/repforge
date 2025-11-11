package domain

type Workout struct {
	Date     string
	Type     string
	Duration int
	Notes    string
}

var Workouts = []Workout{
	{
		Date:     "2025-11-10",
		Type:     "Running",
		Duration: 30,
		Notes:    "Morning jog in the park",
	},
	{
		Date:     "2025-11-09",
		Type:     "Cycling",
		Duration: 45,
		Notes:    "Evening ride with friends",
	},
	{
		Date:     "2025-11-08",
		Type:     "Yoga",
		Duration: 60,
		Notes:    "Relaxing session at home",
	},
}

