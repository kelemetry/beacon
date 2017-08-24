# Beacon

Beacon is a monitoring cli out-of-clouster or in-cluster controller that will broadcast changes to a selected resource type

## Running

```
# if outside of the cluster
go run *.go -kubeconfig=$HOME/.kube/config -logtostderr=true -kind=pods -namespace=default
```

## Current Status
### Supported Resoures
-kind=ingresses
-kind=pods
-kind=services

## Roadmap
   1. Add Complete set of Resources
   2. Add multiple transports
   3. Add multiple signals (name only, name + service, full)