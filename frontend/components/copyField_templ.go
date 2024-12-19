// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func CopyField(id string, text string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"{id}\" class=\"copy-field-container\"><input type=\"text\" value=\"{text}\" id=\"{id}-input\" readonly class=\"copy-field-input\"> <button onclick=\"copyToClipboard(&#39;{id}-input&#39;)\" class=\"copy-field-button\">Copy</button></div><script>\n        function copyToClipboard(inputId) {\n            var copyText = document.getElementById(inputId);\n            copyText.select();\n            copyText.setSelectionRange(0, 99999); // For mobile devices\n            document.execCommand(\"copy\");\n            alert(\"Copied the text: \" + copyText.value);\n        }\n    </script><style>\n        .copy-field-container {\n            display: flex;\n            align-items: center;\n        }\n        .copy-field-input {\n            flex: 1;\n            padding: 8px;\n            margin-right: 10px;\n            border: 1px solid #ccc;\n            border-radius: 3px;\n        }\n        .copy-field-button {\n            padding: 8px 12px;\n            background-color: #007bff;\n            color: white;\n            border: none;\n            border-radius: 3px;\n            cursor: pointer;\n        }\n        .copy-field-button:hover {\n            background-color: #0056b3;\n        }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate