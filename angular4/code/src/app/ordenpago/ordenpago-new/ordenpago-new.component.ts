import { Component, OnInit } from '@angular/core';
import { Ordenpago } from '../../models/ordenpago';
import { OrdenpagoService } from '../../services/ordenpago.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-ordenpago-new',
  templateUrl: './ordenpago-new.component.html',
  styleUrls: []
})
export class OrdenpagoNewComponent implements OnInit {

  ordenpago: Ordenpago;
  display = false;
  constructor(private ordenpagoService: OrdenpagoService, private location: Location) { }

  ngOnInit() {
    this.ordenpago = new Ordenpago();
  }

  guardar(ordenpago: Ordenpago): void {

    this.ordenpagoService.create(ordenpago);
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