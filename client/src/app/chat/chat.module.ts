import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ChatComponent } from './chat/chat.component';
import { SocketService } from './shared';
import { UserModalComponent } from './user-modal/user-modal.component';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
  ],
  declarations: [
    ChatComponent,
    UserModalComponent
  ],
  providers: [
    SocketService,
  ]
})
export class ChatModule { }
