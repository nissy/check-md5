# mackerel-plugin-md5
Check the MD5 checksum

### Example of mackerel-agent.conf

```
[plugin.checks.md5]
command = "/path/to/mackerel-plugin-md5 -c /path/to/md5ck.conf"
```

### Example of md5ck.conf

```
[[Files]] # PHP 5.6.30
PATH = "/usr/bin/php"
SUM = "8139d07ec8dd8aa0aaf76b51b73c86c3"

[[Files]] # CentOS Linux release 7.3.1611
PATH = "/etc/redhat-release"
SUM = "af2cb935515b9d48999ecb2b1f4122e6"
```