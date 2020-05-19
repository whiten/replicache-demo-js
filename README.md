# Replicache JS client testbed

go run ./
Visit http://localhost:8090/

In the Javascript console, run:

replicache.dispatch("test", "open", "", (err, resp) => { console.log(err, resp) })
replicache.dispatch("test", "openTransaction", "{}", (err, resp) => { console.log(err, resp) })
replicache.dispatch("test", "scan", '{"transactionId":1}', (err, resp) => { console.log(err, resp) })
replicache.dispatch("test", "put", '{"transactionId":1, "key": "hello", "value": "world"}', (err, resp) => { console.log(err, resp) })
replicache.dispatch("test", "commitTransaction", '{"transactionId":1}', (err, resp) => { console.log(err, resp) })
replicache.dispatch("test", "openTransaction", "{}", (err, resp) => { console.log(err, resp) })
replicache.dispatch("test", "scan", '{"transactionId":2}', (err, resp) => { console.log(err, resp) })

Refresh, and scan again:

replicache.dispatch("test", "open", "", (err, resp) => { console.log(err, resp) })
replicache.dispatch("test", "openTransaction", "{}", (err, resp) => { console.log(err, resp) })
replicache.dispatch("test", "scan", '{"transactionId":1}', (err, resp) => { console.log(err, resp) })
