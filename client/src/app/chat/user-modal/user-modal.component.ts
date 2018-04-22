import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { UserService } from '../shared';

@Component({
  selector: 'app-user-modal',
  templateUrl: './user-modal.component.html',
  styleUrls: ['./user-modal.component.scss']
})
export class UserModalComponent implements OnInit {

  userName: string;

  constructor(private userService: UserService) { }

  ngOnInit() {
  }

  setUser() {
    this.userService.set(this.userName);
  }

}
