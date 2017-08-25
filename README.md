# Beacon

Beacon is a monitoring cli out-of-clouster or in-cluster controller that will broadcast changes to a selected resource type

A beacon is a combination of a Kubernetes resource, a Transport method and a Signal.

For instance you can set up a beacon to listen to changes to Pod resoures and then using the Stdout transport, send a signal that is just the name.


## Running

```
# if outside of the cluster
go run *.go -kubeconfig=$HOME/.kube/config -logtostderr=true -kind=pods -namespace=default
```

## Current Status
### Supported Resoures
-kind=

deployment
ingresses
pods
services


### Supported Transports
-transport=

stdout

## Roadmap
   1. Add Complete set of Resources
   2. Add multiple transports
   3. Add multiple signals (name only, name + service, full)