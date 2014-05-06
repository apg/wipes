package main

const clientJs = `WipesClient = function(config) {
    var CLOSE_NORMAL = 1000
    , CLOSE_GOING_AWAY = 1001
    , MAX_BACKOFF = 8000;

    var wsUrl = "ws://{{URL}}/_ws"
    , log = config.log || function (msg) {}
    , stopping = false
    , pause = false
    , error = config.error || function (msg) { alert(msg); }
    , handleLine = (typeof config.handleLine === "function") ? config.handleLine : null
    , handleJson = (typeof config.handleJson !== "function") ? null : function(data) {
        config.handleJson(JSON.parse(data));
    }
    , getWS = function () {
        return window.WebSocket || window.MozWebSocket;
    }
    , backoff = 0
    , init = function () {
        var WS = getWS();
        if (WS) {
            var s = new WS(wsUrl);
            s.onmessage = function (e) {
                if (stopping) {
                    s.close();
                    return;
                }
                else if (pause) {
                    return;
                }
                callback = handleJson || handleLine;
                if (callback) {
                    callback(e.data);
                }
            };
            s.onerror = function (e) {
                if (e.code == CLOSE_NORMAL || e.code == CLOSE_GOING_AWAY) {
                    if (backoff === 0) {
                        log("Backing off for " + backoff + "ms");
                        setTimeout(init, backoff);
                        backoff = 500;
                    }
                    else if (backoff >= MAX_BACKOFF) {
                        log("Backing off for " + backoff + "ms");
                        setTimeout(init, backoff);
                        backoff = MAX_BACKOFF;
                    }
                    else {
                        log("Backing off for " + backoff + "ms");
                        setTimeout(init, backoff);
                        backoff *= 2;
                    }
                }
                else {
                    error("WebSocket error: " + e.reason);
                    e.close();
                }
            };
        }
    }

    init();

    return {
        'stop': function () {
            stopping = true;
        }
        , 'pause': function () {
            pause = true;
        }
        , 'unpause': function () {
            pause = false;
        }
    };
};
`