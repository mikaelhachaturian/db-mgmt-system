import type { Handle, RequestEvent } from '@sveltejs/kit'
import { PUBLIC_WATSON_BACKEND_URL } from '$env/static/public'
import { WATSON_AUTH_TOKEN, SUPER_SECRET_SESSION_KEY } from "$env/static/private";
import { error } from '@sveltejs/kit'
import { redirect } from '@sveltejs/kit';
import { handleSession } from 'svelte-kit-cookie-session';

// Attach authorization to each server request (role may have changed)
async function attachUserToRequestEvent(sessionId: number, event: RequestEvent) {
  const res = await fetch(`http://${PUBLIC_WATSON_BACKEND_URL}/api/user/id/${sessionId}`, {
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
  event.locals.session.set({user: resUser.user, loggedIn: true})
}

const unProtectedRoutes: string[] = [
  '/',
  '/login',
  '/auth',
  '/health'
];

// Invoked for each endpoint called and initially for SSR router
export const handle: Handle = handleSession(
	{
		secret: SUPER_SECRET_SESSION_KEY
	}, async ({ event, resolve }) => {
  const { url } = event

  const sessionId = event.locals.session.data.user?.ID

  // before endpoint or page is called
  if (!sessionId && !unProtectedRoutes.includes(url.pathname) || url.pathname == "/"){
    throw redirect(302, "/login")
  }

  if (sessionId)
    await attachUserToRequestEvent(sessionId, event)
    
  const response = await resolve(event)

  // after endpoint or page is called

  return response
});