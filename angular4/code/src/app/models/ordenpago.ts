
import { Registropresupuestal } from './registropresupuestal';

export class Ordenpago {
   _id: string;
  vigencia:	int;
  valor_base:	int;
  unidad_ejecutora:	int;
  forma_pago:	int;
  registro_presupuestal: RegistroPresupuestal[];
}