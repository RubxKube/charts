# kyoo

![Version: 0.1.4](https://img.shields.io/badge/Version-0.1.4-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 4.7.0](https://img.shields.io/badge/AppVersion-4.7.0-informational?style=flat-square)

Kyoo is a media manager and transcoder for your media files.

**Homepage:** <https://github.com/zoriya/Kyoo>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-pause-cafe.fr> |  |

## Source Code

* <https://github.com/zoriya/Kyoo>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | postgresql | 16.2.5 |
| https://charts.bitnami.com/bitnami | rabbitmq | 15.1.0 |
| https://meilisearch.github.io/meilisearch-kubernetes | meilisearch | 0.10.2 |
| https://rubxkube.github.io/common-charts | front(common) | v0.3.11 |
| https://rubxkube.github.io/common-charts | back(common) | v0.3.11 |
| https://rubxkube.github.io/common-charts | transcoder(common) | v0.3.11 |
| https://rubxkube.github.io/common-charts | scanner(common) | v0.3.11 |
| https://rubxkube.github.io/common-charts | autosync(common) | v0.3.11 |
| https://rubxkube.github.io/common-charts | matcher(common) | v0.3.11 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| autosync.cpuLimit | string | `"250m"` |  |
| autosync.enabled | bool | `true` |  |
| autosync.image.pullPolicy | string | `"IfNotPresent"` |  |
| autosync.image.repository | string | `"ghcr.io/zoriya/kyoo_autosync"` |  |
| autosync.image.tag | string | `"4.7.0"` |  |
| autosync.memoryLimit | string | `"300Mi"` |  |
| autosync.name | string | `"autosync"` |  |
| autosync.variables.nonSecret.RABBITMQ_HOST | string | `"kyoo-rabbitmq"` |  |
| autosync.variables.nonSecret.RABBITMQ_PORT | string | `"5672"` |  |
| autosync.variables.secret.existingSecret[0].envName | string | `"RABBITMQ_DEFAULT_USER"` |  |
| autosync.variables.secret.existingSecret[0].key | string | `"rabbitmq_user"` |  |
| autosync.variables.secret.existingSecret[0].name | string | `"kyoo-rabbitmq"` |  |
| autosync.variables.secret.existingSecret[1].envName | string | `"RABBITMQ_DEFAULT_PASS"` |  |
| autosync.variables.secret.existingSecret[1].key | string | `"rabbitmq_password"` |  |
| autosync.variables.secret.existingSecret[1].name | string | `"kyoo-rabbitmq"` |  |
| back.deployment.cpuLimit | string | `"1000m"` |  |
| back.deployment.initContainers[0].env[0].name | string | `"POSTGRES_USER"` |  |
| back.deployment.initContainers[0].env[0].value | string | `"kyoouserherefordb"` |  |
| back.deployment.initContainers[0].env[1].name | string | `"POSTGRES_PASSWORD"` |  |
| back.deployment.initContainers[0].env[1].valueFrom.secretKeyRef.key | string | `"password"` |  |
| back.deployment.initContainers[0].env[1].valueFrom.secretKeyRef.name | string | `"kyoo-postgresql"` |  |
| back.deployment.initContainers[0].env[2].name | string | `"POSTGRES_DB"` |  |
| back.deployment.initContainers[0].env[2].value | string | `"kyoo_back"` |  |
| back.deployment.initContainers[0].env[3].name | string | `"POSTGRES_SERVER"` |  |
| back.deployment.initContainers[0].env[3].value | string | `"kyoo-postgresql"` |  |
| back.deployment.initContainers[0].env[4].name | string | `"POSTGRES_PORT"` |  |
| back.deployment.initContainers[0].env[4].value | string | `"5432"` |  |
| back.deployment.initContainers[0].image | string | `"ghcr.io/zoriya/kyoo_migrations:4.7.0"` |  |
| back.deployment.initContainers[0].name | string | `"migrations"` |  |
| back.deployment.memoryLimit | string | `"2000Mi"` |  |
| back.deployment.port | int | `5000` |  |
| back.enabled | bool | `true` |  |
| back.image.pullPolicy | string | `"IfNotPresent"` |  |
| back.image.repository | string | `"ghcr.io/zoriya/kyoo_back"` |  |
| back.image.tag | string | `"4.7.0"` |  |
| back.name | string | `"back"` |  |
| back.persistence.enabled | bool | `true` |  |
| back.persistence.volumes[0].containerMount | string | `"/metadata"` |  |
| back.persistence.volumes[0].name | string | `"back"` |  |
| back.persistence.volumes[0].pvcClaim | string | `""` |  |
| back.persistence.volumes[0].size | string | `"5Gi"` |  |
| back.persistence.volumes[0].storageClassName | string | `"longhorn"` |  |
| back.service.containerPort | int | `5000` |  |
| back.service.servicePort | int | `5000` |  |
| back.variables.nonSecret.DEFAULT_PERMISSIONS | string | `"overall.read,overall.play"` |  |
| back.variables.nonSecret.MEILI_HOST | string | `"http://kyoo-meilisearch:7700"` |  |
| back.variables.nonSecret.POSTGRES_DB | string | `"kyoo_back"` |  |
| back.variables.nonSecret.POSTGRES_PORT | string | `"5432"` |  |
| back.variables.nonSecret.POSTGRES_SERVER | string | `"kyoo-postgresql"` |  |
| back.variables.nonSecret.PUBLIC_URL | string | `"https://kyoo.mydomain.com"` |  |
| back.variables.nonSecret.RABBITMQ_HOST | string | `"kyoo-rabbitmq"` |  |
| back.variables.nonSecret.RABBITMQ_PORT | string | `"5672"` |  |
| back.variables.nonSecret.REQUIRE_ACCOUNT_VERIFICATION | string | `"true"` |  |
| back.variables.nonSecret.TRANSCODER_URL | string | `"http://transcoder:7666"` |  |
| back.variables.secret.data.POSTGRES_USER | string | `"kyoouserherefordb"` |  |
| back.variables.secret.existingSecret[0].envName | string | `"KYOO_APIKEYS"` |  |
| back.variables.secret.existingSecret[0].key | string | `"kyoo-apikeys"` |  |
| back.variables.secret.existingSecret[0].name | string | `"kyoo-secrets"` |  |
| back.variables.secret.existingSecret[1].envName | string | `"POSTGRES_PASSWORD"` |  |
| back.variables.secret.existingSecret[1].key | string | `"password"` |  |
| back.variables.secret.existingSecret[1].name | string | `"kyoo-postgresql"` |  |
| back.variables.secret.existingSecret[2].envName | string | `"RABBITMQ_DEFAULT_USER"` |  |
| back.variables.secret.existingSecret[2].key | string | `"rabbitmq_user"` |  |
| back.variables.secret.existingSecret[2].name | string | `"kyoo-rabbitmq"` |  |
| back.variables.secret.existingSecret[3].envName | string | `"RABBITMQ_DEFAULT_PASS"` |  |
| back.variables.secret.existingSecret[3].key | string | `"rabbitmq_password"` |  |
| back.variables.secret.existingSecret[3].name | string | `"kyoo-rabbitmq"` |  |
| back.variables.secret.existingSecret[4].envName | string | `"MEILI_MASTER_KEY"` |  |
| back.variables.secret.existingSecret[4].key | string | `"MEILI_MASTER_KEY"` |  |
| back.variables.secret.existingSecret[4].name | string | `"kyoo-meilisearch"` |  |
| config.data.LIBRARY_IGNORE_PATTERN | string | `".*/[dD]ownloads?/.*"` |  |
| config.data.LIBRARY_LANGUAGES | string | `"en"` |  |
| config.enabled | bool | `true` |  |
| config.name | string | `"kyoo-config"` |  |
| extraResources | list | `[]` |  |
| front.deployment.cpuLimit | string | `"250m"` |  |
| front.deployment.memoryLimit | string | `"100Mi"` |  |
| front.deployment.port | int | `8901` |  |
| front.enabled | bool | `true` |  |
| front.image.pullPolicy | string | `"IfNotPresent"` |  |
| front.image.repository | string | `"ghcr.io/zoriya/kyoo_front"` |  |
| front.image.tag | string | `"4.7.0"` |  |
| front.name | string | `"front"` |  |
| front.service.containerPort | int | `8901` |  |
| front.service.enabled | bool | `true` |  |
| front.service.extraLabels | object | `{}` |  |
| front.service.servicePort | int | `8901` |  |
| front.service.type | string | `"ClusterIP"` |  |
| front.variables.nonSecret.KYOO_URL | string | `"http://back:5000"` |  |
| matcher.deployment.args[0] | string | `"matcher"` |  |
| matcher.deployment.cpuLimit | string | `"500m"` |  |
| matcher.deployment.memoryLimit | string | `"300Mi"` |  |
| matcher.enabled | bool | `true` |  |
| matcher.image.pullPolicy | string | `"IfNotPresent"` |  |
| matcher.image.repository | string | `"ghcr.io/zoriya/kyoo_scanner"` |  |
| matcher.image.tag | string | `"4.7.0"` |  |
| matcher.name | string | `"matcher"` |  |
| matcher.variables.configMap.existingConfigMap[0] | string | `"kyoo-config"` |  |
| matcher.variables.nonSecret.KYOO_URL | string | `"http://back:5000"` |  |
| matcher.variables.nonSecret.RABBITMQ_HOST | string | `"kyoo-rabbitmq"` |  |
| matcher.variables.nonSecret.RABBITMQ_PORT | string | `"5672"` |  |
| matcher.variables.secret.data | object | `{}` |  |
| matcher.variables.secret.existingSecret[0].envName | string | `"RABBITMQ_DEFAULT_USER"` |  |
| matcher.variables.secret.existingSecret[0].key | string | `"rabbitmq_user"` |  |
| matcher.variables.secret.existingSecret[0].name | string | `"kyoo-rabbitmq"` |  |
| matcher.variables.secret.existingSecret[1].envName | string | `"RABBITMQ_DEFAULT_PASS"` |  |
| matcher.variables.secret.existingSecret[1].key | string | `"rabbitmq_password"` |  |
| matcher.variables.secret.existingSecret[1].name | string | `"kyoo-rabbitmq"` |  |
| matcher.variables.secret.existingSecret[2].envName | string | `"KYOO_APIKEYS"` |  |
| matcher.variables.secret.existingSecret[2].key | string | `"kyoo-apikeys"` |  |
| matcher.variables.secret.existingSecret[2].name | string | `"kyoo-secrets"` |  |
| matcher.variables.secret.existingSecret[3].envName | string | `"THEMOVIEDB_APIKEY"` |  |
| matcher.variables.secret.existingSecret[3].key | string | `"tmdb-api"` |  |
| matcher.variables.secret.existingSecret[3].name | string | `"kyoo-secrets"` |  |
| matcher.variables.secret.extraExistingSecrets | string | `nil` |  |
| meilisearch.auth.existingMasterKeySecret | string | `"kyoo-meilisearch"` |  |
| meilisearch.enabled | bool | `true` |  |
| meilisearch.environment.MEILI_ENV | string | `"production"` |  |
| meilisearch.persistence.enabled | bool | `true` |  |
| meilisearch.persistence.size | string | `"3Gi"` |  |
| postgresql.auth.username | string | `"kyoouserherefordb"` |  |
| postgresql.back.database | string | `"kyoo_back"` |  |
| postgresql.enabled | bool | `true` |  |
| postgresql.primary.initdb.scripts."kyoo_back.sql" | string | `"CREATE DATABASE {{ .Values.back.database }} WITH OWNER {{ .Values.auth.username }};\n\\connect {{ .Values.back.database }};\nCREATE SCHEMA IF NOT EXISTS data AUTHORIZATION {{ .Values.auth.username }};\n"` |  |
| postgresql.primary.initdb.scripts."kyoo_transcoder.sql" | string | `"CREATE DATABASE {{ .Values.transcoder.database }} WITH OWNER {{ .Values.auth.username }};\n\\connect {{ .Values.transcoder.database }};\nREVOKE ALL ON SCHEMA public FROM PUBLIC;\nCREATE SCHEMA IF NOT EXISTS data AUTHORIZATION {{ .Values.auth.username }};\n"` |  |
| postgresql.primary.initdb.scripts."user.sql" | string | `"ALTER ROLE {{ .Values.auth.username }}\nIN DATABASE {{ .Values.back.database }} SET search_path TO \"$user\", public;\nALTER ROLE {{ .Values.auth.username }}\nIN DATABASE {{ .Values.transcoder.database }} SET search_path TO \"$user\", data;\n"` |  |
| postgresql.primary.persistence.size | string | `"3Gi"` |  |
| postgresql.transcoder.database | string | `"kyoo_transcoder"` |  |
| rabbitmq.auth.existingErlangSecret | string | `"kyoo-rabbitmq"` |  |
| rabbitmq.auth.existingPasswordSecret | string | `"kyoo-rabbitmq"` |  |
| rabbitmq.auth.existingSecretErlangKey | string | `"erlang_cookie"` |  |
| rabbitmq.auth.existingSecretPasswordKey | string | `"rabbitmq_password"` |  |
| rabbitmq.auth.username | string | `"kyoo"` |  |
| rabbitmq.enabled | bool | `true` |  |
| route.deployment.annotations | object | `{}` |  |
| route.deployment.extraLabels | object | `{}` |  |
| route.deployment.replicaCount | int | `1` |  |
| route.enabled | bool | `true` |  |
| route.image.repository | string | `"haproxy"` |  |
| route.image.tag | string | `"latest"` |  |
| route.ingress.annotations | object | `{}` |  |
| route.ingress.enabled | bool | `true` |  |
| route.ingress.extraLabels | object | `{}` |  |
| route.ingress.hostName | string | `"kyoo.mydomain.com"` |  |
| route.ingress.ingressClassName | string | `"nginx"` |  |
| route.ingress.tls.enabled | bool | `false` |  |
| route.ingress.tls.secretName | string | `"kyoo-tls"` |  |
| route.serviceType | string | `"ClusterIP"` |  |
| scanner.deployment.cpuLimit | string | `"250m"` |  |
| scanner.deployment.memoryLimit | string | `"300Mi"` |  |
| scanner.enabled | bool | `true` |  |
| scanner.image.pullPolicy | string | `"IfNotPresent"` |  |
| scanner.image.repository | string | `"ghcr.io/zoriya/kyoo_scanner"` |  |
| scanner.image.tag | string | `"4.7.0"` |  |
| scanner.name | string | `"scanner"` |  |
| scanner.persistence.enabled | bool | `true` |  |
| scanner.persistence.volumes[0].containerMount | string | `"/data"` |  |
| scanner.persistence.volumes[0].name | string | `"media"` |  |
| scanner.persistence.volumes[0].pvcClaim | string | `"media"` |  |
| scanner.persistence.volumes[0].size | string | `"5Gi"` |  |
| scanner.persistence.volumes[0].storageClassName | string | `""` |  |
| scanner.variables.configMap.existingConfigMap[0] | string | `"kyoo-config"` |  |
| scanner.variables.nonSecret.KYOO_URL | string | `"http://back:5000"` |  |
| scanner.variables.nonSecret.POSTGRES_PORT | string | `"5432"` |  |
| scanner.variables.nonSecret.RABBITMQ_HOST | string | `"kyoo-rabbitmq"` |  |
| scanner.variables.nonSecret.RABBITMQ_PORT | string | `"5672"` |  |
| scanner.variables.nonSecret.SCANNER_LIBRARY_ROOT | string | `"/data"` |  |
| scanner.variables.secret.data.POSTGRES_USER | string | `"kyoouserherefordb"` |  |
| scanner.variables.secret.existingSecret[0].envName | string | `"KYOO_APIKEYS"` |  |
| scanner.variables.secret.existingSecret[0].key | string | `"kyoo-apikeys"` |  |
| scanner.variables.secret.existingSecret[0].name | string | `"kyoo-secrets"` |  |
| scanner.variables.secret.existingSecret[1].envName | string | `"POSTGRES_PASSWORD"` |  |
| scanner.variables.secret.existingSecret[1].key | string | `"password"` |  |
| scanner.variables.secret.existingSecret[1].name | string | `"kyoo-postgresql"` |  |
| scanner.variables.secret.existingSecret[2].envName | string | `"RABBITMQ_DEFAULT_USER"` |  |
| scanner.variables.secret.existingSecret[2].key | string | `"rabbitmq_user"` |  |
| scanner.variables.secret.existingSecret[2].name | string | `"kyoo-rabbitmq"` |  |
| scanner.variables.secret.existingSecret[3].envName | string | `"RABBITMQ_DEFAULT_PASS"` |  |
| scanner.variables.secret.existingSecret[3].key | string | `"rabbitmq_password"` |  |
| scanner.variables.secret.existingSecret[3].name | string | `"kyoo-rabbitmq"` |  |
| secrets.kyoo.data.kyoo-apikeys | string | `"your_api_keys"` |  |
| secrets.kyoo.data.tmdb-api | string | `"your_tmdb_api"` |  |
| secrets.kyoo.enabled | bool | `true` |  |
| secrets.kyoo.name | string | `"kyoo-secrets"` |  |
| secrets.meilisearch.data.MEILI_MASTER_KEY | string | `"your_meilisearch_master_key"` |  |
| secrets.meilisearch.enabled | bool | `true` |  |
| secrets.meilisearch.name | string | `"kyoo-meilisearch"` |  |
| secrets.postgres.enabled | bool | `false` |  |
| secrets.rabbitmq.data.erlang_cookie | string | `"rabbitmq_erlang_cookie"` |  |
| secrets.rabbitmq.data.rabbitmq_password | string | `"your_rabbitmq_password"` |  |
| secrets.rabbitmq.data.rabbitmq_user | string | `"kyoo"` |  |
| secrets.rabbitmq.enabled | bool | `true` |  |
| secrets.rabbitmq.name | string | `"kyoo-rabbitmq"` |  |
| transcoder.deployment.cpuLimit | string | `"1000m"` |  |
| transcoder.deployment.emptyDir[0].containerMount | string | `"/cache"` |  |
| transcoder.deployment.emptyDir[0].name | string | `"cache"` |  |
| transcoder.deployment.memoryLimit | string | `"500Mi"` |  |
| transcoder.deployment.port | int | `7666` |  |
| transcoder.enabled | bool | `true` |  |
| transcoder.image.pullPolicy | string | `"IfNotPresent"` |  |
| transcoder.image.repository | string | `"ghcr.io/zoriya/kyoo_transcoder"` |  |
| transcoder.image.tag | string | `"4.7.0"` |  |
| transcoder.name | string | `"transcoder"` |  |
| transcoder.persistence.enabled | bool | `true` |  |
| transcoder.persistence.extraVolumes[0].containerMount | string | `"/data"` |  |
| transcoder.persistence.extraVolumes[0].name | string | `"media"` |  |
| transcoder.persistence.extraVolumes[0].pvcClaim | string | `"media"` |  |
| transcoder.persistence.extraVolumes[0].size | string | `"5Gi"` |  |
| transcoder.persistence.extraVolumes[0].storageClassName | string | `""` |  |
| transcoder.persistence.volumes[0].containerMount | string | `"/metadata"` |  |
| transcoder.persistence.volumes[0].name | string | `"metadata"` |  |
| transcoder.persistence.volumes[0].pvcClaim | string | `""` |  |
| transcoder.persistence.volumes[0].size | string | `"3Gi"` |  |
| transcoder.persistence.volumes[0].storageClassName | string | `""` |  |
| transcoder.service.containerPort | int | `7666` |  |
| transcoder.service.enabled | bool | `true` |  |
| transcoder.service.extraLabels | object | `{}` |  |
| transcoder.service.servicePort | int | `7666` |  |
| transcoder.service.type | string | `"ClusterIP"` |  |
| transcoder.variables.nonSecret.GOCODER_CACHE_ROOT | string | `"/cache"` |  |
| transcoder.variables.nonSecret.GOCODER_HWACCEL | string | `"disabled"` |  |
| transcoder.variables.nonSecret.GOCODER_METADATA_ROOT | string | `"/metadata"` |  |
| transcoder.variables.nonSecret.GOCODER_PREFIX | string | `"/video"` |  |
| transcoder.variables.nonSecret.GOCODER_PRESET | string | `"fast"` |  |
| transcoder.variables.nonSecret.GOCODER_SAFE_PATH | string | `"/data"` |  |
| transcoder.variables.nonSecret.POSTGRES_DB | string | `"kyoo_transcoder"` |  |
| transcoder.variables.nonSecret.POSTGRES_PORT | string | `"5432"` |  |
| transcoder.variables.nonSecret.POSTGRES_SCHEMA | string | `"disabled"` |  |
| transcoder.variables.nonSecret.POSTGRES_SERVER | string | `"kyoo-postgresql"` |  |
| transcoder.variables.nonSecret.RABBITMQ_HOST | string | `"kyoo-rabbitmq"` |  |
| transcoder.variables.nonSecret.RABBITMQ_PORT | string | `"5672"` |  |
| transcoder.variables.secret.data.POSTGRES_USER | string | `"kyoouserherefordb"` |  |
| transcoder.variables.secret.existingSecret[0].envName | string | `"RABBITMQ_DEFAULT_USER"` |  |
| transcoder.variables.secret.existingSecret[0].key | string | `"rabbitmq_user"` |  |
| transcoder.variables.secret.existingSecret[0].name | string | `"kyoo-rabbitmq"` |  |
| transcoder.variables.secret.existingSecret[1].envName | string | `"RABBITMQ_DEFAULT_PASS"` |  |
| transcoder.variables.secret.existingSecret[1].key | string | `"rabbitmq_password"` |  |
| transcoder.variables.secret.existingSecret[1].name | string | `"kyoo-rabbitmq"` |  |
| transcoder.variables.secret.existingSecret[2].envName | string | `"POSTGRES_PASSWORD"` |  |
| transcoder.variables.secret.existingSecret[2].key | string | `"password"` |  |
| transcoder.variables.secret.existingSecret[2].name | string | `"kyoo-postgresql"` |  |
| volumes.media.accessModes[0] | string | `"ReadWriteMany"` |  |
| volumes.media.enabled | bool | `true` |  |
| volumes.media.existingClaim | string | `""` |  |
| volumes.media.extraAnnotations | object | `{}` |  |
| volumes.media.mountPath | string | `"/data"` |  |
| volumes.media.name | string | `"media"` |  |
| volumes.media.resources.requests.storage | string | `"10Gi"` |  |
| volumes.media.storageClassName | string | `""` | If empty, will use the default storage class |

