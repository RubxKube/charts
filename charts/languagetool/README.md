# languagetool

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 6.8](https://img.shields.io/badge/AppVersion-6.8-informational?style=flat-square)

LanguageTool is an open-source proofreading tool for spelling, grammar and style, supporting more than 30 languages.

**Homepage:** <https://languagetool.org/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> | <https://une-tasse-de.cafe> |

## Source Code

* <https://github.com/languagetool-org/languagetool>
* <https://github.com/meyayl/docker-languagetool>
* <https://github.com/rayliuca/grammared-language>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://rubxkube.github.io/common-charts | common | v0.6.4 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| common.deployment.cpuLimit | string | `nil` |  |
| common.deployment.cpuRequest | string | `nil` |  |
| common.deployment.emptyDir[0].containerMount | string | `"/tmp"` |  |
| common.deployment.emptyDir[0].name | string | `"tmp"` |  |
| common.deployment.memoryLimit | string | `nil` |  |
| common.deployment.memoryRequest | string | `nil` |  |
| common.deployment.port | int | `8081` |  |
| common.deployment.strategy.rollingUpdate.maxSurge | string | `"25%"` |  |
| common.deployment.strategy.rollingUpdate.maxUnavailable | string | `"25%"` |  |
| common.deployment.strategy.type | string | `"RollingUpdate"` |  |
| common.hpa.avgCpuUtilization | int | `50` |  |
| common.hpa.enabled | bool | `false` |  |
| common.hpa.maxReplicas | int | `2` |  |
| common.hpa.minReplicas | int | `1` |  |
| common.image.pullPolicy | string | `"IfNotPresent"` |  |
| common.image.repository | string | `"meyay/languagetool"` |  |
| common.image.repositorySettings.isPrivate | bool | `false` |  |
| common.image.repositorySettings.secretName | string | `nil` |  |
| common.image.tag | string | `"6.8"` |  |
| common.ingress.annotations | object | `{}` |  |
| common.ingress.enabled | bool | `false` |  |
| common.ingress.extraLabels | object | `{}` |  |
| common.ingress.hostName | string | `"languagetool.your-domain.com"` |  |
| common.ingress.ingressClassName | string | `""` |  |
| common.ingress.tls.enabled | bool | `false` |  |
| common.ingress.tls.secretName | string | `"secret-languagetool-tls"` |  |
| common.livenessProbe.failureThreshold | int | `3` |  |
| common.livenessProbe.httpGet.path | string | `"/v2/languages"` |  |
| common.livenessProbe.httpGet.port | int | `8081` |  |
| common.livenessProbe.initialDelaySeconds | int | `30` |  |
| common.livenessProbe.periodSeconds | int | `60` |  |
| common.livenessProbe.timeoutSeconds | int | `5` |  |
| common.livenessProbeEnabled | bool | `true` |  |
| common.name | string | `"languagetool"` |  |
| common.persistence.enabled | bool | `true` |  |
| common.persistence.volumes[0].containerMount | string | `"/fasttext"` |  |
| common.persistence.volumes[0].name | string | `"fasttext"` |  |
| common.persistence.volumes[0].pvcClaim | string | `""` |  |
| common.persistence.volumes[0].size | string | `"1Gi"` |  |
| common.persistence.volumes[0].storageClassName | string | `""` |  |
| common.persistence.volumes[1].containerMount | string | `"/ngrams"` |  |
| common.persistence.volumes[1].name | string | `"ngrams"` |  |
| common.persistence.volumes[1].pvcClaim | string | `""` |  |
| common.persistence.volumes[1].size | string | `"5Gi"` |  |
| common.persistence.volumes[1].storageClassName | string | `""` |  |
| common.readinessProbe.failureThreshold | int | `3` |  |
| common.readinessProbe.httpGet.path | string | `"/v2/languages"` |  |
| common.readinessProbe.httpGet.port | int | `8081` |  |
| common.readinessProbe.initialDelaySeconds | int | `10` |  |
| common.readinessProbe.periodSeconds | int | `30` |  |
| common.readinessProbe.timeoutSeconds | int | `5` |  |
| common.readinessProbeEnabled | bool | `true` |  |
| common.service.servicePort | int | `8081` |  |
| common.service.type | string | `"ClusterIP"` |  |
| common.startupProbe.failureThreshold | int | `24` |  |
| common.startupProbe.httpGet.path | string | `"/v2/languages"` |  |
| common.startupProbe.httpGet.port | int | `8081` |  |
| common.startupProbe.periodSeconds | int | `10` |  |
| common.startupProbe.timeoutSeconds | int | `5` |  |
| common.startupProbeEnabled | bool | `true` |  |
| common.tests.classicHttp.enabled | bool | `false` |  |
| common.tests.curlHostHeader.enabled | bool | `true` |  |
| common.tests.curlHostHeader.path | string | `"/v2/languages"` |  |
| common.variables.nonSecret.JAVA_XMS | string | `"256m"` |  |
| common.variables.nonSecret.JAVA_XMX | string | `"1536m"` |  |
| common.variables.nonSecret.LOG_LEVEL | string | `"INFO"` |  |
| common.variables.secret | object | `{}` |  |
| define | int | `8081` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.14.2](https://github.com/norwoodj/helm-docs/releases/v1.14.2)
