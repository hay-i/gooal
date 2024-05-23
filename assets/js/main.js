document.addEventListener("click", function (e) {
  if (e.target && e.target.dataset.dismissable !== undefined) {
    document.querySelector(`#${e.target.dataset.dismissable}`).innerHTML = "";
  }
});
