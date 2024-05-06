import { Injectable } from '@angular/core';


export class Post {

	constructor(public title: string, public body: string, public imageurl: string, public id?: string, public date?: string, public intro?: string) { }

}

@Injectable({
	providedIn: 'root',
})
export class PostService {
	private apiUrl: string = "https://api.rmulnick.dev/posts";
	// test: "http://localhost:5000/posts"

	public post?: Post;

	private headers: HeadersInit = { 'Content-Type': 'application/json' };

	public async getAllPosts(): Promise<Post[]> {
		let response = await fetch(`${this.apiUrl}/`, {
			method: 'get',
			credentials: 'include',
			headers: this.headers,
		})

		return await response.json();
	}

	public async getPostByID(postId: string): Promise<Post> {
		let response = await fetch(`${this.apiUrl}/p?id=${postId}`, {
			method: 'get',
			credentials: 'include',
			headers: this.headers,
		});

		return await response.json();
	}

	public async newPost(post: Post): Promise<string> {
		let response = await fetch(`${this.apiUrl}/edit/new`, {
			method: 'post',
			credentials: 'include',
			headers: this.headers,
			body: JSON.stringify(post),
		})

		return await response.json();
	}

	public async updatePost(post: Post): Promise<string> {
		let request: RequestInit = {
			method: 'put',
			credentials: 'include',
			headers: this.headers,
			body: JSON.stringify(post),
		}

		let response = await fetch(`${this.apiUrl}/edit/update`, request);

		return await response.json();
	}

	public async deletePost(postId: string): Promise<string> {
		let response = await fetch(`${this.apiUrl}/edit/delete?id=${postId}`, {
			method: 'get',
			credentials: 'include',
			headers: this.headers,
		})

		return await response.json();
	}

}