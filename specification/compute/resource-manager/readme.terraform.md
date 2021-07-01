
## terraform

These settings apply only when `--terraform` is specified on the command line.

``` yaml $(terraform)
terraform:
    cli-name: compute
    package-name: compute
clear-output-folder: true
output-folder: $(terraform-output-folder)/compute
```

```yaml $(tag)=='package-2021-03-01' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2021-03-01/compute
```

```yaml $(tag)=='package-2020-12-01' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2020-12-01/compute
```

```yaml $(tag)=='package-2020-10-01-preview' && $(terraform)
namespace: compute
gosdk-folder: services/preview/compute/mgmt/2020-10-01-preview/compute
```

```yaml $(tag)=='profile-hybrid-2020-09-01' && $(terraform)
namespace: compute
gosdk-folder: profiles/2020-09-01/compute/mgmt/compute
```

```yaml $(tag)=='package-2020-06-30' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2020-06-30/compute
```

```yaml $(tag)=='package-2020-06-01' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2020-06-01/compute
```

```yaml $(tag)=='package-2019-12-01' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2019-12-01/compute
```

```yaml $(tag)=='package-2019-07' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2019-07-01/compute
```

```yaml $(tag)=='package-2019-03-01' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2019-03-01/compute
```

```yaml $(tag)=='package-2018-10-01' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2018-10-01/compute
```

```yaml $(tag)=='package-2018-06' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2018-06-01/compute
```

```yaml $(tag)=='package-compute-2018-04' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2018-04-01/compute
```

```yaml $(tag)=='package-compute-2017-12' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2017-12-01/compute
```

```yaml $(tag)=='package-compute-2017-03' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2017-03-30/compute
```

```yaml $(tag)=='package-compute-2016-04-preview' && $(terraform)
namespace: compute
gosdk-folder: services/preview/compute/mgmt/2016-04-30-preview/compute
```

```yaml $(tag)=='package-compute-2016-03' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2016-03-30/compute
```

```yaml $(tag)=='package-compute-2015-06' && $(terraform)
namespace: compute
gosdk-folder: services/compute/mgmt/2015-06-15/compute
```

```yaml $(tag)=='package-skus-2017-09' && $(terraform)
namespace: skus
gosdk-folder: services/compute/mgmt/2017-09-01/skus
```

```yaml $(tag)=='package-container-service-2017-01' && $(terraform)
namespace: containerservice
gosdk-folder: services/containerservice/mgmt/2017-01-31/containerservice
```

```yaml $(tag)=='package-container-service-2016-09' && $(terraform)
namespace: containerservice
gosdk-folder: services/containerservice/mgmt/2016-09-30/containerservice
```

```yaml $(tag)=='package-container-service-2016-03' && $(terraform)
namespace: containerservice
gosdk-folder: services/containerservice/mgmt/2016-03-30/containerservice
```

```yaml $(tag)=='package-container-service-2015-11-preview' && $(terraform)
namespace: containerservice
gosdk-folder: services/preview/containerservice/mgmt/2015-11-01-preview/containerservice
```
