// {{ block "client" . }}
const SUPABASE_URL = '{{ .supabaseUrl }}'
const SUPABASE_KEY = '{{ .supabaseKey }}'
const client = supabase.createClient(SUPABASE_URL, SUPABASE_KEY)
document.client = client
// {{ end }}
