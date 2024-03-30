// {{ block "client" . }}
const SUPABASE_URL = '{{ .supabaseUrl }}'
const SUPABASE_KEY = '{{ .supabaseKey }}'
document.client = supabase.createClient(SUPABASE_URL, SUPABASE_KEY)
// {{ end }}
