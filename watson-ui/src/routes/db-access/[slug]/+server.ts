import { error,json, type HttpError} from '@sveltejs/kit'
import type { RequestHandler } from './$types'
import { PUBLIC_WATSON_BACKEND_URL } from '$env/static/public'
import { WATSON_AUTH_TOKEN } from "$env/static/private";


const createUser = async (username, db, duration) => {
	const createUser = await fetch(`http://${PUBLIC_WATSON_BACKEND_URL}/api/db-access`, {
			method: 'POST',
			body: JSON.stringify({
        username: username,
        db: db,
        duration: duration
      }),
			headers: {
				'Content-Type': 'application/json',
				Token: WATSON_AUTH_TOKEN
			}
		});
    if (!createUser.ok) {
			const msg = await createUser.json();
			throw error(400, {
				message: msg.message
			});
		}

    const activeTempUser = await fetch(
			`http://${PUBLIC_WATSON_BACKEND_URL}/api/user/access-status`,
			{
				method: 'POST',
				body: JSON.stringify({
					email: username,
					dbname: db,
					duration: duration,
					action: true,
				}),
				headers: {
					'Content-Type': 'application/json',
					Token: WATSON_AUTH_TOKEN
				}
			}
		);
		if (!activeTempUser.ok) {
			const msg = await activeTempUser.json();
			throw error(400, {
				message: msg.message
			});
		}
    return {
      activeTempUser: await activeTempUser.json(),
      createdUser: await createUser.json()
    } 
}

const deleteUser = async (username) => {
	const deleteUser = await fetch(`http://${PUBLIC_WATSON_BACKEND_URL}/api/db-access`, {
		method: 'DELETE',
		body: JSON.stringify({ username: username, duration: '0s' }),
		headers: {
			'Content-Type': 'application/json',
			Token: WATSON_AUTH_TOKEN
		}
	});
	if (!deleteUser.ok) {
		const msg = await deleteUser.json();
		throw error(400, msg.message);
	}

	const activeTempUser = await fetch(
		`http://${PUBLIC_WATSON_BACKEND_URL}/api/user/access-status`,
		{
			method: 'POST',
			body: JSON.stringify({ email: username, action: false }),
			headers: {
				'Content-Type': 'application/json',
				Token: WATSON_AUTH_TOKEN
			}
		}
	);
	if (!activeTempUser.ok) {
		const msg = await activeTempUser.json();
		throw error(400, {
			message: msg.message
		});
	}
}

// create user
export const POST: RequestHandler = async event => {
	const { slug } = event.params
  try {
		switch (slug) {
			case 'create':
				const { username, db, duration } = await event.request.json()
				const {activeTempUser, createdUser } = await createUser(username, db, duration)
				return json({
					activeTempUser: activeTempUser,
					createdUser: createdUser
				})
				default:
					throw error(404, 'Invalid endpoint.')
			
		}
  } catch (err) {
    console.log(err)
    let message = (err as HttpError).body.message
    throw error(500, message)
  }
}

// delete user
export const DELETE: RequestHandler = async event => {
	const { slug } = event.params
  try {
		switch (slug) {
				case 'delete': 
					const { username } = await event.request.json()
					await deleteUser(username)
					return json({
						message: "ok",
					})
				default:
					throw error(404, 'Invalid endpoint.')
		}
  } catch (err) {
    console.log(err)
    let message = (err as HttpError).body.message
    throw error(500, message)
  }
}