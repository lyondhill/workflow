This is a simple temporal usage as a DAG based workflow engine. 

Running:
get termporal up and running
```
➜  ~ git clone https://github.com/temporalio/docker-compose.git
cd  docker-compose
docker-compose up
```

compile the worker and the starter:
```
go build ./cmd/worker
go build ./cmd/starter
```

start the worker:
`./worker`

run the starter:
`./starter -dagFile simpleDAG.json`

there are 2 more sample dag files `complexDAG.json` and `superComplexDAG.json`

