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

const getDBs = async (token: string) => {
  const allDBs = await fetch(`http://${PUBLIC_WATSON_BACKEND_URL}/api/db-access`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Token': token
      }
  })
  if (!allDBs.ok){
    const msg = await allDBs.json()
    throw error(400, {
      message: msg.message
    })
  }
  const result = await allDBs.json()
  return result["db-list"]
}

export const load: PageServerLoad = ( async ({locals}) => {
  return {
    users: await getUsers(WATSON_AUTH_TOKEN),
    dbs: await getDBs(WATSON_AUTH_TOKEN),
    token: WATSON_AUTH_TOKEN,
    session: locals.session.data
  };
}) satisfies PageServerLoad;
