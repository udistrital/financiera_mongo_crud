import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { DisponibilidadApropiacion } from '../models/disponibilidadapropiacion';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class DisponibilidadApropiacionService {

  private serviceURL = 'http://localhost:8081/v1/disponibilidadapropiacion';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getDisponibilidadApropiacions(): Promise<DisponibilidadApropiacion[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as DisponibilidadApropiacion[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getDisponibilidadApropiacion(id: string): Promise<DisponibilidadApropiacion> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as DisponibilidadApropiacion)
      .catch(this.handleError);
  }


  update(disponibilidadapropiacion: DisponibilidadApropiacion): Promise<DisponibilidadApropiacion> {
    const url = `${this.serviceURL}/${ disponibilidadapropiacion._id}`;
    return this.http
      .put(url, JSON.stringify(disponibilidadapropiacion), {headers: this.headers})
      .toPromise()
      .then(() => disponibilidadapropiacion)
      .catch(this.handleError);
  }


  create(disponibilidadapropiacion: DisponibilidadApropiacion): Promise<DisponibilidadApropiacion> {
    return this.http
      .post(this.serviceURL, JSON.stringify(disponibilidadapropiacion), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as DisponibilidadApropiacion)
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