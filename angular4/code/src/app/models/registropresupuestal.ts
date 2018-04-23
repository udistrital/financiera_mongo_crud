
import { Disponibilidadapropiacion } from './disponibilidadapropiacion';

export class Registropresupuestal {
   _id: string;
  vigencia:	int;
  fecha_registro:	string;
  estado:	string;
  numero_registro_presupuestal:	int;
  solicitud:	int;
  disponibilidad_apropiacion: DisponibilidadApropiacion[];
}