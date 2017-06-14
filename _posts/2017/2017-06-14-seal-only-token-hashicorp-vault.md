---
title: 'How to create a seal only token for Hashicorp Vault'
author: Jacob Tomlinson
layout: post
category: Security
thumbnail: vault
tags:
  - hashicorp
  - vault
  - token
  - security
  - linux
  - infrastructure
---

## Introduction

When using [Hashicorp's Vault][vault] you may want to have an authentication token which only has permissions to seal the vault. This can then be used in an emergency situation to seal the vault, perhaps through a chatbot.

### The policy

The seal only policy is fairly simple. Just create a `.hcl` policy file with the following contents:

```
path "/sys/seal" {
  policy = "sudo"
}
```

### Create the policy

Create a new policy in vault using the policy file you just created.

```
vault policy-write seal-only /path/to/my/policy.hcl
```

### Generate a token

You can now generate tokens which only have the seal permission. You must do this with a root key or a user with `sudo` permissions on `auth/token/create`.

```
vault token-create -orphan -policy="seal-only"
```

This will print out a new token with seal only permissions.

```
Key            	Value
---            	-----
token          	abcdefgh-1234-5678-abcd-zyxwvutrspqo
token_accessor 	abcdefgh-1234-5678-abcd-zyxwvutrspqo
token_duration 	168h0m0s
token_renewable	true
token_policies 	[default seal-only]
```

### Renewing

As you can see this token will expire after 7 days. If this token is being used by a bot or similar system you probably want to implement some scheduled process to renew the token's lease.

[vault]: https://www.vaultproject.io
