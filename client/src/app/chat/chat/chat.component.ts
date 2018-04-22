import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { SocketService, Message, Event, Action } from '../shared';
import { UserService } from 'app/shared';

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
  Action = Action;

  constructor(
    private socketService: SocketService,
    private userService: UserService,
    private router: Router,
  ) { }

  ngOnInit() {
    this.initConnection();
    this.user = this.userService.getSnapshot();
  }

  initConnection() {
    this.socketService.initSocket();
    this.conn = this.socketService.onMessage().subscribe(
      (message: Message) => this.messages.unshift(message)
    );
    this.socketService.onOpen().subscribe(
      () => {
        console.log('connected');
        this.sendNotification({action: Action.JOIN});
      }
    );
    this.socketService.onClose().subscribe(
      () => {
        console.log('disconnected');
      }
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

  sendNotification(opt: {params?: any, action: Action}) {
    let message: Message;
    if (opt.action === Action.JOIN || opt.action === Action.LEAVE) {
      message = new Message({
        action: opt.action,
        userName: this.user,
      });
    } else if (opt.action === Action.RENAME) {
      message = new Message({
        action: opt.action,
        body: {
          userName: this.user,
          previousUserName: opt.params.previousUserName,
        },
      });
    }
    message.timestamp = (+ new Date()).toString();
    this.socketService.send(message);
  }

  logout() {
    this.userService.set(null);
    this.sendNotification({action: Action.LEAVE});
    this.router.navigate(['/']);
  }

  ngOnDestroy() {
    if (this.conn) {
      this.conn.unsubscribe();
    }
    this.socketService.closeSocket();
  }

}
