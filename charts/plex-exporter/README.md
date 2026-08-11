# plex-exporter

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: latest](https://img.shields.io/badge/AppVersion-latest-informational?style=flat-square)

Prometheus exporter for Plex Media Server metrics (libraries, sessions, transcoding).

**Homepage:** <https://github.com/jsclayton/prometheus-plex-exporter>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> |  |

## Source Code

* <https://github.com/jsclayton/prometheus-plex-exporter>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://rubxkube.github.io/common-charts | common | v0.6.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| common.deployment.cpuRequest | string | `"10m"` |  |
| common.deployment.memoryLimit | string | `"128Mi"` |  |
| common.deployment.memoryRequest | string | `"32Mi"` |  |
| common.deployment.minReplicas | int | `1` |  |
| common.deployment.podSecurityContext.runAsNonRoot | bool | `true` |  |
| common.deployment.podSecurityContext.runAsUser | int | `65534` |  |
| common.deployment.podSecurityContext.seccompProfile.type | string | `"RuntimeDefault"` |  |
| common.deployment.port | int | `9000` |  |
| common.deployment.securityContext.allowPrivilegeEscalation | bool | `false` |  |
| common.deployment.securityContext.capabilities.drop[0] | string | `"ALL"` |  |
| common.deployment.securityContext.readOnlyRootFilesystem | bool | `true` |  |
| common.deployment.strategy.rollingUpdate.maxSurge | string | `"25%"` |  |
| common.deployment.strategy.rollingUpdate.maxUnavailable | string | `"25%"` |  |
| common.deployment.strategy.type | string | `"RollingUpdate"` |  |
| common.image.pullPolicy | string | `"Always"` |  |
| common.image.repository | string | `"ghcr.io/jsclayton/prometheus-plex-exporter"` |  |
| common.image.tag | string | `"latest"` |  |
| common.ingress.enabled | bool | `false` |  |
| common.name | string | `"plex-exporter"` |  |
| common.persistence.enabled | bool | `false` |  |
| common.service.containerPort | int | `9000` |  |
| common.service.enabled | bool | `true` |  |
| common.service.servicePort | int | `9000` |  |
| common.service.type | string | `"ClusterIP"` |  |
| common.variables.secret.existingSecret | list | `[]` |  |
| define | int | `9000` |  |

