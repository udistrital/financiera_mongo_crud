import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { OrdenPago } from '../models/ordenpago';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class OrdenPagoService {

  private serviceURL = 'http://localhost:8081/v1/ordenpago';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getOrdenPagos(): Promise<OrdenPago[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as OrdenPago[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getOrdenPago(id: string): Promise<OrdenPago> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as OrdenPago)
      .catch(this.handleError);
  }


  update(ordenpago: OrdenPago): Promise<OrdenPago> {
    const url = `${this.serviceURL}/${ ordenpago._id}`;
    return this.http
      .put(url, JSON.stringify(ordenpago), {headers: this.headers})
      .toPromise()
      .then(() => ordenpago)
      .catch(this.handleError);
  }


  create(ordenpago: OrdenPago): Promise<OrdenPago> {
    return this.http
      .post(this.serviceURL, JSON.stringify(ordenpago), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as OrdenPago)
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