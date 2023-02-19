import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = ({ locals, url }) => {
  return {
    url: url.pathname,
    session: locals.session.data
  }
}