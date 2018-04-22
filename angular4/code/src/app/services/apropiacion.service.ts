import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Apropiacion } from '../models/apropiacion';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class ApropiacionService {

  private serviceURL = 'http://localhost:8081/v1/apropiacion';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getApropiacions(): Promise<Apropiacion[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Apropiacion[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getApropiacion(id: string): Promise<Apropiacion> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Apropiacion)
      .catch(this.handleError);
  }


  update(apropiacion: Apropiacion): Promise<Apropiacion> {
    const url = `${this.serviceURL}/${ apropiacion._id}`;
    return this.http
      .put(url, JSON.stringify(apropiacion), {headers: this.headers})
      .toPromise()
      .then(() => apropiacion)
      .catch(this.handleError);
  }


  create(apropiacion: Apropiacion): Promise<Apropiacion> {
    return this.http
      .post(this.serviceURL, JSON.stringify(apropiacion), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Apropiacion)
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