import { error, json } from '@sveltejs/kit'
import type { RequestHandler } from './$types'

export const POST: RequestHandler = async (event) => {
	const { slug } = event.params

	try {
		switch (slug) {
			case 'logout':
				await event.locals.session.destroy()
				return json({ message: 'Logout successful.' })
			default:
				throw error(404, 'Invalid endpoint.')
		}
	} catch (err) {
		throw error(503, 'Could not communicate with database.')
	}
}