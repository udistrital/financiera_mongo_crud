import { Component, OnInit } from '@angular/core';
import { Ordenpago } from '../../models/ordenpago';
import { OrdenpagoService } from '../../services/ordenpago.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-ordenpago-edit',
  templateUrl: './ordenpago-edit.component.html',
  styleUrls: []
})
export class OrdenpagoEditComponent implements OnInit {

  ordenpago: Ordenpago = new Ordenpago();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private ordenpagoService: OrdenpagoService) {

  }

  actualizar(ordenpago: Ordenpago): void {
    this.ordenpagoService.update(ordenpago).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.ordenpagoService.getOrdenpago(params['id']))
      .subscribe(ordenpago => this.ordenpago = ordenpago);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}