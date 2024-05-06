import { Routes } from '@angular/router';
import { ResumeComponent } from './resume/resume.component';
import { HomeComponent } from './home/home.component';
import { BlogComponent } from './blog/blog.component';
import { ContactComponent } from './contact/contact.component';
import { PostComponent } from './blog/post/post.component';
import { LoginComponent } from './login/login.component';
import { LogoutComponent } from './logout/logout.component';

export const routes: Routes = [
    { path: '', component: HomeComponent },

    { path: 'resume', component: ResumeComponent },

    { path: 'contact', component: ContactComponent },

    { path: 'blog', component: BlogComponent },

    { path: 'post/:id', component: PostComponent },

    { path: 'login', component: LoginComponent },
    { path: 'logout', component: LogoutComponent },
];
