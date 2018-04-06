import { Component, OnInit } from '@angular/core';
import { RubroService } from '../../services/rubro.service';
import { Rubro } from '../../models/rubro';
import { Router} from '@angular/router';
import { GlobalsComponent } from '../../globals/globals.component';
import { ConfirmationService } from 'primeng/primeng';

@Component({
  selector: 'app-rubro',
  templateUrl: './rubro-view.component.html',
  styleUrls: []
})
export class RubroComponent implements OnInit {

  rubros: Rubro[];
  rubro: Rubro;

  constructor(private rubroService: RubroService,
      private router: Router, private globals: GlobalsComponent,
      private confirmationService: ConfirmationService) {
      this.globals = globals;
  }

  ngOnInit(): void {
    this.rubroService.getRubros().then(rubros => this.rubros = rubros);
  }

  newRubro(): void {

    this.router.navigate(['/rubro/new']).then(() => null);
    this.globals.currentModule = 'Rubro';
  }

  editar(rubro: Rubro): void {
    this.rubro = rubro;
    this.router.navigate(['/rubro/edit', this.rubro._id ]);
  }

  borrar(rubro: Rubro): void {
    this.confirmationService.confirm({
      message: 'Esta seguro que quiere borrar rubro?',
      accept: () => {
        this.rubroService.delete(rubro._id)
          .then(response => this.rubroService.getRubros().then(rubros => this.rubros = rubros));
      }
    });
  }
}