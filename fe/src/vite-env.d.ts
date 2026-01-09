interface ViteTypeOptions {
  strictImportMetaEnv: unknown;
}

interface ImportMetaEnv {
  readonly VITE_MINESWEEPER_API_URL: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
