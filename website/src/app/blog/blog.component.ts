import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { ActivatedRoute, Router, RouterModule, RouterOutlet } from '@angular/router';
import { Post, PostService } from './posts.service';


@Component({
  selector: 'app-blog',
  standalone: true,
  imports: [CommonModule, RouterOutlet, RouterModule],
  templateUrl: './blog.component.html',
  styleUrl: './blog.component.scss'
})
export class BlogComponent {
  posts?: Post[];

  constructor(private router: Router, private route: ActivatedRoute, private postService: PostService) { }

  public ngOnInit(): void {
    this.postService.getAllPosts().then((data) => {
      this.posts = data;
    })
  }

  public onClickPost(post: string): void {
    this.router.navigate(["post", post], { relativeTo: this.route })
  }
}
