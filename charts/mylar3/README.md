# mylar3

![Version: 0.0.1](https://img.shields.io/badge/Version-0.0.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.8.0](https://img.shields.io/badge/AppVersion-0.8.0-informational?style=flat-square)

Mylar3 is an automated Comic Book (cbr/cbz) downloader program

**Homepage:** <https://mylarcomics.com/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> |  |

## Source Code

* <https://github.com/mylar3/mylar3>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://rubxkube.github.io/common-charts | common | v0.3.12 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| common.deployment.args | list | `[]` |  |
| common.deployment.cpuLimit | string | `nil` |  |
| common.deployment.cpuRequest | string | `nil` |  |
| common.deployment.emptyDir | list | `[]` |  |
| common.deployment.initContainers | list | `[]` |  |
| common.deployment.memoryLimit | string | `nil` |  |
| common.deployment.memoryRequest | string | `nil` |  |
| common.deployment.minReplicas | int | `1` |  |
| common.deployment.port | int | `8090` |  |
| common.deployment.strategy.type | string | `"Recreate"` |  |
| common.hpa.avgCpuUtilization | int | `50` |  |
| common.hpa.enabled | bool | `false` |  |
| common.hpa.maxReplicas | int | `2` |  |
| common.hpa.minReplicas | int | `1` |  |
| common.image.pullPolicy | string | `"Always"` |  |
| common.image.repository | string | `"lscr.io/linuxserver/mylar3"` |  |
| common.image.repositorySettings.isPrivate | bool | `false` |  |
| common.image.repositorySettings.secretName | string | `nil` |  |
| common.image.tag | string | `"0.8.0"` |  |
| common.ingress.annotations | object | `{}` |  |
| common.ingress.enabled | bool | `false` |  |
| common.ingress.extraLabels | object | `{}` |  |
| common.ingress.hostName | string | `"toto.lan"` |  |
| common.ingress.ingressClassName | string | `"istio"` |  |
| common.ingress.tls.enabled | bool | `true` |  |
| common.ingress.tls.secretName | string | `""` |  |
| common.livenessProbe | object | `{}` |  |
| common.livenessProbeEnabled | bool | `false` |  |
| common.name | string | `"mylar3"` |  |
| common.persistence.enabled | bool | `true` |  |
| common.persistence.extraVolumes | list | `[]` |  |
| common.persistence.volumes[0].containerMount | string | `"/config"` |  |
| common.persistence.volumes[0].name | string | `"config"` |  |
| common.persistence.volumes[0].size | string | `"2Gi"` |  |
| common.persistence.volumes[1].containerMount | string | `"/comics"` |  |
| common.persistence.volumes[1].name | string | `"comics"` |  |
| common.persistence.volumes[1].size | string | `"5Gi"` |  |
| common.persistence.volumes[2].containerMount | string | `"/downloads"` |  |
| common.persistence.volumes[2].name | string | `"downloads"` |  |
| common.persistence.volumes[2].size | string | `"2Gi"` |  |
| common.readinessProbe | object | `{}` |  |
| common.readinessProbeEnabled | bool | `false` |  |
| common.service.containerPort | int | `8090` |  |
| common.service.enabled | bool | `true` |  |
| common.service.extraLabels | object | `{}` |  |
| common.service.servicePort | int | `8090` |  |
| common.service.type | string | `"ClusterIP"` |  |
| common.startupProbe | object | `{}` |  |
| common.startupProbeEnabled | bool | `false` |  |
| common.tests.classicHttp.enabled | bool | `false` |  |
| common.tests.curlHostHeader.enabled | bool | `false` |  |
| common.tests.curlHostHeader.path | string | `""` |  |
| common.variables.configMap.existingConfigMap | list | `[]` |  |
| common.variables.nonSecret.PGID | int | `1000` |  |
| common.variables.nonSecret.PUID | int | `1000` |  |
| common.variables.nonSecret.TZ | string | `"Europe/Paris"` |  |
| common.variables.secret.data | object | `{}` |  |
| common.variables.secret.existingSecret | list | `[]` |  |
| common.variables.secret.extraExistingSecrets | list | `[]` |  |
| define | int | `8090` |  |

