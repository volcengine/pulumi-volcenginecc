
# Volcenginecc Resource Provider

---
[![Volcengine SKILL & CLI Survey](img_new.png)](https://www.volcengine.com/survey/q/v2/7662656599489085475/q36U8r81/5ffc/channel=doc)

To improve your experience with Volcengine skill, we invite you to participate in our [Survey with Prizes](https://www.volcengine.com/survey/q/v2/7662656599489085475/q36U8r81/5ffc/channel=doc) .

---

The Volcenginecc resource provider for Pulumi lets you use volcengine resources in your cloud programs.  To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @volcengine/pulumi-volcenginecc

or `yarn`:

    $ yarn add @volcengine/pulumi-volcenginecc

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_volcenginecc

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/volcengine/pulumi-volcengine/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Volcenginecc


## Configuration Reference

In addition to generic `provider` arguments
(e.g. `alias` and `version`), the following arguments are supported in the Volcenginecc
provider configuration:

### Optional

- `accessKey` (String) The Access Key for Volcengine Provider. It must be provided, but it can also be sourced from the `VOLCENGINE_ACCESS_KEY` environment variable
- `secretKey` (String) The Secret Key for Volcengine Provider. It must be provided, but it can also be sourced from the `VOLCENGINE_SECRET_KEY` environment variable
- `sessionToken` (String) The Session Token for Volcengine Provider, used together with AccessKey and SecretKey when authenticating with temporary STS credentials. It can also be sourced from the `VOLCENGINE_SESSION_TOKEN` environment variable
- `profile` (String) The Profile for Volcengine Provider. It can be used as an alternative authentication method to AK/SK, and can also be sourced from the `VOLCENGINE_PROFILE` environment variable
- `file_path` (String) The File Path for Volcengine Provider. It specifies the path to the profile configuration file. If not specified, the default path `~/.volcengine` will be used, and can also be sourced from the `VOLCENGINE_FILE_PATH` environment variable
- `assumeRole` (Attributes) An `assume_role` block (documented below). Only one `assume_role` block may be in the configuration. (see [below for nested schema](#nestedatt--assume_role))
- `customerHeaders` (String) CUSTOMER HEADERS for Volcengine Provider. The customer_headers field uses commas (,) to separate multiple headers, and colons (:) to separate each header key from its corresponding value.
- `disableSsl` (Boolean) Disable SSL for Volcengine Provider
- `endpoints` (Attributes) An `endpoints` block (documented below). Only one `endpoints` block may be in the configuration. (see [below for nested schema](#nestedatt--endpoints))
- `proxyUrl` (String) PROXY URL for Volcengine Provider
- `region` (String) The Region for Volcengine Provider. It must be provided, but it can also be sourced from the `VOLCENGINE_REGION` environment variable

### Credentials and state encryption

The provider marks `accessKey`, `secretKey`, and `sessionToken` as Pulumi secrets,
including values supplied through environment variables. Pulumi encrypts these
provider inputs and outputs in the stack state using the stack's secrets provider.
Environment variables can still be copied into provider state; they do not prevent
credentials from being persisted.

Store stack configuration as secrets as well. Run these commands without a value
argument to enter credentials at the hidden prompt:

```sh
pulumi config set --secret volcenginecc:accessKey
pulumi config set --secret volcenginecc:secretKey
# Only when using temporary STS credentials:
pulumi config set --secret volcenginecc:sessionToken
```

When constructing a provider explicitly, use secret inputs, especially with older
provider versions. For example, in TypeScript:

```typescript
const config = new pulumi.Config("volcenginecc");
const provider = new volcenginecc.Provider("volcenginecc", {
    accessKey: config.requireSecret("accessKey"),
    secretKey: config.requireSecret("secretKey"),
    sessionToken: config.getSecret("sessionToken"),
});
```

Earlier versions did not mark these fields as secrets. To migrate an affected
stack, upgrade both the provider plugin and language SDK to a release containing
this fix, convert any plaintext stack configuration using the commands above,
review `pulumi preview`, and run `pulumi up`. Verify that the provider's credential
fields in both `inputs` and `outputs` contain `ciphertext` objects in a normal
`pulumi stack export`. Do not use `--show-secrets` for this check: that option
intentionally decrypts secrets.

Updating the current state does not erase plaintext credentials from older state
versions, backups, exports, or logs. If credentials were exposed, rotate them and
handle those copies according to your backend's retention and access policies.


<a id="nestedatt--assume_role"></a>

### Nested Schema for `assume_role`

Required:

- `assumeRoleTrn` (String) he TRN of the role to assume.
- `durationSeconds` (Number) The duration of the session when making the AssumeRole call. Its value ranges from 900 to 43200(seconds), and default is 3600 seconds.
  Optional:

- `policy` (String) A more restrictive policy when making the AssumeRole call

<a id="nestedatt--endpoints"></a>

### Nested Schema for `endpoints`

Optional:

- `cloudcontrolapi` (String) Use this to override the default Cloud Control API service endpoint URL
- `sts` (String) Use this to override the default STS service endpoint URL


## Security and privacy

This project takes security seriously.
For vulnerability reporting and supported versions, see [SECURITY.md](SECURITY.md)
