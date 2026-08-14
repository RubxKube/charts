# searxng

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2026.7.28](https://img.shields.io/badge/AppVersion-2026.7.28-informational?style=flat-square)

SearXNG is a free internet metasearch engine which aggregates results from
more than 70 different search services. Users are neither tracked nor profiled.

**Homepage:** <https://searxng.org/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| QJOLY | <github@une-tasse-de.cafe> | <https://une-tasse-de.cafe> |

## Source Code

* <https://github.com/searxng/searxng>

## Requirements

Kubernetes: `>= 1.18`

| Repository | Name | Version |
|------------|------|---------|
| https://rubxkube.github.io/common-charts | common | v0.6.4 |
| https://valkey.io/valkey-helm | valkey | 0.11.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| common.configMap.enabled | bool | `true` | Ships `settings.yml` as a read-only ConfigMap at `/etc/searxng`. |
| common.deployment.emptyDir[0].containerMount | string | `"/var/cache/searxng"` | Backs the image's writable data path (cache) under `readOnlyRootFilesystem`. |
| common.deployment.emptyDir[1].containerMount | string | `"/tmp"` | Backs `/tmp` for granian/python temp writes under `readOnlyRootFilesystem`. |
| common.deployment.podSecurityContext.fsGroup | int | `977` | SearXNG image runs as the `searxng` user (UID/GID 977). |
| common.deployment.securityContext.allowPrivilegeEscalation | bool | `false` | Hardening. |
| common.deployment.securityContext.capabilities.drop[0] | string | `"ALL"` | Drop all Linux capabilities. |
| common.deployment.securityContext.readOnlyRootFilesystem | bool | `true` | Root FS read-only; writable paths backed by emptyDir. |
| common.deployment.securityContext.runAsGroup | int | `977` | GID of the `searxng` user. |
| common.deployment.securityContext.runAsNonRoot | bool | `true` | Run the container as a non-root user. |
| common.deployment.securityContext.runAsUser | int | `977` | UID of the `searxng` user inside the image. |
| common.deployment.strategy.type | string | `"RollingUpdate"` | Stateless frontend; zero-downtime upgrades and HPA-compatible. |
| common.hpa.enabled | bool | `false` | Safe to enable: pod is stateless, no RWO PVC. |
| common.image.pullPolicy | string | `"Always"` | |
| common.image.repository | string | `"searxng/searxng"` | Official SearXNG image. |
| common.image.tag | string | `"2026.7.28-c01178d03"` | Pinned upstream tag. |
| common.ingress.enabled | bool | `false` | Enable ingress to expose SearXNG. |
| common.ingress.hostName | string | `"searxng.example.com"` | Public hostname used for the ingress and `SEARXNG_BASE_URL`. |
| common.persistence.enabled | bool | `false` | No PVC: config is a ConfigMap; cache/tmp use emptyDir. |
| common.service.containerPort | int | `8080` | SearXNG listens on 8080 by default. |
| common.service.servicePort | int | `8080` | |
| common.variables.nonSecret.SEARXNG_BASE_URL | string | `"https://searxng.example.com/"` | Public base URL; must match your ingress scheme + hostName. |
| common.variables.nonSecret.SEARXNG_LIMITER | string | `nil` | Set `"true"` to enable bot protection (requires Valkey). |
| common.variables.nonSecret.SEARXNG_VALKEY_URL | string | `nil` | Valkey URL, e.g. `redis://searxng-valkey:6379/0` (requires `valkey.enabled`). |
| common.startupProbe.httpGet.path | string | `"/healthz"` | SearXNG ships a `/healthz` endpoint. |
| common.startupProbe.failureThreshold | int | `18` | ~180s startup window. |
| common.livenessProbe.failureThreshold | int | `3` | Tolerate transient blips before restart. |
| common.readinessProbe.httpGet.path | string | `"/healthz"` | |
| common.livenessProbe.httpGet.path | string | `"/healthz"` | |
| common.tests.curlHostHeader.enabled | bool | `true` | Hits `/healthz` with the ingress Host header. |
| common.tests.classicHttp.enabled | bool | `false` | |
| valkey.enabled | bool | `false` | Deploy the official Valkey subchart to back the rate limiter (off by default). | |

## Configuration notes

### Secret key

The `server.secret_key` is **not** sourced from the `settings.yml` file on the
PVC. SearXNG reads the `SEARXNG_SECRET` environment variable at runtime and
overrides whatever `secret_key` is in `settings.yml`, so the chart injects it
from a Kubernetes Secret (`<release>-searxng-secret`). Leaving
`common.variables.secret.data.SEARXNG_SECRET` empty auto-generates a 48-char
random secret on first install that stays stable across upgrades (the chart
reuses the existing Secret via a lookup). Set an explicit string to use your
own value, or rotate it by editing the value and restarting the pod.

To use an externally-managed Secret (Sealed Secrets / External Secrets /
Vault), drop `data:` and use `common.variables.secret.existingSecret`.

### Base URL and ingress

Set `common.variables.nonSecret.SEARXNG_BASE_URL` to the **public** URL of
your instance. It must match `common.ingress.hostName` and the scheme must
match `common.ingress.tls.enabled` (the default `https://` matches the default
`tls.enabled: true`); a mismatch causes mixed-content links and redirect
loops. To run behind an ingress, set `common.ingress.enabled: true` and
configure `hostName` and TLS.

### Settings file

SearXNG's `settings.yml` is shipped **declaratively** as a read-only
ConfigMap mounted at `/etc/searxng/settings.yml` (`common.configMap`). The
pod is therefore stateless: no config PVC, `RollingUpdate` strategy, and
horizontal scaling / HPA work without RWO multi-attach conflicts. The default
`settings.yml` uses `use_default_settings: true` and only sets `secret_key`
(overridden by `SEARXNG_SECRET`) and `image_proxy`; extend it in
`common.configMap.data` for settings that have no env override (engines,
autocomplete, formats, ...). The image's writable paths (`/var/cache/searxng`,
`/tmp`) are backed by `emptyDir` so `readOnlyRootFilesystem` can stay on.

