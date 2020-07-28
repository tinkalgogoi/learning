# Install and configure the metrics cluster

Install Prometheus and Grafana:

```sh
    kubectl create ns monitoring

    helm install eod-monitoring stable/prometheus -n monitoring

    kubectl apply -f datasource-cm.yaml
    kubectl apply -f flink-dashboard-cm.yaml
    helm install eod-metrics-ui stable/grafana -f values-grafana.yaml -n monitoring
```

Get your grafana 'admin' user password by running:

```sh
   kubectl get secret -n monitoring eod-metrics-ui-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```

Apply pometheus configuration(by default prometheus will scrape from its own cluster) for scraping the separate AKS cluster containing flink metrics:

```sh
    kubectl apply -f p8s-config-cm.yaml
```

UI of p8s and grafana

```sh
    kubectl port-forward service/eod-monitoring-prometheus-server 9090:80 -n monitoring
    kubectl port-forward service/eod-metrics-ui-grafana 3000:80 -n monitoring
```
