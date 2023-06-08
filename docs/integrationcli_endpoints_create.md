## integrationcli endpoints create

Create an endpoint attachments in the region

### Synopsis

Create an endpoint attachments in the region

```
integrationcli endpoints create [flags]
```

### Options

```
  -d, --description string          Endpoint attachment description
  -h, --help                        help for create
  -n, --name string                 Endpoint attachment name; Ex: sample
  -s, --service-attachment string   Endpoint attachment url; format = projects/*/regions/*/serviceAttachments/*
      --wait                        Waits for the connector to finish, with success or error; default is false
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --api api          Sets the control plane API. Must be one of prod, staging or autopush; default is prod
      --disable-check    Disable check for newer versions
      --no-output        Disable printing all statements to stdout
      --print-output     Control printing of info log statements (default true)
  -p, --proj string      Integration GCP Project name
  -r, --reg string       Integration region name
  -t, --token string     Google OAuth Token
      --verbose          Enable verbose output from integrationcli
```

### SEE ALSO

* [integrationcli endpoints](integrationcli_endpoints.md)	 - Manage endpoint attachments for connections

###### Auto generated by spf13/cobra on 29-Apr-2023