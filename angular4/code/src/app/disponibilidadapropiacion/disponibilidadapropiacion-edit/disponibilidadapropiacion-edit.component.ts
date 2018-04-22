import { Component, OnInit } from '@angular/core';
import { Disponibilidadapropiacion } from '../../models/disponibilidadapropiacion';
import { DisponibilidadapropiacionService } from '../../services/disponibilidadapropiacion.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-disponibilidadapropiacion-edit',
  templateUrl: './disponibilidadapropiacion-edit.component.html',
  styleUrls: []
})
export class DisponibilidadapropiacionEditComponent implements OnInit {

  disponibilidadapropiacion: Disponibilidadapropiacion = new Disponibilidadapropiacion();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private disponibilidadapropiacionService: DisponibilidadapropiacionService) {

  }

  actualizar(disponibilidadapropiacion: Disponibilidadapropiacion): void {
    this.disponibilidadapropiacionService.update(disponibilidadapropiacion).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.disponibilidadapropiacionService.getDisponibilidadapropiacion(params['id']))
      .subscribe(disponibilidadapropiacion => this.disponibilidadapropiacion = disponibilidadapropiacion);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}