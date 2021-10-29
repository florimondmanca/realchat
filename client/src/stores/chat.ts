import { get, writable } from "svelte/store";
import type { Message } from "../models/message";

export const messages = writable<Message[]>([]);

export const users = writable<string[]>([]);

export const addUser = (userName: string) => {
  users.set([...get(users), userName]);
};

export const removeUser = (userName: string) => {
  const _users = get(users);
  const index = _users.indexOf(userName);
  if (index >= 0) {
    _users.splice(index, 1);
    users.set(_users);
  }
};

export const addMessage = (message: Message) => {
  messages.set([...get(messages), message]);
};

export const clearMessages = () => {
  messages.set([]);
  users.set([]);
};
