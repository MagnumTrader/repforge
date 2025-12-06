CREATE TABLE IF NOT EXISTS workoutexercises (
    id INTEGER PRIMARY KEY,
    workout_id INTEGER NOT NULL,
    -- this is the 
    exercise_id INTEGER NOT NULL,

    order_in_workout INTEGER,
    FOREIGN KEY (workout_id) REFERENCES workouts(id) ON DELETE CASCADE,
    FOREIGN KEY (exercise_id) REFERENCES exercises(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sets (
    id INTEGER PRIMARY KEY,
    workoutexercise_id INTEGER,

    weight REAL,
    reps INTEGER,
    FOREIGN KEY (workoutexercise_id) REFERENCES workoutexercises(id) ON DELETE CASCADE
);

INSERT INTO workoutexercises (workout_id, exercise_id, order_in_workout) VALUES (19, 1, 1);
INSERT INTO workoutexercises (workout_id, exercise_id, order_in_workout) VALUES (19, 3, 2);

INSERT INTO sets (workoutexercise_id, weight, reps) VALUES (1, 50, 8);
INSERT INTO sets (workoutexercise_id, weight, reps) VALUES (1, 60, 10);
INSERT INTO sets (workoutexercise_id, weight, reps) VALUES (1, 55, 8);

INSERT INTO sets (workoutexercise_id, weight, reps) VALUES (2, 10, 12);
INSERT INTO sets (workoutexercise_id, weight, reps) VALUES (2, 10, 12);
INSERT INTO sets (workoutexercise_id, weight, reps) VALUES (2, 10, 12);
