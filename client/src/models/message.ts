import type { Action } from "./action";

export interface Message {
  id: number;
  userName: string;
  body: any;
  timestampSeconds: number;
  action?: Action;
}