Note: the SearXNG entrypoint prints non-fatal ownership `WARNING`s for both
`/etc/searxng` and `/var/cache/searxng` (it expects `searxng:searxng`, gets
`root:searxng`). This is expected and cosmetic. Kubernetes creates ConfigMap
projected volumes and `emptyDir` volumes owned by `root`; the pod's
`fsGroup: 977` makes the kubelet set the **group** to 977 and the mode
group-writable, but it cannot change the **owner** from `root`. Since the pod
runs non-root (`runAsUser: 977`, `readOnlyRootFilesystem: true`), nothing can
`chown` them to `977:977`. Access still works: the ConfigMap files are
world-readable (`0644`) and the cache `emptyDir` is group-writable, so UID 977
reads its config and writes its cache fine. The entrypoint just doesn't account
for `fsGroup`-based group access. `/healthz` returns 200 and the cache is
writable in practice.

### Bot protection / rate limiting (optional)

SearXNG's limiter (`server.limiter`) needs a Valkey instance. This chart
bundles the **official Valkey chart** (`valkey-io/valkey-helm`), deployed only
when you opt in (`valkey.enabled: false` by default).

To enable it, set these together (the Valkey service is named `<release>-valkey`,
e.g. `searxng-valkey:6379` for `helm install searxng ...`):

```yaml
valkey:
  enabled: true

common:
  variables:
    nonSecret:
      SEARXNG_LIMITER: "true"
      SEARXNG_VALKEY_URL: "redis://searxng-valkey:6379/0"
  deployment:
    initContainers:
      - name: wait-for-valkey          # copy the full block from values.yaml
        # ...image, securityContext, command, args from values.yaml...
        env:
          - name: SEARXNG_VALKEY_URL
            value: "redis://searxng-valkey:6379/0"   # mirror the value above
```

> **Override heads-up:** Helm **replaces lists** on merge. If you pass
> `initContainers` via `-f`/`--set`, redeclare the **whole** `wait-for-valkey`
> container (copy it from `values.yaml`), not just the `env:` snippet. Editing
> `values.yaml` directly? Just set the two `value:` lines.

Notes:
- **`SEARXNG_VALKEY_URL` must match the release name** (`<release>-valkey:6379`);
  keep it in sync in both `nonSecret` and the initContainer `env`.
- **All three go together.** A wrong URL or a forgotten `valkey.enabled: true`
  doesn't crash: SearXNG logs an error and the limiter stays silently inactive.
- **Startup race.** SearXNG initializes the limiter once at boot and doesn't
  retry Valkey, so on a fresh install it can come up before Valkey and stay
  inactive until a pod restart. The `wait-for-valkey` initContainer blocks pod
  start until Valkey is up (≤60s, then best-effort), fixing this. It no-ops when
  `SEARXNG_VALKEY_URL` is empty, so users who skip the limiter are unaffected.
  It needs the URL too (Kubernetes doesn't share env between containers), hence
  the mirror above.
- **Valkey is only for the limiter.** SearXNG is stateless otherwise, so you can
  run multiple replicas / HPA *without* Valkey. With the limiter on, Valkey keeps
  the rate-limit state consistent across pods.
- To persist the limiter cache across restarts, set `valkey.dataStorage.enabled:
  true` and a `requestedSize`.
- To use an **external/existing** Valkey, leave `valkey.enabled: false` and point
  `SEARXNG_VALKEY_URL` at it.