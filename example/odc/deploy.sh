#!/bin/usr/env bash

# 1. Deploy an oceanbase database

ROOT_PWD=$(openssl rand -base64 12)

kubectl create secret generic -n oceanbase root-password --from-literal password=$ROOT_PWD

kubectl apply -f obcluster.yaml

kubectl wait --for=jsonpath='{.status.status}'=running -n oceanbase obclusters.oceanbase.oceanbase.com obcluster

# 2. Deploy an odc service

POD_IP=$(kubectl get -n oceanbase pod -l ref-obcluster=obcluster -o jsonpath='{.items[0].status.podIP}')

