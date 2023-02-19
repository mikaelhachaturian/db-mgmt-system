import { json } from '@sveltejs/kit'
import type { RequestHandler } from './$types'

// Returns healthcheck
export const GET: RequestHandler = async event => {
  return json({
    message: 'ok.'
  })
}