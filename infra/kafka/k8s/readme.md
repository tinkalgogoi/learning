# confluent install in kafka

Helm install:
[link1](https://docs.confluent.io/5.0.0/installation/installing_cp/cp-helm-charts/docs/index.html)

Demo install:

```sh
    kubectl create ns kafka
    helm install my-confluent-oss confluentinc/cp-helm-charts -f values.yaml -n kafka

    # Get actual all values
    helm get values my-confluent-oss -a  -n kafka -o yaml
```

values for each chart [link1](https://github.com/confluentinc/cp-helm-charts/tree/master/charts)

fix - JMX port env name for rest proxy as KAFKAREST_JMX_PORT
https://github.com/ravenolf/cp-helm-charts/commit/7c49a07d3cf3d8d7c5b4fd0bbe8e726675a71ab4
