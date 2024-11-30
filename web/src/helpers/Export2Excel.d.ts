import * as XLSX from 'xlsx'; // Import if necessary

interface ExportOptions {
  multiHeader?: any[][]; //Type of multiHeader is unclear, needs clarification.
  header?: any[];       //Type of header is unclear, needs clarification.
  data: any[][];         //Data is an array of arrays
  sheetName?: string;
  filename?: string;
  merges?: string[];    //Type of merges unclear, assuming array of strings.
  autoWidth?: boolean;
  bookType?: 'xlsx' | 'xls' | 'csv'; //Explicit types for bookType
}


export function export_json_to_excel(options?: ExportOptions): void;