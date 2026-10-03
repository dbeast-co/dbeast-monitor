import { FieldType } from '@grafana/data';
import dayjs from 'dayjs';

/**
 * parseValues converts values to the given field type.
 */
export const parseValues = (values: any[], type: FieldType): any[] => {
  switch (type) {
    case FieldType.time:
      return parseTimeValues(values);
    case FieldType.string:
      return values.every((_) => typeof _ === 'string')
        ? values
        : values.map((_) => {
            if (_ === null) {
              return _;
            } else if (typeof _ === 'object') {
              return JSON.stringify(_);
            } else {
              return _.toString();
            }
          });
    case FieldType.number:
      return parseNumberValues(values);
    case FieldType.boolean:
      return values.every((_) => typeof _ === 'boolean')
        ? values
        : values.map((_) => {
            if (_ === null) {
              return _;
            }

            switch (_.toString()) {
              case '0':
              case 'false':
              case 'FALSE':
              case 'False':
                return false;
              case '1':
              case 'true':
              case 'TRUE':
              case 'True':
                return true;
              default:
                throw new Error('Found non-boolean values in a field of type boolean: ' + _.toString());
            }
          });
    default:
      throw new Error('Unsupported field type');
  }
};

const parseTimeValues = (values: any[]): any[] => {
  const nonNullValues = values.filter((value) => value !== null && value !== undefined);

  if (nonNullValues.length === 0) {
    return values;
  }

  if (nonNullValues.every((value) => typeof value === 'string')) {
    return values.map((value) => (value === null || value === undefined ? value : dayjs(value).valueOf()));
  }

  if (nonNullValues.every((value) => typeof value === 'number')) {
    const ms = 1_000_000_000_000;

    // If there are no "big" numbers, assume seconds.
    if (nonNullValues.every((value) => value < ms)) {
      return values.map((value) => (value === null || value === undefined ? value : value * 1000.0));
    }

    // ... otherwise assume milliseconds.
    return values;
  }

  throw new Error('Unsupported time property');
};

const parseNumberValues = (values: any[]): any[] => {
  if (values.every((value) => typeof value === 'number')) {
    return values;
  }

  return values.map((value) => {
    if (value === null || value === undefined || typeof value === 'number') {
      return value;
    }
    if (typeof value !== 'string') {
      return NaN;
    }

    const trimmed = value.trim();
    return /^[+-]?(?:\d+\.?\d*|\.\d+)(?:[eE][+-]?\d+)?$/.test(trimmed) ? Number(trimmed) : NaN;
  });
};
