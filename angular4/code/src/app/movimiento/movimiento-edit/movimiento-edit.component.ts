import { Component, OnInit } from '@angular/core';
import { Movimiento } from '../../models/movimiento';
import { MovimientoService } from '../../services/movimiento.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-movimiento-edit',
  templateUrl: './movimiento-edit.component.html',
  styleUrls: []
})
export class MovimientoEditComponent implements OnInit {

  movimiento: Movimiento = new Movimiento();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private movimientoService: MovimientoService) {

  }

  actualizar(movimiento: Movimiento): void {
    this.movimientoService.update(movimiento).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.movimientoService.getMovimiento(params['id']))
      .subscribe(movimiento => this.movimiento = movimiento);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}