import { Component, OnInit } from '@angular/core';
import { ApropiacionesService } from '../../services/apropiaciones.service';
import { Apropiaciones } from '../../models/apropiaciones';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-apropiaciones',
  templateUrl: './apropiaciones-view.component.html',
  styleUrls: []
})
export class ApropiacionesComponent implements OnInit {

  apropiacioness: Apropiaciones[];
  apropiaciones: Apropiaciones;

  constructor(private apropiacionesService: ApropiacionesService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.apropiacionesService.getApropiacioness().then(apropiacioness => this.apropiacioness = apropiacioness);
  }

  newApropiaciones(): void {

    this.router.navigate(['/apropiaciones/new']).then(() => null);
    this.globals.currentModule = 'Apropiaciones';
  }

  editar(apropiaciones: Apropiaciones): void {
    this.apropiaciones = apropiaciones;
    this.router.navigate(['/apropiaciones/edit', this.apropiaciones._id ]);
  }

  borrar(apropiaciones: Apropiaciones): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar apropiaciones?',
      accept: () => {
        this.apropiacionesService.delete(apropiaciones._id)
          .then(response => this.apropiacionesService.getApropiacioness().then(apropiacioness => this.apropiacioness = apropiacioness));
      }
    });
  }
}