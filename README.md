# Pulumi AWS Hosted Zone

A simple Pulumi program to provision a hosted zone in Route 53.

## Required config values

- aws:region
- domain

These can be set with the following shell command:

```sh
pulumi config set KEY VALUE
```

For example:

```sh
pulumi config set aws:region us-west-2
```
