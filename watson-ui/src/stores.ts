import { writable } from 'svelte/store'

let TempUsers: User[] = [];

export const loggedInStore = writable<boolean>();
export const userInStore = writable<User>();

export const activeTempUsers =  writable(TempUsers);