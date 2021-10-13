import { derived, writable } from "svelte/store";

export const user = writable<string>("John Doe");

export const signIn = (value: string) => user.set(value);

export const isLoggedIn = derived(
  user,
  ($user, set: (value: boolean) => void) => set($user != null)
);

export const logout = () => user.set(null);
