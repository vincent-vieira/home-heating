import { Observable, Subscription } from 'rxjs';

import { ItemWriter } from './data/item-writer';
import { ItemReader } from './data/item-reader';
import { Measurement } from './data/measurement';

export class App {

  constructor(
    private reader: ItemReader<Measurement, void, void>,
    private writer: ItemWriter<Measurement, void, void>
  ) {}

  public setup(): Observable<Measurement> {
    return Observable
      .fromPromise(this.writer.setup(process.env))
      .withLatestFrom(Observable.fromPromise(this.reader.setup(process.env)))
      .switchMap(() => this.reader.bindMeasurement())
      .finally(() => {
        this.writer.teardown();
        this.reader.teardown();
      });
  }
}