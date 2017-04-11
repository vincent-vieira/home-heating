import { Manageable } from './manageable';
import { Observable } from 'rxjs';

export interface ItemReader<Input, Setup, Teardown> extends Manageable<Setup, Teardown> {
  bindMeasurement(): Observable<Input>;
}