/**
 * Set the user box to the current user
 */
function setUserBox () {
  const token = localStorage.getItem('supabase.auth.token')
  if (token.length > 0) {
    const userBox = document.getElementById('user-box')
    if (userBox) {
      const user = localStorage.getItem('supabase.auth.user')
      userBox.innerHTML = user
    }
  }
}

/**
 * Remove the user token from local storage
 */
function logout () {
  localStorage.removeItem('supabase.auth.user')
  localStorage.removeItem('supabase.auth.token')

  const userBox = document.getElementById('user-box')
  if (userBox) {
    userBox.innerHTML = ''
  }
}

/**
 * Save the user token to local storage
 *
 * @param {string} user - the user to store
 * @param {string} token - the token to store
 */
function login (user, token) {
  if (token.length > 0) {
    localStorage.setItem('supabase.auth.user', user)
    localStorage.setItem('supabase.auth.token', token)
  }
  const userBox = document.getElementById('user-box')
  if (userBox) {
    userBox.innerHTML = user
  }
}
