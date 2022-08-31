---
author: Jacob Tomlinson
date: 2017-06-14T00:00:00+00:00
categories:
  - Security
tags:
  - hashicorp
  - vault
  - token
  - security
  - linux
  - infrastructure
thumbnail: vault
title: How to create a seal only token for Hashicorp Vault
aliases:
  - /2017/06/14/seal-only-token-hashicorp-vault/
---


## Introduction

When using [Hashicorp's Vault][vault] you may want to have an authentication token which only has permissions to seal the vault. This can then be used in an emergency situation to seal the vault, [perhaps through a chatbot][opsdroid-skill-vault].

### The policy

The seal only policy is fairly simple. Just create a `.hcl` policy file with the following contents:

```hcl
path "sys/seal" {
  capabilities = ["update", "sudo"]
}
```

### Create the policy

Create a new policy in vault using the policy file you just created.

```console
$ vault policy-write seal-only /path/to/my/policy.hcl
```

### Generate a token

You can now generate tokens which only have the seal permission. You must do this with a root key or a user with `sudo` permissions on `auth/token/create`.
This will print out a new token with seal only permissions.

```console
$ vault token-create -orphan -policy="seal-only"
Key            	Value
---            	-----
token          	abcdefgh-1234-5678-abcd-zyxwvutrspqo
token_accessor 	abcdefgh-1234-5678-abcd-zyxwvutrspqo
token_duration 	168h0m0s
token_renewable	true
token_policies 	[default seal-only]
```

### Renewing

As you can see this token will expire after 7 days. If this token is being used by a bot or similar system you probably want to implement some scheduled process to [renew the token's lease][vault-renew-token].

[opsdroid-skill-vault]: https://github.com/opsdroid/skill-vault
[vault]: https://www.vaultproject.io
[vault-renew-token]: https://www.vaultproject.io/docs/auth/token.html#auth-token-renew-self
