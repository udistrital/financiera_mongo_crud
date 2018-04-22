import { Component, OnInit } from '@angular/core';
import { DisponibilidadApropiacionService } from '../../services/disponibilidadapropiacion.service';
import { DisponibilidadApropiacion } from '../../models/disponibilidadapropiacion';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-disponibilidadapropiacion',
  templateUrl: './disponibilidadapropiacion-view.component.html',
  styleUrls: []
})
export class DisponibilidadApropiacionComponent implements OnInit {

  disponibilidadapropiacions: DisponibilidadApropiacion[];
  disponibilidadapropiacion: DisponibilidadApropiacion;

  constructor(private disponibilidadapropiacionService: DisponibilidadApropiacionService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.disponibilidadapropiacionService.getDisponibilidadApropiacions().then(disponibilidadapropiacions => this.disponibilidadapropiacions = disponibilidadapropiacions);
  }

  newDisponibilidadApropiacion(): void {

    this.router.navigate(['/disponibilidadapropiacion/new']).then(() => null);
    this.globals.currentModule = 'DisponibilidadApropiacion';
  }

  editar(disponibilidadapropiacion: DisponibilidadApropiacion): void {
    this.disponibilidadapropiacion = disponibilidadapropiacion;
    this.router.navigate(['/disponibilidadapropiacion/edit', this.disponibilidadapropiacion._id ]);
  }

  borrar(disponibilidadapropiacion: DisponibilidadApropiacion): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar disponibilidadapropiacion?',
      accept: () => {
        this.disponibilidadapropiacionService.delete(disponibilidadapropiacion._id)
          .then(response => this.disponibilidadapropiacionService.getDisponibilidadApropiacions().then(disponibilidadapropiacions => this.disponibilidadapropiacions = disponibilidadapropiacions));
      }
    });
  }
}