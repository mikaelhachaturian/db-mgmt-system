import {PUBLIC_GOOGLE_CLIENT_ID} from "$env/static/public";
import { error , type HttpError} from "@sveltejs/kit";
import { goto } from '$app/navigation'
import { loggedInStore, userInStore } from "../stores";

export function renderGoogleButton() {
  const btn = document.getElementById('googleButton')
  if (btn) {
    //@ts-ignore
    google.accounts.id.renderButton(btn, { theme: "outline", size: "large", shape:"square" })
  }
}

export function initializeGoogleAccounts() {
  //@ts-ignore
  google.accounts.id.initialize({
    client_id: PUBLIC_GOOGLE_CLIENT_ID,
    callback: googleCallback
  })

  //@ts-ignore
  async function googleCallback(response: google.accounts.id.CredentialResponse) {
    try {
      const res = await fetch('/auth', {
          method: 'POST',
          body: JSON.stringify({token: response.credential}),
          headers: {
            'Content-Type': 'application/json'
          }
      })

      const {user} = await res.json()
      
      loggedInStore.set(true)
      userInStore.set(user) // for layout to update accordingly

      goto("/db-access", { invalidateAll: true })
    }
    catch (err){
      let message = (err as HttpError).body.message
      throw error(500, message)
    }
  }
}
