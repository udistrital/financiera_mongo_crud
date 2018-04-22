import { Component, OnInit } from '@angular/core';
import { RegistroPresupuestalService } from '../../services/registropresupuestal.service';
import { RegistroPresupuestal } from '../../models/registropresupuestal';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-registropresupuestal',
  templateUrl: './registropresupuestal-view.component.html',
  styleUrls: []
})
export class RegistroPresupuestalComponent implements OnInit {

  registropresupuestals: RegistroPresupuestal[];
  registropresupuestal: RegistroPresupuestal;

  constructor(private registropresupuestalService: RegistroPresupuestalService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.registropresupuestalService.getRegistroPresupuestals().then(registropresupuestals => this.registropresupuestals = registropresupuestals);
  }

  newRegistroPresupuestal(): void {

    this.router.navigate(['/registropresupuestal/new']).then(() => null);
    this.globals.currentModule = 'RegistroPresupuestal';
  }

  editar(registropresupuestal: RegistroPresupuestal): void {
    this.registropresupuestal = registropresupuestal;
    this.router.navigate(['/registropresupuestal/edit', this.registropresupuestal._id ]);
  }

  borrar(registropresupuestal: RegistroPresupuestal): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar registropresupuestal?',
      accept: () => {
        this.registropresupuestalService.delete(registropresupuestal._id)
          .then(response => this.registropresupuestalService.getRegistroPresupuestals().then(registropresupuestals => this.registropresupuestals = registropresupuestals));
      }
    });
  }
}