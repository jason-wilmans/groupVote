import {Component, OnInit} from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {
  title = 'app';

  constructor(private http: HttpClient) { }

  ngOnInit() {
    this.http.get<Object>('http://localhost:8080/test').subscribe(message => {
      console.log("message: " , message)
    });

  }
}
