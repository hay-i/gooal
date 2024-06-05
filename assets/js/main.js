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

      // Hacky once again
      // This basically turns a data-delete-row into a delete button.
      // TODO: This should really be done by HTMX but the reason this is hacky
      // is because for some reason the event for sortable is overriding
      // hx-delete
      itemEl
        .querySelector("[data-delete-row='true']")
        .addEventListener("click", async function () {
          const deletionResponse = await fetch("/templates/questions/delete", {
            method: "DELETE",
            headers: { "HX-Request": "true" },
          });
          const deletedHtml = await deletionResponse.text();
          itemEl.outerHTML = deletedHtml;
        });
    },
  });
});
