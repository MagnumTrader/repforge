let source;

function init() {
    source = new EventSource("/hotreload");

    // A file has been updated
    source.onmessage = (_) => {
        location.reload()
    };

    // This happens when the server errors
    source.onerror = (_) => {
        const interval = 200
        const startTime = Date.now();
        const maxWait = 8000

        const ping = setInterval(async () => {
            try {
                const res = await fetch("/health", { method: "GET" });
                if (res.ok) {
                    console.log("Server back! Reloading...");
                    clearInterval(ping);
                    location.reload();
                }
            } catch {
                if (Date.now() - startTime > maxWait) {
                    console.warn(`Server did not respond within 8 seconds`);
                    clearInterval(ping);
                    location.reload();
                }
            }
        }, interval);
    };
}

// This is for closing the connection to the go server
// Had some issues with the browser getting conflicted and not 
// being able to open a new connection
window.addEventListener("beforeunload", () => {
    if (source) source.close();
});

init();
