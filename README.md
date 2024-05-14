# golang-basics-finalproject
A project for udacity's Go course. This is an ultra simple API that passed the test.
Is not intended to be used for anything beyond Udacity's GO lang course, and is expected to pass the minimum requirements only.

# Explanation of folders
/project is the where the main project is. Udacity evaluators shoulds use that folder. /simple can be ignored for Udacity Purposes

## How To Run
This project is intended to to be run in a docker and requires the environment variable API_PORT set to 3000.

Assuming that Docker is installed and running, do the following:

1. Build the Docker image in /project with "docker run build -t udacitytag/joelsapi . " (or whatever tags you choose)
2. Alternately, download the docker image from https://hub.docker.com/r/gonzaga626/apigo
3. Run the Docker image with the following paramaters in CLI 

docker run --env=API_PORT=3000 -p 3000:3000 -d udacitytag/joelsapi:latest

After that the docker container should run. API tests can proceed from there, on http://localhost:3000