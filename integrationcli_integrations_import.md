## integrationcli integrations import

Import all Integration flows from folder

### Synopsis

Import all Integration flows from folder

```
integrationcli integrations import [flags]
```

### Options

```
  -f, --folder string   Folder to import Integration flows
  -h, --help            help for import
```

### Options inherited from parent commands

```
  -a, --account string       Path Service Account private key in JSON
      --apigee-integration   Use Apigee Integration; default is false (Application Integration)
      --disable-check        Disable check for newer versions
      --no-output            Disable printing API responses from the control plane
  -p, --proj string          Integration GCP Project name
  -r, --reg string           Integration region name
  -t, --token string         Google OAuth Token
```

### SEE ALSO

* [integrationcli integrations](integrationcli_integrations.md)	 - Manage integrations in a GCP project

###### Auto generated by spf13/cobra on 26-Oct-2022