import { get, writable } from "svelte/store";
import type { Message } from "../models/message";

export const messages = writable<Message[]>([]);

export const users = writable<string[]>([]);

export const addMessage = (message: Message) => {
  messages.set([...get(messages), message]);

  if (message.type === "JOIN") {
    users.set([...get(users), message.data.userName]);
  }

  if (message.type === "LEAVE") {
    const _users = get(users);
    const index = _users.indexOf(message.data.userName);
    _users.splice(index, 1);
    users.set(_users);
  }
};

export const clearMessages = () => {
  messages.set([]);
  users.set([]);
};
