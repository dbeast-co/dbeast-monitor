{{/*
    Checks whether the user is attempting to store secrets in plaintext
    in the grafana.ini configmap
*/}}
{{/* grafana.assertNoLeakedSecrets checks for sensitive keys in values */}}
{{- define "grafana.assertNoLeakedSecrets" -}}
      {{- $sensitiveKeysYaml := `
sensitiveKeys:
- path: ["database", "password"]
- path: ["smtp", "password"]
- path: ["security", "secret_key"]
- path: ["security", "admin_password"]
- path: ["auth.basic", "password"]
- path: ["auth.ldap", "bind_password"]
- path: ["auth.google", "client_secret"]
- path: ["auth.github", "client_secret"]
- path: ["auth.gitlab", "client_secret"]
- path: ["auth.generic_oauth", "client_secret"]
- path: ["auth.okta", "client_secret"]
- path: ["auth.azuread", "client_secret"]
- path: ["auth.grafana_com", "client_secret"]
- path: ["auth.grafananet", "client_secret"]
- path: ["azure", "user_identity_client_secret"]
- path: ["unified_alerting", "ha_redis_password"]
- path: ["metrics", "basic_auth_password"]
- path: ["external_image_storage.s3", "secret_key"]
- path: ["external_image_storage.webdav", "password"]
- path: ["external_image_storage.azure_blob", "account_key"]
` | fromYaml -}}
  {{- if $.Values.assertNoLeakedSecrets -}}
      {{- $grafanaIni := index .Values "grafana.ini" -}}
      {{- range $_, $secret := $sensitiveKeysYaml.sensitiveKeys -}}
        {{- $currentMap := $grafanaIni -}}
        {{- $shouldContinue := true -}}
        {{- range $index, $elem := $secret.path -}}
          {{- if and $shouldContinue (hasKey $currentMap $elem) -}}
            {{- if eq (len $secret.path) (add1 $index) -}}
              {{- if not (regexMatch "\\$(?:__(?:env|file|vault))?{[^}]+}" (index $currentMap $elem)) -}}
                {{- fail (printf "Sensitive key '%s' should not be defined explicitly in values. Use variable expansion instead. You can disable this client-side validation by changing the value of assertNoLeakedSecrets." (join "." $secret.path)) -}}
              {{- end -}}
            {{- else -}}
              {{- $currentMap = index $currentMap $elem -}}
            {{- end -}}
          {{- else -}}
              {{- $shouldContinue = false -}}
          {{- end -}}
        {{- end -}}
      {{- end -}}
  {{- end -}}
{{- end -}}
