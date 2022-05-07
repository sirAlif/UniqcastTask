# UniqcastTask
My solution for Uniqcast backend task

## Running the app

```bash
# start the nats server
$ docker-compose up -d

# run the service
$ cd mp4Service && go run .

# run the app (another terminal)
$ cd app && npm install && npm start
