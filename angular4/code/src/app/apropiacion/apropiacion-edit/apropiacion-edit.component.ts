import { Component, OnInit } from '@angular/core';
import { Apropiacion } from '../../models/apropiacion';
import { ApropiacionService } from '../../services/apropiacion.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-apropiacion-edit',
  templateUrl: './apropiacion-edit.component.html',
  styleUrls: []
})
export class ApropiacionEditComponent implements OnInit {

  apropiacion: Apropiacion = new Apropiacion();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private apropiacionService: ApropiacionService) {

  }

  actualizar(apropiacion: Apropiacion): void {
    this.apropiacionService.update(apropiacion).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.apropiacionService.getApropiacion(params['id']))
      .subscribe(apropiacion => this.apropiacion = apropiacion);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}