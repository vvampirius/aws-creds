# aws-creds

Just simple tool to get AWS SecretAccessKey from keychain

# Usage

1. Create password item in keychain with name of AccessKeyId.
2. [Configure AWS Command Line Interface](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html) (~/.aws/config) like:

```
[default]
region = us-east-1
output = text
credential_process = /Users/user/.bin/aws-creds keychain login.keychain AUIAJJSKGVUYVASMEMBQ
```

