
import { Movimiento } from './movimiento';
import { Rubro } from './rubro';

export class Apropiacion {
   _id: string;
  vigencia:	int;
  mes:	int;
  valor_inicial:	int;
  movimiento: Movimiento[];
  rubro: Rubro[];
}