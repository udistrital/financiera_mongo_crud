import { Component, OnInit } from '@angular/core';
import { Disponibilidadapropiacion } from '../../models/disponibilidadapropiacion';
import { DisponibilidadapropiacionService } from '../../services/disponibilidadapropiacion.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-disponibilidadapropiacion-new',
  templateUrl: './disponibilidadapropiacion-new.component.html',
  styleUrls: []
})
export class DisponibilidadapropiacionNewComponent implements OnInit {

  disponibilidadapropiacion: Disponibilidadapropiacion;
  display = false;
  constructor(private disponibilidadapropiacionService: DisponibilidadapropiacionService, private location: Location) { }

  ngOnInit() {
    this.disponibilidadapropiacion = new Disponibilidadapropiacion();
  }

  guardar(disponibilidadapropiacion: Disponibilidadapropiacion): void {

    this.disponibilidadapropiacionService.create(disponibilidadapropiacion);
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