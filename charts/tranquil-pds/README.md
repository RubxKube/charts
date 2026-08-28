# tranquil-pds

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: latest](https://img.shields.io/badge/AppVersion-latest-informational?style=flat-square)

Tranquil PDS (AT Protocol Personal Data Server, Rust) with a bundled PostgreSQL and optional OIDC SSO.

User handles are served as subdomains of `ingress.host`, so the chart issues a
**wildcard** TLS certificate (`*.host`) via cert-manager. All application secrets
(database password, JWT/DPoP/master keys, OIDC client credentials, and the image
pull `dockerconfigjson`) are pulled from an external-secrets store — the chart
creates the `Secret`s, never the values.

**Homepage:** <https://tangled.org/tranquil.farm/tranquil-pds>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> |  |

## Source Code

* <https://tangled.org/tranquil.farm/tranquil-pds>

## Requirements

The target cluster must provide:

* [external-secrets](https://external-secrets.io/) with a reachable store (a
  `ClusterSecretStore` named `vault-backend` by default).
* [cert-manager](https://cert-manager.io/) with a `ClusterIssuer` able to solve
  **DNS-01** (required to issue the wildcard certificate).

The secret referenced by `vault.key` in the store must contain: `db_password`,
`jwt_secret`, `dpop_secret`, `master_key`, `dockerconfigjson`, and — when
`sso.enabled` — `authentik_id` and `authentik_secret`.

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| image | string | `"atcr.io/tranquil.farm/tranquil-pds:latest"` | Container image for the PDS. |
| imagePullPolicy | string | `"IfNotPresent"` | Image pull policy. |
| imagePullSecret | string | `"atcr-pull"` | Name of the imagePullSecret (materialized from the store). |
| secretName | string | `"tranquil-pds-secrets"` | Name of the app Secret (materialized from the store). |
| storageClass | string | `""` | StorageClass for the blob and database volumes (empty = cluster default). |
| app.blobStorage | string | `"10Gi"` | Storage size for the blob PVC. |
| app.resources | object | requests 100m/256Mi, limit 1Gi | Resources for the PDS container. |
| db.image | string | `"postgres:18-alpine"` | PostgreSQL image. |
| db.storage | string | `"10Gi"` | Storage size for the database PVC. |
| db.resources | object | requests 50m/128Mi, limit 512Mi | Resources for the PostgreSQL container. |
| ingress.host | string | `"pds.example.com"` | Public hostname; handles are served as `*.host`. |
| ingress.className | string | `"nginx"` | IngressClass name. |
| ingress.clusterIssuer | string | `"letsencrypt-prod"` | cert-manager ClusterIssuer for the wildcard certificate. |
| ingress.tlsSecretName | string | `"tranquil-pds-tls"` | Name of the TLS Secret produced by the Certificate. |
| sso.enabled | bool | `true` | Enable OIDC single sign-on. |
| sso.issuer | string | `""` | OIDC issuer URL (the exact `iss` validated by the PDS). |
| sso.displayName | string | `"SSO"` | Display name of the SSO button. |
| secretStore.name | string | `"vault-backend"` | Name of the external-secrets store. |
| secretStore.kind | string | `"ClusterSecretStore"` | Kind of the store. |
| secretStore.refreshInterval | string | `"1h"` | Refresh interval for the ExternalSecrets. |
| vault.key | string | `"tranquil-pds"` | Remote key (secret path) in the store. |
