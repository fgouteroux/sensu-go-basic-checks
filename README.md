# sensu-go-basic-checks

This repo is for packaging sensu-go-basic-checks asset.

Sources:
- [go-check-plugins](https://github.com/mackerelio/go-check-plugins)
- [sensu-go-cpu-check](https://github.com/asachs01/sensu-go-cpu-check)
- [sensu-go-memory-check](https://github.com/asachs01/sensu-go-memory-check)
- [sensu-go-systemd-check](https://github.com/sardinasystems/sensu-go-systemd-check)

See individual plugin documentation in the upstream repository.

- [check-cpu](./upstream/sensu-go-cpu-check/README.md)
- [check-memory](./upstream/sensu-go-memory-check/README.md)
- [check-systemd](./upstream/sensu-go-systemd-check/README.md)
- [check-cert-file](./upstream/go-check-plugins/check-cert-file/README.md)
- [check-disk](./upstream/go-check-plugins/check-disk/README.md)
- [check-file-age](./upstream/go-check-plugins/check-file-age/README.md)
- [check-file-size](./upstream/go-check-plugins/check-file-size/README.md)
- [check-http](./upstream/go-check-plugins/check-http/README.md)
- [check-load](./upstream/go-check-plugins/check-load/README.md)
- [check-log](./upstream/go-check-plugins/check-log/README.md)
- [check-ntpoffset](./upstream/go-check-plugins/check-ntpoffset/README.md)
- [check-ntservice](./upstream/go-check-plugins/check-ntservice/README.md)
- [check-ping](./upstream/go-check-plugins/check-ping/README.md)
- [check-procs](./upstream/go-check-plugins/check-procs/README.md)
- [check-smtp](./upstream/go-check-plugins/check-smtp/README.md)
- [check-ssl-cert](./upstream/go-check-plugins/check-ssl-cert/README.md)
- [check-tcp](./upstream/go-check-plugins/check-tcp/README.md)
- [check-uptime](./upstream/go-check-plugins/check-uptime/README.md)
- [check-windows-eventlog](./upstream/go-check-plugins/check-windows-eventlog/README.md)

## Release build

Some go sources are not both compatible linux/windows.

To build linux asset:
```sh
./linux_build.sh
```

To build windows asset:
```sh
./windows_build.sh
```
