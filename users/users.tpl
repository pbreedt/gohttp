<!DOCTYPE html>
<html>
<body>
    <div>
        {{ range . }}
            {{- template "user" . }}
        {{ end }}
    </div>
</body>
</html>
