# Hex Plugin - Local

Plugin to run commands locally.

```
{
  "rule": "local rule example",
  "match": "disk space",
  "actions": [
    {
      "type": "hex-local",
      "command": "df -h",
      "config": {
        "env": "MY_PATH=/tmp; MY_VAR=xyz",
        "dir": "/tmp"
      }
    }
  ]
}
```
