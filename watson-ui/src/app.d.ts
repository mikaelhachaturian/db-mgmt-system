// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
declare namespace App {
	// interface Error {}
	interface Locals {
    session: import('svelte-kit-cookie-session').Session<SessionData>;
	}
	interface PageData {
    session: SessionData;
  }
  interface PageServerData {
    session: SessionData;
  }
	// interface Platform {}
}

interface UserProperties {
  ID: number
  CreatedAt?: string // ISO-8601 datetime
  UpdatedAt?: string // ISO-8601 datetime
  firstname?: string
  lastname?: string
  email?: string
  duration?: string
  picture?: string
  dbname?: string
}

type User = UserProperties | undefined | null

interface SessionData {
	user: User;
  loggedIn: boolean
}