document.addEventListener("click", function (e) {
  if (e.target && e.target.dataset.dismissable !== undefined) {
    document.querySelector(`#${e.target.dataset.dismissable}`).innerHTML = "";
  }
});

document.addEventListener("DOMContentLoaded", function () {
  new Sortable(document.querySelector("[data-sortable-left]"), {
    group: {
      name: "shared",
      pull: "clone",
      put: false,
    },
    animation: 150,
    ghostClass: "blue-background-class",
  });

  new Sortable(document.querySelector("[data-sortable-right]"), {
    group: {
      name: "shared",
      pull: false,
      put: true,
    },
    animation: 150,
    ghostClass: "blue-background-class",
    onAdd: async (evt) => {
      const itemEl = evt.item;
      const inputType = itemEl.getAttribute("data-type");
      // Hacky way to get the response to render without `Base` page
      const response = await fetch(
        `/templates/builder?inputType=${inputType}`,
        { headers: { "HX-Request": "true" } },
      );
      const formGroupHtml = await response.text();
      itemEl.innerHTML = formGroupHtml;
    },
  });
});
