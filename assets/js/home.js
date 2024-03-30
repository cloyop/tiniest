document.body.addEventListener("htmx:beforeSwap", function (evt) {
  const resText = evt.detail.xhr.responseText;
  const requestTrigger = evt.detail.requestConfig.elt;
  if (evt.detail.xhr.status > 199 && evt.detail.xhr.status < 300) {
    if (requestTrigger.id === "feeds") {
      requestTrigger.subject.value = ""
      requestTrigger.description.value = ""
      alert("Feedback sent");
    }
    return;
  }
  if (requestTrigger.id === "feeds") {
    requestTrigger.parentElement.previousSibling.innerHTML = resText;
    setTimeout(() => {
      requestTrigger.parentElement.previousSibling.innerHTML = "";
    }, 2500);
  }
  const div = document.createElement("div");
  div.style.width = "25rem";
  div.innerHTML = resText;
  document.getElementById("e-b").appendChild(div);
  setTimeout(() => div.remove(), 2500);
});

document.onclick = async (e) => {
  if (e.target.dataset.cp !== undefined) {
    navigator.clipboard.writeText(e.target.dataset.cp);
    e.target.textContent = "✔";
    setTimeout(() => {
      e.target.textContent = "📄";
    }, 800);
    return;
  }
};
