# Sopse

**Sopse** (*Stephen's Obsessive Pair Storage Engine*) is an ephemeral key-value storage API, written in Go 1.25 by Stephen Malone.

- See [`changes.md`][ch] for the complete changelog.
- See [`license.md`][li] for the open-source license (BSD-3).

## Installation

You can install Sopse using your Go tools...

```text
go install github.com/stvmln86/sopse@latest
```

...or download the [latest binary release][re] for your platform.

## Configuration

Sopse is configured entirely through command-line flags. 

### Essential Flags

Name    | Description                 | Default 
------- | --------------------------- | -------
`-addr` | The address to serve on.    | `"127.0.0.1:8000"`
`-logs` | The logging stream or file. | `""` (stdout)
`-path` | The database file path.     | `"./sopse.db"`

### Rate Limiting Flags

Name        | Description                           | Default 
----------- | ------------------------------------- | -------
`-rateName` | Maximum key name length.              | `64`
`-rateBody` | Maximum key value length.             | `65536`
`-rateUser` | Maximum number of keys per user.      | `256`
`-rateHits` | Maximum number of user hits per hour. | `100`

[ch]: https://github.com/stvmln86/sopse/blob/main/changes.md
[li]: https://github.com/stvmln86/sopse/blob/main/license.md
[re]: https://github.com/stvmln86/sopse/releases/latest
