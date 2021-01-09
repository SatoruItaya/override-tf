# override-tf
The override-tf creates `override.tf` as below,

```
terraform {
    backend "local" {
    }
}
```
Generally, Terraform uses remote storage(for example, S3) as a backend for tfstate.
But it would make some difficulties if tfstate were updated by some developers.
The `override.tf` overrides backend configuration temporary, and backend for tfstate will be local.

# Install
```
$ go instlall
```

# How to use

1. Initialization(Only for first time)
```
$ terraform init
```

2. Pull tfstate from remote to local(A file name is arbitrary.)
```
$ terraform state pull > tmp.tfstate
```

3. Create `override.tf`
```
$ override-tf
```

4. Override backend config
```
$ terraform init -reconfigure
```

A message `Successfully configured the backend "local"!` will be shown.
Then you can use local tfstate file with `-state` option like below,

```
$ terraform apply -state=tmp.tfstate
```
