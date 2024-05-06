import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../login/users.service';

@Component({
  selector: 'app-logout',
  standalone: true,
  imports: [],
  templateUrl: './logout.component.html',
  styleUrl: './logout.component.scss'
})
export class LogoutComponent {

  constructor(private router: Router, private service: UserService) { }

  public ngOnInit(): void {
    // this isn't an ideal way to handle page reload/redirect... but it works for now
    if (sessionStorage.getItem("logged_in")) {
      sessionStorage.removeItem("logged_in");
      this.service.logout();
      window.location.reload();
    }

    this.router.navigate(['']);
  }

}
