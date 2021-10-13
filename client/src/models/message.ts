import type { Action } from "./action";

export interface Message {
  id: number;
  userName: string;
  body: any;
  timestampSeconds: number;
  action?: Action;
}

export interface ISendDetail {
  body: string;
}

export interface IMessagePayload {
  userName: string;
  body: string;
}
