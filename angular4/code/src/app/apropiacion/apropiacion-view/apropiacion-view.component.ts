import { Component, OnInit } from '@angular/core';
import { ApropiacionService } from '../../services/apropiacion.service';
import { Apropiacion } from '../../models/apropiacion';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-apropiacion',
  templateUrl: './apropiacion-view.component.html',
  styleUrls: []
})
export class ApropiacionComponent implements OnInit {

  apropiacions: Apropiacion[];
  apropiacion: Apropiacion;

  constructor(private apropiacionService: ApropiacionService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.apropiacionService.getApropiacions().then(apropiacions => this.apropiacions = apropiacions);
  }

  newApropiacion(): void {

    this.router.navigate(['/apropiacion/new']).then(() => null);
    this.globals.currentModule = 'Apropiacion';
  }

  editar(apropiacion: Apropiacion): void {
    this.apropiacion = apropiacion;
    this.router.navigate(['/apropiacion/edit', this.apropiacion._id ]);
  }

  borrar(apropiacion: Apropiacion): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar apropiacion?',
      accept: () => {
        this.apropiacionService.delete(apropiacion._id)
          .then(response => this.apropiacionService.getApropiacions().then(apropiacions => this.apropiacions = apropiacions));
      }
    });
  }
}