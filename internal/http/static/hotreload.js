(async () => {
    while (true) {
        await new Promise (r => setTimeout(r, 100))
        try {
            response = await fetch("/version")
        } catch {
            continue
        }

        serverVer = await response.text();

        // (Id of version).data-version 
        currentVer = version.dataset.version
        if (serverVer !== currentVer) {
            console.log("not equal" + currentVer + ' ' + serverVer)
            location.reload()
        }
    }
})()
