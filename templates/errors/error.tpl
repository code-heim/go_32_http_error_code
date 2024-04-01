{{ define "errors/error.tpl"}}
    {{ template "layouts/header.tpl" .}}
        {{ if .Message }}
            <div style="color: white; background-color: blue; padding: 20px; margin: 10px 0; border-radius: 5px;">
                {{ .Message }}
            </div>
        {{ end }}
        {{ if .Error }}
            <div style="color: white; background-color: red; padding: 20px; margin: 10px 0; border-radius: 5px;">
                {{ .Error }}
            </div>
        {{ end }}
    {{ template "layouts/footer.tpl" .}}
{{end}}
