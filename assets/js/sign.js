document.body.addEventListener("htmx:beforeSwap", function (evt) {
  if (evt.detail.xhr.status > 199 && evt.detail.xhr.status < 300) return;
  const div = document.createElement("div");
  div.style.position = "Absolute";
  div.style.top = "0";
  div.style.padding = ".5rem";

  div.style.width = "25rem";
  div.innerHTML = evt.detail.xhr.responseText;
  document.getElementById("main").insertAdjacentElement("afterbegin", div);
  setTimeout(() => div.remove(), 2500);
});
