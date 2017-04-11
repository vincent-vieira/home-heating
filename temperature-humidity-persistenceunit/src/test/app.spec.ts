import { App } from '../app/app';
import { Observable } from 'rxjs';

describe('Application', () => {
  let noopReader = jasmine.createSpyObj('reader', ['bindMeasurement', 'setup', 'teardown']);
  let noopWriter = jasmine.createSpyObj('writer', ['write', 'setup', 'teardown']);


  describe('setup', () => {
    let sampleValue = {
      date: new Date().getTime() * 1000,
      temperature: 23,
      humidity: 30.1
    };

    beforeEach(() => {
      noopReader.setup.and.returnValue(new Promise<void>((resolve) => resolve()));
      noopReader.teardown.and.returnValue(new Promise<void>((resolve) => resolve()));
      noopReader.bindMeasurement.and.returnValue(Observable.of(sampleValue));

      noopWriter.setup.and.returnValue(new Promise<void>((resolve) => resolve()));
      noopWriter.teardown.and.returnValue(new Promise<void>((resolve) => resolve()));
    });

    it('should properly follow the measurement flow', (done) => {
      new App(noopReader, noopWriter).setup().subscribe(() => {
        expect(noopWriter.setup).toHaveBeenCalled();
        expect(noopReader.setup).toHaveBeenCalled();
        done();
      });
    });
  });

  describe('data flow', () => {
    beforeEach(() => {
      noopReader.setup.and.returnValue(new Promise<void>((resolve) => resolve()));
      noopReader.teardown.and.returnValue(new Promise<void>((resolve) => resolve()));
      noopReader.bindMeasurement.and.returnValue(Observable.interval(500).map(clockTime => {
        return {
          date: new Date().getTime() * 1000,
          temperature: 23 + (clockTime / 1000),
          humidity: 30.1
        };
      }));

      noopWriter.setup.and.returnValue(new Promise<void>((resolve) => resolve()));
      noopWriter.teardown.and.returnValue(new Promise<void>((resolve) => resolve()));
    });

    it('should properly read a value once', (done) => {
      new App(noopReader, noopWriter).setup().subscribe(measurement => {
        expect(measurement.date).toBeDefined();
        expect(measurement.temperature).toBeDefined();
        expect(measurement.humidity).toEqual(30.1);

        done();
      });
    });

    it('should properly read values multiple times', (done) => {
      let count = 0;
      new App(noopReader, noopWriter).setup().subscribe(measurement => {
        expect(measurement.date).toBeDefined();
        expect(measurement.temperature).toBeDefined();
        expect(measurement.humidity).toEqual(30.1);

        ++count;
        if(count === 3){
          done();
        }
      });
    });
  })
});