// excelDownload.d.ts
import type { DefineComponent } from 'vue';

export interface ExcelDownloadProps {
  fields: {};
  data: any[]; // Or a more specific type if known
}

declare module '*.vue' {
  import type { DefineComponent } from 'vue';
  const component: DefineComponent<{}, {}, any>;
  export default component;
}

const ExcelDownload: DefineComponent<ExcelDownloadProps, {}, any>;
export default ExcelDownload;