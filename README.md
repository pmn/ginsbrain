# Ginsbrain

A brain for Ginsbot!

### Running
You'll need to add the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables to use this (this uses S3 for persistence)

* Clone this repository
* Go into the directory and run `go deps`
* Run `go build && ./ginsbrain` to run it


### testing
There are no tests. Good luck and godspeed. 

#### Add a memory

curl -X POST -H "Content-Type: application/json" -d '{"text": "cooooooool"}' http://localhost:8080/memories
