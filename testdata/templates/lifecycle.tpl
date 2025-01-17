{{- define "lifecycle_1" }}
{
"alias": "{{ template "alias1" }}",
"description": "{{ template "description" }}",
"id": "{{ template "id1" }}",
"index": 1,
"name": "{{ template "name1" }}"
}
{{ end }}

{{- define "lifecycle-pre-alpha" }}
{
"alias": "pre-alpha",
"description": "Service is only being used/tested by the people building it.",
"id": "Z2lkOi8vb3BzbGV2ZWwvTGlmZWN5Y2xlLzQyNg",
"index": 1,
"name": "Pre-alpha"
}
{{ end }}

{{- define "lifecycle-alpha" }}
{
"alias": "alpha",
"description": "Service is supporting features used by others at the company, or a very small set of friendly customers.",
"id": "Z2lkOi8vb3BzbGV2ZWwvTGlmZWN5Y2xlLzQyNw",
"index": 2,
"name": "Alpha"
}
{{ end }}

{{- define "lifecycle-beta" }}
{
"alias": "beta",
"description": "Service is supporting features used by a larger set of beta customers.",
"id": "Z2lkOi8vb3BzbGV2ZWwvTGlmZWN5Y2xlLzQyOA",
"index": 3,
"name": "Beta"
}
{{ end }}

{{- define "lifecycle-ga" }}
{
"alias": "generally_available",
"description": "Service is supporting features available to be used by all customers, and should be fully stable.",
"id": "Z2lkOi8vb3BzbGV2ZWwvTGlmZWN5Y2xlLzQyOQ",
"index": 4,
"name": "Generally Available"
}
{{ end }}

{{- define "lifecycle-eol" }}
{
"alias": "end-of-life",
"description": "Service is being retired. Might have reduced or no support.",
"id": "Z2lkOi8vb3BzbGV2ZWwvTGlmZWN5Y2xlLzQzMA",
"index": 5,
"name": "End-of-life"
}
{{ end }}