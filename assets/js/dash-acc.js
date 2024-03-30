const div = document.createElement("div");
div.classList = "top-0 absolute flex justify-center w-full md:w-4/6 lg:w-3/6";
document.body.addEventListener("htmx:beforeSwap", function (evt) {
  const status = evt.detail.xhr.status;
  if (status > 199 && status < 300) return;
  div.innerHTML = evt.detail.xhr.responseText;
  document.body.appendChild(div);
  setTimeout(() => div.remove(), 2000);
});
document.body.addEventListener("htmx:afterSwap", function (evt) {
  const status = evt.detail.xhr.status;
  console.log(status);
  if (status > 399) return;
  const triggerer = evt.detail.requestConfig.elt;
  if (triggerer.id === "ub") alert("Updated Succesfully");
});


