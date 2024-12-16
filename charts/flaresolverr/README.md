# flaresolverr

![Version: 0.0.1](https://img.shields.io/badge/Version-0.0.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: v3.3.21](https://img.shields.io/badge/AppVersion-v3.3.21-informational?style=flat-square)

FlareSolverr is a proxy server that solves Cloudflare's Javascript challenges.

**Homepage:** <https://github.com/FlareSolverr/FlareSolverr>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> |  |

## Source Code

* <https://github.com/FlareSolverr/FlareSolverr>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://rubxkube.github.io/common-charts | common | v0.3.12 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| common.deployment.cpuLimit | string | `nil` |  |
| common.deployment.cpuRequest | string | `nil` |  |
| common.deployment.memoryLimit | string | `nil` |  |
| common.deployment.memoryRequest | string | `nil` |  |
| common.deployment.port | int | `8191` |  |
| common.deployment.strategy.rollingUpdate.maxSurge | string | `"25%"` |  |
| common.deployment.strategy.rollingUpdate.maxUnavailable | string | `"25%"` |  |
| common.deployment.strategy.type | string | `"RollingUpdate"` |  |
| common.hpa.avgCpuUtilization | int | `50` |  |
| common.hpa.enabled | bool | `false` |  |
| common.hpa.maxReplicas | int | `2` |  |
| common.hpa.minReplicas | int | `1` |  |
| common.image.pullPolicy | string | `"Always"` |  |
| common.image.repository | string | `"ghcr.io/flaresolverr/flaresolverr"` |  |
| common.image.repositorySettings.isPrivate | bool | `false` |  |
| common.image.repositorySettings.secretName | string | `nil` |  |
| common.image.tag | string | `"v3.3.21"` |  |
| common.ingress.annotations | object | `{}` |  |
| common.ingress.enabled | bool | `false` |  |
| common.ingress.extraLabels | object | `{}` |  |
| common.ingress.hostName | string | `"toto.lan"` |  |
| common.ingress.ingressClassName | string | `"istio"` |  |
| common.ingress.tls.enabled | bool | `false` |  |
| common.ingress.tls.secretName | string | `"secret-flaresolverr-tls"` |  |
| common.livenessProbe.failureThreshold | int | `1` |  |
| common.livenessProbe.initialDelaySeconds | int | `30` |  |
| common.livenessProbe.periodSeconds | int | `60` |  |
| common.livenessProbe.tcpSocket.port | int | `8191` |  |
| common.livenessProbe.timeoutSeconds | int | `3` |  |
| common.livenessProbeEnabled | bool | `true` |  |
| common.name | string | `"flaresolverr"` |  |
| common.persistence.enabled | bool | `false` |  |
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
| common.readinessProbe.tcpSocket.port | int | `8191` |  |
| common.readinessProbe.timeoutSeconds | int | `3` |  |
| common.readinessProbeEnabled | bool | `true` |  |
| common.service.servicePort | int | `8191` |  |
| common.service.type | string | `"ClusterIP"` |  |
| common.startupProbe.failureThreshold | int | `20` |  |
| common.startupProbe.httpGet.path | string | `"/"` |  |
| common.startupProbe.httpGet.port | int | `8191` |  |
| common.startupProbe.periodSeconds | int | `10` |  |
| common.startupProbe.timeoutSeconds | int | `1` |  |
| common.startupProbeEnabled | bool | `false` |  |
| common.tests.classicHttp.enabled | bool | `false` |  |
| common.tests.curlHostHeader.enabled | bool | `true` |  |
| common.tests.curlHostHeader.path | string | `"/"` |  |
| common.variables.nonSecret.LOG_LEVEL | string | `"info"` |  |
| common.variables.secret | object | `{}` |  |
| define | int | `8191` |  |

