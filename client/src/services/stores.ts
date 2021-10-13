import { writable } from "svelte/store";

export const user = writable<string>("John Doe");
