# joplin

![Version: 1.1.16](https://img.shields.io/badge/Version-1.1.16-informational?style=flat-square) ![AppVersion: 2.14-beta](https://img.shields.io/badge/AppVersion-2.14--beta-informational?style=flat-square)

Joplin is an open source note-taking app. Capture your thoughts and securely access them from any device.

**Homepage:** <https://joplinapp.org/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> | <https://une-tasse-de.cafe> |

## Source Code

* <https://github.com/laurent22/joplin>
* <https://github.com/QJoly/helm-charts>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | postgresql | 13.2.25 |
| https://rubxkube.github.io/common-charts | common | v0.2.2 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| common.app.containerPort | int | `22300` |  |
| common.app.servicePort | int | `80` |  |
| common.deployment.cpuLimit | string | `nil` |  |
| common.deployment.cpuRequest | string | `nil` |  |
| common.deployment.memoryLimit | string | `nil` |  |
| common.deployment.memoryRequest | string | `nil` |  |
| common.deployment.strategy.rollingUpdate.maxSurge | string | `"25%"` |  |
| common.deployment.strategy.rollingUpdate.maxUnavailable | string | `"25%"` |  |
| common.deployment.strategy.type | string | `"RollingUpdate"` |  |
| common.hpa.avgCpuUtilization | int | `50` |  |
| common.hpa.enabled | bool | `false` |  |
| common.hpa.maxReplicas | int | `2` |  |
| common.hpa.minReplicas | int | `1` |  |
| common.image.pullPolicy | string | `"Always"` |  |
| common.image.repository | string | `"joplin/server"` |  |
| common.image.repositorySettings.isPrivate | bool | `false` |  |
| common.image.repositorySettings.secretName | string | `nil` |  |
| common.image.tag | string | `"2.14-beta"` |  |
| common.ingress.certResolver | string | `"letsencrypt"` |  |
| common.ingress.enabled | bool | `false` |  |
| common.ingress.entrypoint | string | `"websecure"` |  |
| common.ingress.hostName | string | `"joplin.une-tasse-de.cafe"` |  |
| common.ingress.ingressClassName | string | `"istio"` |  |
| common.ingress.isIngressRoute | bool | `true` |  |
| common.ingress.tls.enabled | bool | `true` |  |
| common.ingress.tls.secretName | string | `""` |  |
| common.livenessProbe.failureThreshold | int | `1` |  |
| common.livenessProbe.initialDelaySeconds | int | `30` |  |
| common.livenessProbe.periodSeconds | int | `60` |  |
| common.livenessProbe.tcpSocket.port | int | `22300` |  |
| common.livenessProbe.timeoutSeconds | int | `3` |  |
| common.livenessProbeEnabled | bool | `true` |  |
| common.persistence.enabled | bool | `false` |  |
| common.readinessProbe.failureThreshold | int | `2` |  |
| common.readinessProbe.initialDelaySeconds | int | `30` |  |
| common.readinessProbe.periodSeconds | int | `30` |  |
| common.readinessProbe.tcpSocket.port | int | `22300` |  |
| common.readinessProbe.timeoutSeconds | int | `3` |  |
| common.readinessProbeEnabled | bool | `true` |  |
| common.startupProbe.failureThreshold | int | `20` |  |
| common.startupProbe.httpGet.path | string | `"/"` |  |
| common.startupProbe.httpGet.port | int | `22300` |  |
| common.startupProbe.periodSeconds | int | `10` |  |
| common.startupProbe.timeoutSeconds | int | `1` |  |
| common.startupProbeEnabled | bool | `false` |  |
| common.tests.classicHttp.enabled | bool | `false` |  |
| common.tests.curlHostHeader.enabled | bool | `true` |  |
| common.tests.curlHostHeader.path | string | `"/api/ping"` |  |
| common.variables.nonSecret.APP_BASE_URL | string | `"https://joplin.une-tasse-de.cafe"` |  |
| common.variables.nonSecret.DB_CLIENT | string | `"pg"` |  |
| common.variables.nonSecret.POSTGRES_DATABASE | string | `"joplin"` |  |
| common.variables.nonSecret.POSTGRES_HOST | string | `"joplin-postgresql"` |  |
| common.variables.nonSecret.POSTGRES_PASSWORD | string | `"joplinpass"` |  |
| common.variables.nonSecret.POSTGRES_PORT | string | `"5432"` |  |
| common.variables.nonSecret.POSTGRES_USER | string | `"joplinuser"` |  |
| common.variables.nonSecret.MAILER_ENABLED | int | `1` |  |
| common.variables.nonSecret.MAILER_HOST | string | `"mail.une-tasse-de.cafe"` |  |
| common.variables.nonSecret.MAILER_PORT | int | `25` |  |
| common.variables.nonSecret.MAILER_SECURITY | string | `"none"` |  |
| common.variables.nonSecret.MAILER_NOREPLY_NAME | string | `"Joplin Server"` |  |
| common.variables.nonSecret.MAILER_NOREPLY_EMAIL | string | `"noreply@une-tasse-de.cafe"` |  |
| common.variables.secret | object | `{}` |  |
| define | int | `22300` |  |
| postgresql.auth.database | string | `"joplin"` |  |
| postgresql.auth.password | string | `"joplinpass"` |  |
| postgresql.auth.postgresPassword | string | `"changeme"` |  |
| postgresql.auth.username | string | `"joplinuser"` |  |
| postgresql.enabled | bool | `true` |  |
| postgresql.persistence.enabled | bool | `true` |  |
| postgresql.persistence.existingClaim | string | `""` |  |
| postgresql.persistence.storageClass | string | `""` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.3](https://github.com/norwoodj/helm-docs/releases/v1.11.3)
