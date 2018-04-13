import { Component, OnInit } from '@angular/core';
import { Rubro } from '../../models/rubro';
import { RubroService } from '../../services/rubro.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-rubro-new',
  templateUrl: './rubro-new.component.html',
  styleUrls: []
})
export class RubroNewComponent implements OnInit {

  rubro: Rubro;
  display = false;
  constructor(private rubroService: RubroService, private location: Location) { }

  ngOnInit() {
    this.rubro = new Rubro();
  }

  guardar(rubro: Rubro): void {

    this.rubroService.create(rubro);
    this.display = true;

  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}