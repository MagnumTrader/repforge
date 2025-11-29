# Repforge

TODO:
Wow for adjusting an entity
- Do the change in domain Entity
- Db migrations
- Repo implementation update to fetch new fields
- Service arguments for creating/filtering if needed
- Rendering if applicable

- Domain Entity - add/modify field
- Database migration (ALTER TABLE)
- Repo - update SELECT/INSERT/UPDATE queries
- Service - add validation/defaults if needed
- Service tests - update fixtures
- Handler - parse new field from request (if user-provided)
- UI - add form input or display field


--- No footshooting guide:
- Did you run templ generate after changing .templ files? ( Air handles this )
- Did you add the new field to ALL queries (SELECT, INSERT, UPDATE)?
- Did you update in_mem.go for local testing?
- Did existing tests break? (good - they caught something!)

lets go back to basics?

# Todos
- [x] Deletion of exercises
- [ ] Exercise form
- [ ] Edit exercise
- [ ] New Exercise
- [ ] Add link to exercise
- [ ] Add db repo for exercises
- [ ] Add workoutExercises which is a subcomponent of every workout
- [ ] Add tests for service..
- [ ] Add ctx in db calls and service?
- [ ] add Sets (number of reps + weights of exercises) then have a workout that we relate that rep and weight/time to 
- [ ] Move to generic repo for entities
- [x] Exercise details
- [x] Exercise objects added
- [x] List of exercise route
- [x] Nicer font :) lets go with ubuntu font
- [x] Add edit workout screen ( same as new but populated with the current data)
      Document the process of how we best do it
      1. Handler should have function for workouts/:id/edit
      2. UI Template for the service to render the overlay
      3. enter data
      4. Update request is sent to workouts/:id
      5. collect that data and push it throught the service
      6. write to db
      7. return success
- [x] Add html content type function and return them from functions
- [x] Cleaning... 
      Think about creating a service for the handler. the handler takes a service, 
      that contains the repo and handles the business Logic,
- [x] Add delete button to list
- [x] Add delete route for workout
- [x] Add new workout screen
        This should not be a screen, i want to have a popup 
        that covers the entire screen where you can add a workout.
    - [x] Add popup when clicking add workout make the screen darker and absolute pos
    - [x] Add content styling so that we can display a form in the middle of Overlay
    - [x] add form for all the fields in a workout
    - [x] submit the form to the server, printing it
    - [x] insert it into the DB
- [x] Add sql folder structure
- [x] Add table for workouts AS IS
- [x] Add functions for creating a workout and store in db
- [x] remove notes in list view
# Ideas
- [ ] Could the system have coaches that distributes workout programs
