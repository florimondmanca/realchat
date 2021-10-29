export interface Message {
  id: number;
  timestampSeconds: number;
  userName: string;
  body: string;
}

export interface ISendDetail {
  body: string;
}
