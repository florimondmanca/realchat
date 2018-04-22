import { Action } from './action';

export class Message {
  userName?: string;
  body?: any;
  timestamp?: string;
  action?: Action;

  constructor(options?: {
    userName?: string,
    body?: any,
    timestamp?: string,
    action?: Action,
  }) {
    this.userName = options.userName || '';
    this.body = options.body;
    this.timestamp = options.timestamp;
    this.action = options.action;
  }

  get date(): Date {
    if (!this.timestamp) return null;
    return new Date(parseInt(this.timestamp));
  }
}
