# Sum Up Coding Task
This project is an HTTP task processing app written in Go 1.15. This application sorts tasks based on their prerequisites to create a proper execution order.

## Technologies used

1. Go 1.15
2. Docker

## Run app outside of Docker

In order to run the go app navigate to src/ and then run:
```
go build -o main 
./main
```
This will compile and run the main file.
## Run app in Docker

This project contains a `Dockerfile` so it could be run as a container. In order to do that first
an image has to be built:
```bash
docker build -t task-sorter .
```
After the command finishes the container can be started with:
```bash
docker run -p 4000:4000 task-sorter
```

## Testing endpoints
There are two endpoints available:
```bigquery
/tasks/json -> to receive json representation of the sorted tasks
/tasks/bash -> to receive a sequence of the sorted bash commands of the tasks
```
You would need to make a POST request and provide a list of tasks in the body, like so:
```bigquery
{
    "tasks":[
        {
            "name":"task-1",
            "command":"touch /tmp/file1"
        },
        {
            "name":"task-2",
            "command":"cat /tmp/file1",
            "requires":[
                "task-3"
            ]
        },
        {
            "name":"task-3",
            "command":"echo 'Hello World!' > /tmp/file1",
            "requires":[
                "task-1"
            ]
        },
        {
            "name":"task-4",
            "command":"rm /tmp/file1",
            "requires":[
                "task-2",
                "task-3"
            ]
        }
    ]
}
```
To hit the endpoint you can paste this into the terminal
 ```bash
curl --location --request POST 'localhost:4000/tasks/json' \
--header 'Content-Type: application/json' \
--data-raw '{"tasks":[{"name":"task-1","command":"touch /tmp/file1"},{"name":"task-2","command":"cat /tmp/file1","requires":["task-3"]},{"name":"task-3","command":"echo '\''Hello World!'\'' > /tmp/file1","requires":["task-1"]},{"name":"task-4","command":"rm /tmp/file1","requires":["task-2","task-3"]}]}'
 ```

## Run tests
To run the tests, navigate to src/ and run `go test`

## Kubernetes

The application is K8s deployable, deployment is configured in `k8s-deployment.yaml` To deploy the service you can follow guide here: https://www.callicoder.com/deploy-containerized-go-app-kubernetes/#creating-a-kubernetes-deployment
