document.body.addEventListener("htmx:beforeSwap", function (evt) {
  const status = evt.detail.xhr.status;
  if (status > 199 && status < 300) return;
  const re = document.getElementById("re");
  re.innerHTML += evt.detail.xhr.responseText;
  setTimeout(() => (re.innerHTML = ""), 2000);
});
