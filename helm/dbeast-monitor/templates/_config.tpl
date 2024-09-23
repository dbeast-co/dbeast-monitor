{{/*
 Generate config map data
 */}}
{{- define "grafana.configData" -}}
{{ include "grafana.assertNoLeakedSecrets" . }}
{{- with .Values.plugins }}
plugins: {{ join "," . }}
{{- end }}
grafana.ini: |
{{- range $elem, $elemVal := index .Values.grafana.config "grafana.ini" }}
  {{- if not (kindIs "map" $elemVal) }}
  {{- if kindIs "invalid" $elemVal }}
  {{ $elem }} =
  {{- else if kindIs "string" $elemVal }}
  {{ $elem }} = {{ tpl $elemVal $ }}
  {{- else }}
  {{ $elem }} = {{ $elemVal }}
  {{- end }}
  {{- end }}
{{- end }}
{{- range $key, $value := index .Values.grafana.config "grafana.ini" }}
  {{- if kindIs "map" $value }}
  [{{ $key }}]
  {{- range $elem, $elemVal := $value }}
  {{- if kindIs "invalid" $elemVal }}
  {{ $elem }} =
  {{- else if kindIs "string" $elemVal }}
  {{ $elem }} = {{ tpl $elemVal $ }}
  {{- else }}
  {{ $elem }} = {{ $elemVal }}
  {{- end }}
  {{- end }}
  {{- end }}
{{- end }}
{{- if .Values.grafana.config.ldap.enabled }}
ldap.toml: |
{{ .Values.grafana.config.ldap.config | nindent 2 }}
{{- end }}
{{- end -}}
