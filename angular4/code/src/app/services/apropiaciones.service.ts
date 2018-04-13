import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Apropiaciones } from '../models/apropiaciones';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class ApropiacionesService {

  private serviceURL = 'http://localhost:8081/v1/apropiaciones';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getApropiacioness(): Promise<Apropiaciones[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Apropiaciones[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getApropiaciones(id: string): Promise<Apropiaciones> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Apropiaciones)
      .catch(this.handleError);
  }


  update(apropiaciones: Apropiaciones): Promise<Apropiaciones> {
    const url = `${this.serviceURL}/${ apropiaciones._id}`;
    return this.http
      .put(url, JSON.stringify(apropiaciones), {headers: this.headers})
      .toPromise()
      .then(() => apropiaciones)
      .catch(this.handleError);
  }


  create(apropiaciones: Apropiaciones): Promise<Apropiaciones> {
    return this.http
      .post(this.serviceURL, JSON.stringify(apropiaciones), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Apropiaciones)
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