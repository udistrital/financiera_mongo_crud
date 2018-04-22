import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { RegistroPresupuestal } from '../models/registropresupuestal';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class RegistroPresupuestalService {

  private serviceURL = 'http://localhost:8081/v1/registropresupuestal';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getRegistroPresupuestals(): Promise<RegistroPresupuestal[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as RegistroPresupuestal[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getRegistroPresupuestal(id: string): Promise<RegistroPresupuestal> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as RegistroPresupuestal)
      .catch(this.handleError);
  }


  update(registropresupuestal: RegistroPresupuestal): Promise<RegistroPresupuestal> {
    const url = `${this.serviceURL}/${ registropresupuestal._id}`;
    return this.http
      .put(url, JSON.stringify(registropresupuestal), {headers: this.headers})
      .toPromise()
      .then(() => registropresupuestal)
      .catch(this.handleError);
  }


  create(registropresupuestal: RegistroPresupuestal): Promise<RegistroPresupuestal> {
    return this.http
      .post(this.serviceURL, JSON.stringify(registropresupuestal), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as RegistroPresupuestal)
      .catch(this.handleError);
  }

  delete(id: string): Promise<void> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.delete(url, {headers: this.headers})
      .toPromise()
      .then(() => null)
      .catch(this.handleError);
  }

}