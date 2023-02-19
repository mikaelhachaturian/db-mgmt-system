import { WATSON_AUTH_TOKEN } from "$env/static/private";
import {PUBLIC_WATSON_BACKEND_URL} from "$env/static/public";
import type { PageServerLoad } from './$types';
import { error } from "@sveltejs/kit";

const getUsers = async (token: string) => {
  const activeTempUser = await fetch(`http://${PUBLIC_WATSON_BACKEND_URL}/api/user/all-temp-active`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Token': token
      }
  })
  if (!activeTempUser.ok){
    const msg = await activeTempUser.json()
    throw error(400, {
      message: msg.message
    })
  }
  const result = await activeTempUser.json()
  return result.users
}

export const load: PageServerLoad = ( async ({locals}) => {
  return {
    users: await getUsers(WATSON_AUTH_TOKEN),
    token: WATSON_AUTH_TOKEN,
    session: locals.session.data
  };
}) satisfies PageServerLoad;
