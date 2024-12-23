# stirling-pdf

![Version: 0.0.1](https://img.shields.io/badge/Version-0.0.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.36.5](https://img.shields.io/badge/AppVersion-0.36.5-informational?style=flat-square)

Stirling-PDF is a robust, locally hosted web-based PDF manipulation tool using Docker. It enables you to carry out various operations on PDF files, including splitting, merging, converting, reorganizing, adding images, rotating, compressing, and more

**Homepage:** <https://github.com/Yooooomi/your_spotify>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> | <https://une-tasse-de.cafe> |

## Source Code

* <https://github.com/Yooooomi/your_spotify>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://rubxkube.github.io/common-charts | common | v0.3.12 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| common.deployment.port | int | `8080` |  |
| common.image.pullPolicy | string | `"IfNotPresent"` |  |
| common.image.repository | string | `"ghcr.io/stirling-tools/stirling-pdf"` |  |
| common.image.tag | string | `"0.36.5"` |  |
| common.ingress.annotations | object | `{}` |  |
| common.ingress.enabled | bool | `true` |  |
| common.ingress.extraLabels | object | `{}` |  |
| common.ingress.hostName | string | `"stirling-pdf.your-domain.com"` |  |
| common.ingress.ingressClassName | string | `"nginx"` |  |
| common.ingress.tls.enabled | bool | `true` |  |
| common.ingress.tls.secretName | string | `"stirling-pdf"` |  |
| common.livenessProbe | object | `{}` |  |
| common.livenessProbeEnabled | bool | `false` |  |
| common.name | string | `"stirling-pdf"` |  |
| common.persistence.enabled | bool | `true` |  |
| common.persistence.volumes[0].containerMount | string | `"/usr/share/tessdata"` |  |
| common.persistence.volumes[0].name | string | `"trainingdata"` |  |
| common.persistence.volumes[0].pvcClaim | string | `""` |  |
| common.persistence.volumes[0].size | string | `"10Gi"` |  |
| common.persistence.volumes[0].storageClassName | string | `""` |  |
| common.persistence.volumes[1].containerMount | string | `"/configs"` |  |
| common.persistence.volumes[1].name | string | `"config"` |  |
| common.persistence.volumes[1].pvcClaim | string | `""` |  |
| common.persistence.volumes[1].size | string | `"1Gi"` |  |
| common.persistence.volumes[1].storageClassName | string | `""` |  |
| common.persistence.volumes[2].containerMount | string | `"/logs"` |  |
| common.persistence.volumes[2].name | string | `"log"` |  |
| common.persistence.volumes[2].pvcClaim | string | `""` |  |
| common.persistence.volumes[2].size | string | `"1Gi"` |  |
| common.persistence.volumes[2].storageClassName | string | `""` |  |
| common.readinessProbe | object | `{}` |  |
| common.readinessProbeEnabled | bool | `false` |  |
| common.service.containerPort | int | `8080` |  |
| common.service.servicePort | int | `8080` |  |
| common.startupProbe | object | `{}` |  |
| common.startupProbeEnabled | bool | `false` |  |
| common.variables.nonSecret.DOCKER_ENABLE_SECURITY | bool | `false` |  |
| common.variables.nonSecret.INSTALL_BOOK_AND_ADVANCED_HTML_OPS | bool | `true` |  |
| common.variables.nonSecret.LANGS | string | `"en_GB"` |  |

