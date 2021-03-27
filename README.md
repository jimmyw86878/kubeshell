## Description

The `kubectl` shell pod. Users can execute `kubectl` commands in this container. This is designed for `Openshift OKD 3.11` platform.

## Prerequisites

* Golang 1.12 up
* Docker
* Make

## Spec

1. HTTP for port 5666

2. The http server just verifies that the pod is health. URL: "http://pod-ip:5666/kubeshell/v1/health"

## Bring kubeshell in Openshift

All needed yaml files already exist in `yamls` folder. Assume we bring kubeshell in `default` namespace.

1. Create serviceAccount

Default serviceAccount name will be `privilegeduser`. Feel free to change it inside `kubeshell-deployment.yml` and `role.yml`.

```
oc create serviceaccount privilegeduser

oc adm policy add-scc-to-user privileged -z privilegeduser
```

2. Create cluster role and role binding

```
kubectl create -f role.yml
```

3. Prepare kube config file for your K8S cluster 

The config will be used for connecting K8S cluster. It should be mounted in kubeshell. In `kubeshell-deployment.yml`, hostpath is `/home/config` by default. Feel free to change it.

3. Create kubeshell

Please prepare `kubeshell` image first. Below section will show that how to build the image.

```
kubectl create -f kubeshell-deployment.yml
```

## How to

* Build

```
make name=your_image_name tag=your_image_tag all
```