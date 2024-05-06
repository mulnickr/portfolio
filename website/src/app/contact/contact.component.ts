import { Component, ElementRef, ViewChild } from '@angular/core';

export { }; declare global { interface Window { Calendly: any; } }

@Component({
  selector: 'app-contact',
  standalone: true,
  imports: [],
  templateUrl: './contact.component.html',
  styleUrl: './contact.component.scss'
})
export class ContactComponent {
  @ViewChild('container') container?: ElementRef;

  public email: string = "mulnickr@gmail.com";

  public ngAfterViewInit(): void {
    if (this.container) {
      window.Calendly.initInlineWidget({
        url: 'https://calendly.com/mulnickr',
        parentElement: this.container.nativeElement
      });
    } else {
      console.log("'container' not found.")
    }

  }

}
