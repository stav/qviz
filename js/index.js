document.addEventListener('DOMContentLoaded', event => {
  console.log('htmx version:', htmx.version)
  // htmx.logAll()
  document.body.addEventListener('htmx:beforeSwap', function (evt) {
    if (evt.detail.xhr.status === 422) {
      evt.detail.shouldSwap = true
      evt.detail.isError = false
    }
  })
})

const SUPABASE_URL = 'https://---.supabase.co'
const SUPABASE_KEY = 'eyJhbGciOiJ---Ql2hB8U9mhbJk'
const client = supabase.createClient(SUPABASE_URL, SUPABASE_KEY)
document.client = client
