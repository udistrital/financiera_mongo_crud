import { Component, OnInit } from '@angular/core';
import { Apropiacion } from '../../models/apropiacion';
import { ApropiacionService } from '../../services/apropiacion.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-apropiacion-new',
  templateUrl: './apropiacion-new.component.html',
  styleUrls: []
})
export class ApropiacionNewComponent implements OnInit {

  apropiacion: Apropiacion;
  display = false;
  constructor(private apropiacionService: ApropiacionService, private location: Location) { }

  ngOnInit() {
    this.apropiacion = new Apropiacion();
  }

  guardar(apropiacion: Apropiacion): void {

    this.apropiacionService.create(apropiacion);
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