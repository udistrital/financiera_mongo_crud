import { Component, OnInit } from '@angular/core';
import { Rubro } from '../../models/rubro';
import { RubroService } from '../../services/rubro.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-rubro-edit',
  templateUrl: './rubro-edit.component.html',
  styleUrls: []
})
export class RubroEditComponent implements OnInit {

  rubro: Rubro = new Rubro();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private rubroService: RubroService) {

  }

  actualizar(rubro: Rubro): void {
    this.rubroService.update(rubro).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.rubroService.getRubro(params['id']))
      .subscribe(rubro => this.rubro = rubro);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}