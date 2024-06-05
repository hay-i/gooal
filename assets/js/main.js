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
    onEnd: async (evt) => {
      const itemEl = evt.item;
      var newOrder = 1;
      if (itemEl.previousSibling) {
        newOrder =
          parseInt(
            itemEl.previousSibling.querySelector("input").id.split("-")[2],
          ) + 1;
      }
      const label = itemEl.querySelector("label");
      const input = itemEl.querySelector("input");

      const currentInputValues = input.id.split("-");

      // TODO: For some reason these aren't applying
      const newFormValue = `${currentInputValues[0]}-${currentInputValues[1]}-${newOrder}`;
      input.id = newFormValue;
      input.name = newFormValue;
      label.htmlFor = newFormValue;
      updateSubsequentOrderItems(itemEl, newOrder);
    },
    onAdd: async (evt) => {
      const itemEl = evt.item;
      var newOrder = 1;
      if (itemEl.previousSibling) {
        newOrder =
          parseInt(
            itemEl.previousSibling.querySelector("input").id.split("-")[2],
          ) + 1;
      }
      updateSubsequentOrderItems(itemEl, newOrder);

      const inputType = itemEl.getAttribute("data-type");
      // Hacky way to get the response to render without `Base` page
      const response = await fetch(
        `/templates/builder?inputType=${inputType}&order=${newOrder}`,
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

// TODO: Maybe this should be done directly with HTMX
// but it seems like overkill to return the whole form again every time
function updateSubsequentOrderItems(itemEl, order) {
  var newOrder = order;
  if (!itemEl.parentElement) return;
  console.log(newOrder);

  const elementsInQuestion = itemEl.parentElement.children;
  for (let i = 0; i < elementsInQuestion.length; i++) {
    const element = elementsInQuestion[i];

    if (!element.querySelector("input") || !element.querySelector("label")) {
      continue;
    }

    const label = element.querySelector("label");
    const input = element.querySelector("input");

    const currentInputValues = input.id.split("-");

    if (parseInt(currentInputValues[2]) >= newOrder) {
      const newFormValue = `${currentInputValues[0]}-${currentInputValues[1]}-${parseInt(currentInputValues[2]) + 1}`;
      input.id = newFormValue;
      input.name = newFormValue;
      label.htmlFor = newFormValue;
      newOrder++;
    }
  }
}
