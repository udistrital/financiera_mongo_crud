import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';
import { Rubro } from '../models/rubro';

import 'rxjs/add/operator/toPromise';


@Injectable()
export class RubroService {

  private serviceURL = 'http://localhost:8081/v1/rubro';
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  getRubros(): Promise<Rubro[]> {
    return this.http.get(this.serviceURL)
      .toPromise()
      .then(response => response.json() as Rubro[])
      .catch(this.handleError)

  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getRubro(id: string): Promise<Rubro> {
    const url = `${this.serviceURL}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json()[0] as Rubro)
      .catch(this.handleError);
  }


  update(rubro: Rubro): Promise<Rubro> {
    const url = `${this.serviceURL}/${ rubro._id}`;
    return this.http
      .put(url, JSON.stringify(rubro), {headers: this.headers})
      .toPromise()
      .then(() => rubro)
      .catch(this.handleError);
  }


  create(rubro: Rubro): Promise<Rubro> {
    return this.http
      .post(this.serviceURL, JSON.stringify(rubro), {headers: this.headers})
      .toPromise()
      .then(res => res.json().data as Rubro)
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