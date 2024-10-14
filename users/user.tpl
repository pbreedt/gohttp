{{- define "user" }}
        <div style="border: 1px dashed darkgrey; padding: 5px">
            <span><strong>ID:</strong>&nbsp;{{ .ID }}</span>
            <span><strong>Name:</strong>&nbsp;{{ .Name }}</span>
            <span><strong>Age:</strong>&nbsp;{{ .Age }}</span>
        </div>
{{- end }}