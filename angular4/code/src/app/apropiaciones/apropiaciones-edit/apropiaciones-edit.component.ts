import { Component, OnInit } from '@angular/core';
import { Apropiaciones } from '../../models/apropiaciones';
import { ApropiacionesService } from '../../services/apropiaciones.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-apropiaciones-edit',
  templateUrl: './apropiaciones-edit.component.html',
  styleUrls: []
})
export class ApropiacionesEditComponent implements OnInit {

  apropiaciones: Apropiaciones = new Apropiaciones();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private apropiacionesService: ApropiacionesService) {

  }

  actualizar(apropiaciones: Apropiaciones): void {
    this.apropiacionesService.update(apropiaciones).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.apropiacionesService.getApropiaciones(params['id']))
      .subscribe(apropiaciones => this.apropiaciones = apropiaciones);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}