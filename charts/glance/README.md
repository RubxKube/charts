# glance

![Version: 0.0.1](https://img.shields.io/badge/Version-0.0.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.36.5](https://img.shields.io/badge/AppVersion-0.36.5-informational?style=flat-square)

A self-hosted dashboard that puts all your feeds in one place

**Homepage:** <https://github.com/glanceapp/glance>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> | <https://une-tasse-de.cafe> |

## Source Code

* <https://github.com/glanceapp/glance>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://rubxkube.github.io/common-charts | common | v0.4.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| common.configMap.data[0].data[0].content."glance.yml" | string | `"pages:\n  - name: Home\n    columns:\n      - size: small\n        widgets:\n          - type: calendar\n\n          - type: rss\n            limit: 10\n            collapse-after: 3\n            cache: 3h\n            feeds:\n              - url: https://ciechanow.ski/atom.xml\n              - url: https://www.joshwcomeau.com/rss.xml\n                title: Josh Comeau\n              - url: https://samwho.dev/rss.xml\n              - url: https://awesomekling.github.io/feed.xml\n              - url: https://ishadeed.com/feed.xml\n                title: Ahmad Shadeed\n\n          - type: twitch-channels\n            channels:\n              - cuistops\n\n      - size: full\n        widgets:\n          - type: hacker-news\n\n          - type: videos\n            channels:\n              - UCR-DXc1voovS8nhAvccRZhg # Jeff Geerling\n              - UCv6J_jJa8GJqFwQNgNrMuww # ServeTheHome\n              - UCOk-gHyjcWZNj3Br4oxwh0A # Techno Tim\n\n          - type: reddit\n            subreddit: selfhosted\n\n      - size: small\n        widgets:\n          - type: weather\n            location: London, United Kingdom\n\n          - type: markets\n            markets:\n              - symbol: SPY\n                name: S&P 500\n              - symbol: BTC-USD\n                name: Bitcoin\n              - symbol: NVDA\n                name: NVIDIA\n              - symbol: AAPL\n                name: Apple\n              - symbol: MSFT\n                name: Microsoft\n              - symbol: GOOGL\n                name: Google\n              - symbol: AMD\n                name: AMD\n              - symbol: RDDT\n                name: Reddit\n"` |  |
| common.configMap.data[0].mountPath | string | `"/mnt"` |  |
| common.configMap.data[0].name | string | `"config"` |  |
| common.configMap.enabled | bool | `true` |  |
| common.deployment.args[0] | string | `"--config"` |  |
| common.deployment.args[1] | string | `"/mnt/glance.yml"` |  |
| common.deployment.port | int | `8080` |  |
| common.image.pullPolicy | string | `"IfNotPresent"` |  |
| common.image.repository | string | `"glanceapp/glance"` |  |
| common.image.tag | string | `"v0.6.3"` |  |
| common.ingress.annotations | object | `{}` |  |
| common.ingress.enabled | bool | `true` |  |
| common.ingress.extraLabels | object | `{}` |  |
| common.ingress.hostName | string | `"glance.your-domain.com"` |  |
| common.ingress.ingressClassName | string | `""` |  |
| common.ingress.tls.enabled | bool | `true` |  |
| common.ingress.tls.secretName | string | `"glance"` |  |
| common.livenessProbe | object | `{}` |  |
| common.livenessProbeEnabled | bool | `false` |  |
| common.name | string | `"glance"` |  |
| common.persistence.enabled | bool | `true` |  |
| common.persistence.volumes | list | `[]` |  |
| common.readinessProbe | object | `{}` |  |
| common.readinessProbeEnabled | bool | `false` |  |
| common.service.containerPort | int | `8080` |  |
| common.service.servicePort | int | `8080` |  |
| common.startupProbe | object | `{}` |  |
| common.startupProbeEnabled | bool | `false` |  |

