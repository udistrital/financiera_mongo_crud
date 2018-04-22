import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CommonModule } from '@angular/common';
import { RegistropresupuestalComponent } from '../registropresupuestal/registropresupuestal-view/registropresupuestal-view.component';
import { RegistropresupuestalNewComponent } from '../registropresupuestal/registropresupuestal-new/registropresupuestal-new.component';
import { RegistropresupuestalEditComponent } from '../registropresupuestal/registropresupuestal-edit/registropresupuestal-edit.component';


const routes: Routes = [
  { path: 'registropresupuestal', component: RegistropresupuestalComponent },
  { path: 'registropresupuestal/new', component: RegistropresupuestalNewComponent },
  { path: 'registropresupuestal/edit/:id', component: RegistropresupuestalEditComponent },
];

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forRoot(routes)
  ],
  exports: [RouterModule],
  declarations: []
})
export class RoutingModule { }