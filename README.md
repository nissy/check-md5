# mackerel-plugin-md5
Check the MD5 checksum

### Example of mackerel-agent.conf

```
[plugin.checks.md5]
command = "/path/to/mackerel-plugin-md5 -c /path/to/md5ck.conf"
```

### Example of md5ck.conf

```
[[Files]]
PATH = "/usr/bin/php"
HASH = "8139d07ec8dd8aa0aaf76b51b73c86c3"

[[Files]]
PATH = "/etc/redhat-release"
HASH = "af2cb935515b9d48999ecb2b1f4122e6"
```

### Save hash config

```
$ /path/to/mackerel-plugin-md5 -s -c /path/to/md5ck.conf
```

### Help

```
Usage: mackerel-plugin-md5 [options]
  -c string
        set cfgiguration file (default "md5ck.conf")
  -h    this help
  -s    save cfgiguration file md5 hash
  -v    show this build version
```