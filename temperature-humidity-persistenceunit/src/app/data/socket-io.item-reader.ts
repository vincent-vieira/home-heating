import { ItemReader } from './item-reader';
import { Measurement } from './measurement';
import { Subject, Observable } from 'rxjs';
import { connect } from 'socket.io-client';

//For Promise type inference, see https://github.com/Microsoft/TypeScript/issues/14871
export class SocketIOItemReader implements ItemReader<Measurement, void, void> {

  private socket: SocketIOClient.Socket;
  private subject: Subject<Measurement> = new Subject<Measurement>();

  setup(environment): Promise<void> {
    this.socket = connect(<string>environment.SOCKET_IO_REMOTE);
    return new Promise<void>((resolve) => resolve());
  }

  bindMeasurement(): Observable<Measurement> {
    this.socket.on('new-measure', (data) => this.subject.next(data));
    return this.subject;
  }

  teardown(): Promise<void> {
    this.socket.disconnect();
    return new Promise<void>((resolve) => resolve());
  }
}