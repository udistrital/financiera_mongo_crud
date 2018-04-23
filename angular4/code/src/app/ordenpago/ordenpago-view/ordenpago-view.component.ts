import { Component, OnInit } from '@angular/core';
import { OrdenPagoService } from '../../services/ordenpago.service';
import { OrdenPago } from '../../models/ordenpago';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-ordenpago',
  templateUrl: './ordenpago-view.component.html',
  styleUrls: []
})
export class OrdenPagoComponent implements OnInit {

  ordenpagos: OrdenPago[];
  ordenpago: OrdenPago;

  constructor(private ordenpagoService: OrdenPagoService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.ordenpagoService.getOrdenPagos().then(ordenpagos => this.ordenpagos = ordenpagos);
  }

  newOrdenPago(): void {

    this.router.navigate(['/ordenpago/new']).then(() => null);
    this.globals.currentModule = 'OrdenPago';
  }

  editar(ordenpago: OrdenPago): void {
    this.ordenpago = ordenpago;
    this.router.navigate(['/ordenpago/edit', this.ordenpago._id ]);
  }

  borrar(ordenpago: OrdenPago): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar ordenpago?',
      accept: () => {
        this.ordenpagoService.delete(ordenpago._id)
          .then(response => this.ordenpagoService.getOrdenPagos().then(ordenpagos => this.ordenpagos = ordenpagos));
      }
    });
  }
}