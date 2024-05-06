import { Component } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Post, PostService } from '../posts.service';
import { CommonModule } from '@angular/common';
import { MarkdownComponent, MarkdownModule } from 'ngx-markdown';

@Component({
  selector: 'app-post',
  standalone: true,
  imports: [CommonModule, MarkdownComponent],
  templateUrl: './post.component.html',
  styleUrl: './post.component.scss'
})
export class PostComponent {

  postId?: string;
  post?: Post;

  editing: boolean = false;
  initialBody?: string;

  isLoggedIn: boolean = false;

  constructor(public router: Router, public route: ActivatedRoute, private postService: PostService) { }

  public ngOnInit(): void {
    // validate better?
    this.isLoggedIn = sessionStorage.getItem("logged_in") != null;
    this.postId = this.route.snapshot.params['id'];

    this.postService.getPostByID(this.postId!).then(
      (data) => {
        this.post = data;
        this.initialBody = data.body;
      },
      (data) => {
        console.log(`Error: ${data}`);
      }
    );

    console.log(`Current Post ID: ${this.postId}`);
  }

  public onEdit(): void {
    this.editing = true;
  }

  public onDelete(): void {
    this.router.navigate(["blog"])
    this.postService.deletePost(this.postId!);
  }

  public onView(): void {
    this.editing = false;
  }

  public onUpdate(): void {
    this.editing = false;
    this.postService.updatePost(this.post!);
    this.initialBody = this.post!.body;
  }

  public onTitleChange(event: any): void {
    const element = event.currentTarget as HTMLInputElement;
    this.post!.title = element.value;
  }

  public onBodyChange(event: any): void {
    const element = event.currentTarget as HTMLInputElement;
    this.post!.body = element.value;
  }

  public onImageChange(event: any): void {
    const element = event.currentTarget as HTMLInputElement;
    this.post!.imageurl = element.value;
  }

}
