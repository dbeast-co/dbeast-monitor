import { css } from '@emotion/css';
import { GrafanaTheme2 } from '@grafana/data';

export const getDataSourceListStyles = (theme: GrafanaTheme2) => {
  return {
    container: css`
      margin: 0 !important;
    `,

    header: css`
      display: flex;
      justify-content: center;
      width: 100%;
      margin-block-end: 2rem;

      h1 {
        color: ${theme.colors.text.primary};
        font-size: ${theme.typography.h1.fontSize};
        font-weight: ${theme.typography.h1.fontWeight};
        margin: 0;
      }
    `,

    cardSection: css`
      display: flex;
      justify-content: center;
      width: 100%;
    `,

    cardList: css`
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(500px, 1fr));
      grid-gap: 7rem;
      width: 100%;
      list-style: none;
      margin: 0;
      padding: 0;

      &[data-col="1"] {
        grid-template-columns: repeat(1, 1fr);
      }
    `,

    cardItemWrapper: css`
      display: flex;
      justify-content: center;
      width: 100%;
    `
  };
};
