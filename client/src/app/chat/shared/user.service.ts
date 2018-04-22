import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Observer } from 'rxjs/Observer';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';

@Injectable()
export class UserService {

  private key = "chatroom-user";
  private user$: BehaviorSubject<string> = new BehaviorSubject(
    sessionStorage.getItem(this.key)
  );

  constructor() {
    this.user$.subscribe(user => sessionStorage.setItem(this.key, user));
  }

  set(user: string) {
    this.user$.next(user);
  }

  get(): Observable<string> {
    return this.user$.asObservable();
  }
}
