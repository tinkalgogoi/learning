# confluent install in kafka

Helm install:
[link1](https://docs.confluent.io/5.0.0/installation/installing_cp/cp-helm-charts/docs/index.html)

Demo install:
```sh
    kubectl create ns kafka
    helm install my-confluent-oss confluentinc/cp-helm-charts -f values.yaml -n kafka
```