import { Component, OnInit } from '@angular/core';
import { Movimiento } from '../../models/movimiento';
import { MovimientoService } from '../../services/movimiento.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-movimiento-new',
  templateUrl: './movimiento-new.component.html',
  styleUrls: []
})
export class MovimientoNewComponent implements OnInit {

  movimiento: Movimiento;
  display = false;
  constructor(private movimientoService: MovimientoService, private location: Location) { }

  ngOnInit() {
    this.movimiento = new Movimiento();
  }

  guardar(movimiento: Movimiento): void {

    this.movimientoService.create(movimiento);
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