import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Movimiento } from '../models/movimiento';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class MovimientoService {

  private serviceURL = 'http://localhost:8081/v1/movimiento';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getMovimientos(): Promise<Movimiento[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Movimiento[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getMovimiento(id: string): Promise<Movimiento> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Movimiento)
      .catch(this.handleError);
  }


  update(movimiento: Movimiento): Promise<Movimiento> {
    const url = `${this.serviceURL}/${ movimiento._id}`;
    return this.http
      .put(url, JSON.stringify(movimiento), {headers: this.headers})
      .toPromise()
      .then(() => movimiento)
      .catch(this.handleError);
  }


  create(movimiento: Movimiento): Promise<Movimiento> {
    return this.http
      .post(this.serviceURL, JSON.stringify(movimiento), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Movimiento)
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