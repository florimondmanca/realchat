interface MessageCommon<D> {
  id: number;
  data: D;
  timestampSeconds: number;
}

export interface ChatMessage
  extends MessageCommon<{ userName: string; body: string }> {
  type: "CHAT";
}

export interface JoinMessage extends MessageCommon<{ userName: string }> {
  type: "JOIN";
}

export interface LeaveMessage extends MessageCommon<{ userName: string }> {
  type: "LEAVE";
}

export type Message = ChatMessage | JoinMessage | LeaveMessage;

export interface ISendDetail {
  body: string;
}

export interface IMessagePayload {
  userName: string;
  body: string;
}
