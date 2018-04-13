import { Component, OnInit } from '@angular/core';
import { MovimientoService } from '../../services/movimiento.service';
import { Movimiento } from '../../models/movimiento';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-movimiento',
  templateUrl: './movimiento-view.component.html',
  styleUrls: []
})
export class MovimientoComponent implements OnInit {

  movimientos: Movimiento[];
  movimiento: Movimiento;

  constructor(private movimientoService: MovimientoService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.movimientoService.getMovimientos().then(movimientos => this.movimientos = movimientos);
  }

  newMovimiento(): void {

    this.router.navigate(['/movimiento/new']).then(() => null);
    this.globals.currentModule = 'Movimiento';
  }

  editar(movimiento: Movimiento): void {
    this.movimiento = movimiento;
    this.router.navigate(['/movimiento/edit', this.movimiento._id ]);
  }

  borrar(movimiento: Movimiento): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar movimiento?',
      accept: () => {
        this.movimientoService.delete(movimiento._id)
          .then(response => this.movimientoService.getMovimientos().then(movimientos => this.movimientos = movimientos));
      }
    });
  }
}