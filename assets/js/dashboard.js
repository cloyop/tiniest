const overAllClasses =
  "w-full h-full flex flex-col items-center justify-center z-20 bg-black-opac absolute top-0 right-0 left-0 bottom-0 gap-4";
document.addEventListener("click", async (e) => {
  const trigger = e.target;
  if (trigger.dataset.overall !== undefined) {
    return trigger.remove();
  }
  if (trigger.dataset.xx !== undefined) {
    return document.querySelector("[data-overall]")?.remove();
  }
  if (trigger.dataset.form !== undefined) {
    if (trigger.dataset.form === "open") {
      document.getElementById("f-c").classList.remove("hidden");
      document.getElementById("menu").classList.add("hidden");
    } else {
      document.getElementById("f-c").classList.add("hidden");
      resetUrlForm(document.getElementById("url-form"));
    }
    return;
  }
  if (trigger.dataset.menu !== undefined) {
    document.getElementById("menu").classList.toggle("hidden");
    return;
  }
  if (trigger.dataset.qr !== undefined) {
    document.getElementById("menu").classList.add("hidden");
    const div2 = document.createElement("div");
    div2.classList = "w-5/6 md:w-3/6 flex justify-end";
    div2.innerHTML += `<button class="cursor-pointer" data-xx>❌</button>`;
    const img = document.createElement("img");
    img.src = trigger.src;
    img.alt = "Qr Image";
    img.style.width = "16rem";
    const div = document.createElement("div");
    div.classList = overAllClasses;
    div.dataset.overall = "";
    div.appendChild(div2);
    div.appendChild(img);
    document.getElementById("main").appendChild(div);
  }
  if (trigger.dataset.cp !== undefined) {
    navigator.clipboard.writeText(trigger.dataset.cp);
    const div = document.createElement("div");
    div.classList =
      "absolute top-4 left-0 right-0 h-8 flex justify-center items-center";
    div.textContent = "Copied!";
    document.body.appendChild(div);
    setTimeout(() => div.remove(), 2500);
    return;
  }
  if (trigger.id == "p-i") {
    const pi = trigger.closest("form").password;
    if (trigger.checked) {
      pi.parentElement.classList.remove("hidden");
      pi.required = true;
    } else {
      pi.parentElement.classList.add("hidden");
      pi.required = false;
    }
    return;
  }
});
document.body.addEventListener("htmx:beforeSwap", function (evt) {
  const requestTrigger = evt.detail.requestConfig.elt;
  const responseText = evt.detail.xhr.responseText;
  const status = evt.detail.xhr.status;
  const itsOk = status > 199 && status < 400;
  if (!itsOk) {
    if (status === 406) {
      setTimeout(() => alert("max created links capped"), 500);
      return;
    }

    if (requestTrigger.id === "url-form") {
      requestTrigger.previousSibling.innerHTML = responseText;
      setTimeout(() => (requestTrigger.previousSibling.innerHTML = ""), 2500);
      return;
    }
    if (
      requestTrigger.dataset.unlock !== undefined ||
      requestTrigger.dataset.lock !== undefined
    ) {
      alert("internal error, try again in a few seconds");
      return;
    }
    if (
      requestTrigger.dataset.locker !== undefined ||
      requestTrigger.dataset.unlocker !== undefined
    ) {
      let lockereb = document.getElementById("lockerEB");
      lockereb.innerHTML = responseText;
      setTimeout(() => (lockereb.innerHTML = ""), 2500);
      return;
    }
    return;
  }
  if (itsOk) {
    if (
      requestTrigger.dataset.locker !== undefined ||
      requestTrigger.dataset.unlocker !== undefined ||
      requestTrigger.dataset.lock !== undefined ||
      requestTrigger.dataset.unlock !== undefined
    ) {
      return document.querySelector("[data-overall]")?.remove();
    }
    if (requestTrigger.id === "url-form") {
      resetUrlForm(requestTrigger);
      const nl = document.getElementById("nl");
      nl?.remove();
      return;
    }
    if (requestTrigger.dataset.pd !== undefined) {
      const urls = document.getElementById("urls");
      if (urls?.childNodes.length === 1)
        setTimeout(() => {
          const div = document.createElement("div");
          div.id = "nl";
          div.textContent = "No Links Created";
          div.classList = "p-2";
          urls.insertAdjacentElement("beforebegin", div);
        }, 100);
      return;
    }
    if (requestTrigger.id === "filter-reset") {
      return document.getElementById("filter-form").reset();
    }

    return;
  }
});
document.addEventListener("keydown", (e) => {
  if (e.key === "Escape" && e.code === "Escape") {
    if (!document.getElementById("f-c").classList.contains("hidden")) {
      document.getElementById("f-c").classList.add("hidden");
      return;
    }
    document.querySelector("[data-overall]")?.remove();
    return;
  }
});
const resetUrlForm = (requestTrigger) => {
  document.getElementById("f-c").classList.add("hidden");
  requestTrigger.reset();
  requestTrigger.title.value = "";
  requestTrigger.password.parentElement.classList.add("hidden");
  requestTrigger.password.required = false;
};
