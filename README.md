# go-starter

An golang starter project with viper, cobra, and github actions integrations

## Deploy

## CLI Flags

```text
The go-starter command

Usage:
  go-starter [flags]

Flags:
  -c, --config string                 Config file (default "config.yml")
      --data-directory string         Directory that holds data (default "deploy/data/")
      --debug                         Enable debugging information
  -h, --help                          help for go-starter
      --log-to-file                   Enable writing log output to file, specified by log-to-file-path
      --log-to-file-compress          Enable compressing log output files
      --log-to-file-max-age int       The maxium number of days to store log output in a file (default 28)
      --log-to-file-max-backups int   The maxium number of rotated logs files to keep (default 3)
      --log-to-file-max-size int      The maximum size of outputed log files in megabytes (default 500)
      --log-to-file-path string       The file to write log output to (default "/tmp/go-starter.log")
      --log-to-stdout                 Enable writing log output to stdout (default true)
      --time-format string            The time format to use for general log messages (default "2006/01/02 - 15:04:05")
  -v, --version                       version for go-starter
  -m, --mork                          Enable mork
```
