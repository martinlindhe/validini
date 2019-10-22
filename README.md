# About

Command line tool to validate INI syntax of input files.


## Installation

Windows and macOS binaries are available under [Releases](https://github.com/martinlindhe/validini/releases)

Or install from source:

    go get -u github.com/martinlindhe/validini


## Usage

Exit code will be 0 if file is good.

    $ validini file.ini
    OK: file.ini

    $ curl http://site.com/file.ini | validini
    OK: -


## License

Under [MIT](LICENSE)
