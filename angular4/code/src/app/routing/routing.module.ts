import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CommonModule } from '@angular/common';
import { OrdenpagoComponent } from '../ordenpago/ordenpago-view/ordenpago-view.component';
import { OrdenpagoNewComponent } from '../ordenpago/ordenpago-new/ordenpago-new.component';
import { OrdenpagoEditComponent } from '../ordenpago/ordenpago-edit/ordenpago-edit.component';


const routes: Routes = [
  { path: 'ordenpago', component: OrdenpagoComponent },
  { path: 'ordenpago/new', component: OrdenpagoNewComponent },
  { path: 'ordenpago/edit/:id', component: OrdenpagoEditComponent },
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