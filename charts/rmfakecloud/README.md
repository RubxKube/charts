# rmfakecloud

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: latest](https://img.shields.io/badge/AppVersion-latest-informational?style=flat-square)

Self-hosted replacement for the reMarkable tablet cloud sync service.

**Homepage:** <https://github.com/ddvk/rmfakecloud>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> |  |

## Source Code

* <https://github.com/ddvk/rmfakecloud>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://rubxkube.github.io/common-charts | common | v0.6.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| common.deployment.cpuRequest | string | `"50m"` |  |
| common.deployment.memoryLimit | string | `"512Mi"` |  |
| common.deployment.memoryRequest | string | `"128Mi"` |  |
| common.deployment.minReplicas | int | `1` |  |
| common.deployment.port | int | `3000` |  |
| common.deployment.strategy.type | string | `"Recreate"` |  |
| common.image.pullPolicy | string | `"Always"` |  |
| common.image.repository | string | `"ddvk/rmfakecloud"` |  |
| common.image.tag | string | `"latest"` |  |
| common.ingress.enabled | bool | `false` |  |
| common.name | string | `"rmfakecloud"` |  |
| common.persistence.enabled | bool | `true` |  |
| common.persistence.volumes[0].containerMount | string | `"/data"` |  |
| common.persistence.volumes[0].name | string | `"data"` |  |
| common.persistence.volumes[0].size | string | `"10Gi"` |  |
| common.readinessProbe.httpGet.path | string | `"/"` |  |
| common.readinessProbe.httpGet.port | int | `3000` |  |
| common.readinessProbe.initialDelaySeconds | int | `10` |  |
| common.readinessProbe.periodSeconds | int | `15` |  |
| common.readinessProbe.timeoutSeconds | int | `5` |  |
| common.readinessProbeEnabled | bool | `true` |  |
| common.service.containerPort | int | `3000` |  |
| common.service.enabled | bool | `true` |  |
| common.service.servicePort | int | `3000` |  |
| common.service.type | string | `"ClusterIP"` |  |
| common.variables.nonSecret | object | `{}` |  |
| common.variables.secret.existingSecret | list | `[]` |  |
| define | int | `3000` |  |

