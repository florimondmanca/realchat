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
  redirectUrl: string;
  fromGuard: boolean = false;

  constructor() {
    this.user$.subscribe(user => sessionStorage.setItem(this.key, user));
  }

  set(user: string) {
    this.user$.next(user);
  }

  get(): Observable<string> {
    return this.user$.asObservable();
  }

  get isSet(): boolean {
    const user: string = this.getSnapshot();
    return (user !== null) && (user !== "null");
  }

  getSnapshot(): string {
    return this.user$.getValue();
  }
}
