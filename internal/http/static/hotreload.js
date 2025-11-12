let source;

function init() {
  source = new EventSource("/hotreload");

  source.onmessage = (event) => {
    location.reload()
  };

  source.onerror = (error) => {
    console.log(error);
  };
}

window.addEventListener("beforeunload", () => {
  if (source) source.close();
});

init();
