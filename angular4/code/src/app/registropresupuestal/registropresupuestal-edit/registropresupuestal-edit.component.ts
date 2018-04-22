import { Component, OnInit } from '@angular/core';
import { Registropresupuestal } from '../../models/registropresupuestal';
import { RegistropresupuestalService } from '../../services/registropresupuestal.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-registropresupuestal-edit',
  templateUrl: './registropresupuestal-edit.component.html',
  styleUrls: []
})
export class RegistropresupuestalEditComponent implements OnInit {

  registropresupuestal: Registropresupuestal = new Registropresupuestal();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private registropresupuestalService: RegistropresupuestalService) {

  }

  actualizar(registropresupuestal: Registropresupuestal): void {
    this.registropresupuestalService.update(registropresupuestal).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.registropresupuestalService.getRegistropresupuestal(params['id']))
      .subscribe(registropresupuestal => this.registropresupuestal = registropresupuestal);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}