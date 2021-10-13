import { get, writable } from "svelte/store";
import type { Message } from "../models/message";

export const messages = writable<Message[]>([]);

export const addMessage = (message: Message) => {
  messages.set([...get(messages), message]);
};

export const clearMessages = () => messages.set([]);
