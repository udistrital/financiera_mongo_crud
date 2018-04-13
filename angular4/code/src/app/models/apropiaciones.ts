
import { Movimiento } from './movimiento';

export class Apropiaciones {
   _id: string;
  vigencia:	int;
  valor_inicial:	float64;
  movimientos: Movimiento[];
}