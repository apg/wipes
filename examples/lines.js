(function ($) {
    Stream = (function(config) {
        var wsUrl = config.wsURL
        , log = config.log || function (msg) {}
        , stopping = false
        , error = config.error || function (msg) {
            alert(msg);
        }
        , handleLine = function (line) {
            if (typeof config.callback === "function") {
                config.callback(line);
            }
            else {
                error("callback is not a function");
            }
        }
        , supportsWS = function () {
            return window.WebSocket || window.MozWebSocket;
        }
        , buffer = '';

        var WS = supportsWS();
        if (WS) {
            var s = new WS(wsUrl);
            s.onmessage = function (e) {
                if (stopping) {
                    s.close();
                    return;
                }
                var lines = (buffer + e.data).split('\n');
                for (var i = 0; i < lines.length - 1; i++) {
                  handleLine(lines[i]);
                }
                buffer = lines[lines.length - 1];
            };
        }
        else {
            error("Need websockets!");
        }
        return {
            stop: function() {
                stopping = true;
            }
        };
    });
})(jQuery);

(function ($) {
    new Stream({
        wsURL: "ws://localhost:1234/_ws"
        , callback: function(line) {
            $('#container').text(line);
        }
    });
})(jQuery);
