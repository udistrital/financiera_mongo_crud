
import { Apropiaciones } from './apropiaciones';

export class Rubro {
   _id: string;
  codigo:	string;
  Nombre:	string;
  entidad:	string;
  descripcion:	string;
  unidad_ejecutora:	int;
  apropiaciones: Apropiaciones[];
}