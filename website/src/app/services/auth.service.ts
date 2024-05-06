import { Injectable } from "@angular/core";


@Injectable()
export class AuthService {

	public canActivate(): boolean {
		return false;
	}

	// login auth properties go here

}