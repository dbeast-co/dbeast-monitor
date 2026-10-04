# DBeast Grafana API data source

The DBeast Grafana API data source queries the internal HTTP API of the Grafana instance where the plugin is running. It turns JSON responses into Grafana data frames, which can be used in dashboards and dashboard variables.

The data source is registered as `dbeast-grafanaapi-datasource`. It requires Grafana 10.1 or later and does not need data source connection settings.

## Querying Grafana

1. Add the **DBeast Grafana API** data source to a panel.
2. In **Path**, enter a Grafana API endpoint path, for example `/api/datasources` or `/api/dashboards`.
3. In **Fields**, add one or more JSONPath or JSONata expressions to select values from the response.
4. Optionally add query-string parameters in **Params**, and set **Cache Time** to control how long responses are cached.

Requests are GET-only and are made to the same Grafana instance using Grafana's backend request service. The user must have permission to access the selected endpoint. Absolute URLs, protocol-relative URLs, backslashes, and parent-directory path traversal are rejected.

### Example

To query `/api/datasources`, add fields such as:

| Expression | Language | Alias |
| --- | --- | --- |
| `$[*].name` | JSONPath | `name` |
| `$[*].type` | JSONPath | `type` |

These expressions select names and types from the datasource array returned by Grafana. The exact JSONPath depends on the response shape returned by the endpoint. All selected fields in a query must produce the same number of values.

## Field selection

Each field supports:

- **JSONPath** (default) or **JSONata** expressions.
- **Auto**, **String**, **Number**, **Time**, or **Boolean** field types. Auto detects a type from the selected values; explicit types convert values where supported.
- An optional **Alias**. If omitted, the field name is inferred from the JSONPath expression, or from the result position for JSONata.

Grafana dashboard variables are interpolated in the API path, query parameter keys and values, and field expressions. The plugin also supports the time macros `$__unixEpochFrom()`, `$__unixEpochTo()`, `$__isoFrom()`, and `$__isoTo()` in interpolated strings.

## Query options

- **Params** adds key/value pairs to the request query string. Keys and values are URL-encoded.
- **Cache Time** caches each response for the selected duration. Set it to `0s` to disable caching. Concurrent identical requests are shared.
- **Metric** (experimental) uses a selected field's values as the display name of the query result.

For dashboard variables, the first field supplies the variable values by default. The **Variable text** and **Variable value** options can select separate fields for labels and values.

## Health check

The data source health check requests `/api/health` to verify connectivity to Grafana.

For more information about DBeast Monitor, see the [project documentation](https://github.com/dbeast-co/dbeast-monitor/wiki).
