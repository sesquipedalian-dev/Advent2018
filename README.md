# Advent of Code 2018
Implementation of the challenges in Advent of Code 2018 (https://adventofcode.com/2018), as a go http server.

## Running
* go build
* advent.exe -addr :<port> - run server on the provided port (default to 1718)
* POST to /<n>, where n is the number of the challenge you wish to check.  Put the challenge input in the post body as plain text. N starts at 1. Puzzle output returned as plain text.
* GET / returns a HTML page summarizing the app. 

## Organization
* separate file per challenge / route