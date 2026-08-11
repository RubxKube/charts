# trmnl-server

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: unstable](https://img.shields.io/badge/AppVersion-unstable-informational?style=flat-square)

Self-hosted server for TRMNL e-ink displays, installed from source since upstream publishes no image.

**Homepage:** <https://github.com/ohAnd/trmnlServer>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> |  |

## Source Code

* <https://github.com/ohAnd/trmnlServer>

## Requirements

Kubernetes: `>= 1.18`

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| certificate.enabled | bool | `false` |  |
| certificate.issuerRef.kind | string | `"ClusterIssuer"` |  |
| certificate.issuerRef.name | string | `""` |  |
| certificate.secretName | string | `""` |  |
| generator.activeDeadlineSeconds | int | `120` |  |
| generator.configMapName | string | `""` |  |
| generator.enabled | bool | `false` |  |
| generator.env | object | `{}` |  |
| generator.pipPackages | string | `"requests pillow"` |  |
| generator.resources.limits.memory | string | `"256Mi"` |  |
| generator.resources.requests.cpu | string | `"50m"` |  |
| generator.resources.requests.memory | string | `"128Mi"` |  |
| generator.schedule | string | `"*/15 * * * *"` |  |
| generator.screens | list | `[]` |  |
| generator.scriptPath | string | `"/script/generate_image.py"` |  |
| generator.secretEnv | list | `[]` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"python"` |  |
| image.tag | string | `"3.14-slim"` |  |
| ingressRoute.enabled | bool | `false` |  |
| ingressRoute.entryPoint | string | `"websecure"` |  |
| ingressRoute.serversTransport.enabled | bool | `true` |  |
| ingressRoute.serversTransport.insecureSkipVerify | bool | `true` |  |
| ingressRoute.tls.secretName | string | `""` |  |
| name | string | `"trmnl-server"` |  |
| persistence.existingClaim | string | `""` |  |
| persistence.size | string | `"2Gi"` |  |
| persistence.storageClassName | string | `""` |  |
| server.extraEnv | object | `{}` |  |
| server.hostname | string | `"trmnl.example.com"` |  |
| server.patch.configMapName | string | `""` |  |
| server.patch.enabled | bool | `false` |  |
| server.patch.files | list | `[]` |  |
| server.port | int | `443` |  |
| server.resources.limits.memory | string | `"256Mi"` |  |
| server.resources.requests.cpu | string | `"50m"` |  |
| server.resources.requests.memory | string | `"128Mi"` |  |
| server.uploadTokenSecret.key | string | `"token"` |  |
| server.uploadTokenSecret.name | string | `"trmnl-upload-token"` |  |
| source.gitImage | string | `"alpine/git:2.47.2"` |  |
| source.repo | string | `"https://github.com/ohAnd/trmnlServer.git"` |  |
| source.revision | string | `"5e2a83bc8ab78214c1f1132048d65945888117ec"` |  |

