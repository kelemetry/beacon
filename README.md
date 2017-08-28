# Beacon

This code is in alpha and is under active development.

Beacon is a monitoring cli out-of-clouster or in-cluster controller that will broadcast changes to a selected resource type

A beacon is a combination of a Kubernetes resource, a Transport method and a Signal.

For instance you can set up a beacon to listen to changes to Pod resoures and then using the Stdout transport, send a signal that is just the name.

| K8s Resources | Keywords |
|---|---|
| Deployments | de, deploy, deployment, deployments |
| Ingresses | in, ing, ingresse, ingresses |
| Pods | po, pod, pods|
| Services | se, svc, service, services|

-transport=<keyword>
| Transports | Keywords |
|---|---|
| Stdout | stdout |
| NATS | nats |

-signal
| Signal | Keywords |
|---|---|
| Status | status |

## Running

```
# if outside of the cluster
go run *.go -kubeconfig=$HOME/.kube/config -logtostderr=true -kind=pods -transport=nats -namespace=default
```

## Roadmap
   1. Add Complete set of Resources
   2. Add multiple transports
   3. Add multiple signals (name only, name + service, full)