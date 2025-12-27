# Sopse

**Sopse** (*Stephen's Obsessive Pair Storage Engine*) is a public ephemeral key-value storage API, written in Go 1.25 by Stephen Malone.

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

Name       | Description                | Default 
---------- | -------------------------- | -------
`-addr`    | The address to serve on.   | `"127.0.0.1:8000"`
`-path`    | The database path to use.  | `"./sopse.db"`
`-bodyMax` | Maximum pair body size.    | `4096`
`-nameMax` | Maximum pair name size     | `64`

## Contributing

Please submit all bug reports and feature requests to the [issue tracker][is], thank you.

[ch]: https://github.com/stvmln86/sopse/blob/main/changes.md
[is]: https://github.com/stvmln86/sopse/issues
[li]: https://github.com/stvmln86/sopse/blob/main/license.md
[re]: https://github.com/stvmln86/sopse/releases/latest
