document.addEventListener("DOMContentLoaded", (event) => {
  console.log('htmx version:', htmx.version);
  htmx.logAll()
  document.body.addEventListener('htmx:beforeSwap', function (evt) {
    if (evt.detail.xhr.status === 422) {
      evt.detail.shouldSwap = true;
      evt.detail.isError = false;
    }
  })
})
