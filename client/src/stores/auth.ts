import { derived, writable } from "svelte/store";

const _storedUser = localStorage.getItem("user");
export const user = writable<string>(_storedUser ? _storedUser : null);

user.subscribe((value) => {
  if (value === null) {
    localStorage.removeItem("user");
  } else {
    localStorage.setItem("user", value);
  }
});

export const signIn = (value: string) => user.set(value);

export const isLoggedIn = derived(
  user,
  ($user, set: (value: boolean) => void) => set($user != null)
);

export const logout = () => user.set(null);
