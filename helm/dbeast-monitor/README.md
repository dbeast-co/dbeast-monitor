# DBeast-Monitor Helm Chart

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square)

A Helm chart for deploying Grafana and Logstash as part of the DBeast-monitor stack in Kubernetes.

## Introduction

This chart bootstraps Grafana and Logstash deployments on a Kubernetes cluster using the Helm package manager.

## Prerequisites

- Kubernetes 1.12+
- Helm 3.0+

## Installation

### Add Helm Repository

```sh
helm repo add dbeast https://dbeast-co.github.io/dbeast-monitor-helm
helm repo update
```

### Install Chart

```sh
helm install dbeast-monitor dbeast/dbeast-monitor
```

The command deploys Grafana and Logstash on the Kubernetes cluster with default configurations. The configurations can
be customized using the `values.yaml` file.

## Uninstallation

To uninstall/delete the `dbeast-monitor` deployment:

```sh
helm uninstall dbeast-monitor
```

This command removes all the Kubernetes components associated with the chart and deletes the release.

## Custom Values File Example

To customize the deployment while keeping the rest of the values as defaults, you can create your own
`custom-values.yaml` file. Below is an example of a partial custom values file:

```yaml
ingress:
  enabled: true
  hosts:
    - dbeast.mycompany.com
  paths:
    - /
  ingressClassName: nginx
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
    nginx.ingress.kubernetes.io/ssl-redirect: "false"

logstash:
  resources:
    limits:
      cpu: 2
      memory: 6Gi
    requests:
      cpu: 2
      memory: 2Gi
  env:
    - name: LS_HEAP_SIZE
      value: "1g"
```

To use the custom values file along with the default values, run the following command:

```sh
helm install dbeast-monitor dbeast/dbeast-monitor -f custom-values.yaml
```

This command merges the values from `custom-values.yaml` with the default values specified in `values.yaml`, allowing
you to override only the parts you need.

## Configuration

The following table lists the configurable parameters of the chart and their default values.

### Grafana Configuration

| Parameter                                         | Description                                   | Default                             |
|---------------------------------------------------|-----------------------------------------------|-------------------------------------|
| `grafana.image`                                   | Grafana Docker image                          | `grafana/grafana-oss:9.5.10-ubuntu` |
| `grafana.service.type`                            | Kubernetes service type                       | `ClusterIP`                         |
| `grafana.service.port`                            | Kubernetes port where service is exposed      | `3000`                              |
| `grafana.resources.requests.cpu`                  | Grafana CPU resource requests                 | `1`                                 |
| `grafana.resources.requests.memory`               | Grafana memory resource requests              | `2Gi`                               |
| `grafana.persistence.grafana_logs.accessModes`    | Access mode for Grafana logs volume           | `ReadWriteOnce`                     |
| `grafana.persistence.grafana_logs.size`           | Size of persistent volume for Grafana logs    | `1Gi`                               |
| `grafana.persistence.grafana_storage.accessModes` | Access mode for Grafana storage volume        | `ReadWriteOnce`                     |
| `grafana.persistence.grafana_storage.size`        | Size of persistent volume for Grafana storage | `1Gi`                               |
| `grafana.persistence.grafana_plugins.accessModes` | Access mode for Grafana plugins volume        | `ReadWriteOnce`                     |
| `grafana.persistence.grafana_plugins.size`        | Size of persistent volume for Grafana plugins | `200Mi`                             |
| `grafana.pluginUrl`                               | URL to install Grafana plugins                | `plugin URL related field`          |
| `grafana.env`                                     | List of environment variables for Grafana     | `GF_LOG_LEVEL: "info"`              |
| `grafana.config.grafana.ini`                      | Ini file configuration for Grafana            | `details as per provided yaml`      |
| `grafana.config.ldap`                             | LDAP configuration for Grafana                | `details as per provided yaml`      |

### Logstash Configuration

| Parameter                                          | Description                                    | Default                         |
|----------------------------------------------------|------------------------------------------------|---------------------------------|
| `logstash.image`                                   | Logstash Docker image                          | `logstash:8.14.3`               |
| `logstash.service.type`                            | Kubernetes service type                        | `ClusterIP`                     |
| `logstash.service.port`                            | Kubernetes port where service is exposed       | `9600`                          |
| `logstash.resources.requests.cpu`                  | Logstash CPU resource requests                 | `1`                             |
| `logstash.resources.requests.memory`               | Logstash memory resource requests              | `2Gi`                           |
| `logstash.persistence.logstash_config.accessModes` | Access mode for Logstash config volume         | `ReadWriteOnce`                 |
| `logstash.persistence.logstash_config.size`        | Size of persistent volume for Logstash config  | `10Mi`                          |
| `logstash.persistence.logstash_lib.accessModes`    | Access mode for Logstash library volume        | `ReadWriteOnce`                 |
| `logstash.persistence.logstash_lib.size`           | Size of persistent volume for Logstash library | `10Mi`                          |
| `logstash.persistence.logstash_logs.accessModes`   | Access mode for Logstash logs volume           | `ReadWriteOnce`                 |
| `logstash.persistence.logstash_logs.size`          | Size of persistent volume for Logstash logs    | `1Gi`                           |
| `logstash.env`                                     | List of environment variables for Logstash     | `LS_JAVA_OPTS: "-Xms1g -Xmx1g"` |
| `logstash.config.logstash.yml`                     | Configuration for Logstash                     | `details as per provided yaml`  |

### Ingress Configuration

| Parameter                  | Description                | Default                                             |
|----------------------------|----------------------------|-----------------------------------------------------|
| `ingress.enabled`          | Enable Ingress for Grafana | `false`                                             |
| `ingress.ingressClassName` | Ingress class name         | `nginx`                                             |
| `ingress.annotations`      | Ingress annotations        | `nginx.ingress.kubernetes.io/rewrite-target: /      |
|                            |                            | nginx.ingress.kubernetes.io/ssl-redirect: "false"}` |
| `ingress.hosts`            | Ingress hostnames          | `["dbeast.local"]`                                  |
| `ingress.tls`              | Ingress TLS configuration  | `[]`                                                |
