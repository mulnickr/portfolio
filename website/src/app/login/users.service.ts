import { Injectable } from "@angular/core";

class User {
	constructor(public email: string, public username: string, public createdat: string) { }
}

export class LoginResponse {
	constructor(public message: string, public token?: string) { }
}


@Injectable({
	providedIn: 'root',
})
export class UserService {
	private apiUrl: string = "https://api.rmulnick.dev/users";
	// test: "http://localhost:5000/users"

	public user?: User;

	public async login(email: string, password: string): Promise<LoginResponse> {
		let response = await fetch(`${this.apiUrl}/login`, {
			method: 'post',
			credentials: 'include',
			body: JSON.stringify({ "email": email, "password": password }),
		});

		return await response.json();
	}

	public async logout(): Promise<string> {
		let response = await fetch(`${this.apiUrl}/logout`, {
			method: 'get',
			credentials: 'include',
		});

		return await response.json();
	}

}