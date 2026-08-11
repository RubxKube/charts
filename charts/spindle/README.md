# spindle

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: v1.16.1-alpha](https://img.shields.io/badge/AppVersion-v1.16.1--alpha-informational?style=flat-square)

Spindle is the self-hosted CI runner of Tangled, the git platform built on ATProto.

**Homepage:** <https://docs.tangled.org/spindles>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> |  |

## Source Code

* <https://tangled.org/tangled.org/core>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://rubxkube.github.io/common-charts | common | v0.6.2 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| common.deployment.additionalContainers[0].args[0] | string | `"dockerd"` |  |
| common.deployment.additionalContainers[0].args[1] | string | `"--host=unix:///var/run/docker.sock"` |  |
| common.deployment.additionalContainers[0].args[2] | string | `"--host=tcp://127.0.0.1:2375"` |  |
| common.deployment.additionalContainers[0].args[3] | string | `"--tls=false"` |  |
| common.deployment.additionalContainers[0].env[0].name | string | `"DOCKER_TLS_CERTDIR"` |  |
| common.deployment.additionalContainers[0].env[0].value | string | `""` |  |
| common.deployment.additionalContainers[0].image | string | `"docker:29.7.2-dind"` |  |
| common.deployment.additionalContainers[0].imagePullPolicy | string | `"IfNotPresent"` |  |
| common.deployment.additionalContainers[0].name | string | `"dind"` |  |
| common.deployment.additionalContainers[0].readinessProbe.exec.command[0] | string | `"docker"` |  |
| common.deployment.additionalContainers[0].readinessProbe.exec.command[1] | string | `"-H"` |  |
| common.deployment.additionalContainers[0].readinessProbe.exec.command[2] | string | `"tcp://127.0.0.1:2375"` |  |
| common.deployment.additionalContainers[0].readinessProbe.exec.command[3] | string | `"info"` |  |
| common.deployment.additionalContainers[0].readinessProbe.initialDelaySeconds | int | `10` |  |
| common.deployment.additionalContainers[0].readinessProbe.periodSeconds | int | `15` |  |
| common.deployment.additionalContainers[0].resources.limits.memory | string | `"10Gi"` |  |
| common.deployment.additionalContainers[0].resources.requests.cpu | string | `"100m"` |  |
| common.deployment.additionalContainers[0].resources.requests.memory | string | `"512Mi"` |  |
| common.deployment.additionalContainers[0].securityContext.privileged | bool | `true` |  |
| common.deployment.additionalContainers[0].volumeMounts[0].mountPath | string | `"/var/lib/docker"` |  |
| common.deployment.additionalContainers[0].volumeMounts[0].name | string | `"docker-data"` |  |
| common.deployment.cpuLimit | string | `nil` |  |
| common.deployment.cpuRequest | string | `"100m"` |  |
| common.deployment.emptyDir[0].name | string | `"docker-data"` |  |
| common.deployment.memoryLimit | string | `"1Gi"` |  |
| common.deployment.memoryRequest | string | `"256Mi"` |  |
| common.deployment.port | int | `6555` |  |
| common.deployment.strategy.type | string | `"Recreate"` |  |
| common.hpa.avgCpuUtilization | int | `50` |  |
| common.hpa.enabled | bool | `false` |  |
| common.hpa.maxReplicas | int | `2` |  |
| common.hpa.minReplicas | int | `1` |  |
| common.image.pullPolicy | string | `"IfNotPresent"` |  |
| common.image.repository | string | `"ghcr.io/qjoly/spindle"` |  |
| common.image.repositorySettings.isPrivate | bool | `false` |  |
| common.image.repositorySettings.secretName | string | `nil` |  |
| common.image.tag | string | `"v1.16.1-alpha"` |  |
| common.ingress.enabled | bool | `false` |  |
| common.ingress.hostName | string | `"spindle.example.com"` |  |
| common.ingress.ingressClassName | string | `nil` |  |
| common.ingress.tls.enabled | bool | `true` |  |
| common.ingress.tls.secretName | string | `""` |  |
| common.livenessProbe.failureThreshold | int | `3` |  |
| common.livenessProbe.httpGet.path | string | `"/"` |  |
| common.livenessProbe.httpGet.port | int | `6555` |  |
| common.livenessProbe.initialDelaySeconds | int | `20` |  |
| common.livenessProbe.periodSeconds | int | `60` |  |
| common.livenessProbe.timeoutSeconds | int | `3` |  |
| common.livenessProbeEnabled | bool | `true` |  |
| common.name | string | `"spindle"` |  |
| common.persistence.enabled | bool | `true` |  |
| common.persistence.volumes[0].containerMount | string | `"/var/lib/spindle"` |  |
| common.persistence.volumes[0].name | string | `"data"` |  |
| common.persistence.volumes[0].pvcClaim | string | `""` |  |
| common.persistence.volumes[0].size | string | `"5Gi"` |  |
| common.persistence.volumes[0].storageClassName | string | `""` |  |
| common.persistence.volumes[1].containerMount | string | `"/var/log/spindle"` |  |
| common.persistence.volumes[1].name | string | `"logs"` |  |
| common.persistence.volumes[1].pvcClaim | string | `""` |  |
| common.persistence.volumes[1].size | string | `"2Gi"` |  |
| common.persistence.volumes[1].storageClassName | string | `""` |  |
| common.readinessProbe.failureThreshold | int | `3` |  |
| common.readinessProbe.httpGet.path | string | `"/"` |  |
| common.readinessProbe.httpGet.port | int | `6555` |  |
| common.readinessProbe.initialDelaySeconds | int | `10` |  |
| common.readinessProbe.periodSeconds | int | `15` |  |
| common.readinessProbe.timeoutSeconds | int | `3` |  |
| common.readinessProbeEnabled | bool | `true` |  |
| common.service.containerPort | int | `6555` |  |
| common.service.enabled | bool | `true` |  |
| common.service.extraLabels | object | `{}` |  |
| common.service.servicePort | int | `6555` |  |
| common.service.type | string | `"ClusterIP"` |  |
| common.startupProbe | object | `{}` |  |
| common.startupProbeEnabled | bool | `false` |  |
| common.tests.classicHttp.enabled | bool | `true` |  |
| common.tests.curlHostHeader.enabled | bool | `true` |  |
| common.tests.curlHostHeader.path | string | `"/xrpc/sh.tangled.owner"` |  |
| common.variables.nonSecret.DOCKER_HOST | string | `"tcp://127.0.0.1:2375"` |  |
| common.variables.nonSecret.SPINDLE_NIXERY_PIPELINES_MAX_CONCURRENT_WORKFLOWS | int | `2` |  |
| common.variables.nonSecret.SPINDLE_NIXERY_PIPELINES_MAX_JOB_MEMORY_MB | int | `4096` |  |
| common.variables.nonSecret.SPINDLE_NIXERY_PIPELINES_WORKFLOW_TIMEOUT | string | `"15m"` |  |
| common.variables.nonSecret.SPINDLE_SERVER_DB_PATH | string | `"/var/lib/spindle/spindle.db"` |  |
| common.variables.nonSecret.SPINDLE_SERVER_HOSTNAME | string | `"spindle.example.com"` | REQUIRED. Public hostname of this spindle, and it must match the one registered on the appview: the knot signs its requests for that audience. |
| common.variables.nonSecret.SPINDLE_SERVER_LOG_DIR | string | `"/var/log/spindle"` |  |
| common.variables.nonSecret.SPINDLE_SERVER_OWNER | string | `"did:plc:000000000000000000000000"` | REQUIRED. ATProto DID of the owner. Resolve yours with `curl "https://public.api.bsky.app/xrpc/com.atproto.identity.resolveHandle?handle=<handle>"` |
| common.variables.nonSecret.SPINDLE_SERVER_REPO_DIR | string | `"/var/lib/spindle/repos"` |  |
| common.variables.nonSecret.SPINDLE_SERVER_TAP_DB_PATH | string | `"/var/lib/spindle/tap.db"` |  |
| common.variables.secret | object | `{}` |  |
| define | int | `6555` |  |

