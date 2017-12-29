# pingdom-operator

Kubernetes Operator to manage and maintain Pingdom Checks

[![CircleCI](https://circleci.com/gh/Nalum/pingdom-operator/tree/master.svg?style=svg&circle-token=1c5e69402faff9279ac1686e4e40b5812a397a96)](https://circleci.com/gh/Nalum/pingdom-operator/tree/master)
[![Docker Repository on Quay](https://quay.io/repository/nalum/pingdom-operator/status "Docker Repository on Quay")](https://quay.io/repository/nalum/pingdom-operator)

# Getting Started

In order to use the Pingdom Operator you will need to create the required [CRD](https://kubernetes.io/docs/tasks/access-kubernetes-api/extend-api-custom-resource-definitions/). You can use the CRD definition found [here](artifacts/examples/crd.yaml) to create the required CRD on your cluster.

You will also find examples for the required Secret and Deployment as well as an example HTTP Check in that folder.
