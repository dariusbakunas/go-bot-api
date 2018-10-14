## go-bot-api - Rest API for controlling SSC-32U

### Local dev:

* Install dependencies:

    ```% go get ./...```
    
* Run:

    ```% go run main.go```
    
### Endpoints:

* `/servos/:id?angle=ANGLE` id - servo/pin number on SSC board and angle is between 0 and 180

Current pulse range is hardcoded for HS-422 servo