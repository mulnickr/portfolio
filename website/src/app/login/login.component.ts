import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginResponse, UserService } from './users.service';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [FormsModule, ReactiveFormsModule, CommonModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss'
})
export class LoginComponent {
  loginForm = this.formBuilder.group({
    email: '',
    password: '',
  });

  constructor(private formBuilder: FormBuilder, private router: Router, private service: UserService) { }

  public ngOnInit(): void {
    if (sessionStorage.getItem("logged_in")) {
      this.router.navigate(["/"]);
    }
  }

  public onSubmit(): void {
    console.log(`Logging in... ${this.loginForm.getRawValue().email}`);

    let email: string | null = this.loginForm.getRawValue().email;
    let password: string | null = this.loginForm.getRawValue().password;

    if (!email || !password) {
      alert("Check username or password and try again.");
      return;
    }

    this.service.login(email, password).then(
      (data) => {
        this.handleLoginSuccess(data);
      },
      (error) => {
        console.log(`Error: ${error}`);
      }
    );
  }

  private handleLoginSuccess(data: LoginResponse): void {
    console.log(`response: ${data.message}`);
    sessionStorage.setItem("logged_in", "true");
    this.loginForm.reset();
    window.location.reload();
  }
}
