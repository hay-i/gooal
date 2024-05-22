// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Drag_n_Drop() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1>Build your questionnaire</h1><div class=\"grid\"><aside class=\"left\"><form class=\"sortable sortable-grid\" hx-post=\"/items\" hx-trigger=\"end\"><div class=\"draggable-input\" data-type=\"text\"><input type=\"hidden\" name=\"item\" value=\"1\">Text Input</div><div class=\"draggable-input\" data-type=\"number\"><input type=\"hidden\" name=\"item\" value=\"2\">Number Input</div><div class=\"draggable-input\" data-type=\"range\"><input type=\"hidden\" name=\"item\" value=\"3\">Range</div><div class=\"draggable-input\" data-type=\"select\"><input type=\"hidden\" name=\"item\" value=\"4\">Select Input</div><div class=\"draggable-input\" data-type=\"radio\"><input type=\"hidden\" name=\"item\" value=\"5\">Radio Button</div></form></aside><main class=\"right sortable\"></main></div><style>\n    .grid{\n        display: grid;\n        grid-template-columns: repeat(2, 1fr);\n        gap: 2rem;\n    }\n\n    main main{\n        border: 2px solid red;\n    }\n\n    .draggable-input{\n        padding: 1rem;\n        border: 1px solid #ececec;\n        margin-bottom: 1rem;\n        cursor: move;\n       display: flex; \n       background-color: #ececec;\n    }\n\n\n    .sortable-grid{\n        display: grid;\n        grid-template-columns: repeat(2, 1fr);\n        gap: 2rem;\n    }\n\n    aside{\n        border: 1px solid blue;\n    }\n\n\n    .form-group .input-wrapper {\n        display: flex;\n        align-items: center;\n    }\n\n    .edit-options {\n        margin-left: 10px;\n        cursor: pointer;\n        color: blue;\n    }\n    .options-container {\n        display: none;\n        margin-top: 10px;\n    }\n\n    .options-container.visible {\n        display: block;\n    }\n\n    .options-container input[type=\"text\"] {\n        display: block;\n        margin-bottom: 5px;\n    }\n    </style><script>\n        document.addEventListener(\"DOMContentLoaded\", () => {\n            const createInputElement = (type) => {\n                let inputEl;\n\n                switch(type) {\n                    case 'text':\n                        inputEl = document.createElement('input');\n                        inputEl.type = 'text';\n                        inputEl.placeholder = 'Text Input';\n                        break;\n                    case 'number':\n                        inputEl = document.createElement('input');\n                        inputEl.type = 'number';\n                        inputEl.placeholder = 'Number Input';\n                        break;\n                    case 'range':\n                        inputEl = document.createElement('input');\n                        inputEl.type = 'range';\n                        break;\n                    case 'select':\n                        inputEl = document.createElement('select');\n                        ['Option 1', 'Option 2'].forEach(text => {\n                            const option = document.createElement('option');\n                            option.value = text.toLowerCase().replace(' ', '');\n                            option.text = text;\n                            inputEl.add(option);\n                        });\n                        break;\n                    case 'radio':\n                        inputEl = document.createElement('input');\n                        inputEl.type = 'radio';\n                        inputEl.name = 'radio';\n                        break;\n                    default:\n                        inputEl = null;\n                }\n\n                return inputEl;\n            };\n\n            const createFormGroup = (inputType) => {\n                const formGroup = document.createElement('div');\n                formGroup.className = 'form-group';\n\n                const label = document.createElement('label');\n                label.contentEditable = true;\n                label.textContent = 'Label';\n\n                const inputWrapper = document.createElement('div');\n                inputWrapper.className = 'input-wrapper';\n\n                const inputEl = createInputElement(inputType);\n                inputWrapper.appendChild(inputEl);\n\n                if (inputType === 'select') {\n                    const editOptions = document.createElement('span');\n                    editOptions.className = 'edit-options';\n                    editOptions.textContent = 'Edit Options';\n\n                    const optionsContainer = document.createElement('div');\n                    optionsContainer.className = 'options-container';\n\n                    const addOptionInput = document.createElement('input');\n                    addOptionInput.type = 'text';\n                    addOptionInput.placeholder = 'New Option';\n\n                    const addOptionButton = document.createElement('button');\n                    addOptionButton.type = 'button';\n                    addOptionButton.textContent = 'Add Option';\n\n                    optionsContainer.appendChild(addOptionInput);\n                    optionsContainer.appendChild(addOptionButton);\n\n                    editOptions.addEventListener('click', () => {\n                        optionsContainer.classList.toggle('visible');\n                    });\n\n                    addOptionButton.addEventListener('click', () => {\n                        const newOption = document.createElement('option');\n                        newOption.value = addOptionInput.value.toLowerCase().replace(' ', '');\n                        newOption.text = addOptionInput.value;\n                        inputEl.add(newOption);\n                        addOptionInput.value = '';\n                    });\n\n                    inputWrapper.appendChild(editOptions);\n                    inputWrapper.appendChild(optionsContainer);\n                }\n\n                formGroup.appendChild(label);\n                formGroup.appendChild(inputWrapper);\n\n                return formGroup;\n            };\n\n            const leftSortable = new Sortable(document.querySelector(\".left .sortable\"), {\n                group: {\n                    name: \"shared\",\n                    pull: \"clone\",\n                    put: false\n                },\n                animation: 150,\n                ghostClass: 'blue-background-class'\n            });\n\n            const rightSortable = new Sortable(document.querySelector(\".right\"), {\n                group: {\n                    name: \"shared\",\n                    pull: false,\n                    put: true\n                },\n                animation: 150,\n                ghostClass: 'blue-background-class',\n                onAdd: (evt) => {\n                    const itemEl = evt.item;\n                    const inputType = itemEl.getAttribute('data-type');\n                    const formGroup = createFormGroup(inputType);\n                    itemEl.innerHTML = '';\n                    itemEl.appendChild(formGroup);\n                }\n            });\n        });\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}