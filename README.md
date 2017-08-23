# Beacon

Beacon is a monitoring cli or incluster controller that will broadcast changes to a selected resource type

## Running

```
# if outside of the cluster
go run *.go -kubeconfig=$HOME/.kube/config -logtostderr=true -kind=pods -namespace=default
```

## Current Status
### Supported Resoures
#### Pods
-kind=pods

#### Services
-kind=services

## Roadmap
   1. Add Complete set of Resources
   2. Add multiple transports
   3. Add multiple signals (name only, name + service, full)