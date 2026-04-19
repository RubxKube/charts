# Auto-Increment Version Workflow

This repository automatically manages chart versions so contributors never have to bump them manually.

## How it works

Two GitHub Actions workflows handle versioning:

- **`increase_version.yml`** — runs on every pull request to preview the version change.
- **`release.yml`** — runs on every merge to `main` and applies the version change before publishing the chart.

Both workflows follow the same logic:

### Version resolution logic

For each chart touched by the commit:

1. **Detect the latest released version** by looking at git tags matching the chart name (e.g. `my-chart-1.2.3`).
   - If no tag exists yet, the version is assumed to be `0.0.0`.

2. **Compare the minor version** between the latest git tag and the `version` field in `Chart.yaml`.
   - If they differ (the PR manually bumped the minor), the workflow resets the patch to `0` and uses that version (`x.y.0`).
   - If they are the same, the workflow increments the **patch** version by 1 (`x.y.z → x.y.(z+1)`).

3. The resolved version is written back into `Chart.yaml`.

### Example

| Situation | Tag version | Chart.yaml version | Result |
|---|---|---|---|
| Normal patch | `my-chart-1.2.5` | `1.2.5` | `1.2.6` |
| First release | *(no tag)* | `0.1.0` | `0.1.0` |
| Minor bump by contributor | `my-chart-1.2.5` | `1.3.0` | `1.3.0` |

## Bumping the minor or major version

To bump the minor (or major) version, simply update `Chart.yaml` manually in your PR:

```yaml
version: 1.3.0  # was 1.2.x
```

The workflow will detect the minor version change and keep `1.3.0` without further incrementing the patch.

## Release pipeline

After merging to `main`, the `release.yml` workflow:

1. Applies the version increment (same logic as above).
2. Lints YAML files.
3. Runs Helm chart tests against a Kind cluster.
4. Publishes the chart via [`helm/chart-releaser-action`](https://github.com/helm/chart-releaser-action), which creates a GitHub Release and a git tag.
5. Regenerates `README.md` and `index.html` on the `gh-pages` branch.
