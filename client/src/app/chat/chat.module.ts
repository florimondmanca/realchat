import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ChatComponent } from './chat/chat.component';
import { SocketService } from './shared';
import { MomentModule } from 'ngx-moment';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    MomentModule,
  ],
  declarations: [
    ChatComponent,
  ],
  providers: [
    SocketService,
  ]
})
export class ChatModule { }
