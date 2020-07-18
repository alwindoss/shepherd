import { Component, OnInit, ViewChild } from '@angular/core';
import { MatMenuTrigger } from '@angular/material/menu';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.css']
})
export class MenuComponent implements OnInit {
  constructor() { }
  @ViewChild(MatMenuTrigger) trigger: MatMenuTrigger;

  ngOnInit(): void {
  }

  someMethod() {
    console.log("Helloi theer")
    this.trigger.openMenu();
  }
}
