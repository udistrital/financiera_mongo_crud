import { Component, OnInit } from '@angular/core';
import { Apropiaciones } from '../../models/apropiaciones';
import { ApropiacionesService } from '../../services/apropiaciones.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-apropiaciones-new',
  templateUrl: './apropiaciones-new.component.html',
  styleUrls: []
})
export class ApropiacionesNewComponent implements OnInit {

  apropiaciones: Apropiaciones;
  display = false;
  constructor(private apropiacionesService: ApropiacionesService, private location: Location) { }

  ngOnInit() {
    this.apropiaciones = new Apropiaciones();
  }

  guardar(apropiaciones: Apropiaciones): void {

    this.apropiacionesService.create(apropiaciones);
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