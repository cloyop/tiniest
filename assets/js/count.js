(() => {
  const counterB = document.getElementById("counter");
  let counter = 119;
  const int = setInterval(() => {
    if (counter > 60) {
      const cr = counter - 60;
      counterB.textContent =
        cr == 60 ? "2:00" : "01:" + `${cr >= 10 ? cr : "0" + cr}`;
    }
    if (counter <= 60 && counter > 0) {
      const cr = counter;
      counterB.textContent = "00:" + `${cr >= 10 ? cr : "0" + cr}`;
    }
    if (counter <= 0) {
      counterB.textContent = "re-send code";
      counterB.removeAttribute("disabled");
      clearInterval(int);
      return;
    }
    counter--;
  }, 1000);
})();
