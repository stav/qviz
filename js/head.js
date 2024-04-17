const userKey = 'supabase.auth.user';
const tokenKey = 'supabase.auth.token';

/**
 * Save the user token to local storage
 *
 * @param {string} user - the user to store
 * @param {string} token - the token to store
 */
function login (user, token) {
  if (token.length > 0) {
    localStorage.setItem(userKey, user)
    localStorage.setItem(tokenKey, token)
  }
  const userBox = document.getElementById('user-box')
  if (userBox) {
    userBox.innerHTML = user
  }
}

/**
 * Remove the user token from local storage
 */
function logout () {
  localStorage.removeItem(userKey)
  localStorage.removeItem(tokenKey)

  const userBox = document.getElementById('user-box')
  if (userBox) {
    userBox.innerHTML = ''
  }
}

/**
 * Set the user box to the current user
 */
function setUserBox () {
  const token = localStorage.getItem(tokenKey)
  if (token?.length) {
    const userBox = document.getElementById('user-box')
    if (userBox) {
      const user = localStorage.getItem(userKey)
      userBox.innerHTML = user
    }
  }
}
