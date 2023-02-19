import { error, json , type HttpError} from '@sveltejs/kit'
import type { RequestHandler } from './$types'
import { OAuth2Client } from 'google-auth-library'
import { PUBLIC_GOOGLE_CLIENT_ID,PUBLIC_WATSON_BACKEND_URL } from '$env/static/public'
import { WATSON_AUTH_TOKEN } from "$env/static/private";

// Verify JWT per https://developers.google.com/identity/gsi/web/guides/verify-google-id-token
async function getGoogleUserFromJWT(token: string): Promise<Partial<User>> {
  try {
    const client = new OAuth2Client(PUBLIC_GOOGLE_CLIENT_ID)
    const ticket = await client.verifyIdToken({
      idToken: token,
      audience: PUBLIC_GOOGLE_CLIENT_ID
    });
    const payload = ticket.getPayload()
    if (!payload) throw error(500, 'Google authentication did not get the expected payload')
    let user = {
      firstname: payload['given_name'] || 'UnknownFirstName',
      lastname: payload['family_name'] || 'UnknownLastName',
      email: payload['email'],
      picture: payload['picture'],
    }
    return user
  } catch (err) {
    debugger;
    let message = ''
    if (err instanceof Error) message = err.message
    throw error(500,`Google user could not be authenticated: ${message}`)
  }
}

// check if user exists
const checkIfUserExists = async (user: Partial<User>) => {
  try {
    const u = await getUser(user)
    return true
  } catch (err) {
    return false
  }
}

// get user
const getUser = async (user: Partial<User>) => {
  const res = await fetch(`http://${PUBLIC_WATSON_BACKEND_URL}/api/user/${user?.email}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      "Token": WATSON_AUTH_TOKEN
  }})
  if (!res.ok){
    const msg = await res.json()
    throw error(500, {
      message: msg.message
    })
  }

  const resUser = await res.json()
  return resUser.user
}

// register user
async function registerUser(user: Partial<User>): Promise<User> {
  const c = await checkIfUserExists(user)
  if (!c){
    const create = await fetch(`http://${PUBLIC_WATSON_BACKEND_URL}/api/user`, {
      method: 'POST',
      body: JSON.stringify({user: user}),
      headers: {
        'Content-Type': 'application/json',
        "Token": WATSON_AUTH_TOKEN
    }})
    if (!create.ok){
      const msg = await create.json()
      throw error(500, {
        message: msg.message
      })
    }
  }

  const returnedUser = await getUser(user)
  return returnedUser
}


// Returns local user if Google user authenticated (and authorized our app)
export const POST: RequestHandler = async event => {
  try {
    const { token } = await event.request.json()
    const user = await getGoogleUserFromJWT(token)
    const userSession = await registerUser(user)

    // Prevent hooks.server.ts's handler() from deleting cookie thinking no one has authenticated
    await event.locals.session.set({user: userSession, loggedIn: true})
    
    return json({
      message: 'Successful Google Sign-In.',
      user: userSession
    })
  } catch (err) {
    console.log(err)
    let message = (err as HttpError).body.message
    throw error(500, message)
  }
}