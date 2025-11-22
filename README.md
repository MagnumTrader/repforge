# repforge

TODO:
I was actually thinking about something the other day

but what was it,


lets go back to basics?

- [ ] Add edit workout screen ( same as new but populated with the current data)
      Document the process of how we best do it
      
      1. Handler should have function for workouts/:id/edit
      2. UI Template for the service to render the overlay
      3. enter data
      4. Update request is sent to workouts/:id
      5. collect that data and push it throught the service
      6. write to db
      7. return success

- [ ] Add html content type function and return them from functions
- [ ] Add tests for service..
- [ ] More fields on workout, what is the flow of expanding the app?
- [ ] Nicer font :) lets go with ubuntu font
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
