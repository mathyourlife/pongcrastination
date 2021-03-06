package main

const homePage = `<!DOCTYPE html>
<html lang="en">
<head>
<title>Chat Example</title>
<script type="text/javascript">
window.onload = function () {
    var conn;
    var msg = document.getElementById("msg");
    var log = document.getElementById("log");

    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    document.getElementById("sendMsg").onclick = function () {
      if (!conn) {
          return false;
      }
      if (!msg.value) {
          return false;
      }
      conn.send(msg.value);
      msg.value = "";
      return false;
    };

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>CONNECTION TERMINATED</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                var item = document.createElement("div");
                item.innerText = messages[i];
                appendLog(item);
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
};
</script>
<style type="text/css">
html {
    overflow: hidden;
}

body {
    overflow: hidden;
    padding: 0;
    margin: 0;
    width: 100%;
    height: 100%;
    background: navy;
}

#container {
    margin: 10px;
}

#log {
    background: Red;
    margin: 10px;
    padding: 0.5em 0.5em 0.5em 0.5em;
    top: 0.5em;
    left: 0.5em;
    right: 0.5em;
    bottom: 3em;
    overflow: auto;
    height: 40em;
}

#form {
    padding: 0 0.5em 0 0.5em;
    margin: 0;
    position: absolute;
    bottom: 1em;
    left: 0px;
    width: 100%;
    overflow: hidden;
}

</style>
</head>
<body>
<div id="container">
  <div id="log"></div>
  <div>
    <button id="sendMsg" onclick="sendMsg()">Post</button>
    <input type="text" id="msg" size="64" autofocus />
  </div>
</div>
</body>
</html>
`
