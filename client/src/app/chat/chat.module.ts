import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ChatComponent } from './chat/chat.component';
import { SocketService, UserService } from './shared';
import { UserModalComponent } from './user-modal/user-modal.component';
import { MomentModule } from 'ngx-moment';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    MomentModule,
  ],
  declarations: [
    ChatComponent,
    UserModalComponent
  ],
  providers: [
    SocketService,
    UserService,
  ]
})
export class ChatModule { }
