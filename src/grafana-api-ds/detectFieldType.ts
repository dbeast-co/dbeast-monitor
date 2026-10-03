import { FieldType } from '@grafana/data';

import moment from 'moment'; // eslint-disable-line no-restricted-imports

/**
 * Detects the field type from an array of values.
 */
export const detectFieldType = (values: any[]): FieldType => {
  // Inspect each value once to avoid repeatedly traversing large result sets.
  let hasValue = false;
  let allISO = true;
  let allNumbers = true;
  let allBooleans = true;

  for (const value of values) {
    if (value === null) {
      continue;
    }

    hasValue = true;
    if (!moment(value, [moment.defaultFormat, 'YYYY-MM-DD'], true).isValid()) {
      allISO = false;
    }
    if (typeof value !== 'number') {
      allNumbers = false;
    }
    if (typeof value !== 'boolean') {
      allBooleans = false;
    }
  }

  if (!hasValue) {
    return FieldType.string;
  }

  if (allISO) {
    return FieldType.time;
  }

  if (allNumbers) {
    return FieldType.number;
  }

  if (allBooleans) {
    return FieldType.boolean;
  }

  return FieldType.string;
};
