import { FieldType } from '@grafana/data';
import { parseValues } from './parseValues';

test('parse numbers', () => {
  const values = [2005, 2006];

  expect(() => parseValues(values, FieldType.boolean)).toThrow();
  expect(parseValues(values, FieldType.number)).toStrictEqual([2005, 2006]);
  expect(parseValues(values, FieldType.string)).toStrictEqual(['2005', '2006']);

  // Numbers are assumed to be epoch time in seconds.
  expect(parseValues(values, FieldType.time)).toStrictEqual([2005000, 2006000]);
});

test('parse numbers from strings', () => {
  const values = ['2005', '2006'];

  expect(() => parseValues(values, FieldType.boolean)).toThrow();
  expect(parseValues(values, FieldType.number)).toStrictEqual([2005, 2006]);
  expect(parseValues(values, FieldType.string)).toStrictEqual(['2005', '2006']);

  // Values get parsed as ISO 8601 strings.
  // TODO: Default to UTC if no time zone is given.
  // expect(parseValues(values, FieldType.time)).toStrictEqual([1104537600000, 1136073600000]);
});

test('parse booleans', () => {
  const values = [false, true, false];

  expect(parseValues(values, FieldType.boolean)).toStrictEqual([false, true, false]);
  expect(parseValues(values, FieldType.number)).toStrictEqual([NaN, NaN, NaN]);
  expect(parseValues(values, FieldType.string)).toStrictEqual(['false', 'true', 'false']);
  expect(() => parseValues(values, FieldType.time)).toThrow();
});

test('parse booleans from strings', () => {
  const values = ['false', 'true', 'false'];

  expect(parseValues(values, FieldType.boolean)).toStrictEqual([false, true, false]);
  expect(parseValues(values, FieldType.number)).toStrictEqual([NaN, NaN, NaN]);
  expect(parseValues(values, FieldType.string)).toStrictEqual(['false', 'true', 'false']);
  expect(parseValues(values, FieldType.time)).toStrictEqual([NaN, NaN, NaN]);
});

test('parse nullable strings', () => {
  const values = ['foo', 'bar', null];

  expect(() => parseValues(values, FieldType.boolean)).toThrow();
  expect(parseValues(values, FieldType.number)).toStrictEqual([NaN, NaN, null]);
  expect(parseValues(values, FieldType.string)).toStrictEqual(['foo', 'bar', null]);
  expect(parseValues(values, FieldType.time)).toStrictEqual([NaN, NaN, null]);
});

test('parse nullable numbers', () => {
  const values = [2005, null, 2006];

  expect(() => parseValues(values, FieldType.boolean)).toThrow();
  expect(parseValues(values, FieldType.number)).toStrictEqual([2005, null, 2006]);
  expect(parseValues(values, FieldType.string)).toStrictEqual(['2005', null, '2006']);

  // Numbers are assumed to be epoch time in seconds.
  expect(parseValues(values, FieldType.time)).toStrictEqual([2005000, null, 2006000]);
});

test('parse time values with missing entries', () => {
  expect(parseValues([null, undefined], FieldType.time)).toStrictEqual([null, undefined]);
  expect(parseValues(['2021-05-17T20:48:09.000Z', undefined], FieldType.time)).toStrictEqual([
    1621284489000,
    undefined,
  ]);
  expect(parseValues([2005, undefined, 2006], FieldType.time)).toStrictEqual([2005000, undefined, 2006000]);
});

test('parse numeric fields strictly', () => {
  expect(parseValues(['12', '1.5e2', '-.25'], FieldType.number)).toStrictEqual([12, 150, -0.25]);
  expect(parseValues(['12oops', '  ', true, undefined], FieldType.number)).toStrictEqual([
    NaN,
    NaN,
    NaN,
    undefined,
  ]);
});

test('parse nullable booleans', () => {
  const values = [null, true, false];

  expect(parseValues(values, FieldType.boolean)).toStrictEqual([null, true, false]);
  expect(parseValues(values, FieldType.number)).toStrictEqual([null, NaN, NaN]);
  expect(parseValues(values, FieldType.string)).toStrictEqual([null, 'true', 'false']);
  expect(() => parseValues(values, FieldType.time)).toThrow();
});

test('parse objects', () => {
  const values = [[1, 2, 3], { '123': 123 }, null];
  expect(parseValues(values, FieldType.string)).toStrictEqual(['[1,2,3]', '{"123":123}', null]);
});
