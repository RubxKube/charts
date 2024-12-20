# jackett

![Version: 1.3.0](https://img.shields.io/badge/Version-1.3.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.22.1085](https://img.shields.io/badge/AppVersion-0.22.1085-informational?style=flat-square)

Jackett works as a proxy server: it translates queries from apps into tracker-site-specific http queries.

**Homepage:** <https://github.com/Jackett/Jackett>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> |  |

## Source Code

* <https://github.com/Jackett/Jackett>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://rubxkube.github.io/common-charts | common | v0.3.12 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| common.define | int | `9117` |  |
| common.deployment.cpuLimit | string | `nil` |  |
| common.deployment.cpuRequest | string | `nil` |  |
| common.deployment.memoryLimit | string | `nil` |  |
| common.deployment.memoryRequest | string | `nil` |  |
| common.deployment.port | int | `9117` |  |
| common.deployment.strategy.rollingUpdate.maxSurge | string | `"25%"` |  |
| common.deployment.strategy.rollingUpdate.maxUnavailable | string | `"25%"` |  |
| common.deployment.strategy.type | string | `"RollingUpdate"` |  |
| common.hpa.avgCpuUtilization | int | `50` |  |
| common.hpa.enabled | bool | `false` |  |
| common.hpa.maxReplicas | int | `2` |  |
| common.hpa.minReplicas | int | `1` |  |
| common.image.pullPolicy | string | `"Always"` |  |
| common.image.repository | string | `"linuxserver/jackett"` |  |
| common.image.repositorySettings.isPrivate | bool | `false` |  |
| common.image.repositorySettings.secretName | string | `nil` |  |
| common.image.tag | string | `"0.22.1085"` |  |
| common.ingress.annotations | object | `{}` |  |
| common.ingress.enabled | bool | `false` |  |
| common.ingress.extraLabels | object | `{}` |  |
| common.ingress.hostName | string | `"toto.lan"` |  |
| common.ingress.ingressClassName | string | `"istio"` |  |
| common.ingress.tls.enabled | bool | `false` |  |
| common.ingress.tls.secretName | string | `"secret-jackett-tls"` |  |
| common.livenessProbe.failureThreshold | int | `1` |  |
| common.livenessProbe.initialDelaySeconds | int | `30` |  |
| common.livenessProbe.periodSeconds | int | `60` |  |
| common.livenessProbe.tcpSocket.port | int | `9117` |  |
| common.livenessProbe.timeoutSeconds | int | `3` |  |
| common.livenessProbeEnabled | bool | `true` |  |
| common.name | string | `"jackett"` |  |
| common.persistence.enabled | bool | `true` |  |
| common.persistence.volumes[0].containerMount | string | `"/config/"` |  |
| common.persistence.volumes[0].name | string | `"config"` |  |
| common.persistence.volumes[0].pvcClaim | string | `""` |  |
| common.persistence.volumes[0].size | string | `"2Gi"` |  |
| common.persistence.volumes[0].storageClassName | string | `""` |  |
| common.persistence.volumes[1].containerMount | string | `"/downloads/"` |  |
| common.persistence.volumes[1].name | string | `"downloads"` |  |
| common.persistence.volumes[1].pvcClaim | string | `""` |  |
| common.persistence.volumes[1].size | string | `"1Gi"` |  |
| common.persistence.volumes[1].storageClassName | string | `""` |  |
| common.readinessProbe.failureThreshold | int | `2` |  |
| common.readinessProbe.initialDelaySeconds | int | `30` |  |
| common.readinessProbe.periodSeconds | int | `30` |  |
| common.readinessProbe.tcpSocket.port | int | `9117` |  |
| common.readinessProbe.timeoutSeconds | int | `3` |  |
| common.readinessProbeEnabled | bool | `true` |  |
| common.service.servicePort | int | `9117` |  |
| common.service.type | string | `"ClusterIP"` |  |
| common.startupProbe.failureThreshold | int | `20` |  |
| common.startupProbe.httpGet.path | string | `"/"` |  |
| common.startupProbe.httpGet.port | int | `9117` |  |
| common.startupProbe.periodSeconds | int | `10` |  |
| common.startupProbe.timeoutSeconds | int | `1` |  |
| common.startupProbeEnabled | bool | `false` |  |
| common.tests.classicHttp.enabled | bool | `false` |  |
| common.tests.curlHostHeader.enabled | bool | `true` |  |
| common.tests.curlHostHeader.path | string | `"/"` |  |
| common.variables.nonSecret.AUTO_UPDATE | bool | `true` |  |
| common.variables.nonSecret.PGID | int | `1000` |  |
| common.variables.nonSecret.PUID | int | `1000` |  |
| common.variables.nonSecret.TZ | string | `"Etc/UTC"` |  |
| common.variables.secret | object | `{}` |  |
| define | int | `9117` |  |

