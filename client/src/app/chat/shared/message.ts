export class Message {
  userName?: string;
  body?: string;
  timestamp?: string;

  constructor(options?: {
    userName?: string,
    body?: string,
    timestamp?: string
  }) {
    this.userName = options.userName || '';
    this.body = options.body || '';
    this.timestamp = options.timestamp;
  }

  get date(): Date {
    if (!this.timestamp) return null;
    return new Date(parseInt(this.timestamp));
  }
}
