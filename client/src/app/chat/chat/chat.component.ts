import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { UserService, SocketService, User, Message, Event } from '../shared';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements OnInit, OnDestroy {

  user: string;
  conn: Subscription;
  messages: Message[] = [];
  message: string;
  

  constructor(
    private socketService: SocketService,
    private userService: UserService,
  ) { }

  ngOnInit() {
    this.initConnection();
    this.userService.get().subscribe(
      (user) => this.user = user === "null" ? null : user
    );
  }

  initConnection() {
    this.socketService.initSocket();
    this.conn = this.socketService.onMessage().subscribe(
      (message: Message) => this.messages.push(message)
    );
    this.socketService.onOpen().subscribe(
      () => console.log('connected')
    );

    this.socketService.onClose().subscribe(
      () => console.log('disconnected')
    );
  }

  sendMessage() {
    if (!this.message) return;
    this.socketService.send(new Message({
      userName: this.user,
      body: this.message,
      timestamp: (+ new Date()).toString(),
    }));
    this.message = null;
  }

  logout() {
    this.userService.set(null);
  }

  ngOnDestroy() {
    if (this.conn) {
      this.conn.unsubscribe();
    }
    this.socketService.closeSocket();
  }

}
