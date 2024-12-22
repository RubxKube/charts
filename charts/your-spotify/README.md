# your-spotify

![Version: 0.0.1](https://img.shields.io/badge/Version-0.0.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.12.0](https://img.shields.io/badge/AppVersion-1.12.0-informational?style=flat-square)

YourSpotify is a self-hosted application that tracks what you listen and offers you a dashboard to explore statistics about it! It's composed of a web server which polls the Spotify API every now and then and a web application on which you can explore your statistics.

**Homepage:** <https://github.com/Yooooomi/your_spotify>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> |  |

## Source Code

* <https://github.com/Yooooomi/your_spotify>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | mongodb | 16.4.0 |
| https://rubxkube.github.io/common-charts | server(common) | v0.3.12 |
| https://rubxkube.github.io/common-charts | web(common) | v0.3.12 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| mongodb.auth.enabled | bool | `false` |  |
| mongodb.enabled | bool | `true` |  |
| server.deployment.port | int | `8080` |  |
| server.image.pullPolicy | string | `"Always"` |  |
| server.image.repository | string | `"yooooomi/your_spotify_server"` |  |
| server.image.tag | string | `"1.12.0"` |  |
| server.ingress.annotations | object | `{}` |  |
| server.ingress.enabled | bool | `true` |  |
| server.ingress.extraLabels | object | `{}` |  |
| server.ingress.hostName | string | `"spotty-api.your-domain.com"` |  |
| server.ingress.ingressClassName | string | `"nginx"` |  |
| server.ingress.tls.enabled | bool | `true` |  |
| server.ingress.tls.secretName | string | `"spotty-api"` |  |
| server.livenessProbe | object | `{}` |  |
| server.livenessProbeEnabled | bool | `false` |  |
| server.name | string | `"server"` |  |
| server.persistence.enabled | bool | `false` |  |
| server.readinessProbe | object | `{}` |  |
| server.readinessProbeEnabled | bool | `false` |  |
| server.service.containerPort | int | `8080` |  |
| server.service.servicePort | int | `8080` |  |
| server.startupProbe | object | `{}` |  |
| server.startupProbeEnabled | bool | `false` |  |
| server.variables.nonSecret.API_ENDPOINT | string | `"https://spotty-api.your-domain.com"` |  |
| server.variables.nonSecret.CLIENT_ENDPOINT | string | `"https://spotty-api.your-domain.com"` |  |
| server.variables.nonSecret.MONGO_ENDPOINT | string | `"mongodb://your-spotify-mongodb:27017/your_spotify"` |  |
| server.variables.secret.data.SPOTIFY_PUBLIC | string | `""` |  |
| server.variables.secret.data.SPOTIFY_SECRET | string | `""` |  |
| web.deployment.port | int | `3000` |  |
| web.image.pullPolicy | string | `"Always"` |  |
| web.image.repository | string | `"yooooomi/your_spotify_client"` |  |
| web.image.tag | string | `"1.12.0"` |  |
| web.ingress.annotations | object | `{}` |  |
| web.ingress.enabled | bool | `true` |  |
| web.ingress.extraLabels | object | `{}` |  |
| web.ingress.hostName | string | `"spotty.your-domain.com"` |  |
| web.ingress.ingressClassName | string | `"nginx"` |  |
| web.ingress.tls.enabled | bool | `true` |  |
| web.ingress.tls.secretName | string | `"spotty"` |  |
| web.livenessProbe | object | `{}` |  |
| web.livenessProbeEnabled | bool | `false` |  |
| web.name | string | `"web"` |  |
| web.persistence.enabled | bool | `false` |  |
| web.readinessProbe | object | `{}` |  |
| web.readinessProbeEnabled | bool | `false` |  |
| web.service.containerPort | int | `3000` |  |
| web.service.servicePort | int | `3000` |  |
| web.startupProbe | object | `{}` |  |
| web.startupProbeEnabled | bool | `false` |  |
| web.variables.nonSecret.API_ENDPOINT | string | `"https://spotty-api.your-domain.com"` |  |

