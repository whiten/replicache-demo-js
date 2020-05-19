# Replicache JS client testbed


### Start the server:

In a checkout of https://github.com/rocicorp/replicache-client, run `make build/repjs`, and copy the result to the `web/` directory in this repo.

`go run ./`

Visit http://localhost:8090/

### In the Javascript console, run:

replicache.dispatch("test", "open", "", (err, resp) => { console.log(err, resp) })

replicache.dispatch("test", "openTransaction", "{}", (err, resp) => { console.log(err, resp) })

replicache.dispatch("test", "scan", '{"transactionId":1}', (err, resp) => { console.log(err, resp) })

replicache.dispatch("test", "put", '{"transactionId":1, "key": "hello", "value": "world"}', (err, resp) => { console.log(err, resp) })

replicache.dispatch("test", "commitTransaction", '{"transactionId":1}', (err, resp) => { console.log(err, resp) })

replicache.dispatch("test", "openTransaction", "{}", (err, resp) => { console.log(err, resp) })

replicache.dispatch("test", "scan", '{"transactionId":2}', (err, resp) => { console.log(err, resp) })

### Refresh, and scan again:

replicache.dispatch("test", "open", "", (err, resp) => { console.log(err, resp) })

replicache.dispatch("test", "openTransaction", "{}", (err, resp) => { console.log(err, resp) })

replicache.dispatch("test", "scan", '{"transactionId":1}', (err, resp) => { console.log(err, resp) })
