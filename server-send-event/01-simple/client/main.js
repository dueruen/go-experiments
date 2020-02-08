function onLoaded() {
    var source = new EventSource("/sse/dashboard");
    source.onmessage = function(event) {
        console.log("OnMessage called:");
        console.dir(event);
        document.getElementById("counter").innerHTML = event.data;
    }
}