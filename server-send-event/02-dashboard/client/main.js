function onLoaded() {
    var source = new EventSource("/sse/dashboard");
    source.onmessage = function (event) {
        console.log("OnMessage called:");
        console.dir(event);
        var dashboard = JSON.parse(event.data);
        var items = dashboard["inventory"]["items"]
        document.getElementById("biprice").innerHTML = items["bicycle"].price;
        document.getElementById("biquantity").innerHTML = items["bicycle"].quantity;
        document.getElementById("bprice").innerHTML = items["book"].price;
        document.getElementById("bquantity").innerHTML = items["book"].quantity;
        document.getElementById("rccprice").innerHTML = items["rccar"].price;
        document.getElementById("rccquantity").innerHTML = items["rccar"].quantity;
    }
}