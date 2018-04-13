import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CommonModule } from '@angular/common';
import { RubroComponent } from '../rubro/rubro-view/rubro-view.component';
import { RubroNewComponent } from '../rubro/rubro-new/rubro-new.component';
import { RubroEditComponent } from '../rubro/rubro-edit/rubro-edit.component';


const routes: Routes = [
  { path: 'rubro', component: RubroComponent },
  { path: 'rubro/new', component: RubroNewComponent },
  { path: 'rubro/edit/:id', component: RubroEditComponent },
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