# PolicyHub CLI

This is the home of the PolicyHub CLI, a CLI tool that makes Rego policies searchable.

## Goals

Policies are everywhere. Compliance policies, security policies, policies that define organisational best practices. The Open Policy Agent project provided a single policy language, Rego, that can be used to automate policy enforcement. However currently there is no existing mechanism that allows you to search for specific Rego policies.

For example you might be looking for a set of policies that validate Kubernetes security best practices as a starting point for your organisations Kubernetes policies. Or you might be looking for a set of Microservice Authorization policies. Right now you have to hope that your google search points you in the right direction.

The PolicyHub CLI aims to make policies searchable. We provide a standard format for policy creators to share their policies. Users of the CLI can search our registry for specific tags or descriptions, hopefully finding the policy they where looking for.

## Searching policies

To search our registry, you can use the `search` command:

```bash
> policy-hub search

+---------------------------+---------------------------------+--------------------------------+
|           NAME            |           MAINTAINERS           |             LABELS             |
+---------------------------+---------------------------------+--------------------------------+
| deprek8ion                | https://github.com/swade1987    | k8s, kubernetes, gatekeeper    |
| contrib.k8s_node_selector | https://github.com/tsandall     | kubernetes, k8s, node_selector |
| redhat-cop.rego-policies  | https://github.com/garethahealy | k8s, kubernetes, gatekeeper    |
| konstraint                | https://github.com/garethahealy | k8s, kubernetes, gatekeeper    |
+---------------------------+---------------------------------+--------------------------------+
```

## Downloading policies

To download a policy, use the `pull` command:

```bash
> policy-hub pull konstraint
```

## Contributing

Join us make policies more searchable!

- We accept contributions to our registry.
- Use [GitHub Issues](https://github.com/policy-hub/policy-hub-cli/issues) to file bugs or propose new features.
- Create a [Pull Request](https://github.com/policy-hub/policy-hub-cli/pulls) and contribute to the project.
