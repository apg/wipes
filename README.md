# wipes - Pipe text to a WebSocket server

## Summary

`wipes` is a simple Go program that reads from STDIN and demuxes
lines to open WebSockets. For convenience, it also serves files over
HTTP (defaulting to the local directory) such that you can talk to
the WS server easily.

## Installation

    $ go get github.com/gorilla/websocket
    $ go build

## Usage

The idea of wipes is that it's part of a command pipeline. Thus:

    $ tail -f /var/log/messages | wipes -addr :8080

Will start an HTTP server on 8080 and serve files from the current
directory. It will also make available the output of
`tail -f /var/log/messages` available via WebSockets at
[ws://localhost:8080/_ws]().

It's possible to serve a different directory:

    $ tail -f /var/log/messages | wipes -static /var/www/wsapp

`wipes` will exit with status 0 when there's nothing left to read.

## Examples

Anything that provides a constant stream of data makes sense to use
here.  In the `examples` directory, there's a simple HTML doc with
a WS client that just displays a line at a time. To test it out, we'll
make use of Meetup's streaming API, specifically, the `/2/rsvps`
endpoint.

    $ curl http://stream.meetup.com/2/rsvps | wipes -static examples

Then, open your browser to `http://localhost:8080/lines.html`. The
result should be lines of JSON formatted text.

## See Also

### websocketd

Like inetd, but for WebSockets. Turn any application that uses
STDIN/STDOUT into a WebSocket server.

[websocketd](https://github.com/joewalnes/websocketd)

## Contributing and Feedback

It's possible that `wipes` has bugs, and/or does something
it shouldn't be, and/or has lots of room for improvement. If you'd
like to fix or contribute something, please fork and submit a pull
request, or open an issue.

If you have any other feedback, feel free to email me at the below
address.

## Authors

Andrew Gwozdziewycz <web@apgwoz.com>

## Copyright

Copyright 2014, Andrew Gwozdziewycz, <web@apgwoz.com>

Licensed under the GNU GPLv3. See LICENSE for more details.
