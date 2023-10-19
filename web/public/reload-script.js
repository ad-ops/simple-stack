// This script will connect to server and if server is stopped
// it will try to reconnect and then reload the page
(function () {
    let serverStopped = false;
    let reload = undefined;
    window.addEventListener('load', () => {
        reload = new EventSource("/reload");
        reload.onerror = (event) => {
            console.log("Reload event source faile");
            serverStopped = true;
        }
        reload.onopen = (event) => {
            if (serverStopped) {
                console.log("Server is back online. Reloading page");
                window.location.reload();
            }
            console.log("Connected to server with fresh connection");
        }
    })

    window.addEventListener("unload", () => {
        if (reload) reload.close();
    });
})()