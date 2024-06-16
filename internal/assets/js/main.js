["DOMContentLoaded", "htmx:afterSwap"].forEach((eventType) => {
  document.addEventListener(eventType, function () {
    addRowRemovalListeners();
    addLeftSortableListener();
    addRightSortableListener();
  });
});

function addRowRemovalListeners() {
  document
    .querySelectorAll("[data-delete-row='true']")
    .forEach((deleteButton) => {
      deleteButton.addEventListener("click", async function () {
        const itemEl = deleteButton.parentElement;
        deleteRow(itemEl);
      });
    });
}

function addLeftSortableListener() {
  new Sortable(document.querySelector("[data-sortable-left]"), {
    group: {
      name: "shared",
      pull: "clone",
      put: false,
    },
    animation: 150,
    ghostClass: "blue-background-class",
  });
}

function addRightSortableListener() {
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
      updateItemsOrders(itemEl);
    },
    onAdd: async (evt) => {
      const itemEl = evt.item;
      const inputType = itemEl.getAttribute("data-type");
      // Hacky way to get the response to render without `Base` page
      const response = await fetch(
        `/templates/get-input?inputType=${inputType}&order=${evt.newDraggableIndex}`,
        { headers: { "HX-Request": "true" } },
      );
      const formGroupHtml = await response.text();
      itemEl.innerHTML = formGroupHtml;

      updateItemsOrders(itemEl);

      // Hacky once again
      // This basically turns a data-delete-row into a delete button.
      // TODO: This should really be done by HTMX but the reason this is hacky
      // is because for some reason the event for sortable is overriding
      // hx-delete
      itemEl
        .querySelector("[data-delete-row='true']")
        .addEventListener("click", async function () {
          const deletionResponse = await fetch("/templates/delete-input", {
            method: "DELETE",
            headers: { "HX-Request": "true" },
          });
          const deletedHtml = await deletionResponse.text();
          itemEl.outerHTML = deletedHtml;
        });
    },
  });
}

async function deleteRow(itemEl) {
  const deletionResponse = await fetch("/templates/delete-input", {
    method: "DELETE",
    headers: { "HX-Request": "true" },
  });
  const deletedHtml = await deletionResponse.text();
  itemEl.outerHTML = deletedHtml;
}

// TODO: Maybe this should be done directly with HTMX
// but it seems like overkill to return the whole form again every time
function updateItemsOrders(itemEl) {
  if (!itemEl.parentElement) return;

  const elementsInQuestion = itemEl.parentElement.children;
  for (let i = 0; i < elementsInQuestion.length; i++) {
    const element = elementsInQuestion[i];

    if (!element.querySelector("input") || !element.querySelector("label")) {
      continue;
    }

    const label = element.querySelector("label");
    const input = element.querySelector("input");

    const currentInputValues = input.id.split("-");

    const newFormValue = `${currentInputValues[0]}-${currentInputValues[1]}-${i}`;
    input.id = newFormValue;
    input.name = newFormValue;
    label.htmlFor = newFormValue;
  }
}
