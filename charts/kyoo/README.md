## Kyoo 

Kyoo is a self-hosted media server focused on video content (Movies, Series & Anime). It is an alternative to Jellyfin or Plex.

It aims to have a low amount of maintenance needed (no folder structure required nor manual metadata edits). Media not being scanned correctly (even with weird names) is considered a bug.

### Features of the chart

- Each deployment is independent of the others, it allows:
  - To have different versions of each service running at the same time (for testing purposes, or just migrating from one version to another)
  - To have different configurations for each service
  - To deploy only the services you need and use services from another cluster (like a database, or the transcoder service)
- Each service managed its own volumes (except for the `media` volume, which is shared between multiple services), see more below

### Considerations for GitOps integration

The bitnami chart is not really GitOps friendly, as it generates a lot of resources (like secrets, configmaps, etc.) with hooks. Meaning that if you update your GitOps application (either with ArgoCD or Flux), it will re-generate these resources, and you will lose the changes made by the chart (for example, the PostgreSQL password).

To avoid this, you can yourself the secrets to avoid the chart generating them.

```yaml
postgresql:
  enabled: true
  auth:
    username: kyoouserherefordb
    password: "your_postgresql_password" # <-- Set this to a random password
```

### Use NFS for the media volume

The `media` volume is a shared volume between the services that need to access the media files. It is mounted in the `/media` and can (or may be not) be created within the chart. If you want to use an existing volume, you can specify it in the `volume.media.existingClaim` value and update values as follows :

```yaml
volume:
  # Disable media volume creation
  media:
    enabled: false

extraResources:
  - apiVersion: v1
    kind: PersistentVolume
    metadata:
      name: nfs
    spec:
      capacity:
        storage: 50Gi
      volumeMode: Filesystem
      accessModes:
        - ReadOnlyMany
      persistentVolumeReclaimPolicy: Retain
      storageClassName: local
      mountOptions:
        - hard
        - nfsvers=4.1
      nfs:
        path: /mnt/Stockage/Plex
        server: 192.168.1.90
  - apiVersion: v1
    kind: PersistentVolumeClaim
    metadata:
      name: media
    spec:
      accessModes:
        - ReadOnlyMany
      volumeMode: Filesystem
      resources:
        requests:
          storage: 50Gi
      storageClassName: local
```



### Production usage

Please be aware that this chart is still in development and may have some bugs. It is not recommended to use it in production yet.

But if you want to, here are some things to consider:
- This chart does not handle any Network Policies, you should set them up yourself.
- **PostgreSQL** generate a random password on each deployment (you can still custom the username).
- You should set up custom passwords for RabbitMQ.

```yaml
postgresql:
  auth:
    username: kyoouserherefordb # Change this
secrets:
  kyoo:
    enabled: true
    name: "kyoo-secrets"
    data:
      kyoo-apikeys: "your_api_keys"
  rabbitmq:
    data:
      # Shoudl match .rabbitmq.auth.username
      rabbitmq_user: your_rabbitmq_user
      # custom both of these
      rabbitmq_password: your_rabbitmq_password
      erlang_cookie: rabbitmq_erlang_cookie

rabbitmq:
  auth:
    username: your_rabbitmq_user
```

Few notes:
- You can always use existing secrets for the RabbitMQ, PostgreSQL and Kyoo API keys.
- The common-chart template does not fit at 100% with the Kyoo chart, so you may need to adjust some values (however, this setup should work out of the box).

