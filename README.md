## go-bot-api - Rest API for controlling SSC-32U

### Local dev:

* Install dependencies:

    ```% go get ./...```
    
* Run:

    ```% go run main.go```
    
### Endpoints:

* `/servos/:id?angle=[angle]&time=[time]` id - servo/pin number on SSC board, angle is between 0 and 180 and time -  time in microseconds to travel from the current position to the desired position (max: 65525)

Current pulse range is hardcoded for HS-422 servo