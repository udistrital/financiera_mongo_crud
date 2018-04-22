import { Component, OnInit } from '@angular/core';
import { Registropresupuestal } from '../../models/registropresupuestal';
import { RegistropresupuestalService } from '../../services/registropresupuestal.service';
import { Location } from '@angular/common';


@Component({
  selector: 'app-registropresupuestal-new',
  templateUrl: './registropresupuestal-new.component.html',
  styleUrls: []
})
export class RegistropresupuestalNewComponent implements OnInit {

  registropresupuestal: Registropresupuestal;
  display = false;
  constructor(private registropresupuestalService: RegistropresupuestalService, private location: Location) { }

  ngOnInit() {
    this.registropresupuestal = new Registropresupuestal();
  }

  guardar(registropresupuestal: Registropresupuestal): void {

    this.registropresupuestalService.create(registropresupuestal);
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