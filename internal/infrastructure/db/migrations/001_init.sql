CREATE TABLE IF NOT EXISTS workouts 
(id INTEGER PRIMARY KEY, date string, type string, duration INTEGER, notes string);

CREATE TABLE IF NOT EXISTS exercises
(id INTEGER PRIMARY KEY, name string NOT NULL, category string)

