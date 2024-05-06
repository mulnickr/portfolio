import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { RouterLinkActive, RouterOutlet, RouterLink } from '@angular/router';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, RouterLink, RouterLinkActive, CommonModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'website';

  public isLoggedIn: boolean = false;

  public ngOnInit(): void {
    // validate tokens as well...
    this.isLoggedIn = sessionStorage.getItem("logged_in") != null;
    console.log(`Logged In??: ${this.isLoggedIn}`);
  }

}
