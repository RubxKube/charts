# pr-gitops-grafana-annotation

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: latest](https://img.shields.io/badge/AppVersion-latest-informational?style=flat-square)

Turns GitHub pull request webhooks into Grafana annotations, to line deploys up with your dashboards.

**Homepage:** <https://github.com/qjoly/pr-gitops-grafana-annotation>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> |  |

## Source Code

* <https://github.com/qjoly/pr-gitops-grafana-annotation>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://rubxkube.github.io/common-charts | common | v0.6.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| common.deployment.cpuRequest | string | `"10m"` |  |
| common.deployment.memoryLimit | string | `"64Mi"` |  |
| common.deployment.memoryRequest | string | `"32Mi"` |  |
| common.deployment.minReplicas | int | `1` |  |
| common.deployment.podSecurityContext.runAsNonRoot | bool | `true` |  |
| common.deployment.podSecurityContext.runAsUser | int | `65534` |  |
| common.deployment.podSecurityContext.seccompProfile.type | string | `"RuntimeDefault"` |  |
| common.deployment.port | int | `8080` |  |
| common.deployment.strategy.type | string | `"Recreate"` |  |
| common.image.pullPolicy | string | `"Always"` |  |
| common.image.repository | string | `"ghcr.io/qjoly/pr-gitops-grafana-annotation"` |  |
| common.image.tag | string | `"latest"` |  |
| common.ingress.enabled | bool | `false` |  |
| common.name | string | `"pr-gitops-grafana-annotation"` |  |
| common.persistence.enabled | bool | `false` |  |
| common.readinessProbe.httpGet.path | string | `"/healthz"` |  |
| common.readinessProbe.httpGet.port | int | `8080` |  |
| common.readinessProbe.initialDelaySeconds | int | `5` |  |
| common.readinessProbe.periodSeconds | int | `15` |  |
| common.readinessProbe.timeoutSeconds | int | `5` |  |
| common.readinessProbeEnabled | bool | `true` |  |
| common.service.containerPort | int | `8080` |  |
| common.service.enabled | bool | `true` |  |
| common.service.servicePort | int | `8080` |  |
| common.service.type | string | `"ClusterIP"` |  |
| common.variables.nonSecret.LISTEN_ADDR | string | `":8080"` |  |
| common.variables.secret.existingSecret | list | `[]` |  |
| define | int | `8080` |  |

